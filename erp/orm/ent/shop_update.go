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
	"github.com/twiglab/crm/erp/orm/ent/predicate"
	"github.com/twiglab/crm/erp/orm/ent/shop"
)

// ShopUpdate is the builder for updating Shop entities.
type ShopUpdate struct {
	config
	hooks    []Hook
	mutation *ShopMutation
}

// Where appends a list predicates to the ShopUpdate builder.
func (su *ShopUpdate) Where(ps ...predicate.Shop) *ShopUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *ShopUpdate) SetUpdateTime(t time.Time) *ShopUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetBizClass1 sets the "biz_class_1" field.
func (su *ShopUpdate) SetBizClass1(s string) *ShopUpdate {
	su.mutation.SetBizClass1(s)
	return su
}

// SetNillableBizClass1 sets the "biz_class_1" field if the given value is not nil.
func (su *ShopUpdate) SetNillableBizClass1(s *string) *ShopUpdate {
	if s != nil {
		su.SetBizClass1(*s)
	}
	return su
}

// ClearBizClass1 clears the value of the "biz_class_1" field.
func (su *ShopUpdate) ClearBizClass1() *ShopUpdate {
	su.mutation.ClearBizClass1()
	return su
}

// SetBizClassName1 sets the "biz_class_name_1" field.
func (su *ShopUpdate) SetBizClassName1(s string) *ShopUpdate {
	su.mutation.SetBizClassName1(s)
	return su
}

// SetNillableBizClassName1 sets the "biz_class_name_1" field if the given value is not nil.
func (su *ShopUpdate) SetNillableBizClassName1(s *string) *ShopUpdate {
	if s != nil {
		su.SetBizClassName1(*s)
	}
	return su
}

// ClearBizClassName1 clears the value of the "biz_class_name_1" field.
func (su *ShopUpdate) ClearBizClassName1() *ShopUpdate {
	su.mutation.ClearBizClassName1()
	return su
}

// SetBizClass2 sets the "biz_class_2" field.
func (su *ShopUpdate) SetBizClass2(s string) *ShopUpdate {
	su.mutation.SetBizClass2(s)
	return su
}

// SetNillableBizClass2 sets the "biz_class_2" field if the given value is not nil.
func (su *ShopUpdate) SetNillableBizClass2(s *string) *ShopUpdate {
	if s != nil {
		su.SetBizClass2(*s)
	}
	return su
}

// ClearBizClass2 clears the value of the "biz_class_2" field.
func (su *ShopUpdate) ClearBizClass2() *ShopUpdate {
	su.mutation.ClearBizClass2()
	return su
}

// SetBizClassName2 sets the "biz_class_name_2" field.
func (su *ShopUpdate) SetBizClassName2(s string) *ShopUpdate {
	su.mutation.SetBizClassName2(s)
	return su
}

// SetNillableBizClassName2 sets the "biz_class_name_2" field if the given value is not nil.
func (su *ShopUpdate) SetNillableBizClassName2(s *string) *ShopUpdate {
	if s != nil {
		su.SetBizClassName2(*s)
	}
	return su
}

// ClearBizClassName2 clears the value of the "biz_class_name_2" field.
func (su *ShopUpdate) ClearBizClassName2() *ShopUpdate {
	su.mutation.ClearBizClassName2()
	return su
}

// Mutation returns the ShopMutation object of the builder.
func (su *ShopUpdate) Mutation() *ShopMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ShopUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ShopUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShopUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShopUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *ShopUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := shop.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ShopUpdate) check() error {
	if v, ok := su.mutation.BizClass1(); ok {
		if err := shop.BizClass1Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_1", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_1": %w`, err)}
		}
	}
	if v, ok := su.mutation.BizClassName1(); ok {
		if err := shop.BizClassName1Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_name_1", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_name_1": %w`, err)}
		}
	}
	if v, ok := su.mutation.BizClass2(); ok {
		if err := shop.BizClass2Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_2", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_2": %w`, err)}
		}
	}
	if v, ok := su.mutation.BizClassName2(); ok {
		if err := shop.BizClassName2Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_name_2", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_name_2": %w`, err)}
		}
	}
	return nil
}

