package create_order

import (
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
)

type s struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewService(rep repository.Repository, pub domain_event.Publisher) service.CreateOrder {
	return &s{
		rep: rep,
		pub: pub,
	}
}
