//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/tkame123/ddd-sample/app/billilng_api/adapter/database"
	"github.com/tkame123/ddd-sample/app/billilng_api/adapter/message"
	"github.com/tkame123/ddd-sample/app/billilng_api/di/provider"
	"github.com/tkame123/ddd-sample/app/billilng_api/usecase"
)

var providerCommandConsumerSet = wire.NewSet(
	message.NewCommandConsumer,
	message.NewEventPublisher,
	usecase.NewService,
	database.NewRepository,

	provider.NewENV,
	provider.NewAWSConfig,
	provider.NewConsumerConfig,
	provider.NewPublisherConfig,
	provider.NewSQSClient,
	provider.NewSNSClient,
)

func InitializeCommandConsumer() (*message.CommandConsumer, func(), error) {
	wire.Build(providerCommandConsumerSet)
	return nil, nil, nil
}
