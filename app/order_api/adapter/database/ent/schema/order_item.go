package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique(),
		field.Int32("sortNo"),
		field.UUID("item_id", uuid.UUID{}),
		field.UUID("order_id", uuid.UUID{}),
		field.Int64("price").
			Default(0),
		field.Int32("quantity").
			Default(0),
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

func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("orderItems").
			Field("order_id").
			Unique().Required(),
	}
}
