// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/order"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/orderitem"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/predicate"
)

// OrderItemUpdate is the builder for updating OrderItem entities.
type OrderItemUpdate struct {
	config
	hooks    []Hook
	mutation *OrderItemMutation
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiu *OrderItemUpdate) Where(ps ...predicate.OrderItem) *OrderItemUpdate {
	oiu.mutation.Where(ps...)
	return oiu
}

// SetSortNo sets the "sortNo" field.
func (oiu *OrderItemUpdate) SetSortNo(i int32) *OrderItemUpdate {
	oiu.mutation.ResetSortNo()
	oiu.mutation.SetSortNo(i)
	return oiu
}

// SetNillableSortNo sets the "sortNo" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableSortNo(i *int32) *OrderItemUpdate {
	if i != nil {
		oiu.SetSortNo(*i)
	}
	return oiu
}

// AddSortNo adds i to the "sortNo" field.
func (oiu *OrderItemUpdate) AddSortNo(i int32) *OrderItemUpdate {
	oiu.mutation.AddSortNo(i)
	return oiu
}

// SetPrice sets the "price" field.
func (oiu *OrderItemUpdate) SetPrice(i int64) *OrderItemUpdate {
	oiu.mutation.ResetPrice()
	oiu.mutation.SetPrice(i)
	return oiu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillablePrice(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetPrice(*i)
	}
	return oiu
}

// AddPrice adds i to the "price" field.
func (oiu *OrderItemUpdate) AddPrice(i int64) *OrderItemUpdate {
	oiu.mutation.AddPrice(i)
	return oiu
}

// SetQuantity sets the "quantity" field.
func (oiu *OrderItemUpdate) SetQuantity(i int32) *OrderItemUpdate {
	oiu.mutation.ResetQuantity()
	oiu.mutation.SetQuantity(i)
	return oiu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableQuantity(i *int32) *OrderItemUpdate {
	if i != nil {
		oiu.SetQuantity(*i)
	}
	return oiu
}

// AddQuantity adds i to the "quantity" field.
func (oiu *OrderItemUpdate) AddQuantity(i int32) *OrderItemUpdate {
	oiu.mutation.AddQuantity(i)
	return oiu
}

// SetOwnerID sets the "owner" edge to the Order entity by ID.
func (oiu *OrderItemUpdate) SetOwnerID(id int) *OrderItemUpdate {
	oiu.mutation.SetOwnerID(id)
	return oiu
}

// SetOwner sets the "owner" edge to the Order entity.
func (oiu *OrderItemUpdate) SetOwner(o *Order) *OrderItemUpdate {
	return oiu.SetOwnerID(o.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiu *OrderItemUpdate) Mutation() *OrderItemMutation {
	return oiu.mutation
}

// ClearOwner clears the "owner" edge to the Order entity.
func (oiu *OrderItemUpdate) ClearOwner() *OrderItemUpdate {
	oiu.mutation.ClearOwner()
	return oiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oiu *OrderItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oiu.sqlSave, oiu.mutation, oiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiu *OrderItemUpdate) SaveX(ctx context.Context) int {
	affected, err := oiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oiu *OrderItemUpdate) Exec(ctx context.Context) error {
	_, err := oiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiu *OrderItemUpdate) ExecX(ctx context.Context) {
	if err := oiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oiu *OrderItemUpdate) check() error {
	if oiu.mutation.OwnerCleared() && len(oiu.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrderItem.owner"`)
	}
	return nil
}

func (oiu *OrderItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt))
	if ps := oiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiu.mutation.SortNo(); ok {
		_spec.SetField(orderitem.FieldSortNo, field.TypeInt32, value)
	}
	if value, ok := oiu.mutation.AddedSortNo(); ok {
		_spec.AddField(orderitem.FieldSortNo, field.TypeInt32, value)
	}
	if value, ok := oiu.mutation.Price(); ok {
		_spec.SetField(orderitem.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.AddedPrice(); ok {
		_spec.AddField(orderitem.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt32, value)
	}
	if value, ok := oiu.mutation.AddedQuantity(); ok {
		_spec.AddField(orderitem.FieldQuantity, field.TypeInt32, value)
	}
	if oiu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OwnerTable,
			Columns: []string{orderitem.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OwnerTable,
			Columns: []string{orderitem.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oiu.mutation.done = true
	return n, nil
}

// OrderItemUpdateOne is the builder for updating a single OrderItem entity.
type OrderItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderItemMutation
}

// SetSortNo sets the "sortNo" field.
func (oiuo *OrderItemUpdateOne) SetSortNo(i int32) *OrderItemUpdateOne {
	oiuo.mutation.ResetSortNo()
	oiuo.mutation.SetSortNo(i)
	return oiuo
}

// SetNillableSortNo sets the "sortNo" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableSortNo(i *int32) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetSortNo(*i)
	}
	return oiuo
}

// AddSortNo adds i to the "sortNo" field.
func (oiuo *OrderItemUpdateOne) AddSortNo(i int32) *OrderItemUpdateOne {
	oiuo.mutation.AddSortNo(i)
	return oiuo
}

// SetPrice sets the "price" field.
func (oiuo *OrderItemUpdateOne) SetPrice(i int64) *OrderItemUpdateOne {
	oiuo.mutation.ResetPrice()
	oiuo.mutation.SetPrice(i)
	return oiuo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillablePrice(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetPrice(*i)
	}
	return oiuo
}

// AddPrice adds i to the "price" field.
func (oiuo *OrderItemUpdateOne) AddPrice(i int64) *OrderItemUpdateOne {
	oiuo.mutation.AddPrice(i)
	return oiuo
}

// SetQuantity sets the "quantity" field.
func (oiuo *OrderItemUpdateOne) SetQuantity(i int32) *OrderItemUpdateOne {
	oiuo.mutation.ResetQuantity()
	oiuo.mutation.SetQuantity(i)
	return oiuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableQuantity(i *int32) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetQuantity(*i)
	}
	return oiuo
}

