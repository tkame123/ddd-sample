// Code generated by ent, DO NOT EDIT.

package cancelordersagastate

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the cancelordersagastate type in the database.
	Label = "cancel_order_saga_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCurrent holds the string denoting the current field in the database.
	FieldCurrent = "current"
	// FieldTicketID holds the string denoting the ticket_id field in the database.
	FieldTicketID = "ticket_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the cancelordersagastate in the database.
	Table = "cancel_order_saga_states"
)

// Columns holds all SQL columns for cancelordersagastate fields.
var Columns = []string{
	FieldID,
	FieldCurrent,
	FieldTicketID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Current defines the type for the "current" enum field.
type Current string

// Current values.
const (
	CurrentCancelPending               Current = "CancelPending"
	CurrentCancelingTicket             Current = "CancelingTicket"
	CurrentCancelingCard               Current = "CancelingCard"
	CurrentCancellationConfirmingOrder Current = "CancellationConfirmingOrder"
	CurrentOrderCanceled               Current = "OrderCanceled"
	CurrentCancellationRejectingOrder  Current = "CancellationRejectingOrder"
	CurrentOrderCancellationRejected   Current = "OrderCancellationRejected"
)

func (c Current) String() string {
	return string(c)
}

// CurrentValidator is a validator for the "current" field enum values. It is called by the builders before save.
func CurrentValidator(c Current) error {
	switch c {
	case CurrentCancelPending, CurrentCancelingTicket, CurrentCancelingCard, CurrentCancellationConfirmingOrder, CurrentOrderCanceled, CurrentCancellationRejectingOrder, CurrentOrderCancellationRejected:
		return nil
	default:
		return fmt.Errorf("cancelordersagastate: invalid enum value for current field: %q", c)
	}
}

// OrderOption defines the ordering options for the CancelOrderSagaState queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCurrent orders the results by the current field.
func ByCurrent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrent, opts...).ToFunc()
}

// ByTicketID orders the results by the ticket_id field.
func ByTicketID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTicketID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
