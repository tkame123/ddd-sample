// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tkame123/ddd-sample/app/order_api/adapter/database/ent/cancelordersagastate"
)

// CancelOrderSagaStateCreate is the builder for creating a CancelOrderSagaState entity.
type CancelOrderSagaStateCreate struct {
	config
	mutation *CancelOrderSagaStateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCurrent sets the "current" field.
func (cossc *CancelOrderSagaStateCreate) SetCurrent(c cancelordersagastate.Current) *CancelOrderSagaStateCreate {
	cossc.mutation.SetCurrent(c)
	return cossc
}

// SetTicketID sets the "ticket_id" field.
func (cossc *CancelOrderSagaStateCreate) SetTicketID(u uuid.UUID) *CancelOrderSagaStateCreate {
	cossc.mutation.SetTicketID(u)
	return cossc
}

// SetNillableTicketID sets the "ticket_id" field if the given value is not nil.
func (cossc *CancelOrderSagaStateCreate) SetNillableTicketID(u *uuid.UUID) *CancelOrderSagaStateCreate {
	if u != nil {
		cossc.SetTicketID(*u)
	}
	return cossc
}

// SetCreatedAt sets the "created_at" field.
func (cossc *CancelOrderSagaStateCreate) SetCreatedAt(t time.Time) *CancelOrderSagaStateCreate {
	cossc.mutation.SetCreatedAt(t)
	return cossc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cossc *CancelOrderSagaStateCreate) SetNillableCreatedAt(t *time.Time) *CancelOrderSagaStateCreate {
	if t != nil {
		cossc.SetCreatedAt(*t)
	}
	return cossc
}

// SetUpdatedAt sets the "updated_at" field.
func (cossc *CancelOrderSagaStateCreate) SetUpdatedAt(t time.Time) *CancelOrderSagaStateCreate {
	cossc.mutation.SetUpdatedAt(t)
	return cossc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cossc *CancelOrderSagaStateCreate) SetNillableUpdatedAt(t *time.Time) *CancelOrderSagaStateCreate {
	if t != nil {
		cossc.SetUpdatedAt(*t)
	}
	return cossc
}

// SetID sets the "id" field.
func (cossc *CancelOrderSagaStateCreate) SetID(u uuid.UUID) *CancelOrderSagaStateCreate {
	cossc.mutation.SetID(u)
	return cossc
}

// Mutation returns the CancelOrderSagaStateMutation object of the builder.
func (cossc *CancelOrderSagaStateCreate) Mutation() *CancelOrderSagaStateMutation {
	return cossc.mutation
}

// Save creates the CancelOrderSagaState in the database.
func (cossc *CancelOrderSagaStateCreate) Save(ctx context.Context) (*CancelOrderSagaState, error) {
	cossc.defaults()
	return withHooks(ctx, cossc.sqlSave, cossc.mutation, cossc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cossc *CancelOrderSagaStateCreate) SaveX(ctx context.Context) *CancelOrderSagaState {
	v, err := cossc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cossc *CancelOrderSagaStateCreate) Exec(ctx context.Context) error {
	_, err := cossc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cossc *CancelOrderSagaStateCreate) ExecX(ctx context.Context) {
	if err := cossc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cossc *CancelOrderSagaStateCreate) defaults() {
	if _, ok := cossc.mutation.CreatedAt(); !ok {
		v := cancelordersagastate.DefaultCreatedAt()
		cossc.mutation.SetCreatedAt(v)
	}
	if _, ok := cossc.mutation.UpdatedAt(); !ok {
		v := cancelordersagastate.DefaultUpdatedAt()
		cossc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cossc *CancelOrderSagaStateCreate) check() error {
	if _, ok := cossc.mutation.Current(); !ok {
		return &ValidationError{Name: "current", err: errors.New(`ent: missing required field "CancelOrderSagaState.current"`)}
	}
	if v, ok := cossc.mutation.Current(); ok {
		if err := cancelordersagastate.CurrentValidator(v); err != nil {
			return &ValidationError{Name: "current", err: fmt.Errorf(`ent: validator failed for field "CancelOrderSagaState.current": %w`, err)}
		}
	}
	if _, ok := cossc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CancelOrderSagaState.created_at"`)}
	}
	if _, ok := cossc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CancelOrderSagaState.updated_at"`)}
	}
	return nil
}

func (cossc *CancelOrderSagaStateCreate) sqlSave(ctx context.Context) (*CancelOrderSagaState, error) {
	if err := cossc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cossc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cossc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cossc.mutation.id = &_node.ID
	cossc.mutation.done = true
	return _node, nil
}

func (cossc *CancelOrderSagaStateCreate) createSpec() (*CancelOrderSagaState, *sqlgraph.CreateSpec) {
	var (
		_node = &CancelOrderSagaState{config: cossc.config}
		_spec = sqlgraph.NewCreateSpec(cancelordersagastate.Table, sqlgraph.NewFieldSpec(cancelordersagastate.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = cossc.conflict
	if id, ok := cossc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cossc.mutation.Current(); ok {
		_spec.SetField(cancelordersagastate.FieldCurrent, field.TypeEnum, value)
		_node.Current = value
	}
	if value, ok := cossc.mutation.TicketID(); ok {
		_spec.SetField(cancelordersagastate.FieldTicketID, field.TypeUUID, value)
		_node.TicketID = &value
	}
	if value, ok := cossc.mutation.CreatedAt(); ok {
		_spec.SetField(cancelordersagastate.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cossc.mutation.UpdatedAt(); ok {
		_spec.SetField(cancelordersagastate.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CancelOrderSagaState.Create().
//		SetCurrent(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CancelOrderSagaStateUpsert) {
//			SetCurrent(v+v).
//		}).
//		Exec(ctx)
func (cossc *CancelOrderSagaStateCreate) OnConflict(opts ...sql.ConflictOption) *CancelOrderSagaStateUpsertOne {
	cossc.conflict = opts
	return &CancelOrderSagaStateUpsertOne{
		create: cossc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CancelOrderSagaState.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cossc *CancelOrderSagaStateCreate) OnConflictColumns(columns ...string) *CancelOrderSagaStateUpsertOne {
	cossc.conflict = append(cossc.conflict, sql.ConflictColumns(columns...))
	return &CancelOrderSagaStateUpsertOne{
		create: cossc,
	}
}

type (
	// CancelOrderSagaStateUpsertOne is the builder for "upsert"-ing
	//  one CancelOrderSagaState node.
	CancelOrderSagaStateUpsertOne struct {
		create *CancelOrderSagaStateCreate
	}

	// CancelOrderSagaStateUpsert is the "OnConflict" setter.
	CancelOrderSagaStateUpsert struct {
		*sql.UpdateSet
	}
)

// SetCurrent sets the "current" field.
func (u *CancelOrderSagaStateUpsert) SetCurrent(v cancelordersagastate.Current) *CancelOrderSagaStateUpsert {
	u.Set(cancelordersagastate.FieldCurrent, v)
	return u
}

// UpdateCurrent sets the "current" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsert) UpdateCurrent() *CancelOrderSagaStateUpsert {
	u.SetExcluded(cancelordersagastate.FieldCurrent)
	return u
}

// SetTicketID sets the "ticket_id" field.
func (u *CancelOrderSagaStateUpsert) SetTicketID(v uuid.UUID) *CancelOrderSagaStateUpsert {
	u.Set(cancelordersagastate.FieldTicketID, v)
	return u
}

// UpdateTicketID sets the "ticket_id" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsert) UpdateTicketID() *CancelOrderSagaStateUpsert {
	u.SetExcluded(cancelordersagastate.FieldTicketID)
	return u
}

// ClearTicketID clears the value of the "ticket_id" field.
func (u *CancelOrderSagaStateUpsert) ClearTicketID() *CancelOrderSagaStateUpsert {
	u.SetNull(cancelordersagastate.FieldTicketID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CancelOrderSagaStateUpsert) SetCreatedAt(v time.Time) *CancelOrderSagaStateUpsert {
	u.Set(cancelordersagastate.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsert) UpdateCreatedAt() *CancelOrderSagaStateUpsert {
	u.SetExcluded(cancelordersagastate.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CancelOrderSagaStateUpsert) SetUpdatedAt(v time.Time) *CancelOrderSagaStateUpsert {
	u.Set(cancelordersagastate.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsert) UpdateUpdatedAt() *CancelOrderSagaStateUpsert {
	u.SetExcluded(cancelordersagastate.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CancelOrderSagaState.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(cancelordersagastate.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CancelOrderSagaStateUpsertOne) UpdateNewValues() *CancelOrderSagaStateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(cancelordersagastate.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CancelOrderSagaState.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CancelOrderSagaStateUpsertOne) Ignore() *CancelOrderSagaStateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CancelOrderSagaStateUpsertOne) DoNothing() *CancelOrderSagaStateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CancelOrderSagaStateCreate.OnConflict
// documentation for more info.
func (u *CancelOrderSagaStateUpsertOne) Update(set func(*CancelOrderSagaStateUpsert)) *CancelOrderSagaStateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CancelOrderSagaStateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCurrent sets the "current" field.
func (u *CancelOrderSagaStateUpsertOne) SetCurrent(v cancelordersagastate.Current) *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.SetCurrent(v)
	})
}

// UpdateCurrent sets the "current" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsertOne) UpdateCurrent() *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.UpdateCurrent()
	})
}

// SetTicketID sets the "ticket_id" field.
func (u *CancelOrderSagaStateUpsertOne) SetTicketID(v uuid.UUID) *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.SetTicketID(v)
	})
}

// UpdateTicketID sets the "ticket_id" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsertOne) UpdateTicketID() *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.UpdateTicketID()
	})
}

// ClearTicketID clears the value of the "ticket_id" field.
func (u *CancelOrderSagaStateUpsertOne) ClearTicketID() *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.ClearTicketID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *CancelOrderSagaStateUpsertOne) SetCreatedAt(v time.Time) *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsertOne) UpdateCreatedAt() *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CancelOrderSagaStateUpsertOne) SetUpdatedAt(v time.Time) *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsertOne) UpdateUpdatedAt() *CancelOrderSagaStateUpsertOne {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *CancelOrderSagaStateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CancelOrderSagaStateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CancelOrderSagaStateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CancelOrderSagaStateUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CancelOrderSagaStateUpsertOne.ID is not supported by MySQL driver. Use CancelOrderSagaStateUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CancelOrderSagaStateUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CancelOrderSagaStateCreateBulk is the builder for creating many CancelOrderSagaState entities in bulk.
type CancelOrderSagaStateCreateBulk struct {
	config
	err      error
	builders []*CancelOrderSagaStateCreate
	conflict []sql.ConflictOption
}

// Save creates the CancelOrderSagaState entities in the database.
func (cosscb *CancelOrderSagaStateCreateBulk) Save(ctx context.Context) ([]*CancelOrderSagaState, error) {
	if cosscb.err != nil {
		return nil, cosscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cosscb.builders))
	nodes := make([]*CancelOrderSagaState, len(cosscb.builders))
	mutators := make([]Mutator, len(cosscb.builders))
	for i := range cosscb.builders {
		func(i int, root context.Context) {
			builder := cosscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CancelOrderSagaStateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cosscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cosscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cosscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cosscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cosscb *CancelOrderSagaStateCreateBulk) SaveX(ctx context.Context) []*CancelOrderSagaState {
	v, err := cosscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cosscb *CancelOrderSagaStateCreateBulk) Exec(ctx context.Context) error {
	_, err := cosscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cosscb *CancelOrderSagaStateCreateBulk) ExecX(ctx context.Context) {
	if err := cosscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CancelOrderSagaState.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CancelOrderSagaStateUpsert) {
//			SetCurrent(v+v).
//		}).
//		Exec(ctx)
func (cosscb *CancelOrderSagaStateCreateBulk) OnConflict(opts ...sql.ConflictOption) *CancelOrderSagaStateUpsertBulk {
	cosscb.conflict = opts
	return &CancelOrderSagaStateUpsertBulk{
		create: cosscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CancelOrderSagaState.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cosscb *CancelOrderSagaStateCreateBulk) OnConflictColumns(columns ...string) *CancelOrderSagaStateUpsertBulk {
	cosscb.conflict = append(cosscb.conflict, sql.ConflictColumns(columns...))
	return &CancelOrderSagaStateUpsertBulk{
		create: cosscb,
	}
}

// CancelOrderSagaStateUpsertBulk is the builder for "upsert"-ing
// a bulk of CancelOrderSagaState nodes.
type CancelOrderSagaStateUpsertBulk struct {
	create *CancelOrderSagaStateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CancelOrderSagaState.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(cancelordersagastate.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CancelOrderSagaStateUpsertBulk) UpdateNewValues() *CancelOrderSagaStateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(cancelordersagastate.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CancelOrderSagaState.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CancelOrderSagaStateUpsertBulk) Ignore() *CancelOrderSagaStateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CancelOrderSagaStateUpsertBulk) DoNothing() *CancelOrderSagaStateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CancelOrderSagaStateCreateBulk.OnConflict
// documentation for more info.
func (u *CancelOrderSagaStateUpsertBulk) Update(set func(*CancelOrderSagaStateUpsert)) *CancelOrderSagaStateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CancelOrderSagaStateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCurrent sets the "current" field.
func (u *CancelOrderSagaStateUpsertBulk) SetCurrent(v cancelordersagastate.Current) *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.SetCurrent(v)
	})
}

// UpdateCurrent sets the "current" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsertBulk) UpdateCurrent() *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.UpdateCurrent()
	})
}

// SetTicketID sets the "ticket_id" field.
func (u *CancelOrderSagaStateUpsertBulk) SetTicketID(v uuid.UUID) *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.SetTicketID(v)
	})
}

// UpdateTicketID sets the "ticket_id" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsertBulk) UpdateTicketID() *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.UpdateTicketID()
	})
}

// ClearTicketID clears the value of the "ticket_id" field.
func (u *CancelOrderSagaStateUpsertBulk) ClearTicketID() *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.ClearTicketID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *CancelOrderSagaStateUpsertBulk) SetCreatedAt(v time.Time) *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsertBulk) UpdateCreatedAt() *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CancelOrderSagaStateUpsertBulk) SetUpdatedAt(v time.Time) *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CancelOrderSagaStateUpsertBulk) UpdateUpdatedAt() *CancelOrderSagaStateUpsertBulk {
	return u.Update(func(s *CancelOrderSagaStateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *CancelOrderSagaStateUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CancelOrderSagaStateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CancelOrderSagaStateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CancelOrderSagaStateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}