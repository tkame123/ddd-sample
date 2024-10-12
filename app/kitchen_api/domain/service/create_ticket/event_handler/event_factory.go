package event_handler

import (
	"encoding/json"
	"errors"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/lib/event"
)

type factory struct {
	raw event.RawEvent
}

func NewCreateTicketServiceEventFactory(raw event.RawEvent) (domain_event.EventFactory, error) {
	return &factory{raw: raw}, nil
}

func (f *factory) Event() (event.Event, error) {
	ev, err := event.NewGeneralEventFromRaw(f.raw)
	if err != nil {
		return nil, err
	}

	if !IsCreateTicketEvent(ev.Name()) {
		return nil, errors.New("invalid event name")
	}

	// CreateTicketはパラメータの抽出が必要なのでOriginを復元する
	if ev.Name() == event.CommandName_TicketCreate {
		return f.createTicketEvent()
	}

	return ev, nil
}

func (f *factory) createTicketEvent() (event.Event, error) {
	var e model.TicketCreateCommand
	if err := json.Unmarshal(f.raw.Origin, &e); err != nil {
		return nil, err
	}
	return &e, nil
}
