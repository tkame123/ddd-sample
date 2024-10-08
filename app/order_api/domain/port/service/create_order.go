package service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type ApproveOrderInput struct {
	OrderID model.OrderID
}

type ApproveOrderOutput struct {
	OrderID model.OrderID
}

type CreateOrderInput struct {
	Items []*model.OrderItemRequest
}

type CreateOrderOutput struct {
	OrderID model.OrderID
}

type RejectOrderInput struct {
	OrderID model.OrderID
}

type RejectOrderOutput struct {
	OrderID model.OrderID
}

type CreateOrder interface {
	ApproveOrder(ctx context.Context, input ApproveOrderInput) (*ApproveOrderOutput, error)
	CreateOrder(ctx context.Context, input CreateOrderInput) (*CreateOrderOutput, error)
	RejectOrder(ctx context.Context, input RejectOrderInput) (*RejectOrderOutput, error)
}
