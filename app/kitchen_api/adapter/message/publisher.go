package message

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/tkame123/ddd-sample/app/kitchen_api/di/provider"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/proto/message"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

type topicArn = string

type EventPublisher struct {
	Client   *sns.Client
	topicMap map[message.Type]provider.TopicArn
}

func NewEventPublisher(cfg *provider.PublisherConfig, client *sns.Client) domain_event.Publisher {
	return &EventPublisher{
		Client:   client,
		topicMap: cfg.TopicMap,
	}
}

func (s *EventPublisher) PublishMessages(ctx context.Context, events []*message.Message) {
	for _, e := range events {
		if err := s.PublishMessage(ctx, e); err != nil {
			// TODO Transactional Outbox Patternの導入までは通知して手動対応ってなるのだろうか。。。
			log.Printf("failed to publish message %v", err)
		}
	}
}

func (s *EventPublisher) PublishMessage(ctx context.Context, event *message.Message) error {
	arn, ok := s.topicMap[event.Subject.Type]
	if !ok {
		return fmt.Errorf("topic not found: %s", event.Subject.Type)
	}

	mes, err := protojson.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	input := sns.PublishInput{
		TopicArn: aws.String(arn),
		Message:  aws.String(string(mes)),
	}
	_, err = s.Client.Publish(ctx, &input)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w arn: %s", err, arn)
	}
	return nil
}
