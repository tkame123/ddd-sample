// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/processedmessage"
)

// ProcessedMessage is the model entity for the ProcessedMessage schema.
type ProcessedMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MessageID holds the value of the "message_id" field.
	MessageID string `json:"message_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProcessedMessage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case processedmessage.FieldID:
			values[i] = new(sql.NullInt64)
		case processedmessage.FieldMessageID:
			values[i] = new(sql.NullString)
		case processedmessage.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProcessedMessage fields.
func (pm *ProcessedMessage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case processedmessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pm.ID = int(value.Int64)
		case processedmessage.FieldMessageID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message_id", values[i])
			} else if value.Valid {
				pm.MessageID = value.String
			}
		case processedmessage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pm.CreatedAt = value.Time
			}
		default:
			pm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProcessedMessage.
// This includes values selected through modifiers, order, etc.
func (pm *ProcessedMessage) Value(name string) (ent.Value, error) {
	return pm.selectValues.Get(name)
}

// Update returns a builder for updating this ProcessedMessage.
// Note that you need to call ProcessedMessage.Unwrap() before calling this method if this ProcessedMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (pm *ProcessedMessage) Update() *ProcessedMessageUpdateOne {
	return NewProcessedMessageClient(pm.config).UpdateOne(pm)
}

// Unwrap unwraps the ProcessedMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pm *ProcessedMessage) Unwrap() *ProcessedMessage {
	_tx, ok := pm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProcessedMessage is not a transactional entity")
	}
	pm.config.driver = _tx.drv
	return pm
}

// String implements the fmt.Stringer.
func (pm *ProcessedMessage) String() string {
	var builder strings.Builder
	builder.WriteString("ProcessedMessage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pm.ID))
	builder.WriteString("message_id=")
	builder.WriteString(pm.MessageID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pm.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProcessedMessages is a parsable slice of ProcessedMessage.
type ProcessedMessages []*ProcessedMessage
