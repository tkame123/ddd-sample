package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/cancel_order_saga"
	"time"
)

// CancelOrderSagaState holds the schema definition for the CancelOrderSagaState entity.
type CancelOrderSagaState struct {
	ent.Schema
}

// Fields of the CancelOrderSagaState.
func (CancelOrderSagaState) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("id is orderID").
			Unique(),
		field.Enum("current").
			Values(
				cancel_order_saga.CancelOrderSagaStep_CancelPending,
				cancel_order_saga.CancelOrderSagaStep_CancelingTicket,
				cancel_order_saga.CancelOrderSagaStep_CancelingCard,
				cancel_order_saga.CancelOrderSagaStep_CancellationConfirmingOrder,
				cancel_order_saga.CancelOrderSagaStep_OrderCanceled,
				cancel_order_saga.CancelOrderSagaStep_CancellationRejectingOrder,
				cancel_order_saga.CancelOrderSagaStep_OrderCancellationRejected,
			),
		field.UUID("ticket_id", uuid.UUID{}).
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the CancelOrderSagaState.
func (CancelOrderSagaState) Edges() []ent.Edge {
	return nil
}
