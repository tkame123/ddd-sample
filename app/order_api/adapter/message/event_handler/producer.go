package event_handler

import (
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/external_service"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/repository"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

type HandlerProducer struct {
	rep            repository.Repository
	orderCreateSVC service.CreateOrder
	orderCancelSVC service.CancelOrder
	kitchenAPI     external_service.KitchenAPI
	billingAPI     external_service.BillingAPI
}

func NewHandlerProducer(
	rep repository.Repository,
	orderCreateSVC service.CreateOrder,
	orderCancelSVC service.CancelOrder,
	kitchenAPI external_service.KitchenAPI,
	billingAPI external_service.BillingAPI,
) *HandlerProducer {
	return &HandlerProducer{
		rep:            rep,
		orderCreateSVC: orderCreateSVC,
		orderCancelSVC: orderCancelSVC,
		kitchenAPI:     kitchenAPI,
		billingAPI:     billingAPI,
	}
}

func (fp *HandlerProducer) GetHandler(m *message.Message) (domain_event.EventHandler, error) {
	switch m.Subject.Type {
	case
		message.Type_TYPE_EVENT_ORDER_CREATED,
		message.Type_TYPE_EVENT_ORDER_APPROVED,
		message.Type_TYPE_EVENT_ORDER_REJECTED,
		message.Type_TYPE_EVENT_TICKET_CREATED,
		message.Type_TYPE_EVENT_TICKET_CREATION_FAILED,
		message.Type_TYPE_EVENT_TICKET_APPROVED,
		message.Type_TYPE_EVENT_TICKET_REJECTED,
		message.Type_TYPE_EVENT_CARD_AUTHORIZED,
		message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED:
		return NewCreateOrderSagaEventHandler(fp.rep, fp.orderCreateSVC, fp.kitchenAPI, fp.billingAPI), nil
	}

	return nil, errors.New("handler not found for event type " + m.Subject.Type.String())
}
