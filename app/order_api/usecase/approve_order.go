package usecase

import (
	"context"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
)

type ApproveOrder struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewApproveOrder(rep repository.Repository, pub domain_event.Publisher) *CreateOrder {
	return &CreateOrder{
		rep: rep,
		pub: pub,
	}
}

type ApproveOrderInput struct {
}

type ApproveOrderOutput struct {
	OrderID model.OrderID
}

func (c *ApproveOrder) Execute(ctx context.Context, input ApproveOrderInput) (*ApproveOrderOutput, error) {
	panic("implement me")
}
