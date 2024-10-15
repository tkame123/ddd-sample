package event_handler

import (
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/billilng_api/domain/port/service"
	"github.com/tkame123/ddd-sample/proto/message"
)

var EventMap = map[message.Type]func(svc service.CreateBill) domain_event.EventHandler{
	message.Type_TYPE_COMMAND_CARD_AUTHORIZE: NewCardAuthorizeWhenCardAuthorizeHandler,
}
