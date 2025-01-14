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
	"github.com/twiglab/crm/card/orm/ent/chargerecord"
	"github.com/twiglab/crm/card/orm/ent/predicate"
)

// ChargeRecordUpdate is the builder for updating ChargeRecord entities.
type ChargeRecordUpdate struct {
	config
	hooks    []Hook
	mutation *ChargeRecordMutation
}

// Where appends a list predicates to the ChargeRecordUpdate builder.
func (cru *ChargeRecordUpdate) Where(ps ...predicate.ChargeRecord) *ChargeRecordUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetUpdateTime sets the "update_time" field.
func (cru *ChargeRecordUpdate) SetUpdateTime(t time.Time) *ChargeRecordUpdate {
	cru.mutation.SetUpdateTime(t)
	return cru
}

// SetPayCode sets the "pay_code" field.
func (cru *ChargeRecordUpdate) SetPayCode(s string) *ChargeRecordUpdate {
	cru.mutation.SetPayCode(s)
	return cru
}

// SetNillablePayCode sets the "pay_code" field if the given value is not nil.
func (cru *ChargeRecordUpdate) SetNillablePayCode(s *string) *ChargeRecordUpdate {
	if s != nil {
		cru.SetPayCode(*s)
	}
	return cru
}

// SetPayTs sets the "pay_ts" field.
func (cru *ChargeRecordUpdate) SetPayTs(i int64) *ChargeRecordUpdate {
	cru.mutation.ResetPayTs()
	cru.mutation.SetPayTs(i)
	return cru
}

// SetNillablePayTs sets the "pay_ts" field if the given value is not nil.
func (cru *ChargeRecordUpdate) SetNillablePayTs(i *int64) *ChargeRecordUpdate {
	if i != nil {
		cru.SetPayTs(*i)
	}
	return cru
}

// AddPayTs adds i to the "pay_ts" field.
func (cru *ChargeRecordUpdate) AddPayTs(i int64) *ChargeRecordUpdate {
	cru.mutation.AddPayTs(i)
	return cru
}

// SetDeduct sets the "deduct" field.
func (cru *ChargeRecordUpdate) SetDeduct(i int64) *ChargeRecordUpdate {
	cru.mutation.ResetDeduct()
	cru.mutation.SetDeduct(i)
	return cru
}

// SetNillableDeduct sets the "deduct" field if the given value is not nil.
func (cru *ChargeRecordUpdate) SetNillableDeduct(i *int64) *ChargeRecordUpdate {
	if i != nil {
		cru.SetDeduct(*i)
	}
	return cru
}

// AddDeduct adds i to the "deduct" field.
func (cru *ChargeRecordUpdate) AddDeduct(i int64) *ChargeRecordUpdate {
	cru.mutation.AddDeduct(i)
	return cru
}

// SetStatus sets the "status" field.
func (cru *ChargeRecordUpdate) SetStatus(i int) *ChargeRecordUpdate {
	cru.mutation.ResetStatus()
	cru.mutation.SetStatus(i)
	return cru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cru *ChargeRecordUpdate) SetNillableStatus(i *int) *ChargeRecordUpdate {
	if i != nil {
		cru.SetStatus(*i)
	}
	return cru
}

// AddStatus adds i to the "status" field.
func (cru *ChargeRecordUpdate) AddStatus(i int) *ChargeRecordUpdate {
	cru.mutation.AddStatus(i)
	return cru
}

