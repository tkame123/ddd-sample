package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/service/create_order_saga"
	"time"
)

// CreateOrderSagaState holds the schema definition for the CreateOrderSagaState entity.
type CreateOrderSagaState struct {
	ent.Schema
}

// Fields of the CreateOrderSagaState.
func (CreateOrderSagaState) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("id is orderID").
			Unique(),
		field.Enum("current").
			Values(
				create_order_saga.CreateOrderSagaStep_ApprovalPending,
				create_order_saga.CreateOrderSagaStep_CreatingTicket,
				create_order_saga.CreateOrderSagaStep_AuthorizingCard,
				create_order_saga.CreateOrderSagaStep_ApprovingTicket,
				create_order_saga.CreateOrderSagaStep_ApprovingOrder,
				create_order_saga.CreateOrderSagaStep_OrderApproved,
				create_order_saga.CreateOrderSagaStep_RejectingTicket,
				create_order_saga.CreateOrderSagaStep_RejectingOrder,
				create_order_saga.CreateOrderSagaStep_OrderRejected,
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

// Edges of the CreateOrderSagaState.
func (CreateOrderSagaState) Edges() []ent.Edge {
	return nil
}
