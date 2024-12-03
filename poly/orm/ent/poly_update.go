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
	"github.com/twiglab/crm/poly/orm/ent/poly"
	"github.com/twiglab/crm/poly/orm/ent/predicate"
)

// PolyUpdate is the builder for updating Poly entities.
type PolyUpdate struct {
	config
	hooks    []Hook
	mutation *PolyMutation
}

// Where appends a list predicates to the PolyUpdate builder.
func (pu *PolyUpdate) Where(ps ...predicate.Poly) *PolyUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetOperator sets the "operator" field.
func (pu *PolyUpdate) SetOperator(s string) *PolyUpdate {
	pu.mutation.SetOperator(s)
	return pu
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableOperator(s *string) *PolyUpdate {
	if s != nil {
		pu.SetOperator(*s)
	}
	return pu
}

// SetRuleCode sets the "rule_code" field.
func (pu *PolyUpdate) SetRuleCode(s string) *PolyUpdate {
	pu.mutation.SetRuleCode(s)
	return pu
}

// SetNillableRuleCode sets the "rule_code" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableRuleCode(s *string) *PolyUpdate {
	if s != nil {
		pu.SetRuleCode(*s)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *PolyUpdate) SetName(s string) *PolyUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableName(s *string) *PolyUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDesc sets the "desc" field.
func (pu *PolyUpdate) SetDesc(s string) *PolyUpdate {
	pu.mutation.SetDesc(s)
	return pu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableDesc(s *string) *PolyUpdate {
	if s != nil {
		pu.SetDesc(*s)
	}
	return pu
}

// SetBudget sets the "budget" field.
func (pu *PolyUpdate) SetBudget(i int64) *PolyUpdate {
	pu.mutation.ResetBudget()
	pu.mutation.SetBudget(i)
	return pu
}

// SetNillableBudget sets the "budget" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableBudget(i *int64) *PolyUpdate {
	if i != nil {
		pu.SetBudget(*i)
	}
	return pu
}

// AddBudget adds i to the "budget" field.
func (pu *PolyUpdate) AddBudget(i int64) *PolyUpdate {
	pu.mutation.AddBudget(i)
	return pu
}

// SetStartTime sets the "start_time" field.
func (pu *PolyUpdate) SetStartTime(t time.Time) *PolyUpdate {
	pu.mutation.SetStartTime(t)
	return pu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableStartTime(t *time.Time) *PolyUpdate {
	if t != nil {
		pu.SetStartTime(*t)
	}
	return pu
}

// SetEndTime sets the "end_time" field.
func (pu *PolyUpdate) SetEndTime(t time.Time) *PolyUpdate {
	pu.mutation.SetEndTime(t)
	return pu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableEndTime(t *time.Time) *PolyUpdate {
	if t != nil {
		pu.SetEndTime(*t)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *PolyUpdate) SetStatus(i int) *PolyUpdate {
	pu.mutation.ResetStatus()
	pu.mutation.SetStatus(i)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableStatus(i *int) *PolyUpdate {
	if i != nil {
		pu.SetStatus(*i)
	}
	return pu
}

// AddStatus adds i to the "status" field.
func (pu *PolyUpdate) AddStatus(i int) *PolyUpdate {
	pu.mutation.AddStatus(i)
	return pu
}

// SetType sets the "type" field.
func (pu *PolyUpdate) SetType(i int) *PolyUpdate {
	pu.mutation.ResetType()
	pu.mutation.SetType(i)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PolyUpdate) SetNillableType(i *int) *PolyUpdate {
	if i != nil {
		pu.SetType(*i)
	}
	return pu
}

// AddType adds i to the "type" field.
func (pu *PolyUpdate) AddType(i int) *PolyUpdate {
	pu.mutation.AddType(i)
	return pu
}

// Mutation returns the PolyMutation object of the builder.
func (pu *PolyUpdate) Mutation() *PolyMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PolyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PolyUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PolyUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PolyUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PolyUpdate) check() error {
	if v, ok := pu.mutation.Operator(); ok {
		if err := poly.OperatorValidator(v); err != nil {
			return &ValidationError{Name: "operator", err: fmt.Errorf(`ent: validator failed for field "Poly.operator": %w`, err)}
		}
	}
	if v, ok := pu.mutation.RuleCode(); ok {
		if err := poly.RuleCodeValidator(v); err != nil {
			return &ValidationError{Name: "rule_code", err: fmt.Errorf(`ent: validator failed for field "Poly.rule_code": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Name(); ok {
		if err := poly.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Poly.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Desc(); ok {
		if err := poly.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf(`ent: validator failed for field "Poly.desc": %w`, err)}
		}
	}
	return nil
}

func (pu *PolyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(poly.Table, poly.Columns, sqlgraph.NewFieldSpec(poly.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Operator(); ok {
		_spec.SetField(poly.FieldOperator, field.TypeString, value)
	}
	if value, ok := pu.mutation.RuleCode(); ok {
		_spec.SetField(poly.FieldRuleCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(poly.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Desc(); ok {
		_spec.SetField(poly.FieldDesc, field.TypeString, value)
	}
	if value, ok := pu.mutation.Budget(); ok {
		_spec.SetField(poly.FieldBudget, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedBudget(); ok {
		_spec.AddField(poly.FieldBudget, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.StartTime(); ok {
		_spec.SetField(poly.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.EndTime(); ok {
		_spec.SetField(poly.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(poly.FieldStatus, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.AddField(poly.FieldStatus, field.TypeInt, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(poly.FieldType, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedType(); ok {
		_spec.AddField(poly.FieldType, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poly.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PolyUpdateOne is the builder for updating a single Poly entity.
type PolyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PolyMutation
}

// SetOperator sets the "operator" field.
func (puo *PolyUpdateOne) SetOperator(s string) *PolyUpdateOne {
	puo.mutation.SetOperator(s)
	return puo
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableOperator(s *string) *PolyUpdateOne {
	if s != nil {
		puo.SetOperator(*s)
	}
	return puo
}

// SetRuleCode sets the "rule_code" field.
func (puo *PolyUpdateOne) SetRuleCode(s string) *PolyUpdateOne {
	puo.mutation.SetRuleCode(s)
	return puo
}

// SetNillableRuleCode sets the "rule_code" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableRuleCode(s *string) *PolyUpdateOne {
	if s != nil {
		puo.SetRuleCode(*s)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *PolyUpdateOne) SetName(s string) *PolyUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableName(s *string) *PolyUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDesc sets the "desc" field.
func (puo *PolyUpdateOne) SetDesc(s string) *PolyUpdateOne {
	puo.mutation.SetDesc(s)
	return puo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableDesc(s *string) *PolyUpdateOne {
	if s != nil {
		puo.SetDesc(*s)
	}
	return puo
}

// SetBudget sets the "budget" field.
func (puo *PolyUpdateOne) SetBudget(i int64) *PolyUpdateOne {
	puo.mutation.ResetBudget()
	puo.mutation.SetBudget(i)
	return puo
}

// SetNillableBudget sets the "budget" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableBudget(i *int64) *PolyUpdateOne {
	if i != nil {
		puo.SetBudget(*i)
	}
	return puo
}

// AddBudget adds i to the "budget" field.
func (puo *PolyUpdateOne) AddBudget(i int64) *PolyUpdateOne {
	puo.mutation.AddBudget(i)
	return puo
}

// SetStartTime sets the "start_time" field.
func (puo *PolyUpdateOne) SetStartTime(t time.Time) *PolyUpdateOne {
	puo.mutation.SetStartTime(t)
	return puo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableStartTime(t *time.Time) *PolyUpdateOne {
	if t != nil {
		puo.SetStartTime(*t)
	}
	return puo
}

// SetEndTime sets the "end_time" field.
func (puo *PolyUpdateOne) SetEndTime(t time.Time) *PolyUpdateOne {
	puo.mutation.SetEndTime(t)
	return puo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableEndTime(t *time.Time) *PolyUpdateOne {
	if t != nil {
		puo.SetEndTime(*t)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *PolyUpdateOne) SetStatus(i int) *PolyUpdateOne {
	puo.mutation.ResetStatus()
	puo.mutation.SetStatus(i)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableStatus(i *int) *PolyUpdateOne {
	if i != nil {
		puo.SetStatus(*i)
	}
	return puo
}

// AddStatus adds i to the "status" field.
func (puo *PolyUpdateOne) AddStatus(i int) *PolyUpdateOne {
	puo.mutation.AddStatus(i)
	return puo
}

// SetType sets the "type" field.
func (puo *PolyUpdateOne) SetType(i int) *PolyUpdateOne {
	puo.mutation.ResetType()
	puo.mutation.SetType(i)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PolyUpdateOne) SetNillableType(i *int) *PolyUpdateOne {
	if i != nil {
		puo.SetType(*i)
	}
	return puo
}

// AddType adds i to the "type" field.
func (puo *PolyUpdateOne) AddType(i int) *PolyUpdateOne {
	puo.mutation.AddType(i)
	return puo
}

// Mutation returns the PolyMutation object of the builder.
func (puo *PolyUpdateOne) Mutation() *PolyMutation {
	return puo.mutation
}

// Where appends a list predicates to the PolyUpdate builder.
func (puo *PolyUpdateOne) Where(ps ...predicate.Poly) *PolyUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PolyUpdateOne) Select(field string, fields ...string) *PolyUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Poly entity.
func (puo *PolyUpdateOne) Save(ctx context.Context) (*Poly, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PolyUpdateOne) SaveX(ctx context.Context) *Poly {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PolyUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PolyUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PolyUpdateOne) check() error {
	if v, ok := puo.mutation.Operator(); ok {
		if err := poly.OperatorValidator(v); err != nil {
			return &ValidationError{Name: "operator", err: fmt.Errorf(`ent: validator failed for field "Poly.operator": %w`, err)}
		}
	}
	if v, ok := puo.mutation.RuleCode(); ok {
		if err := poly.RuleCodeValidator(v); err != nil {
			return &ValidationError{Name: "rule_code", err: fmt.Errorf(`ent: validator failed for field "Poly.rule_code": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Name(); ok {
		if err := poly.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Poly.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Desc(); ok {
		if err := poly.DescValidator(v); err != nil {
			return &ValidationError{Name: "desc", err: fmt.Errorf(`ent: validator failed for field "Poly.desc": %w`, err)}
		}
	}
	return nil
}

func (puo *PolyUpdateOne) sqlSave(ctx context.Context) (_node *Poly, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(poly.Table, poly.Columns, sqlgraph.NewFieldSpec(poly.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Poly.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, poly.FieldID)
		for _, f := range fields {
			if !poly.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != poly.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Operator(); ok {
		_spec.SetField(poly.FieldOperator, field.TypeString, value)
	}
	if value, ok := puo.mutation.RuleCode(); ok {
		_spec.SetField(poly.FieldRuleCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(poly.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Desc(); ok {
		_spec.SetField(poly.FieldDesc, field.TypeString, value)
	}
	if value, ok := puo.mutation.Budget(); ok {
		_spec.SetField(poly.FieldBudget, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedBudget(); ok {
		_spec.AddField(poly.FieldBudget, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.StartTime(); ok {
		_spec.SetField(poly.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.EndTime(); ok {
		_spec.SetField(poly.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(poly.FieldStatus, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.AddField(poly.FieldStatus, field.TypeInt, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(poly.FieldType, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedType(); ok {
		_spec.AddField(poly.FieldType, field.TypeInt, value)
	}
	_node = &Poly{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poly.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
