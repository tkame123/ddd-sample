//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database"
	connect "github.com/tkame123/ddd-sample/app/order_api/adapter/gateway/api"
	"github.com/tkame123/ddd-sample/di/provider"
)

var providerOrderAPIServerSet = wire.NewSet(
	connect.NewServer,
	database.NewRepository,

	provider.NewConfig,
	provider.NewOrderApiDB,
)

func InitializeOrderAPIServer() (connect.Server, func(), error) {
	wire.Build(providerOrderAPIServerSet)
	return connect.Server{}, nil, nil
}