// Mutation returns the ChargeRecordMutation object of the builder.
func (cru *ChargeRecordUpdate) Mutation() *ChargeRecordMutation {
	return cru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *ChargeRecordUpdate) Save(ctx context.Context) (int, error) {
	cru.defaults()
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *ChargeRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *ChargeRecordUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *ChargeRecordUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cru *ChargeRecordUpdate) defaults() {
	if _, ok := cru.mutation.UpdateTime(); !ok {
		v := chargerecord.UpdateDefaultUpdateTime()
		cru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *ChargeRecordUpdate) check() error {
	if v, ok := cru.mutation.PayCode(); ok {
		if err := chargerecord.PayCodeValidator(v); err != nil {
			return &ValidationError{Name: "pay_code", err: fmt.Errorf(`ent: validator failed for field "ChargeRecord.pay_code": %w`, err)}
		}
	}
	return nil
}

func (cru *ChargeRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(chargerecord.Table, chargerecord.Columns, sqlgraph.NewFieldSpec(chargerecord.FieldID, field.TypeInt))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.UpdateTime(); ok {
		_spec.SetField(chargerecord.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cru.mutation.PayCode(); ok {
		_spec.SetField(chargerecord.FieldPayCode, field.TypeString, value)
	}
	if value, ok := cru.mutation.PayTs(); ok {
		_spec.SetField(chargerecord.FieldPayTs, field.TypeInt64, value)
	}
	if value, ok := cru.mutation.AddedPayTs(); ok {
		_spec.AddField(chargerecord.FieldPayTs, field.TypeInt64, value)
	}
	if value, ok := cru.mutation.Deduct(); ok {
		_spec.SetField(chargerecord.FieldDeduct, field.TypeInt64, value)
	}
	if value, ok := cru.mutation.AddedDeduct(); ok {
		_spec.AddField(chargerecord.FieldDeduct, field.TypeInt64, value)
	}
	if value, ok := cru.mutation.Status(); ok {
		_spec.SetField(chargerecord.FieldStatus, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedStatus(); ok {
		_spec.AddField(chargerecord.FieldStatus, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chargerecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// ChargeRecordUpdateOne is the builder for updating a single ChargeRecord entity.
type ChargeRecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChargeRecordMutation
}

// SetUpdateTime sets the "update_time" field.
func (cruo *ChargeRecordUpdateOne) SetUpdateTime(t time.Time) *ChargeRecordUpdateOne {
	cruo.mutation.SetUpdateTime(t)
	return cruo
}

// SetPayCode sets the "pay_code" field.
func (cruo *ChargeRecordUpdateOne) SetPayCode(s string) *ChargeRecordUpdateOne {
	cruo.mutation.SetPayCode(s)
	return cruo
}

// SetNillablePayCode sets the "pay_code" field if the given value is not nil.
func (cruo *ChargeRecordUpdateOne) SetNillablePayCode(s *string) *ChargeRecordUpdateOne {
	if s != nil {
		cruo.SetPayCode(*s)
	}
	return cruo
}

// SetPayTs sets the "pay_ts" field.
func (cruo *ChargeRecordUpdateOne) SetPayTs(i int64) *ChargeRecordUpdateOne {
	cruo.mutation.ResetPayTs()
	cruo.mutation.SetPayTs(i)
	return cruo
}

// SetNillablePayTs sets the "pay_ts" field if the given value is not nil.
func (cruo *ChargeRecordUpdateOne) SetNillablePayTs(i *int64) *ChargeRecordUpdateOne {
	if i != nil {
		cruo.SetPayTs(*i)
	}
	return cruo
}

// AddPayTs adds i to the "pay_ts" field.
func (cruo *ChargeRecordUpdateOne) AddPayTs(i int64) *ChargeRecordUpdateOne {
	cruo.mutation.AddPayTs(i)
	return cruo
}

// SetDeduct sets the "deduct" field.
func (cruo *ChargeRecordUpdateOne) SetDeduct(i int64) *ChargeRecordUpdateOne {
	cruo.mutation.ResetDeduct()
	cruo.mutation.SetDeduct(i)
	return cruo
}

// SetNillableDeduct sets the "deduct" field if the given value is not nil.
func (cruo *ChargeRecordUpdateOne) SetNillableDeduct(i *int64) *ChargeRecordUpdateOne {
	if i != nil {
		cruo.SetDeduct(*i)
	}
	return cruo
}

// AddDeduct adds i to the "deduct" field.
func (cruo *ChargeRecordUpdateOne) AddDeduct(i int64) *ChargeRecordUpdateOne {
	cruo.mutation.AddDeduct(i)
	return cruo
}

// SetStatus sets the "status" field.
func (cruo *ChargeRecordUpdateOne) SetStatus(i int) *ChargeRecordUpdateOne {
	cruo.mutation.ResetStatus()
	cruo.mutation.SetStatus(i)
	return cruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cruo *ChargeRecordUpdateOne) SetNillableStatus(i *int) *ChargeRecordUpdateOne {
	if i != nil {
		cruo.SetStatus(*i)
	}
	return cruo
}

// AddStatus adds i to the "status" field.
func (cruo *ChargeRecordUpdateOne) AddStatus(i int) *ChargeRecordUpdateOne {
	cruo.mutation.AddStatus(i)
	return cruo
}

// Mutation returns the ChargeRecordMutation object of the builder.
func (cruo *ChargeRecordUpdateOne) Mutation() *ChargeRecordMutation {
	return cruo.mutation
}

// Where appends a list predicates to the ChargeRecordUpdate builder.
func (cruo *ChargeRecordUpdateOne) Where(ps ...predicate.ChargeRecord) *ChargeRecordUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *ChargeRecordUpdateOne) Select(field string, fields ...string) *ChargeRecordUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated ChargeRecord entity.
func (cruo *ChargeRecordUpdateOne) Save(ctx context.Context) (*ChargeRecord, error) {
	cruo.defaults()
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *ChargeRecordUpdateOne) SaveX(ctx context.Context) *ChargeRecord {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *ChargeRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *ChargeRecordUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cruo *ChargeRecordUpdateOne) defaults() {
	if _, ok := cruo.mutation.UpdateTime(); !ok {
		v := chargerecord.UpdateDefaultUpdateTime()
		cruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *ChargeRecordUpdateOne) check() error {
	if v, ok := cruo.mutation.PayCode(); ok {
		if err := chargerecord.PayCodeValidator(v); err != nil {
			return &ValidationError{Name: "pay_code", err: fmt.Errorf(`ent: validator failed for field "ChargeRecord.pay_code": %w`, err)}
		}
	}
	return nil
}

func (cruo *ChargeRecordUpdateOne) sqlSave(ctx context.Context) (_node *ChargeRecord, err error) {
	if err := cruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(chargerecord.Table, chargerecord.Columns, sqlgraph.NewFieldSpec(chargerecord.FieldID, field.TypeInt))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChargeRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chargerecord.FieldID)
		for _, f := range fields {
			if !chargerecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chargerecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.UpdateTime(); ok {
		_spec.SetField(chargerecord.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cruo.mutation.PayCode(); ok {
		_spec.SetField(chargerecord.FieldPayCode, field.TypeString, value)
	}
	if value, ok := cruo.mutation.PayTs(); ok {
		_spec.SetField(chargerecord.FieldPayTs, field.TypeInt64, value)
	}
	if value, ok := cruo.mutation.AddedPayTs(); ok {
		_spec.AddField(chargerecord.FieldPayTs, field.TypeInt64, value)
	}
	if value, ok := cruo.mutation.Deduct(); ok {
		_spec.SetField(chargerecord.FieldDeduct, field.TypeInt64, value)
	}
	if value, ok := cruo.mutation.AddedDeduct(); ok {
		_spec.AddField(chargerecord.FieldDeduct, field.TypeInt64, value)
	}
	if value, ok := cruo.mutation.Status(); ok {
		_spec.SetField(chargerecord.FieldStatus, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedStatus(); ok {
		_spec.AddField(chargerecord.FieldStatus, field.TypeInt, value)
	}
	_node = &ChargeRecord{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chargerecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}
