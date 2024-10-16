package provider

const (
	maxMessages             = 5
	maxWorkers              = 3
	messageChan             = 10
	pollingWaitTimeSecond   = 20
	VisibilityTimeoutSecond = 30
)

type ConsumerConfig struct {
	Event   consumerConfig
	Command consumerConfig
	Reply   consumerConfig
}

type consumerConfig struct {
	QueueUrl                string
	MaxMessages             int
	MaxWorkers              int
	MessageChan             int
	PollingWaitTimeSecond   int
	VisibilityTimeoutSecond int
}

func NewConsumerConfig(envCfg *EnvConfig) *ConsumerConfig {
	return &ConsumerConfig{
		Event: consumerConfig{
			QueueUrl:                envCfg.SqsUrlBillingEvent,
			MaxMessages:             maxMessages,
			MaxWorkers:              maxWorkers,
			MessageChan:             messageChan,
			PollingWaitTimeSecond:   pollingWaitTimeSecond,
			VisibilityTimeoutSecond: VisibilityTimeoutSecond,
		},
		Command: consumerConfig{
			QueueUrl:                envCfg.SqsUrlBillingCommand,
			MaxMessages:             maxMessages,
			MaxWorkers:              maxWorkers,
			MessageChan:             messageChan,
			PollingWaitTimeSecond:   pollingWaitTimeSecond,
			VisibilityTimeoutSecond: VisibilityTimeoutSecond,
		},
		Reply: consumerConfig{
			QueueUrl:                envCfg.SqsUrlBillingReply,
			MaxMessages:             maxMessages,
			MaxWorkers:              maxWorkers,
			MessageChan:             messageChan,
			PollingWaitTimeSecond:   pollingWaitTimeSecond,
			VisibilityTimeoutSecond: VisibilityTimeoutSecond,
		},
	}
}
