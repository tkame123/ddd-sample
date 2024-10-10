package sns

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/caarlos0/env/v11"
	"log"
)

type Topic = int
type topicArn = string

const (
	Topic_Event_Order_OrderCreated Topic = iota
	Topic_Event_Order_OrderrApproved
	Topic_Event_Order_OrderRejected
	Topic_Command_Order_OrderApprove
	Topic_Command_Order_OrderReject
	Topic_Command_Kitchen_TicketCreate
	Topic_Command_Kitchen_TicketAprove
	Topic_Command_Kitchen_TicketReject
	Topic_Command_Billing_CardAuthorize
)

type arnConfig struct {
	ArnTopicEventOrderOrderCreated      string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_CREATED"`
	ArnTopicEventOrderOrderApproved     string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_APPROVED"`
	ArnTopicEventOrderOrderRejected     string `env:"TOPIC_ARN_EVENT_ORDER_ORDER_REJECTED"`
	ArnTopicCommandOrderOrderApprove    string `env:"TOPIC_ARN_COMMAND_ORDER_ORDER_APPROVE"`
	ArnTopicCommandOrderOrderReject     string `env:"TOPIC_ARN_COMMAND_ORDER_ORDER_REJECT"`
	ArnTopicCommandKitchenTicketCreate  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_CREATE"`
	ArnTopicCommandKitchenTicketApprove string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_APPROVE"`
	ArnTopicCommandKitchenTicketReject  string `env:"TOPIC_ARN_COMMAND_KITCHEN_TICKET_REJECT"`
	ArnTopicCommandBillingCardAuthorize string `env:"TOPIC_ARN_COMMAND_BILLING_CARD_AUTHORIZE"`
}

type Actions struct {
	Client   *sns.Client
	topicMap map[Topic]topicArn
}

func NewActions() *Actions {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	var arnCfg arnConfig
	if err := env.Parse(&arnCfg); err != nil {
		log.Fatalf("unable to parase env, %v", err)
	}

	snsClient := sns.NewFromConfig(cfg)
	topicMap := map[Topic]topicArn{
		Topic_Event_Order_OrderCreated:      arnCfg.ArnTopicEventOrderOrderCreated,
		Topic_Event_Order_OrderrApproved:    arnCfg.ArnTopicEventOrderOrderApproved,
		Topic_Event_Order_OrderRejected:     arnCfg.ArnTopicEventOrderOrderRejected,
		Topic_Command_Order_OrderApprove:    arnCfg.ArnTopicCommandOrderOrderApprove,
		Topic_Command_Order_OrderReject:     arnCfg.ArnTopicCommandOrderOrderReject,
		Topic_Command_Kitchen_TicketCreate:  arnCfg.ArnTopicCommandKitchenTicketCreate,
		Topic_Command_Kitchen_TicketAprove:  arnCfg.ArnTopicCommandKitchenTicketApprove,
		Topic_Command_Kitchen_TicketReject:  arnCfg.ArnTopicCommandKitchenTicketReject,
		Topic_Command_Billing_CardAuthorize: arnCfg.ArnTopicCommandBillingCardAuthorize,
	}

	return &Actions{
		Client:   snsClient,
		topicMap: topicMap,
	}
}

func (a Actions) PublishMessage(ctx context.Context, topic Topic, message string) error {
	arn, ok := a.topicMap[topic]
	if !ok {
		return errors.New("topic not found")
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
