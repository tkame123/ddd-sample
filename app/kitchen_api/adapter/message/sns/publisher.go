package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/tkame123/ddd-sample/app/kitchen_api/di/provider"
	"github.com/tkame123/ddd-sample/proto/message"
	"google.golang.org/protobuf/encoding/protojson"
)

type topicArn = string

type Publisher struct {
	Client   *sns.Client
	topicMap map[message.Type]topicArn
}

func NewPublisher(envCfg *provider.EnvConfig, client *sns.Client) *Publisher {
	topicMap := map[message.Type]topicArn{
		message.Type_TYPE_EVENT_TICKET_CREATED:         envCfg.ArnTopicEventKitchenTicketCreated,
		message.Type_TYPE_EVENT_TICKET_CREATION_FAILED: envCfg.ArnTopicEventKitchenTicketCreationFailed,
		message.Type_TYPE_EVENT_TICKET_APPROVED:        envCfg.ArnTopicEventKitchenTicketApproved,
		message.Type_TYPE_EVENT_TICKET_REJECTED:        envCfg.ArnTopicEventKitchenTicketRejected,
	}

	return &Publisher{
		Client:   client,
		topicMap: topicMap,
	}
}

func (a Publisher) PublishMessage(ctx context.Context, event *message.Message) error {
	arn, ok := a.topicMap[event.Subject.Type]
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
	_, err = a.Client.Publish(ctx, &input)
	if err != nil {
		return err
	}
	return nil
}
