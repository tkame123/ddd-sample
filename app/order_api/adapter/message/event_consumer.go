package message

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga/event_handler"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/lib/sqs_consumer"
	"github.com/tkame123/ddd-sample/proto/message"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type EventConsumer struct {
	cfg        *provider.ConsumerConfig
	sqsClient  *sqs.Client
	queueUrl   string
	rep        repository.Repository
	orderSVC   service.CreateOrder
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI
}

func NewEventConsumer(
	cfg *provider.ConsumerConfig,
	envCfg *provider.EnvConfig,
	sqsClient *sqs.Client,
	rep repository.Repository,
	orderSVC service.CreateOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) *EventConsumer {
	return &EventConsumer{
		cfg:        cfg,
		sqsClient:  sqsClient,
		queueUrl:   envCfg.SqsUrlOrderEvent,
		rep:        rep,
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,
	}
}

func (e *EventConsumer) Run() {
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
	go consumer.PollMessages(
		ctxPolling,
		messagesChan,
		e.cfg.Event.MaxMessages,
		e.cfg.Event.PollingWaitTimeSecond,
		e.cfg.Event.VisibilityTimeoutSecond,
	)

	// ワーカーを並列に起動
	wgWorker := new(sync.WaitGroup)
	for i := 0; i < e.cfg.Event.MaxWorkers; i++ {
		wgWorker.Add(1)
		worker := sqs_consumer.NewWorker(i, wgWorker, e.workerHandler)
		go worker.Start(ctxWorker, messagesChan)
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

func (e *EventConsumer) workerHandler(ctx context.Context, msg *types.Message) error {
	mes, err := event_helper.ParseMessageFromSQS(msg)
	if err != nil {
		return err
	}

	err = e.processEvent(ctx, mes)
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

func (e *EventConsumer) processEvent(ctx context.Context, mes *message.Message) error {
	if !event_handler.IsCreateOrderSagaEvent(mes.Subject.Type) {
		return fmt.Errorf("invalid event type: %s", mes.Subject.Type)
	}

	sagaFactory := func(ctx context.Context, rep repository.Repository, id model.OrderID) (*create_order_saga.CreateOrderSaga, error) {
		state, err := e.rep.CreateOrderSagaStateFindOne(ctx, id)
		if err != nil {
			return nil, err
		}
		saga := create_order_saga.NewCreateOrderSaga(
			state,
			e.rep,
			e.orderSVC,
			e.kitchenAPI,
			e.billingAPI,
		)

		return saga, nil
	}

	err := event_handler.EventMap[mes.Subject.Type](e.rep).Handler(ctx, sagaFactory, mes)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventConsumer) deleteMessage(ctx context.Context, msg *types.Message) error {
	_, err := e.sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(e.queueUrl),
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
