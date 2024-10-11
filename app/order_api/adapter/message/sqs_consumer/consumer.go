package sqs_consumer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SQSConsumer struct {
	sqsClient *sqs.Client
	queueUrl  string
	wg        *sync.WaitGroup
}

func NewSQSConsumer(sqsClient *sqs.Client, queueUrl string, wg *sync.WaitGroup) *SQSConsumer {
	return &SQSConsumer{
		sqsClient: sqsClient,
		queueUrl:  queueUrl,
		wg:        wg,
	}
}

func (c *SQSConsumer) PollMessages(ctx context.Context, maxMessages int, messagesChan chan<- *types.Message) {
	defer c.wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Println("Polling stopped")
			return
		default:
			output, err := c.sqsClient.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(c.queueUrl),
				MaxNumberOfMessages: int32(maxMessages),
				WaitTimeSeconds:     20, // Long Polling
			})
			if err != nil {
				log.Println("Error receiving messages:", err)
				time.Sleep(5 * time.Second) // エラーバックオフ
				continue
			}

			if len(output.Messages) == 0 {
				log.Println("No messages received")
				continue
			}

			// 取得したメッセージをメッセージチャネルに送信
			log.Println("Received messages", len(output.Messages))
			for _, msg := range output.Messages {
				messagesChan <- &msg
			}
		}
	}
}
