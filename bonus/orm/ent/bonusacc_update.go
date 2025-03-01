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
	"github.com/twiglab/crm/bonus/orm/ent/bonusacc"
	"github.com/twiglab/crm/bonus/orm/ent/predicate"
)

// BonusAccUpdate is the builder for updating BonusAcc entities.
type BonusAccUpdate struct {
	config
	hooks    []Hook
	mutation *BonusAccMutation
}

// Where appends a list predicates to the BonusAccUpdate builder.
func (bau *BonusAccUpdate) Where(ps ...predicate.BonusAcc) *BonusAccUpdate {
	bau.mutation.Where(ps...)
	return bau
}

// SetUpdateTime sets the "update_time" field.
func (bau *BonusAccUpdate) SetUpdateTime(t time.Time) *BonusAccUpdate {
	bau.mutation.SetUpdateTime(t)
	return bau
}

// SetStatus sets the "status" field.
func (bau *BonusAccUpdate) SetStatus(i int) *BonusAccUpdate {
	bau.mutation.ResetStatus()
	bau.mutation.SetStatus(i)
	return bau
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bau *BonusAccUpdate) SetNillableStatus(i *int) *BonusAccUpdate {
	if i != nil {
		bau.SetStatus(*i)
	}
	return bau
}

// AddStatus adds i to the "status" field.
func (bau *BonusAccUpdate) AddStatus(i int) *BonusAccUpdate {
	bau.mutation.AddStatus(i)
	return bau
}

// Mutation returns the BonusAccMutation object of the builder.
func (bau *BonusAccUpdate) Mutation() *BonusAccMutation {
	return bau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bau *BonusAccUpdate) Save(ctx context.Context) (int, error) {
	bau.defaults()
	return withHooks(ctx, bau.sqlSave, bau.mutation, bau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bau *BonusAccUpdate) SaveX(ctx context.Context) int {
	affected, err := bau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bau *BonusAccUpdate) Exec(ctx context.Context) error {
	_, err := bau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bau *BonusAccUpdate) ExecX(ctx context.Context) {
	if err := bau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bau *BonusAccUpdate) defaults() {
	if _, ok := bau.mutation.UpdateTime(); !ok {
		v := bonusacc.UpdateDefaultUpdateTime()
		bau.mutation.SetUpdateTime(v)
	}
}

func (bau *BonusAccUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bonusacc.Table, bonusacc.Columns, sqlgraph.NewFieldSpec(bonusacc.FieldID, field.TypeInt))
	if ps := bau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bau.mutation.UpdateTime(); ok {
		_spec.SetField(bonusacc.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := bau.mutation.Status(); ok {
		_spec.SetField(bonusacc.FieldStatus, field.TypeInt, value)
	}
	if value, ok := bau.mutation.AddedStatus(); ok {
		_spec.AddField(bonusacc.FieldStatus, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bonusacc.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bau.mutation.done = true
	return n, nil
}

// BonusAccUpdateOne is the builder for updating a single BonusAcc entity.
type BonusAccUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BonusAccMutation
}

// SetUpdateTime sets the "update_time" field.
func (bauo *BonusAccUpdateOne) SetUpdateTime(t time.Time) *BonusAccUpdateOne {
	bauo.mutation.SetUpdateTime(t)
	return bauo
}

// SetStatus sets the "status" field.
func (bauo *BonusAccUpdateOne) SetStatus(i int) *BonusAccUpdateOne {
	bauo.mutation.ResetStatus()
	bauo.mutation.SetStatus(i)
	return bauo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bauo *BonusAccUpdateOne) SetNillableStatus(i *int) *BonusAccUpdateOne {
	if i != nil {
		bauo.SetStatus(*i)
	}
	return bauo
}

// AddStatus adds i to the "status" field.
func (bauo *BonusAccUpdateOne) AddStatus(i int) *BonusAccUpdateOne {
	bauo.mutation.AddStatus(i)
	return bauo
}

// Mutation returns the BonusAccMutation object of the builder.
func (bauo *BonusAccUpdateOne) Mutation() *BonusAccMutation {
	return bauo.mutation
}

// Where appends a list predicates to the BonusAccUpdate builder.
func (bauo *BonusAccUpdateOne) Where(ps ...predicate.BonusAcc) *BonusAccUpdateOne {
	bauo.mutation.Where(ps...)
	return bauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bauo *BonusAccUpdateOne) Select(field string, fields ...string) *BonusAccUpdateOne {
	bauo.fields = append([]string{field}, fields...)
	return bauo
}

// Save executes the query and returns the updated BonusAcc entity.
func (bauo *BonusAccUpdateOne) Save(ctx context.Context) (*BonusAcc, error) {
	bauo.defaults()
	return withHooks(ctx, bauo.sqlSave, bauo.mutation, bauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bauo *BonusAccUpdateOne) SaveX(ctx context.Context) *BonusAcc {
	node, err := bauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bauo *BonusAccUpdateOne) Exec(ctx context.Context) error {
	_, err := bauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bauo *BonusAccUpdateOne) ExecX(ctx context.Context) {
	if err := bauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bauo *BonusAccUpdateOne) defaults() {
	if _, ok := bauo.mutation.UpdateTime(); !ok {
		v := bonusacc.UpdateDefaultUpdateTime()
		bauo.mutation.SetUpdateTime(v)
	}
}

func (bauo *BonusAccUpdateOne) sqlSave(ctx context.Context) (_node *BonusAcc, err error) {
	_spec := sqlgraph.NewUpdateSpec(bonusacc.Table, bonusacc.Columns, sqlgraph.NewFieldSpec(bonusacc.FieldID, field.TypeInt))
	id, ok := bauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BonusAcc.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bonusacc.FieldID)
		for _, f := range fields {
			if !bonusacc.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bonusacc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bauo.mutation.UpdateTime(); ok {
		_spec.SetField(bonusacc.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := bauo.mutation.Status(); ok {
		_spec.SetField(bonusacc.FieldStatus, field.TypeInt, value)
	}
	if value, ok := bauo.mutation.AddedStatus(); ok {
		_spec.AddField(bonusacc.FieldStatus, field.TypeInt, value)
	}
	_node = &BonusAcc{config: bauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bonusacc.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bauo.mutation.done = true
	return _node, nil
}
