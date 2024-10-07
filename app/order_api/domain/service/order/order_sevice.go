package order

import (
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
)

type Service struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewService(rep repository.Repository, pub domain_event.Publisher) *Service {
	return &Service{
		rep: rep,
		pub: pub,
	}
}
