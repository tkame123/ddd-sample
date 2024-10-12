package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/proto/message"
	"google.golang.org/protobuf/encoding/protojson"
)

type topicArn = string

type Publisher struct {
	Client   *sns.Client
	topicMap map[message.Type]topicArn
}

func NewPublisher(envCfg *provider.EnvConfig, client *sns.Client) *Publisher {
	// TODO: fix
	topicMap := map[message.Type]topicArn{
		message.Type_TYPE_EVENT_ORDER_CREATED:    envCfg.ArnTopicEventOrderOrderCreated,
		message.Type_TYPE_EVENT_ORDER_APPROVED:   envCfg.ArnTopicEventOrderOrderApproved,
		message.Type_TYPE_EVENT_ORDER_REJECTED:   envCfg.ArnTopicEventOrderOrderRejected,
		message.Type_TYPE_COMMAND_TICKET_CREATE:  envCfg.ArnTopicCommandKitchenTicketCreate,
		message.Type_TYPE_COMMAND_TICKET_APPROVE: envCfg.ArnTopicCommandKitchenTicketApprove,
		message.Type_TYPE_COMMAND_TICKET_REJECT:  envCfg.ArnTopicCommandKitchenTicketReject,
		message.Type_TYPE_COMMAND_CARD_AUTHORIZE: envCfg.ArnTopicCommandBillingCardAuthorize,
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
