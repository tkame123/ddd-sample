package sqs_consumer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Worker struct {
	id int
	wg *sync.WaitGroup
}

func NewWorker(id int, wg *sync.WaitGroup) *Worker {
	return &Worker{id: id, wg: wg}
}

func (w *Worker) Start(ctx context.Context, messagesChan <-chan *types.Message) {
	defer w.wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d stopped\n", w.id)
			return
		case msg := <-messagesChan:
			w.processMessage(msg)
		}
	}
}

func (w *Worker) processMessage(msg *types.Message) {
	log.Printf("Worker %d processing message: %s\n", w.id, *msg.Body)
	// メッセージの処理ロジックをここに実装
	time.Sleep(3 * time.Second) // 処理に時間がかかる場合をシミュレーション
}
