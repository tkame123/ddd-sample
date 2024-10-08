package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Order struct {
	ent.Schema
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("orderID", uuid.UUID{}),
		field.Int64("approvalLimit"),
		field.Enum("status").
			Values("OrderStatus_ApprovalPending", "OrderStatus_OrderApproved", "OrderStatus_OrderRejected"),
	}
}
