//go:build wireinject
// +build wireinject

package di

import (
	_ "github.com/lib/pq"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/gateway/consumer"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/gateway/publisher"
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
	publisher.NewEventPublisher,
	sns.NewActions,

	provider.NewENV,
	provider.NewAWSConfig,
	provider.NewOrderApiDB,
	provider.NewSNSClient,
)

var providerEventConsumerSet = wire.NewSet(
	consumer.NewEventConsumer,
	publisher.NewEventPublisher,
	database.NewRepository,
	create_order.NewService,
	proxy.NewBillingAPI,
	proxy.NewKitchenAPI,
	sns.NewActions,

	provider.NewENV,
	provider.NewAWSConfig,
	provider.NewOrderApiDB,
	provider.NewSQSClient,
	provider.NewSNSClient,
)

func InitializeAPIServer() (connect.Server, func(), error) {
	wire.Build(providerServerSet)
	return connect.Server{}, nil, nil
}

func InitializeEventConsumer() (*consumer.EventConsumer, func(), error) {
	wire.Build(providerEventConsumerSet)
	return nil, nil, nil
}
