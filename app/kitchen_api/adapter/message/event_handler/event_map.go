package event_handler

import (
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

var EventMap = map[message.Type]func(svc service.Ticket) domain_event.EventHandler{
	message.Type_TYPE_COMMAND_TICKET_CREATE:  NewTicketCreateWhenTicketCreateHandler,
	message.Type_TYPE_COMMAND_TICKET_APPROVE: NewTicketApproveWhenTicketApproveHandler,
	message.Type_TYPE_COMMAND_TICKET_REJECT:  NewTicketRejectWhenTicketRejectHandler,
	message.Type_TYPE_COMMAND_TICKET_CANCEL:  NewTicketCancelWhenTicketCancelHandler,
}
