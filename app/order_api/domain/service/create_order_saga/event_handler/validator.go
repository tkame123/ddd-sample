package event_handler

import (
	"github.com/tkame123/ddd-sample/proto/message"
)

func IsCreateOrderSagaEvent(tp message.Type) bool {
	switch tp {
	case message.Type_TYPE_EVENT_ORDER_CREATED, message.Type_TYPE_EVENT_ORDER_APPROVED, message.Type_TYPE_EVENT_ORDER_REJECTED:
		return true
	case message.Type_TYPE_EVENT_TICKET_CREATED, message.Type_TYPE_EVENT_TICKET_CREATION_FAILED, message.Type_TYPE_EVENT_TICKET_APPROVED, message.Type_TYPE_EVENT_TICKET_REJECTED:
		return true
	case message.Type_TYPE_EVENT_CARD_AUTHORIZED, message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED:
		return true
	}
	return false
}
