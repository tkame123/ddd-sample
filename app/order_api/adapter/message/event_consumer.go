package message

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message/sqs_consumer"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
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
	rep repository.Repository
}

func NewEventConsumer() *EventConsumer {
	return &EventConsumer{}
}

func (e *EventConsumer) Exec() {
	queueUrl := "https://localhost.localstack.cloud:4566/000000000000/ddd-sample-order-event-queque" //TODO: Constructorへ

	// コンテキストとキャンセル関数を作成
	ctxPolling, ctxPollingCancel := context.WithCancel(context.Background())
	ctxWorker, ctxWorkerCancel := context.WithCancel(context.Background())
	defer ctxPollingCancel()
	defer ctxWorkerCancel()

	// SQS コンシューマを作成
	wgPolling := new(sync.WaitGroup)
	wgPolling.Add(1)
	consumer := sqs_consumer.NewSQSConsumer(queueUrl, wgPolling)

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
	type Body struct {
		Message string `json:"message"`
	}
	var body Body
	if err := json.Unmarshal([]byte(*msg.Body), &body); err != nil {
		return err
	}
	var message event.Dto
	if err := json.Unmarshal([]byte(body.Message), &message); err != nil {
		return err
	}

	// TODO: Event毎にHandlerをCallする
	log.Println("type:", message.Type)
	log.Println("body:", string(message.Origin))

	return nil
}