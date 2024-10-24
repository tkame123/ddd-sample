package service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/model"
)

type Bill interface {
	AuthorizeCard(ctx context.Context, orderID model.OrderID, token any) error
	CancelCard(ctx context.Context, orderID model.OrderID) error
}
