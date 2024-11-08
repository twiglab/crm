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
	"github.com/twiglab/crm/adv/orm/ent/mall"
	"github.com/twiglab/crm/adv/orm/ent/predicate"
)

// MallUpdate is the builder for updating Mall entities.
type MallUpdate struct {
	config
	hooks    []Hook
	mutation *MallMutation
}

// Where appends a list predicates to the MallUpdate builder.
func (mu *MallUpdate) Where(ps ...predicate.Mall) *MallUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdateTime sets the "update_time" field.
func (mu *MallUpdate) SetUpdateTime(t time.Time) *MallUpdate {
	mu.mutation.SetUpdateTime(t)
	return mu
}

// SetMallName sets the "mall_name" field.
func (mu *MallUpdate) SetMallName(s string) *MallUpdate {
	mu.mutation.SetMallName(s)
	return mu
}

// SetNillableMallName sets the "mall_name" field if the given value is not nil.
func (mu *MallUpdate) SetNillableMallName(s *string) *MallUpdate {
	if s != nil {
		mu.SetMallName(*s)
	}
	return mu
}

// SetH3Index6 sets the "h3_index_6" field.
func (mu *MallUpdate) SetH3Index6(s string) *MallUpdate {
	mu.mutation.SetH3Index6(s)
	return mu
}

// SetNillableH3Index6 sets the "h3_index_6" field if the given value is not nil.
func (mu *MallUpdate) SetNillableH3Index6(s *string) *MallUpdate {
	if s != nil {
		mu.SetH3Index6(*s)
	}
	return mu
}

// SetH3Index7 sets the "h3_index_7" field.
func (mu *MallUpdate) SetH3Index7(s string) *MallUpdate {
	mu.mutation.SetH3Index7(s)
	return mu
}

// SetNillableH3Index7 sets the "h3_index_7" field if the given value is not nil.
func (mu *MallUpdate) SetNillableH3Index7(s *string) *MallUpdate {
	if s != nil {
		mu.SetH3Index7(*s)
	}
	return mu
}

// SetH3Index8 sets the "h3_index_8" field.
func (mu *MallUpdate) SetH3Index8(s string) *MallUpdate {
	mu.mutation.SetH3Index8(s)
	return mu
}

// SetNillableH3Index8 sets the "h3_index_8" field if the given value is not nil.
func (mu *MallUpdate) SetNillableH3Index8(s *string) *MallUpdate {
	if s != nil {
		mu.SetH3Index8(*s)
	}
	return mu
}

