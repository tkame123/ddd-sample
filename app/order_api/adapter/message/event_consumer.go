package message

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"github.com/tkame123/ddd-sample/lib/sqs_consumer"
	"github.com/tkame123/ddd-sample/proto/message"
	"google.golang.org/protobuf/encoding/protojson"
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
	envCfg *provider.EnvConfig,
	sqsClient *sqs.Client,
	rep repository.Repository,
	orderSVC service.CreateOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) *EventConsumer {
	return &EventConsumer{
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
	mes, err := parseMessage(msg)
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

func (e *EventConsumer) processEvent(ctx context.Context, mes domain_event.Message) error {
	// CreateOrderSaga以外に活用する段階ではAbstractFactoryを使う形なのかな？
	state, err := e.rep.CreateOrderSagaStateFindOne(ctx, mes.ID())
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
	sagaHandler, err := NewCreateOrderSagaContext(mes.Raw())
	if err != nil {
		return err
	}
	err = sagaHandler.Handler(ctx, saga)
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

func parseMessage(msg *types.Message) (domain_event.Message, error) {
	type Body struct {
		Message string `json:"message"`
	}
	var body Body
	if err := json.Unmarshal([]byte(*msg.Body), &body); err != nil {
		return nil, err
	}
	var m message.Message
	if err := protojson.Unmarshal([]byte(body.Message), &m); err != nil {
		return nil, err
	}

	// CreateOrderSaga以外に活用する段階ではAbstractFactoryを使う形なのかな？
	eventFc, err := NewCreateOrderSagaEventFactory(m.Subject.Type, &m)
	if err != nil {
		return nil, err
	}
	return eventFc.Event()
}
