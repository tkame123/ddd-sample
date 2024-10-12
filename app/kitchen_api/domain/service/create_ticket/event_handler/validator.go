package event_handler

import "github.com/tkame123/ddd-sample/lib/event_helper"

func IsCreateTicketEvent(name event_helper.Name) bool {
	switch name {
	case event_helper.CommandName_TicketCreate, event_helper.CommandName_TicketApprove, event_helper.CommandName_TicketReject:
		return true
	}
	return false
}