// SetMemo sets the "memo" field.
func (mu *MallUpdate) SetMemo(s string) *MallUpdate {
	mu.mutation.SetMemo(s)
	return mu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (mu *MallUpdate) SetNillableMemo(s *string) *MallUpdate {
	if s != nil {
		mu.SetMemo(*s)
	}
	return mu
}

// SetStatus sets the "status" field.
func (mu *MallUpdate) SetStatus(i int) *MallUpdate {
	mu.mutation.ResetStatus()
	mu.mutation.SetStatus(i)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MallUpdate) SetNillableStatus(i *int) *MallUpdate {
	if i != nil {
		mu.SetStatus(*i)
	}
	return mu
}

// AddStatus adds i to the "status" field.
func (mu *MallUpdate) AddStatus(i int) *MallUpdate {
	mu.mutation.AddStatus(i)
	return mu
}

// Mutation returns the MallMutation object of the builder.
func (mu *MallUpdate) Mutation() *MallMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MallUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MallUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MallUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MallUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MallUpdate) defaults() {
	if _, ok := mu.mutation.UpdateTime(); !ok {
		v := mall.UpdateDefaultUpdateTime()
		mu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MallUpdate) check() error {
	if v, ok := mu.mutation.MallName(); ok {
		if err := mall.MallNameValidator(v); err != nil {
			return &ValidationError{Name: "mall_name", err: fmt.Errorf(`ent: validator failed for field "Mall.mall_name": %w`, err)}
		}
	}
	if v, ok := mu.mutation.H3Index6(); ok {
		if err := mall.H3Index6Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_6", err: fmt.Errorf(`ent: validator failed for field "Mall.h3_index_6": %w`, err)}
		}
	}
	if v, ok := mu.mutation.H3Index7(); ok {
		if err := mall.H3Index7Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_7", err: fmt.Errorf(`ent: validator failed for field "Mall.h3_index_7": %w`, err)}
		}
	}
	if v, ok := mu.mutation.H3Index8(); ok {
		if err := mall.H3Index8Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_8", err: fmt.Errorf(`ent: validator failed for field "Mall.h3_index_8": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Memo(); ok {
		if err := mall.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "Mall.memo": %w`, err)}
		}
	}
	return nil
}

func (mu *MallUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(mall.Table, mall.Columns, sqlgraph.NewFieldSpec(mall.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdateTime(); ok {
		_spec.SetField(mall.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := mu.mutation.MallName(); ok {
		_spec.SetField(mall.FieldMallName, field.TypeString, value)
	}
	if value, ok := mu.mutation.H3Index6(); ok {
		_spec.SetField(mall.FieldH3Index6, field.TypeString, value)
	}
	if value, ok := mu.mutation.H3Index7(); ok {
		_spec.SetField(mall.FieldH3Index7, field.TypeString, value)
	}
	if value, ok := mu.mutation.H3Index8(); ok {
		_spec.SetField(mall.FieldH3Index8, field.TypeString, value)
	}
	if value, ok := mu.mutation.Memo(); ok {
		_spec.SetField(mall.FieldMemo, field.TypeString, value)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(mall.FieldStatus, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedStatus(); ok {
		_spec.AddField(mall.FieldStatus, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mall.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MallUpdateOne is the builder for updating a single Mall entity.
type MallUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MallMutation
}

// SetUpdateTime sets the "update_time" field.
func (muo *MallUpdateOne) SetUpdateTime(t time.Time) *MallUpdateOne {
	muo.mutation.SetUpdateTime(t)
	return muo
}

// SetMallName sets the "mall_name" field.
func (muo *MallUpdateOne) SetMallName(s string) *MallUpdateOne {
	muo.mutation.SetMallName(s)
	return muo
}

// SetNillableMallName sets the "mall_name" field if the given value is not nil.
func (muo *MallUpdateOne) SetNillableMallName(s *string) *MallUpdateOne {
	if s != nil {
		muo.SetMallName(*s)
	}
	return muo
}

// SetH3Index6 sets the "h3_index_6" field.
func (muo *MallUpdateOne) SetH3Index6(s string) *MallUpdateOne {
	muo.mutation.SetH3Index6(s)
	return muo
}

// SetNillableH3Index6 sets the "h3_index_6" field if the given value is not nil.
func (muo *MallUpdateOne) SetNillableH3Index6(s *string) *MallUpdateOne {
	if s != nil {
		muo.SetH3Index6(*s)
	}
	return muo
}

// SetH3Index7 sets the "h3_index_7" field.
func (muo *MallUpdateOne) SetH3Index7(s string) *MallUpdateOne {
	muo.mutation.SetH3Index7(s)
	return muo
}

// SetNillableH3Index7 sets the "h3_index_7" field if the given value is not nil.
func (muo *MallUpdateOne) SetNillableH3Index7(s *string) *MallUpdateOne {
	if s != nil {
		muo.SetH3Index7(*s)
	}
	return muo
}

// SetH3Index8 sets the "h3_index_8" field.
func (muo *MallUpdateOne) SetH3Index8(s string) *MallUpdateOne {
	muo.mutation.SetH3Index8(s)
	return muo
}

// SetNillableH3Index8 sets the "h3_index_8" field if the given value is not nil.
func (muo *MallUpdateOne) SetNillableH3Index8(s *string) *MallUpdateOne {
	if s != nil {
		muo.SetH3Index8(*s)
	}
	return muo
}

// SetMemo sets the "memo" field.
func (muo *MallUpdateOne) SetMemo(s string) *MallUpdateOne {
	muo.mutation.SetMemo(s)
	return muo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (muo *MallUpdateOne) SetNillableMemo(s *string) *MallUpdateOne {
	if s != nil {
		muo.SetMemo(*s)
	}
	return muo
}

// SetStatus sets the "status" field.
func (muo *MallUpdateOne) SetStatus(i int) *MallUpdateOne {
	muo.mutation.ResetStatus()
	muo.mutation.SetStatus(i)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MallUpdateOne) SetNillableStatus(i *int) *MallUpdateOne {
	if i != nil {
		muo.SetStatus(*i)
	}
	return muo
}

// AddStatus adds i to the "status" field.
func (muo *MallUpdateOne) AddStatus(i int) *MallUpdateOne {
	muo.mutation.AddStatus(i)
	return muo
}

// Mutation returns the MallMutation object of the builder.
func (muo *MallUpdateOne) Mutation() *MallMutation {
	return muo.mutation
}

// Where appends a list predicates to the MallUpdate builder.
func (muo *MallUpdateOne) Where(ps ...predicate.Mall) *MallUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MallUpdateOne) Select(field string, fields ...string) *MallUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mall entity.
func (muo *MallUpdateOne) Save(ctx context.Context) (*Mall, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MallUpdateOne) SaveX(ctx context.Context) *Mall {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MallUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MallUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MallUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdateTime(); !ok {
		v := mall.UpdateDefaultUpdateTime()
		muo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MallUpdateOne) check() error {
	if v, ok := muo.mutation.MallName(); ok {
		if err := mall.MallNameValidator(v); err != nil {
			return &ValidationError{Name: "mall_name", err: fmt.Errorf(`ent: validator failed for field "Mall.mall_name": %w`, err)}
		}
	}
	if v, ok := muo.mutation.H3Index6(); ok {
		if err := mall.H3Index6Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_6", err: fmt.Errorf(`ent: validator failed for field "Mall.h3_index_6": %w`, err)}
		}
	}
	if v, ok := muo.mutation.H3Index7(); ok {
		if err := mall.H3Index7Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_7", err: fmt.Errorf(`ent: validator failed for field "Mall.h3_index_7": %w`, err)}
		}
	}
	if v, ok := muo.mutation.H3Index8(); ok {
		if err := mall.H3Index8Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_8", err: fmt.Errorf(`ent: validator failed for field "Mall.h3_index_8": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Memo(); ok {
		if err := mall.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "Mall.memo": %w`, err)}
		}
	}
	return nil
}

func (muo *MallUpdateOne) sqlSave(ctx context.Context) (_node *Mall, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(mall.Table, mall.Columns, sqlgraph.NewFieldSpec(mall.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mall.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mall.FieldID)
		for _, f := range fields {
			if !mall.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mall.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdateTime(); ok {
		_spec.SetField(mall.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := muo.mutation.MallName(); ok {
		_spec.SetField(mall.FieldMallName, field.TypeString, value)
	}
	if value, ok := muo.mutation.H3Index6(); ok {
		_spec.SetField(mall.FieldH3Index6, field.TypeString, value)
	}
	if value, ok := muo.mutation.H3Index7(); ok {
		_spec.SetField(mall.FieldH3Index7, field.TypeString, value)
	}
	if value, ok := muo.mutation.H3Index8(); ok {
		_spec.SetField(mall.FieldH3Index8, field.TypeString, value)
	}
	if value, ok := muo.mutation.Memo(); ok {
		_spec.SetField(mall.FieldMemo, field.TypeString, value)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(mall.FieldStatus, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedStatus(); ok {
		_spec.AddField(mall.FieldStatus, field.TypeInt, value)
	}
	_node = &Mall{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mall.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
