package event_handler

import "github.com/tkame123/ddd-sample/lib/event"

func IsCreateTicketEvent(name event.Name) bool {
	switch name {
	case event.CommandName_TicketCreate, event.CommandName_TicketApprove, event.CommandName_TicketReject:
		return true
	}
	return false
}
