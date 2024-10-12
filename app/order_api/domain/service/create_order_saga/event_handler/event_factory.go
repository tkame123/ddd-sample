package event_handler

import (
	"fmt"
	"github.com/tkame123/ddd-sample/app/order_api/domain/port/domain_event"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"github.com/tkame123/ddd-sample/proto/message"
)

// TODO:　不要かも？

type factory struct {
	raw event_helper.RawEvent
}

func NewCreateOrderSagaEventFactory(tp message.Type, raw event_helper.RawEvent) (domain_event.EventFactory, error) {
	if !IsCreateOrderSagaEvent(tp) {
		return nil, fmt.Errorf("invalid event type: %v", tp)
	}

	return &factory{raw: raw}, nil
}

func (f *factory) Event() (*message.Message, error) {
	// MEMO: 現状ID以外を関心をおいてないのでGeneralEventへ復元している
	//return event_helper.NewGeneralEventFromRaw(f.raw)
	panic("implement me")
}