func (su *ShopUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(shop.Table, shop.Columns, sqlgraph.NewFieldSpec(shop.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(shop.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.BizClass1(); ok {
		_spec.SetField(shop.FieldBizClass1, field.TypeString, value)
	}
	if su.mutation.BizClass1Cleared() {
		_spec.ClearField(shop.FieldBizClass1, field.TypeString)
	}
	if value, ok := su.mutation.BizClassName1(); ok {
		_spec.SetField(shop.FieldBizClassName1, field.TypeString, value)
	}
	if su.mutation.BizClassName1Cleared() {
		_spec.ClearField(shop.FieldBizClassName1, field.TypeString)
	}
	if value, ok := su.mutation.BizClass2(); ok {
		_spec.SetField(shop.FieldBizClass2, field.TypeString, value)
	}
	if su.mutation.BizClass2Cleared() {
		_spec.ClearField(shop.FieldBizClass2, field.TypeString)
	}
	if value, ok := su.mutation.BizClassName2(); ok {
		_spec.SetField(shop.FieldBizClassName2, field.TypeString, value)
	}
	if su.mutation.BizClassName2Cleared() {
		_spec.ClearField(shop.FieldBizClassName2, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ShopUpdateOne is the builder for updating a single Shop entity.
type ShopUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShopMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *ShopUpdateOne) SetUpdateTime(t time.Time) *ShopUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetBizClass1 sets the "biz_class_1" field.
func (suo *ShopUpdateOne) SetBizClass1(s string) *ShopUpdateOne {
	suo.mutation.SetBizClass1(s)
	return suo
}

// SetNillableBizClass1 sets the "biz_class_1" field if the given value is not nil.
func (suo *ShopUpdateOne) SetNillableBizClass1(s *string) *ShopUpdateOne {
	if s != nil {
		suo.SetBizClass1(*s)
	}
	return suo
}

// ClearBizClass1 clears the value of the "biz_class_1" field.
func (suo *ShopUpdateOne) ClearBizClass1() *ShopUpdateOne {
	suo.mutation.ClearBizClass1()
	return suo
}

// SetBizClassName1 sets the "biz_class_name_1" field.
func (suo *ShopUpdateOne) SetBizClassName1(s string) *ShopUpdateOne {
	suo.mutation.SetBizClassName1(s)
	return suo
}

// SetNillableBizClassName1 sets the "biz_class_name_1" field if the given value is not nil.
func (suo *ShopUpdateOne) SetNillableBizClassName1(s *string) *ShopUpdateOne {
	if s != nil {
		suo.SetBizClassName1(*s)
	}
	return suo
}

// ClearBizClassName1 clears the value of the "biz_class_name_1" field.
func (suo *ShopUpdateOne) ClearBizClassName1() *ShopUpdateOne {
	suo.mutation.ClearBizClassName1()
	return suo
}

// SetBizClass2 sets the "biz_class_2" field.
func (suo *ShopUpdateOne) SetBizClass2(s string) *ShopUpdateOne {
	suo.mutation.SetBizClass2(s)
	return suo
}

// SetNillableBizClass2 sets the "biz_class_2" field if the given value is not nil.
func (suo *ShopUpdateOne) SetNillableBizClass2(s *string) *ShopUpdateOne {
	if s != nil {
		suo.SetBizClass2(*s)
	}
	return suo
}

// ClearBizClass2 clears the value of the "biz_class_2" field.
func (suo *ShopUpdateOne) ClearBizClass2() *ShopUpdateOne {
	suo.mutation.ClearBizClass2()
	return suo
}

// SetBizClassName2 sets the "biz_class_name_2" field.
func (suo *ShopUpdateOne) SetBizClassName2(s string) *ShopUpdateOne {
	suo.mutation.SetBizClassName2(s)
	return suo
}

// SetNillableBizClassName2 sets the "biz_class_name_2" field if the given value is not nil.
func (suo *ShopUpdateOne) SetNillableBizClassName2(s *string) *ShopUpdateOne {
	if s != nil {
		suo.SetBizClassName2(*s)
	}
	return suo
}

// ClearBizClassName2 clears the value of the "biz_class_name_2" field.
func (suo *ShopUpdateOne) ClearBizClassName2() *ShopUpdateOne {
	suo.mutation.ClearBizClassName2()
	return suo
}

// Mutation returns the ShopMutation object of the builder.
func (suo *ShopUpdateOne) Mutation() *ShopMutation {
	return suo.mutation
}

// Where appends a list predicates to the ShopUpdate builder.
func (suo *ShopUpdateOne) Where(ps ...predicate.Shop) *ShopUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ShopUpdateOne) Select(field string, fields ...string) *ShopUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Shop entity.
func (suo *ShopUpdateOne) Save(ctx context.Context) (*Shop, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ShopUpdateOne) SaveX(ctx context.Context) *Shop {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ShopUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShopUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *ShopUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := shop.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ShopUpdateOne) check() error {
	if v, ok := suo.mutation.BizClass1(); ok {
		if err := shop.BizClass1Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_1", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_1": %w`, err)}
		}
	}
	if v, ok := suo.mutation.BizClassName1(); ok {
		if err := shop.BizClassName1Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_name_1", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_name_1": %w`, err)}
		}
	}
	if v, ok := suo.mutation.BizClass2(); ok {
		if err := shop.BizClass2Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_2", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_2": %w`, err)}
		}
	}
	if v, ok := suo.mutation.BizClassName2(); ok {
		if err := shop.BizClassName2Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_name_2", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_name_2": %w`, err)}
		}
	}
	return nil
}

func (suo *ShopUpdateOne) sqlSave(ctx context.Context) (_node *Shop, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(shop.Table, shop.Columns, sqlgraph.NewFieldSpec(shop.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Shop.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shop.FieldID)
		for _, f := range fields {
			if !shop.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shop.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(shop.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.BizClass1(); ok {
		_spec.SetField(shop.FieldBizClass1, field.TypeString, value)
	}
	if suo.mutation.BizClass1Cleared() {
		_spec.ClearField(shop.FieldBizClass1, field.TypeString)
	}
	if value, ok := suo.mutation.BizClassName1(); ok {
		_spec.SetField(shop.FieldBizClassName1, field.TypeString, value)
	}
	if suo.mutation.BizClassName1Cleared() {
		_spec.ClearField(shop.FieldBizClassName1, field.TypeString)
	}
	if value, ok := suo.mutation.BizClass2(); ok {
		_spec.SetField(shop.FieldBizClass2, field.TypeString, value)
	}
	if suo.mutation.BizClass2Cleared() {
		_spec.ClearField(shop.FieldBizClass2, field.TypeString)
	}
	if value, ok := suo.mutation.BizClassName2(); ok {
		_spec.SetField(shop.FieldBizClassName2, field.TypeString, value)
	}
	if suo.mutation.BizClassName2Cleared() {
		_spec.ClearField(shop.FieldBizClassName2, field.TypeString)
	}
	_node = &Shop{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shop.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
