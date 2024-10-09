package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("sortNo"),
		field.Int64("price").
			Default(0),
		field.Int32("quantity").
			Default(0),
	}
}

func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Order.Type).
			Ref("orderItems").
			Unique().Required(),
	}
}
