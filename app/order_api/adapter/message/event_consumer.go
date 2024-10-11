package message

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message/sqs_consumer"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga/event_handler"
	"github.com/tkame123/ddd-sample/lib/event"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	maxMessages = 5
	maxWorkers  = 3
	messageChan = 10
)

type EventConsumer struct {
	sqsClient  *sqs.Client
	queueUrl   string
	rep        repository.Repository
	orderSVC   service.CreateOrder
	kitchenAPI external_service.KitchenAPI
	billingAPI external_service.BillingAPI
}

func NewEventConsumer(
	rep repository.Repository,
	orderSVC service.CreateOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) *EventConsumer {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	sqsClient := sqs.NewFromConfig(cfg)
	queueUrl := "https://localhost.localstack.cloud:4566/000000000000/ddd-sample-order-event-queque" //TODO: env経由へ

	return &EventConsumer{
		sqsClient:  sqsClient,
		queueUrl:   queueUrl,
		rep:        rep,
		orderSVC:   orderSVC,
		kitchenAPI: kitchenAPI,
		billingAPI: billingAPI,
	}
}

func (e *EventConsumer) Exec() {

	// コンテキストとキャンセル関数を作成
	ctxPolling, ctxPollingCancel := context.WithCancel(context.Background())
	ctxWorker, ctxWorkerCancel := context.WithCancel(context.Background())

	// SQS コンシューマを作成
	wgPolling := new(sync.WaitGroup)
	wgPolling.Add(1)
	consumer := sqs_consumer.NewSQSConsumer(e.sqsClient, e.queueUrl, wgPolling)

	// メッセージをやり取りするチャネルを作成
	messagesChan := make(chan *types.Message, messageChan) // バッファ付きのチャネル

	// Poller を起動
	go consumer.PollMessages(ctxPolling, maxMessages, messagesChan)

	// ワーカーを並列に起動
	wgWorker := new(sync.WaitGroup)
	for i := 0; i < maxWorkers; i++ {
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
	ev, err := parseEvent(msg)
	if err != nil {
		return err
	}

	log.Println("event: ", ev)
	log.Println("eventName: ", ev.Name())
	if ev == nil {
		return errors.New("event is nil")

	}

	err = e.processEvent(ctx, ev)
	if err != nil {
		return err
	}

	err = e.deleteMessage(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventConsumer) processEvent(ctx context.Context, ev event.Event) error {
	// CreateOrderSaga以外に活用する段階ではAbstractFactoryを使う形なのかな？
	state, err := e.rep.CreateOrderSagaStateFindOne(ctx, ev.ID())
	if err != nil {
		return err
	}
	saga := create_order_saga.NewCreateOrderSaga(
		state,
		e.rep,
		e.orderSVC,
		e.kitchenAPI,
		e.billingAPI,
	)
	sagaHandler, err := NewCreateOrderSagaContext(ev, saga)
	if err != nil {
		return err
	}
	err = sagaHandler.Handler(ctx)
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

func parseEvent(msg *types.Message) (event.Event, error) {
	type Body struct {
		Message string `json:"message"`
	}
	var body Body
	if err := json.Unmarshal([]byte(*msg.Body), &body); err != nil {
		return nil, err
	}
	var message event.RawEvent
	if err := json.Unmarshal([]byte(body.Message), &message); err != nil {
		return nil, err
	}

	// CreateOrderSaga以外に活用する段階ではAbstractFactoryを使う形なのかな？
	eventFc, err := event_handler.NewCreateOrderSagaEventFactory(message.Type, message)
	if err != nil {
		return nil, err
	}
	domainEvent, err := eventFc.Event()
	if err != nil {
		return nil, err
	}

	return domainEvent, nil
}
