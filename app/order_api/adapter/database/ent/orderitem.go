// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/order"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/orderitem"
)

// OrderItem is the model entity for the OrderItem schema.
type OrderItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SortNo holds the value of the "sortNo" field.
	SortNo int32 `json:"sortNo,omitempty"`
	// Price holds the value of the "price" field.
	Price int64 `json:"price,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int32 `json:"quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderItemQuery when eager-loading is set.
	Edges        OrderItemEdges `json:"edges"`
	order_id     *int
	selectValues sql.SelectValues
}

// OrderItemEdges holds the relations/edges for other nodes in the graph.
type OrderItemEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Order `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) OwnerOrErr() (*Order, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: order.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID, orderitem.FieldSortNo, orderitem.FieldPrice, orderitem.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case orderitem.ForeignKeys[0]: // order_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderItem fields.
func (oi *OrderItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int(value.Int64)
		case orderitem.FieldSortNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sortNo", values[i])
			} else if value.Valid {
				oi.SortNo = int32(value.Int64)
			}
		case orderitem.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				oi.Price = value.Int64
			}
		case orderitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				oi.Quantity = int32(value.Int64)
			}
		case orderitem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_id", value)
			} else if value.Valid {
				oi.order_id = new(int)
				*oi.order_id = int(value.Int64)
			}
		default:
			oi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderItem.
// This includes values selected through modifiers, order, etc.
func (oi *OrderItem) Value(name string) (ent.Value, error) {
	return oi.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the OrderItem entity.
func (oi *OrderItem) QueryOwner() *OrderQuery {
	return NewOrderItemClient(oi.config).QueryOwner(oi)
}

// Update returns a builder for updating this OrderItem.
// Note that you need to call OrderItem.Unwrap() before calling this method if this OrderItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderItem) Update() *OrderItemUpdateOne {
	return NewOrderItemClient(oi.config).UpdateOne(oi)
}

// Unwrap unwraps the OrderItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderItem) Unwrap() *OrderItem {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderItem is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("sortNo=")
	builder.WriteString(fmt.Sprintf("%v", oi.SortNo))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", oi.Price))
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", oi.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// OrderItems is a parsable slice of OrderItem.
type OrderItems []*OrderItem