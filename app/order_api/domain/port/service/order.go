package service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type OrderService interface {
	ApproveOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
	CreateOrder(ctx context.Context, items []*model.OrderItemRequest) (model.OrderID, error)
	RejectOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
	CancelOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
	CancelConfirmOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
	CancelRejectOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
}
