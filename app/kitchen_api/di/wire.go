//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/tkame123/ddd-sample/app/kitchen_api/adapter/database"
	"github.com/tkame123/ddd-sample/app/kitchen_api/adapter/message"
	"github.com/tkame123/ddd-sample/app/kitchen_api/adapter/message/sns"
	"github.com/tkame123/ddd-sample/app/kitchen_api/di/provider"
	"github.com/tkame123/ddd-sample/app/kitchen_api/usecase/create_ticket"
)

var providerCommandConsumerSet = wire.NewSet(
	message.NewCommandConsumer,
	create_ticket.NewService,
	message.NewEventPublisher,
	database.NewRepository,
	sns.NewPublisher,

	provider.NewENV,
	provider.NewAWSConfig,
	provider.NewConsumerConfig,
	provider.NewSQSClient,
	provider.NewSNSClient,
)

func InitializeCommandConsumer() (*message.CommandConsumer, func(), error) {
	wire.Build(providerCommandConsumerSet)
	return nil, nil, nil
}
