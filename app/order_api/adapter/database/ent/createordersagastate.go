// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/createordersagastate"
)

// CreateOrderSagaState is the model entity for the CreateOrderSagaState schema.
type CreateOrderSagaState struct {
	config `json:"-"`
	// ID of the ent.
	// id is orderID
	ID uuid.UUID `json:"id,omitempty"`
	// Current holds the value of the "current" field.
	Current      createordersagastate.Current `json:"current,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CreateOrderSagaState) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case createordersagastate.FieldCurrent:
			values[i] = new(sql.NullString)
		case createordersagastate.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CreateOrderSagaState fields.
func (coss *CreateOrderSagaState) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case createordersagastate.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				coss.ID = *value
			}
		case createordersagastate.FieldCurrent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field current", values[i])
			} else if value.Valid {
				coss.Current = createordersagastate.Current(value.String)
			}
		default:
			coss.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CreateOrderSagaState.
// This includes values selected through modifiers, order, etc.
func (coss *CreateOrderSagaState) Value(name string) (ent.Value, error) {
	return coss.selectValues.Get(name)
}

// Update returns a builder for updating this CreateOrderSagaState.
// Note that you need to call CreateOrderSagaState.Unwrap() before calling this method if this CreateOrderSagaState
// was returned from a transaction, and the transaction was committed or rolled back.
func (coss *CreateOrderSagaState) Update() *CreateOrderSagaStateUpdateOne {
	return NewCreateOrderSagaStateClient(coss.config).UpdateOne(coss)
}

// Unwrap unwraps the CreateOrderSagaState entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (coss *CreateOrderSagaState) Unwrap() *CreateOrderSagaState {
	_tx, ok := coss.config.driver.(*txDriver)
	if !ok {
		panic("ent: CreateOrderSagaState is not a transactional entity")
	}
	coss.config.driver = _tx.drv
	return coss
}

// String implements the fmt.Stringer.
func (coss *CreateOrderSagaState) String() string {
	var builder strings.Builder
	builder.WriteString("CreateOrderSagaState(")
	builder.WriteString(fmt.Sprintf("id=%v, ", coss.ID))
	builder.WriteString("current=")
	builder.WriteString(fmt.Sprintf("%v", coss.Current))
	builder.WriteByte(')')
	return builder.String()
}

// CreateOrderSagaStates is a parsable slice of CreateOrderSagaState.
type CreateOrderSagaStates []*CreateOrderSagaState