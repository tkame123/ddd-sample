package message

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga/event_handler"
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

func NewCreateOrderSagaEventFactory(tp message.Type, mes *message.Message) (domain_event.MessageFactory, error) {
	if !event_handler.IsCreateOrderSagaEvent(tp) {
		return nil, fmt.Errorf("invalid event type: %v", tp)
	}

	return &factory{mes: mes}, nil
}

func (f *factory) Event() (domain_event.Message, error) {
	// TODO: いい感じにまとめられないのだろうか。。
	switch f.mes.Subject.Type {
	case message.Type_TYPE_EVENT_ORDER_CREATED:
		var c message.EventOrderCreated
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_EVENT_ORDER_APPROVED:
		var c message.EventOrderApproved
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_EVENT_ORDER_REJECTED:
		var c message.EventOrderRejected
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_EVENT_TICKET_CREATED:
		var c message.EventTicketCreated
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_EVENT_TICKET_CREATION_FAILED:
		var c message.EventTicketCreationFailed
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_EVENT_TICKET_APPROVED:
		var c message.EventTicketApproved
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_EVENT_TICKET_REJECTED:
		var c message.EventTicketRejected
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_EVENT_CARD_AUTHORIZED:
		var c message.EventCardAuthorized
		if err := f.mes.Envelope.UnmarshalTo(&c); err != nil {
			return nil, err
		}
		id, err := event_helper.ParseID(c.OrderId)
		if err != nil {
			return nil, err
		}
		return NewMessage(id, f.mes), nil

	case message.Type_TYPE_EVENT_CARD_AUTHORIZATION_FAILED:
		var c message.EventCardAuthorizationFailed
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
