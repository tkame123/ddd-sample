package create_ticket

import (
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
)

type s struct {
	rep repository.Repository
	pub domain_event.Publisher
}

func NewService(rep repository.Repository, pub domain_event.Publisher) service.CreateTicket {
	return &s{
		rep: rep,
		pub: pub,
	}
}