// AddQuantity adds i to the "quantity" field.
func (oiuo *OrderItemUpdateOne) AddQuantity(i int32) *OrderItemUpdateOne {
	oiuo.mutation.AddQuantity(i)
	return oiuo
}

// SetOwnerID sets the "owner" edge to the Order entity by ID.
func (oiuo *OrderItemUpdateOne) SetOwnerID(id int) *OrderItemUpdateOne {
	oiuo.mutation.SetOwnerID(id)
	return oiuo
}

// SetOwner sets the "owner" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) SetOwner(o *Order) *OrderItemUpdateOne {
	return oiuo.SetOwnerID(o.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiuo *OrderItemUpdateOne) Mutation() *OrderItemMutation {
	return oiuo.mutation
}

// ClearOwner clears the "owner" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) ClearOwner() *OrderItemUpdateOne {
	oiuo.mutation.ClearOwner()
	return oiuo
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiuo *OrderItemUpdateOne) Where(ps ...predicate.OrderItem) *OrderItemUpdateOne {
	oiuo.mutation.Where(ps...)
	return oiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oiuo *OrderItemUpdateOne) Select(field string, fields ...string) *OrderItemUpdateOne {
	oiuo.fields = append([]string{field}, fields...)
	return oiuo
}

// Save executes the query and returns the updated OrderItem entity.
func (oiuo *OrderItemUpdateOne) Save(ctx context.Context) (*OrderItem, error) {
	return withHooks(ctx, oiuo.sqlSave, oiuo.mutation, oiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) SaveX(ctx context.Context) *OrderItem {
	node, err := oiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oiuo *OrderItemUpdateOne) Exec(ctx context.Context) error {
	_, err := oiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) ExecX(ctx context.Context) {
	if err := oiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oiuo *OrderItemUpdateOne) check() error {
	if oiuo.mutation.OwnerCleared() && len(oiuo.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OrderItem.owner"`)
	}
	return nil
}

func (oiuo *OrderItemUpdateOne) sqlSave(ctx context.Context) (_node *OrderItem, err error) {
	if err := oiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt))
	id, ok := oiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for _, f := range fields {
			if !orderitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiuo.mutation.SortNo(); ok {
		_spec.SetField(orderitem.FieldSortNo, field.TypeInt32, value)
	}
	if value, ok := oiuo.mutation.AddedSortNo(); ok {
		_spec.AddField(orderitem.FieldSortNo, field.TypeInt32, value)
	}
	if value, ok := oiuo.mutation.Price(); ok {
		_spec.SetField(orderitem.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.AddedPrice(); ok {
		_spec.AddField(orderitem.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt32, value)
	}
	if value, ok := oiuo.mutation.AddedQuantity(); ok {
		_spec.AddField(orderitem.FieldQuantity, field.TypeInt32, value)
	}
	if oiuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OwnerTable,
			Columns: []string{orderitem.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OwnerTable,
			Columns: []string{orderitem.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderItem{config: oiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oiuo.mutation.done = true
	return _node, nil
}