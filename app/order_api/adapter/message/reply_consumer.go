package message

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/lib/sqs_consumer"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ReplyConsumer struct {
	cfg        *provider.ConsumerConfig
	sqsClient  *sqs.Client
	queueUrl   string
	rep        repository.Repository
	orderSVC   service.CreateOrder
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI
}

func NewReplyConsumer(
	cfg *provider.ConsumerConfig,
	sqsClient *sqs.Client,
	rep repository.Repository,
	orderSVC service.CreateOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) *ReplyConsumer {
	return &ReplyConsumer{
		cfg:        cfg,
		sqsClient:  sqsClient,
		queueUrl:   cfg.Reply.QueueUrl,
		rep:        rep,
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,
	}
}

func (e *ReplyConsumer) Run() {
	// コンテキストとキャンセル関数を作成
	ctxPolling, ctxPollingCancel := context.WithCancel(context.Background())
	ctxWorker, ctxWorkerCancel := context.WithCancel(context.Background())

	// SQS コンシューマを作成
	wgPolling := new(sync.WaitGroup)
	wgPolling.Add(1)
	consumer := sqs_consumer.NewSQSConsumer(e.sqsClient, e.queueUrl, wgPolling)

	// メッセージをやり取りするチャネルを作成
	messagesChan := make(chan *types.Message, e.cfg.Event.MessageChan) // バッファ付きのチャネル

	// Poller を起動
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Recovered from panic:", r)
			}
		}()

		consumer.PollMessages(
			ctxPolling,
			messagesChan,
			e.cfg.Event.MaxMessages,
			e.cfg.Event.PollingWaitTimeSecond,
			e.cfg.Event.VisibilityTimeoutSecond,
		)
	}()

	// ワーカーを並列に起動
	wgWorker := new(sync.WaitGroup)
	for i := 0; i < e.cfg.Event.MaxWorkers; i++ {
		wgWorker.Add(1)
		worker := sqs_consumer.NewWorker(i, wgWorker, e.workerHandler)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered from panic:", r)
				}
			}()
			worker.Start(ctxWorker, messagesChan)
		}()
	}

	// SIGINT (Ctrl+C) で停止
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down...")

	// pollingを停止する
	ctxPollingCancel()
	wgPolling.Wait()

	// workerを停止する
	ctxWorkerCancel()
	wgWorker.Wait()
}

func (e *ReplyConsumer) workerHandler(ctx context.Context, msg *types.Message) error {
	rollback := func(ctx context.Context, MessageId string, cause error) error {
		err := e.rep.ProcessedMessageDelete(ctx, MessageId)
		if err != nil {
			return fmt.Errorf("failed to rollback: %w", errors.Join(cause, err))
		}
		return nil
	}

	err := e.rep.ProcessedMessageSave(ctx, *msg.MessageId)
	if ent.IsConstraintError(err) {
		return fmt.Errorf("message already processed: %s", *msg.MessageId)
	} else if err != nil {
		return fmt.Errorf("failed to save processed message: %w", err)
	}

	mes, err := event_helper.ParseMessageFromSQS(msg)
	if err != nil {
		return rollback(ctx, *msg.MessageId, err)
	}

	err = e.processEvent(ctx, mes)
	if err != nil {
		return fmt.Errorf("failed to process event: %w", err)
	}

	err = e.deleteMessage(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	return nil
}

func (e *ReplyConsumer) processEvent(ctx context.Context, mes *message.Message) error {
	err := NewEventHandler(
		e.rep,
		e.orderSVC,
		e.kitchenAPI,
		e.billingAPI).
		Handler(ctx, mes)
	if err != nil {
		return err
	}

	return nil
}

func (e *ReplyConsumer) deleteMessage(ctx context.Context, msg *types.Message) error {
	_, err := e.sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(e.queueUrl),
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
