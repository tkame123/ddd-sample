package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Order struct {
	ent.Schema
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(
			entproto.Methods(entproto.MethodGet),
		),
	}
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("approvalLimit").
			Annotations(
				entproto.Field(2),
			),
	}
}
