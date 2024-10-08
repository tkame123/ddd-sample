package service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type CreateOrder interface {
	ApproveOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
	CreateOrder(ctx context.Context, items []*model.OrderItemRequest) (model.OrderID, error)
	RejectOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
}
