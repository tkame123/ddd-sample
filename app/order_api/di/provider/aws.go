package provider

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewAWSConfig() (*aws.Config, error) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &cfg, nil
}

func NewSNSClient(cfg *aws.Config) (*sns.Client, error) {
	return sns.NewFromConfig(*cfg), nil
}

func NewSQSClient(cfg *aws.Config) (*sqs.Client, error) {
	return sqs.NewFromConfig(*cfg), nil
}
