package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type CreateOrder struct{}

func NewCreateOrder() *CreateOrder {
	return &CreateOrder{}
}

type CreateOrderInput struct {
	Items []*model.OrderItemRequest
}

type CreateOrderOutput struct {
	OrderID model.OrderID
}

func (c *CreateOrder) Execute(ctx context.Context, input CreateOrderInput) (*CreateOrderOutput, error) {
	order, err := model.NewOrder(input.Items)
	if err != nil {
		return nil, err
	}

	// TODO: save repository

	return &CreateOrderOutput{OrderID: order.OrderID}, nil
}
