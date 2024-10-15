package service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/model"
)

type CreateBill interface {
	AuthorizeCard(ctx context.Context, orderID model.OrderID, token any) error
}
