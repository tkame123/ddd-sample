package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/proto/message"
	"google.golang.org/protobuf/encoding/protojson"
)

type topicArn = string

type Publisher struct {
	Client *sns.Client
	// Deprecated: use topicMap2
	topicMap  map[event_helper.Name]topicArn
	topicMap2 map[message.Type]topicArn
}

func NewPublisher(envCfg *provider.EnvConfig, client *sns.Client) *Publisher {
	topicMap := map[event_helper.Name]topicArn{
		event_helper.EventName_OrderCreated:    envCfg.ArnTopicEventOrderOrderCreated,
		event_helper.EventName_OrderApproved:   envCfg.ArnTopicEventOrderOrderApproved,
		event_helper.EventName_OrderRejected:   envCfg.ArnTopicEventOrderOrderRejected,
		event_helper.CommandName_TicketCreate:  envCfg.ArnTopicCommandKitchenTicketCreate,
		event_helper.CommandName_TicketApprove: envCfg.ArnTopicCommandKitchenTicketApprove,
		event_helper.CommandName_TicketReject:  envCfg.ArnTopicCommandKitchenTicketReject,
		event_helper.CommandName_CardAuthorize: envCfg.ArnTopicCommandBillingCardAuthorize,
	}

	topicMap2 := map[message.Type]topicArn{
		message.Type_TYPE_EVENT_ORDER_CREATED: envCfg.ArnTopicEventOrderOrderCreated,
		//event.EventName_OrderApproved:   envCfg.ArnTopicEventOrderOrderApproved,
		//event.EventName_OrderRejected:   envCfg.ArnTopicEventOrderOrderRejected,
		//event.CommandName_TicketCreate:  envCfg.ArnTopicCommandKitchenTicketCreate,
		//event.CommandName_TicketApprove: envCfg.ArnTopicCommandKitchenTicketApprove,
		//event.CommandName_TicketReject:  envCfg.ArnTopicCommandKitchenTicketReject,
		//event.CommandName_CardAuthorize: envCfg.ArnTopicCommandBillingCardAuthorize,
	}

	return &Publisher{
		Client:    client,
		topicMap:  topicMap,
		topicMap2: topicMap2,
	}
}

func (a Publisher) PublishMessage(ctx context.Context, eventName event_helper.Name, message string) error {
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

func (a Publisher) PublishMessage2(ctx context.Context, event *message.Message) error {
	arn, ok := a.topicMap2[event.Subject.Type]
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
