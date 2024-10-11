package sns

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Actions struct {
	Client *sns.Client
}

func NewActions(client *sns.Client) *Actions {
	return &Actions{
		Client: client,
	}
}

func (a Actions) PublishMessage(ctx context.Context, arn string, message string) error {
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
