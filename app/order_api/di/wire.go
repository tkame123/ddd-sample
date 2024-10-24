// go:build wireinject
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database"
	connect "github.com/tkame123/ddd-sample/app/order_api/adapter/gateway/api"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/idempotency"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/message"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/proxy"
	provider "github.com/tkame123/ddd-sample/app/order_api/di/provider"
	"github.com/tkame123/ddd-sample/app/order_api/usecase"
)

var providerServerSet = wire.NewSet(
	connect.NewServer,
	database.NewRepository,
	message.NewEventPublisher,
	idempotency.NewRepository,

	provider.NewENV,
	provider.NewAuthConfig,
	provider.NewCasbinEnforcer,
	provider.NewAWSConfig,
	provider.NewPublisherConfig,
	provider.NewOrderApiDB,
	provider.NewSNSClient,
	provider.NewDynamoClient,
)

var providerEventConsumerSet = wire.NewSet(
	message.NewEventConsumer,
	providerConsumerSet,
)

var providerCommandConsumerSet = wire.NewSet(
	message.NewCommandConsumer,
	providerConsumerSet,
)

var providerReplyConsumerSet = wire.NewSet(
	message.NewReplyConsumer,
	providerConsumerSet,
)

var providerConsumerSet = wire.NewSet(
	message.NewEventPublisher,
	database.NewRepository,
	usecase.NewOrderService,
	proxy.NewBillingAPI,
	proxy.NewKitchenAPI,

	provider.NewENV,
	provider.NewAWSConfig,
	provider.NewConsumerConfig,
	provider.NewPublisherConfig,
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

func InitializeCommandConsumer() (*message.CommandConsumer, func(), error) {
	wire.Build(providerCommandConsumerSet)
	return nil, nil, nil
}

func InitializeReplyConsumer() (*message.ReplyConsumer, func(), error) {
	wire.Build(providerReplyConsumerSet)
	return nil, nil, nil
}
