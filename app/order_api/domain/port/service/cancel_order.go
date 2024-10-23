package service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type CancelOrder interface {
	CancelOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
	CancelConfirmOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
	CancelRejectOrder(ctx context.Context, orderID model.OrderID) (model.OrderID, error)
}
