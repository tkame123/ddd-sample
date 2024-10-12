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
	MaxMessages             int
	MaxWorkers              int
	MessageChan             int
	PollingWaitTimeSecond   int
	VisibilityTimeoutSecond int
}

func NewConsumerConfig() *ConsumerConfig {
	return &ConsumerConfig{
		Event: consumerConfig{
			MaxMessages:             maxMessages,
			MaxWorkers:              maxWorkers,
			MessageChan:             messageChan,
			PollingWaitTimeSecond:   pollingWaitTimeSecond,
			VisibilityTimeoutSecond: VisibilityTimeoutSecond,
		},
		Command: consumerConfig{
			MaxMessages:             maxMessages,
			MaxWorkers:              maxWorkers,
			MessageChan:             messageChan,
			PollingWaitTimeSecond:   pollingWaitTimeSecond,
			VisibilityTimeoutSecond: VisibilityTimeoutSecond,
		},
		Reply: consumerConfig{
			MaxMessages:             maxMessages,
			MaxWorkers:              maxWorkers,
			MessageChan:             messageChan,
			PollingWaitTimeSecond:   pollingWaitTimeSecond,
			VisibilityTimeoutSecond: VisibilityTimeoutSecond,
		},
	}
}
