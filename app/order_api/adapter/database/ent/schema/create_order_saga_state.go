package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
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
				model.CreateOrderSagaStep_ApprovalPending,
				model.CreateOrderSagaStep_CreatingTicket,
				model.CreateOrderSagaStep_AuthorizingCard,
				model.CreateOrderSagaStep_ApprovingTicket,
				model.CreateOrderSagaStep_ApprovingOrder,
				model.CreateOrderSagaStep_OrderApproved,
				model.CreateOrderSagaStep_RejectingTicket,
				model.CreateOrderSagaStep_RejectingOrder,
				model.CreateOrderSagaStep_OrderRejected,
			),
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
