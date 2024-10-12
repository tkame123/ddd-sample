package sns

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Publisher struct {
	Client *sns.Client
}

func NewPublisher(client *sns.Client) *Publisher {
	return &Publisher{
		Client: client,
	}
}

func (a Publisher) PublishMessage(ctx context.Context, arn string, message string) error {
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
