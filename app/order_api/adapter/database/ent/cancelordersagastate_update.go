// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/cancelordersagastate"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/predicate"
)

// CancelOrderSagaStateUpdate is the builder for updating CancelOrderSagaState entities.
type CancelOrderSagaStateUpdate struct {
	config
	hooks    []Hook
	mutation *CancelOrderSagaStateMutation
}

// Where appends a list predicates to the CancelOrderSagaStateUpdate builder.
func (cossu *CancelOrderSagaStateUpdate) Where(ps ...predicate.CancelOrderSagaState) *CancelOrderSagaStateUpdate {
	cossu.mutation.Where(ps...)
	return cossu
}

// SetCurrent sets the "current" field.
func (cossu *CancelOrderSagaStateUpdate) SetCurrent(c cancelordersagastate.Current) *CancelOrderSagaStateUpdate {
	cossu.mutation.SetCurrent(c)
	return cossu
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (cossu *CancelOrderSagaStateUpdate) SetNillableCurrent(c *cancelordersagastate.Current) *CancelOrderSagaStateUpdate {
	if c != nil {
		cossu.SetCurrent(*c)
	}
	return cossu
}

// SetTicketID sets the "ticket_id" field.
func (cossu *CancelOrderSagaStateUpdate) SetTicketID(u uuid.UUID) *CancelOrderSagaStateUpdate {
	cossu.mutation.SetTicketID(u)
	return cossu
}

// SetNillableTicketID sets the "ticket_id" field if the given value is not nil.
func (cossu *CancelOrderSagaStateUpdate) SetNillableTicketID(u *uuid.UUID) *CancelOrderSagaStateUpdate {
	if u != nil {
		cossu.SetTicketID(*u)
	}
	return cossu
}

// ClearTicketID clears the value of the "ticket_id" field.
func (cossu *CancelOrderSagaStateUpdate) ClearTicketID() *CancelOrderSagaStateUpdate {
	cossu.mutation.ClearTicketID()
	return cossu
}

// SetCreatedAt sets the "created_at" field.
func (cossu *CancelOrderSagaStateUpdate) SetCreatedAt(t time.Time) *CancelOrderSagaStateUpdate {
	cossu.mutation.SetCreatedAt(t)
	return cossu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cossu *CancelOrderSagaStateUpdate) SetNillableCreatedAt(t *time.Time) *CancelOrderSagaStateUpdate {
	if t != nil {
		cossu.SetCreatedAt(*t)
	}
	return cossu
}

// SetUpdatedAt sets the "updated_at" field.
func (cossu *CancelOrderSagaStateUpdate) SetUpdatedAt(t time.Time) *CancelOrderSagaStateUpdate {
	cossu.mutation.SetUpdatedAt(t)
	return cossu
}

// Mutation returns the CancelOrderSagaStateMutation object of the builder.
func (cossu *CancelOrderSagaStateUpdate) Mutation() *CancelOrderSagaStateMutation {
	return cossu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cossu *CancelOrderSagaStateUpdate) Save(ctx context.Context) (int, error) {
	cossu.defaults()
	return withHooks(ctx, cossu.sqlSave, cossu.mutation, cossu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cossu *CancelOrderSagaStateUpdate) SaveX(ctx context.Context) int {
	affected, err := cossu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cossu *CancelOrderSagaStateUpdate) Exec(ctx context.Context) error {
	_, err := cossu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cossu *CancelOrderSagaStateUpdate) ExecX(ctx context.Context) {
	if err := cossu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cossu *CancelOrderSagaStateUpdate) defaults() {
	if _, ok := cossu.mutation.UpdatedAt(); !ok {
		v := cancelordersagastate.UpdateDefaultUpdatedAt()
		cossu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cossu *CancelOrderSagaStateUpdate) check() error {
	if v, ok := cossu.mutation.Current(); ok {
		if err := cancelordersagastate.CurrentValidator(v); err != nil {
			return &ValidationError{Name: "current", err: fmt.Errorf(`ent: validator failed for field "CancelOrderSagaState.current": %w`, err)}
		}
	}
	return nil
}

func (cossu *CancelOrderSagaStateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cossu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cancelordersagastate.Table, cancelordersagastate.Columns, sqlgraph.NewFieldSpec(cancelordersagastate.FieldID, field.TypeUUID))
	if ps := cossu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cossu.mutation.Current(); ok {
		_spec.SetField(cancelordersagastate.FieldCurrent, field.TypeEnum, value)
	}
	if value, ok := cossu.mutation.TicketID(); ok {
		_spec.SetField(cancelordersagastate.FieldTicketID, field.TypeUUID, value)
	}
	if cossu.mutation.TicketIDCleared() {
		_spec.ClearField(cancelordersagastate.FieldTicketID, field.TypeUUID)
	}
	if value, ok := cossu.mutation.CreatedAt(); ok {
		_spec.SetField(cancelordersagastate.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cossu.mutation.UpdatedAt(); ok {
		_spec.SetField(cancelordersagastate.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cossu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cancelordersagastate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cossu.mutation.done = true
	return n, nil
}

// CancelOrderSagaStateUpdateOne is the builder for updating a single CancelOrderSagaState entity.
type CancelOrderSagaStateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CancelOrderSagaStateMutation
}

// SetCurrent sets the "current" field.
func (cossuo *CancelOrderSagaStateUpdateOne) SetCurrent(c cancelordersagastate.Current) *CancelOrderSagaStateUpdateOne {
	cossuo.mutation.SetCurrent(c)
	return cossuo
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (cossuo *CancelOrderSagaStateUpdateOne) SetNillableCurrent(c *cancelordersagastate.Current) *CancelOrderSagaStateUpdateOne {
	if c != nil {
		cossuo.SetCurrent(*c)
	}
	return cossuo
}

// SetTicketID sets the "ticket_id" field.
func (cossuo *CancelOrderSagaStateUpdateOne) SetTicketID(u uuid.UUID) *CancelOrderSagaStateUpdateOne {
	cossuo.mutation.SetTicketID(u)
	return cossuo
}

// SetNillableTicketID sets the "ticket_id" field if the given value is not nil.
func (cossuo *CancelOrderSagaStateUpdateOne) SetNillableTicketID(u *uuid.UUID) *CancelOrderSagaStateUpdateOne {
	if u != nil {
		cossuo.SetTicketID(*u)
	}
	return cossuo
}

// ClearTicketID clears the value of the "ticket_id" field.
func (cossuo *CancelOrderSagaStateUpdateOne) ClearTicketID() *CancelOrderSagaStateUpdateOne {
	cossuo.mutation.ClearTicketID()
	return cossuo
}

// SetCreatedAt sets the "created_at" field.
func (cossuo *CancelOrderSagaStateUpdateOne) SetCreatedAt(t time.Time) *CancelOrderSagaStateUpdateOne {
	cossuo.mutation.SetCreatedAt(t)
	return cossuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cossuo *CancelOrderSagaStateUpdateOne) SetNillableCreatedAt(t *time.Time) *CancelOrderSagaStateUpdateOne {
	if t != nil {
		cossuo.SetCreatedAt(*t)
	}
	return cossuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cossuo *CancelOrderSagaStateUpdateOne) SetUpdatedAt(t time.Time) *CancelOrderSagaStateUpdateOne {
	cossuo.mutation.SetUpdatedAt(t)
	return cossuo
}

// Mutation returns the CancelOrderSagaStateMutation object of the builder.
func (cossuo *CancelOrderSagaStateUpdateOne) Mutation() *CancelOrderSagaStateMutation {
	return cossuo.mutation
}

// Where appends a list predicates to the CancelOrderSagaStateUpdate builder.
func (cossuo *CancelOrderSagaStateUpdateOne) Where(ps ...predicate.CancelOrderSagaState) *CancelOrderSagaStateUpdateOne {
	cossuo.mutation.Where(ps...)
	return cossuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cossuo *CancelOrderSagaStateUpdateOne) Select(field string, fields ...string) *CancelOrderSagaStateUpdateOne {
	cossuo.fields = append([]string{field}, fields...)
	return cossuo
}

// Save executes the query and returns the updated CancelOrderSagaState entity.
func (cossuo *CancelOrderSagaStateUpdateOne) Save(ctx context.Context) (*CancelOrderSagaState, error) {
	cossuo.defaults()
	return withHooks(ctx, cossuo.sqlSave, cossuo.mutation, cossuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cossuo *CancelOrderSagaStateUpdateOne) SaveX(ctx context.Context) *CancelOrderSagaState {
	node, err := cossuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cossuo *CancelOrderSagaStateUpdateOne) Exec(ctx context.Context) error {
	_, err := cossuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cossuo *CancelOrderSagaStateUpdateOne) ExecX(ctx context.Context) {
	if err := cossuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cossuo *CancelOrderSagaStateUpdateOne) defaults() {
	if _, ok := cossuo.mutation.UpdatedAt(); !ok {
		v := cancelordersagastate.UpdateDefaultUpdatedAt()
		cossuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cossuo *CancelOrderSagaStateUpdateOne) check() error {
	if v, ok := cossuo.mutation.Current(); ok {
		if err := cancelordersagastate.CurrentValidator(v); err != nil {
			return &ValidationError{Name: "current", err: fmt.Errorf(`ent: validator failed for field "CancelOrderSagaState.current": %w`, err)}
		}
	}
	return nil
}

func (cossuo *CancelOrderSagaStateUpdateOne) sqlSave(ctx context.Context) (_node *CancelOrderSagaState, err error) {
	if err := cossuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cancelordersagastate.Table, cancelordersagastate.Columns, sqlgraph.NewFieldSpec(cancelordersagastate.FieldID, field.TypeUUID))
	id, ok := cossuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CancelOrderSagaState.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cossuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cancelordersagastate.FieldID)
		for _, f := range fields {
			if !cancelordersagastate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cancelordersagastate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cossuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cossuo.mutation.Current(); ok {
		_spec.SetField(cancelordersagastate.FieldCurrent, field.TypeEnum, value)
	}
	if value, ok := cossuo.mutation.TicketID(); ok {
		_spec.SetField(cancelordersagastate.FieldTicketID, field.TypeUUID, value)
	}
	if cossuo.mutation.TicketIDCleared() {
		_spec.ClearField(cancelordersagastate.FieldTicketID, field.TypeUUID)
	}
	if value, ok := cossuo.mutation.CreatedAt(); ok {
		_spec.SetField(cancelordersagastate.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cossuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cancelordersagastate.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &CancelOrderSagaState{config: cossuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cossuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cancelordersagastate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cossuo.mutation.done = true
	return _node, nil
}
