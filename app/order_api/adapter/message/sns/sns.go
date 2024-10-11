package sns

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

type Actions struct {
	Client *sns.Client
}

func NewActions() *Actions {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	snsClient := sns.NewFromConfig(cfg)

	return &Actions{
		Client: snsClient,
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
