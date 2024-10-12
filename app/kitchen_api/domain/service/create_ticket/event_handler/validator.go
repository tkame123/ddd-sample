package event_handler

import (
	"github.com/tkame123/ddd-sample/proto/message"
)

func IsCreateTicketEvent(tp message.Type) bool {
	switch tp {
	case message.Type_TYPE_COMMAND_TICKET_CREATE, message.Type_TYPE_COMMAND_TICKET_APPROVE, message.Type_TYPE_COMMAND_TICKET_REJECT:
		return true
	}
	return false
}
