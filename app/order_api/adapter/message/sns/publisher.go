package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/lib/event"
)

type topicArn = string

type Publisher struct {
	Client   *sns.Client
	topicMap map[event.Name]topicArn
}

func NewPublisher(envCfg *provider.EnvConfig, client *sns.Client) *Publisher {
	topicMap := map[event.Name]topicArn{
		event.EventName_OrderCreated:    envCfg.ArnTopicEventOrderOrderCreated,
		event.EventName_OrderApproved:   envCfg.ArnTopicEventOrderOrderApproved,
		event.EventName_OrderRejected:   envCfg.ArnTopicEventOrderOrderRejected,
		event.CommandName_TicketCreate:  envCfg.ArnTopicCommandKitchenTicketCreate,
		event.CommandName_TicketApprove: envCfg.ArnTopicCommandKitchenTicketApprove,
		event.CommandName_TicketReject:  envCfg.ArnTopicCommandKitchenTicketReject,
		event.CommandName_CardAuthorize: envCfg.ArnTopicCommandBillingCardAuthorize,
	}

	return &Publisher{
		Client:   client,
		topicMap: topicMap,
	}
}

func (a Publisher) PublishMessage(ctx context.Context, eventName event.Name, message string) error {
	arn, ok := a.topicMap[eventName]
	if !ok {
		return fmt.Errorf("topic not found: %s", eventName)
	}

	input := sns.PublishInput{
		TopicArn: aws.String(arn),
		Message:  aws.String(message),
	}
	_, err := a.Client.Publish(ctx, &input)
	if err != nil {
		return err
	}
	return nil
}