//go:build wireinject
// +build wireinject

package di

import (
	_ "github.com/lib/pq"

	"github.com/google/wire"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database"
	connect "github.com/tkame123/ddd-sample/app/order_api/adapter/gateway/api"
	provider "github.com/tkame123/ddd-sample/app/order_api/di/provider"
)

var providerOrderAPIServerSet = wire.NewSet(
	connect.NewServer,
	database.NewRepository,

	provider.NewENV,
	provider.NewOrderApiDB,
)

func InitializeOrderAPIServer() (connect.Server, func(), error) {
	wire.Build(providerOrderAPIServerSet)
	return connect.Server{}, nil, nil
}
