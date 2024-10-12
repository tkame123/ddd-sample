package message

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/service/create_ticket/event_handler"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/proto/message"
)

type m struct {
	id  uuid.UUID
	raw *message.Message
}

func (m *m) ID() uuid.UUID {
	return m.id
}

func (m *m) Raw() *message.Message {
	return m.raw
}

func NewMessage(id uuid.UUID, mes *message.Message) domain_event.Message {
	return &m{id: id, raw: mes}
}

type factory struct {
	mes *message.Message
}

func NewCreateTicketServiceEventFactory(tp message.Type, mes *message.Message) (domain_event.MessageFactory, error) {
	if !event_handler.IsCreateTicketEvent(tp) {
		return nil, fmt.Errorf("invalid event type: %v", tp)
	}

	return &factory{mes: mes}, nil
}

func (f *factory) Event() (domain_event.Message, error) {
	switch f.mes.Subject.Type {
	case message.Type_TYPE_COMMAND_TICKET_CREATE:
		var c message.CommandTicketCreate
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_COMMAND_TICKET_APPROVE:
		var c message.CommandTicketApprove
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_COMMAND_TICKET_REJECT:
		var c message.CommandTicketReject
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	}

	return nil, fmt.Errorf("invalid event type: %v", f.mes.Subject.Type)
}
