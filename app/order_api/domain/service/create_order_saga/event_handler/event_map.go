package event_handler

import (
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/proto/message"
)

var EventMap = map[message.Type]func(rep repository.Repository) domain_event.CreateOrderSagaEventHandler{
	message.Type_TYPE_EVENT_ORDER_CREATED:             NewNextStepSagaWhenOrderCreatedHandler,
	message.Type_TYPE_EVENT_ORDER_APPROVED:            NewNextStepSagaWhenOrderApprovedHandler,
	message.Type_TYPE_EVENT_ORDER_REJECTED:            NewNextStepSagaWhenOrderRejectedHandler,
	message.Type_TYPE_EVENT_TICKET_CREATED:            NewNextStepSagaWhenTicketCreatedHandler,
	message.Type_TYPE_EVENT_TICKET_CREATION_FAILED:    NewNextStepSagaWhenTicketCreationFailedHandler,
	message.Type_TYPE_EVENT_TICKET_APPROVED:           NewNextStepSagaWhenTicketApprovedHandler,
	message.Type_TYPE_EVENT_TICKET_REJECTED:           NewNextStepSagaWhenTicketRejectedHandler,
	message.Type_TYPE_EVENT_CARD_AUTHORIZED:           NewNextStepSagaWhenCardAuthorizedHandler,
	message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED: NewNextStepSagaWhenCardAuthorizeFailedHandler,
}
