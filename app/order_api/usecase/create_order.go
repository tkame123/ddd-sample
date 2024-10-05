package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type CreateOrder struct {
	rep repository.Repository
}

func NewCreateOrder(rep repository.Repository) *CreateOrder {
	return &CreateOrder{
		rep: rep,
	}
}

type CreateOrderInput struct {
	Items []*model.OrderItemRequest
}

type CreateOrderOutput struct {
	OrderID model.OrderID
}

func (c *CreateOrder) Execute(ctx context.Context, input CreateOrderInput) (*CreateOrderOutput, error) {
	order, events, err := model.NewOrder(input.Items)
	if err != nil {
		return nil, err
	}

	if err := c.rep.Order.Save(ctx, order); err != nil {
		return nil, err
	}

	if events != nil {
		// TODO: publish event
	}

	return &CreateOrderOutput{OrderID: order.OrderID}, nil
}
