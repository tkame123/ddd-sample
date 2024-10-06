package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
)

type CreateOrder struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewCreateOrder(rep repository.Repository, pub domain_event.Publisher) *CreateOrder {
	return &CreateOrder{
		rep: rep,
		pub: pub,
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

	c.pub.PublishMessages(ctx, events)

	return &CreateOrderOutput{OrderID: order.OrderID()}, nil
}
