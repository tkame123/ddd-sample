package message

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/tkame123/ddd-sample/app/billilng_api/adapter/message/event_handler"
	"github.com/tkame123/ddd-sample/app/billilng_api/di/provider"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/service"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/lib/sqs_consumer"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type CommandConsumer struct {
	cfg       *provider.ConsumerConfig
	sqsClient *sqs.Client
	queueUrl  string
	svc       service.CreateBill
}

func NewCommandConsumer(
	cfg *provider.ConsumerConfig,
	sqsClient *sqs.Client,
	svc service.CreateBill,
) *CommandConsumer {
	return &CommandConsumer{
		cfg:       cfg,
		sqsClient: sqsClient,
		queueUrl:  cfg.Command.QueueUrl,
		svc:       svc,
	}
}

func (e *CommandConsumer) Run() {
	// コンテキストとキャンセル関数を作成
	ctxPolling, ctxPollingCancel := context.WithCancel(context.Background())
	ctxWorker, ctxWorkerCancel := context.WithCancel(context.Background())

	// SQS コンシューマを作成
	wgPolling := new(sync.WaitGroup)
	wgPolling.Add(1)
	consumer := sqs_consumer.NewSQSConsumer(e.sqsClient, e.queueUrl, wgPolling)

	// メッセージをやり取りするチャネルを作成
	messagesChan := make(chan *types.Message, e.cfg.Command.MessageChan) // バッファ付きのチャネル

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

func (e *CommandConsumer) workerHandler(ctx context.Context, msg *types.Message) error {
	ev, err := event_helper.ParseMessageFromSQS(msg)
	if err != nil {
		return err
	}

	err = e.processEvent(ctx, ev)
	if err != nil {
		return err
	}

	err = e.deleteMessage(ctx, msg)
	if err != nil {
		// TODO Transactional Outbox Patternの導入までは通知して手動対応ってなるのだろうか。。。
		log.Printf("failed to delete message %v", err)
	}

	return nil
}

func (e *CommandConsumer) processEvent(ctx context.Context, mes *message.Message) error {
	if !event_handler.IsCreateBillEvent(mes.Subject.Type) {
		return fmt.Errorf("invalid event: %s", mes.Subject.Type)
	}
	err := event_handler.EventMap[mes.Subject.Type](e.svc).Handler(ctx, mes)
	if err != nil {
		return err
	}

	return nil
}

func (e *CommandConsumer) deleteMessage(ctx context.Context, msg *types.Message) error {
	_, err := e.sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(e.queueUrl),
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}