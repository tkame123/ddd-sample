package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type UpdateOrderItems struct {
	rep repository.Repository
}

func NewUpdateOrderItems(rep repository.Repository) *UpdateOrderItems {
	return &UpdateOrderItems{
		rep: rep,
	}
}

type UpdateOrderItemsInput struct {
	OrderID model.OrderID
	Items   []*model.OrderItemRequest
}

type UpdateOrderItemsOutput struct {
	OrderID model.OrderID
}

func (u *UpdateOrderItems) Execute(ctx context.Context, input UpdateOrderItemsInput) (*UpdateOrderItemsOutput, error) {
	order, err := u.rep.Order.FindOne(ctx, &input.OrderID)
	if err != nil {
		return nil, err
	}

	if err := order.UpdateOrderItems(input.Items); err != nil {
		return nil, err
	}

	if err := u.rep.Order.Save(ctx, order); err != nil {
		return nil, err
	}

	return &UpdateOrderItemsOutput{OrderID: order.OrderID}, nil
}