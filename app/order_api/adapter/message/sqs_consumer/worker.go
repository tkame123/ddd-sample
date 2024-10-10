package sqs_consumer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"log"
	"sync"
)

type Worker struct {
	id      int
	wg      *sync.WaitGroup
	handler func(ctx context.Context, msg *types.Message) error
}

func NewWorker(id int, wg *sync.WaitGroup, handler func(ctx context.Context, msg *types.Message) error) *Worker {
	return &Worker{id: id, wg: wg, handler: handler}
}

func (w *Worker) Start(ctx context.Context, messagesChan <-chan *types.Message) {
	defer w.wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d stopped\n", w.id)
			return
		case msg := <-messagesChan:
			w.processMessage(context.Background(), msg)
		}
	}
}

func (w *Worker) processMessage(ctx context.Context, msg *types.Message) {
	log.Printf("Worker %d processing message: %s\n", w.id, *msg.Body)
	if w.handler != nil {
		if err := w.handler(ctx, msg); err != nil {
			log.Printf("Worker %d failed to process message: %s\n", w.id, err)
		}
	}
}
