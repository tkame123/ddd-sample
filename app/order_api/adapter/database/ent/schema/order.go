package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/domain/model"
)

type Order struct {
	ent.Schema
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique(),
		field.Int64("approvalLimit"),
		field.Enum("status").
			Values(
				model.OrderStatus_ApprovalPending,
				model.OrderStatus_OrderApproved,
				model.OrderStatus_OrderRejected,
			),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orderItems", OrderItem.Type).
			StorageKey(edge.Column("order_id")).
			StructTag("orderID"),
	}
}
