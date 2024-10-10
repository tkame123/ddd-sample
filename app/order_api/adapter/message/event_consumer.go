package message

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message/sqs_consumer"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	maxMessages = 5
	maxWorkers  = 3
)

type EventConsumer struct {
}

func NewEventConsumer() *EventConsumer {
	return &EventConsumer{}
}

func (e *EventConsumer) Exec() {
	queueUrl := "https://localhost.localstack.cloud:4566/000000000000/ddd-sample-order-event-queque"

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
	messagesChan := make(chan *types.Message, 10) // バッファ付きのチャネル

	// Poller を起動
	go consumer.PollMessages(ctxPolling, maxMessages, messagesChan)

	// 3つのワーカーを並列に起動
	wgWorker := new(sync.WaitGroup)
	for i := 0; i < maxWorkers; i++ {
		wgWorker.Add(1)
		worker := sqs_consumer.NewWorker(i, wgWorker)
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
