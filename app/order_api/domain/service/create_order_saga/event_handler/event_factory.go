package event_handler

import (
	"errors"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/lib/event"
)

type factory struct {
	raw event.RawEvent
}

func NewCreateOrderSagaEventFactory(name event.Name, raw event.RawEvent) (domain_event.EventFactory, error) {
	if !IsCreateOrderSagaEvent(name) {
		return nil, errors.New("invalid event name")
	}

	return &factory{raw: raw}, nil
}

func (f *factory) Event() (event.Event, error) {
	// MEMO: 現状ID以外を関心をおいてないのでGeneralEventへ復元している
	return event.NewGeneralEventFromRaw(f.raw)
}
