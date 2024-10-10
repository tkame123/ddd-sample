package model_test

import (
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
	"testing"
)

func TestOrderCreatedEvent_ToBody(t *testing.T) {
	type fields struct {
		OrderID model.OrderID
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				OrderID: uuid.Nil,
			},
			want:    `{"type":"event-order-order_created","origin":{"OrderID":"00000000-0000-0000-0000-000000000000"}}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &model.OrderCreatedEvent{
				OrderID: tt.fields.OrderID,
			}
			got, err := e.ToBody()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBody() got = %v, want %v", got, tt.want)
			}
		})
	}
}
