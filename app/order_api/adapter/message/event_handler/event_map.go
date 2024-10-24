package event_handler

import (
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

var EventMap = map[message.Type]func(
	rep repository.Repository,
	orderSVC service.OrderService,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) domain_event.EventHandler{
	message.Type_TYPE_EVENT_ORDER_CREATED:                NewCallOrderCreateSagaWhenOrderCreatedHandler,
	message.Type_TYPE_EVENT_ORDER_APPROVED:               NewCallOrderCreateSagaWhenOrderApprovedHandler,
	message.Type_TYPE_EVENT_ORDER_REJECTED:               NewCallOrderCreateSagaWhenOrderRejectedHandler,
	message.Type_TYPE_EVENT_TICKET_CREATED:               NewCallOrderCreateSagaWhenTicketCreatedHandler,
	message.Type_TYPE_EVENT_TICKET_CREATION_FAILED:       NewCallOrderCreateSagaWhenTicketCreationFailedHandler,
	message.Type_TYPE_EVENT_TICKET_APPROVED:              NewCallOrderCreateSagaWhenTicketApprovedHandler,
	message.Type_TYPE_EVENT_TICKET_REJECTED:              NewCallOrderCreateSagaWhenTicketRejectedHandler,
	message.Type_TYPE_EVENT_CARD_AUTHORIZED:              NewCallOrderCreateSagaWhenCardAuthorizedHandler,
	message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED:    NewCallOrderCreateSagaWhenCardAuthorizationFailedHandler,
	message.Type_TYPE_EVENT_ORDER_CANCELED:               NewCallOrderCancelSagaWhenOrderCanceledHandler,
	message.Type_TYPE_EVENT_TICKET_CANCELED:              NewCallOrderCancelSagaWhenTicketCanceledHandler,
	message.Type_TYPE_EVENT_CARD_CANCELED:                NewCallOrderCancelSagaWhenCardCanceledHandler,
	message.Type_TYPE_EVENT_ORDER_CANCELLATION_CONFIRMED: NewCallOrderCancelSagaWhenOrderCancellationConfirmedHandler,
	message.Type_TYPE_EVENT_ORDER_CANCELLATION_REJECTED:  NewCallOrderCancelSagaWhenOrderCancellationRejectedHandler,
}
