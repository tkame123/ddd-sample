// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/cancelordersagastate"
)

// CancelOrderSagaState is the model entity for the CancelOrderSagaState schema.
type CancelOrderSagaState struct {
	config `json:"-"`
	// ID of the ent.
	// id is orderID
	ID uuid.UUID `json:"id,omitempty"`
	// Current holds the value of the "current" field.
	Current cancelordersagastate.Current `json:"current,omitempty"`
	// TicketID holds the value of the "ticket_id" field.
	TicketID *uuid.UUID `json:"ticket_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CancelOrderSagaState) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cancelordersagastate.FieldTicketID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case cancelordersagastate.FieldCurrent:
			values[i] = new(sql.NullString)
		case cancelordersagastate.FieldCreatedAt, cancelordersagastate.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case cancelordersagastate.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CancelOrderSagaState fields.
func (coss *CancelOrderSagaState) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cancelordersagastate.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				coss.ID = *value
			}
		case cancelordersagastate.FieldCurrent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field current", values[i])
			} else if value.Valid {
				coss.Current = cancelordersagastate.Current(value.String)
			}
		case cancelordersagastate.FieldTicketID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field ticket_id", values[i])
			} else if value.Valid {
				coss.TicketID = new(uuid.UUID)
				*coss.TicketID = *value.S.(*uuid.UUID)
			}
		case cancelordersagastate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				coss.CreatedAt = value.Time
			}
		case cancelordersagastate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				coss.UpdatedAt = value.Time
			}
		default:
			coss.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CancelOrderSagaState.
// This includes values selected through modifiers, order, etc.
func (coss *CancelOrderSagaState) Value(name string) (ent.Value, error) {
	return coss.selectValues.Get(name)
}

// Update returns a builder for updating this CancelOrderSagaState.
// Note that you need to call CancelOrderSagaState.Unwrap() before calling this method if this CancelOrderSagaState
// was returned from a transaction, and the transaction was committed or rolled back.
func (coss *CancelOrderSagaState) Update() *CancelOrderSagaStateUpdateOne {
	return NewCancelOrderSagaStateClient(coss.config).UpdateOne(coss)
}

// Unwrap unwraps the CancelOrderSagaState entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (coss *CancelOrderSagaState) Unwrap() *CancelOrderSagaState {
	_tx, ok := coss.config.driver.(*txDriver)
	if !ok {
		panic("ent: CancelOrderSagaState is not a transactional entity")
	}
	coss.config.driver = _tx.drv
	return coss
}

// String implements the fmt.Stringer.
func (coss *CancelOrderSagaState) String() string {
	var builder strings.Builder
	builder.WriteString("CancelOrderSagaState(")
	builder.WriteString(fmt.Sprintf("id=%v, ", coss.ID))
	builder.WriteString("current=")
	builder.WriteString(fmt.Sprintf("%v", coss.Current))
	builder.WriteString(", ")
	if v := coss.TicketID; v != nil {
		builder.WriteString("ticket_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(coss.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(coss.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CancelOrderSagaStates is a parsable slice of CancelOrderSagaState.
type CancelOrderSagaStates []*CancelOrderSagaState