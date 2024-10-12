package event_handler_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/model"
	"github.com/tkame123/ddd-sample/app/kitchen_api/domain/service/create_ticket/event_handler"
	"github.com/tkame123/ddd-sample/lib/event_helper"
	"testing"
)

func Test_factory_Event(t *testing.T) {
	type fields struct {
		raw event_helper.RawEvent
	}
	tests := []struct {
		name    string
		fields  fields
		want    event_helper.Event
		wantErr bool
	}{
		{
			name: "Test generalEvent success",
			fields: fields{
				raw: event_helper.RawEvent{
					Type:   event_helper.CommandName_TicketApprove,
					ID:     uuid.Nil.String(),
					Origin: []byte(`{}`),
				},
			},
			want: &event_helper.GeneralEvent{
				Id:     uuid.Nil,
				Type:   event_helper.CommandName_TicketApprove,
				Origin: []byte(`{}`),
			},
			wantErr: false,
		},
		{
			name: "Test generalEvent failed",
			fields: fields{
				raw: event_helper.RawEvent{
					Type:   event_helper.EventName_TicketCreated,
					ID:     uuid.Nil.String(),
					Origin: []byte(`{}`),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test CommandName_TicketCreate success",
			fields: fields{
				raw: event_helper.RawEvent{
					Type:   event_helper.CommandName_TicketCreate,
					ID:     uuid.Nil.String(),
					Origin: []byte(`{"order_id":"00000000-0000-0000-0000-000000000000","items":[{"item_id":"00000000-0000-0000-0000-000000000000","quantity":1}]}`),
				},
			},
			want: &model.TicketCreateCommand{
				OrderID: model.OrderID(uuid.Nil),
				Items: []*model.TicketItemRequest{
					{
						ItemID:   model.ItemID(uuid.Nil),
						Quantity: 1,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := event_handler.NewCreateTicketServiceEventFactory(tt.fields.raw)
			got, err := f.Event()
			if (err != nil) != tt.wantErr {
				t.Errorf("Event() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			g, _ := json.Marshal(got)
			w, _ := json.Marshal(tt.want)
			if string(g) != string(w) {
				t.Errorf("Event() got = %+v, want %+v", string(g), string(w))
			}
		})
	}
}
