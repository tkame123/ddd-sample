//go:build wireinject
// +build wireinject

package di

import (
	_ "github.com/lib/pq"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message/sns"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/proxy"
	"github.com/tkame123/ddd-sample/app/order_api/usecase/create_order"

	"github.com/google/wire"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database"
	connect "github.com/tkame123/ddd-sample/app/order_api/adapter/gateway/api"
	provider "github.com/tkame123/ddd-sample/app/order_api/di/provider"
)

var providerServerSet = wire.NewSet(
	connect.NewServer,
	database.NewRepository,
	message.NewEventPublisher,
	sns.NewPublisher,

	provider.NewENV,
	provider.NewAWSConfig,
	provider.NewOrderApiDB,
	provider.NewSNSClient,
)

var providerEventConsumerSet = wire.NewSet(
	message.NewEventConsumer,
	message.NewEventPublisher,
	database.NewRepository,
	create_order.NewService,
	proxy.NewBillingAPI,
	proxy.NewKitchenAPI,
	sns.NewPublisher,

	provider.NewENV,
	provider.NewAWSConfig,
	provider.NewConsumerConfig,
	provider.NewOrderApiDB,
	provider.NewSQSClient,
	provider.NewSNSClient,
)

func InitializeAPIServer() (connect.Server, func(), error) {
	wire.Build(providerServerSet)
	return connect.Server{}, nil, nil
}

func InitializeEventConsumer() (*message.EventConsumer, func(), error) {
	wire.Build(providerEventConsumerSet)
	return nil, nil, nil
}
