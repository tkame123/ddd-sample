package event_handler

import "github.com/tkame123/ddd-sample/lib/event"

func IsCreateOrderSagaEvent(name event.Name) bool {
	switch name {
	case event.EventName_OrderCreated, event.EventName_OrderApproved, event.EventName_OrderRejected:
		return true
	case event.EventName_TicketCreated, event.EventName_TicketCreationFailed, event.EventName_TicketApproved, event.EventName_TicketRejected:
		return true
	case event.EventName_CardAuthorized, event.EventName_CardAuthorizeFailed:
		return true
	}
	return false
}
