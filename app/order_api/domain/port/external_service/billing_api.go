package external_service

import (
	"context"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
)

type BillingAPI interface {
	AuthorizeCard(ctx context.Context, orderID model.OrderID)
}