package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// ProcessedMessage holds the schema definition for the ProcessedMessage entity.
type ProcessedMessage struct {
	ent.Schema
}

// Fields of the ProcessedMessage.
func (ProcessedMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("message_id").
			Unique().
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the ProcessedMessage.
func (ProcessedMessage) Edges() []ent.Edge {
	return nil
}
