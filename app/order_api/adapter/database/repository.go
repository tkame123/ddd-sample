package database

import "github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"

func NewRepository(order repository.Order, cos repository.CreateOrderSagaState) *repository.Repository {
	return &repository.Repository{Order: order, CreateOrderSagaState: cos}
}
