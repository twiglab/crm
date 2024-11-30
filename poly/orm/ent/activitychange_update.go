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
	"github.com/twiglab/crm/poly/orm/ent/activitychange"
	"github.com/twiglab/crm/poly/orm/ent/predicate"
)

// ActivityChangeUpdate is the builder for updating ActivityChange entities.
type ActivityChangeUpdate struct {
	config
	hooks    []Hook
	mutation *ActivityChangeMutation
}

// Where appends a list predicates to the ActivityChangeUpdate builder.
func (acu *ActivityChangeUpdate) Where(ps ...predicate.ActivityChange) *ActivityChangeUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetOperator sets the "operator" field.
func (acu *ActivityChangeUpdate) SetOperator(s string) *ActivityChangeUpdate {
	acu.mutation.SetOperator(s)
	return acu
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (acu *ActivityChangeUpdate) SetNillableOperator(s *string) *ActivityChangeUpdate {
	if s != nil {
		acu.SetOperator(*s)
	}
	return acu
}

// SetApprover sets the "approver" field.
func (acu *ActivityChangeUpdate) SetApprover(s string) *ActivityChangeUpdate {
	acu.mutation.SetApprover(s)
	return acu
}

// SetNillableApprover sets the "approver" field if the given value is not nil.
func (acu *ActivityChangeUpdate) SetNillableApprover(s *string) *ActivityChangeUpdate {
	if s != nil {
		acu.SetApprover(*s)
	}
	return acu
}

// SetApproveTime sets the "approve_time" field.
func (acu *ActivityChangeUpdate) SetApproveTime(t time.Time) *ActivityChangeUpdate {
	acu.mutation.SetApproveTime(t)
	return acu
}

// SetNillableApproveTime sets the "approve_time" field if the given value is not nil.
func (acu *ActivityChangeUpdate) SetNillableApproveTime(t *time.Time) *ActivityChangeUpdate {
	if t != nil {
		acu.SetApproveTime(*t)
	}
	return acu
}

// SetStatus sets the "status" field.
func (acu *ActivityChangeUpdate) SetStatus(i int) *ActivityChangeUpdate {
	acu.mutation.ResetStatus()
	acu.mutation.SetStatus(i)
	return acu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (acu *ActivityChangeUpdate) SetNillableStatus(i *int) *ActivityChangeUpdate {
	if i != nil {
		acu.SetStatus(*i)
	}
	return acu
}

// AddStatus adds i to the "status" field.
func (acu *ActivityChangeUpdate) AddStatus(i int) *ActivityChangeUpdate {
	acu.mutation.AddStatus(i)
	return acu
}

// SetChangeSummary sets the "change_summary" field.
func (acu *ActivityChangeUpdate) SetChangeSummary(s string) *ActivityChangeUpdate {
	acu.mutation.SetChangeSummary(s)
	return acu
}

// SetNillableChangeSummary sets the "change_summary" field if the given value is not nil.
func (acu *ActivityChangeUpdate) SetNillableChangeSummary(s *string) *ActivityChangeUpdate {
	if s != nil {
		acu.SetChangeSummary(*s)
	}
	return acu
}

// SetChangeReason sets the "change_reason" field.
func (acu *ActivityChangeUpdate) SetChangeReason(s string) *ActivityChangeUpdate {
	acu.mutation.SetChangeReason(s)
	return acu
}

// SetNillableChangeReason sets the "change_reason" field if the given value is not nil.
func (acu *ActivityChangeUpdate) SetNillableChangeReason(s *string) *ActivityChangeUpdate {
	if s != nil {
		acu.SetChangeReason(*s)
	}
	return acu
}

// SetChangeRecord sets the "change_record" field.
func (acu *ActivityChangeUpdate) SetChangeRecord(s string) *ActivityChangeUpdate {
	acu.mutation.SetChangeRecord(s)
	return acu
}

// SetNillableChangeRecord sets the "change_record" field if the given value is not nil.
func (acu *ActivityChangeUpdate) SetNillableChangeRecord(s *string) *ActivityChangeUpdate {
	if s != nil {
		acu.SetChangeRecord(*s)
	}
	return acu
}

// Mutation returns the ActivityChangeMutation object of the builder.
func (acu *ActivityChangeUpdate) Mutation() *ActivityChangeMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *ActivityChangeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, acu.sqlSave, acu.mutation, acu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acu *ActivityChangeUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *ActivityChangeUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *ActivityChangeUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acu *ActivityChangeUpdate) check() error {
	if v, ok := acu.mutation.Operator(); ok {
		if err := activitychange.OperatorValidator(v); err != nil {
			return &ValidationError{Name: "operator", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.operator": %w`, err)}
		}
	}
	if v, ok := acu.mutation.Approver(); ok {
		if err := activitychange.ApproverValidator(v); err != nil {
			return &ValidationError{Name: "approver", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.approver": %w`, err)}
		}
	}
	if v, ok := acu.mutation.ChangeSummary(); ok {
		if err := activitychange.ChangeSummaryValidator(v); err != nil {
			return &ValidationError{Name: "change_summary", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.change_summary": %w`, err)}
		}
	}
	if v, ok := acu.mutation.ChangeReason(); ok {
		if err := activitychange.ChangeReasonValidator(v); err != nil {
			return &ValidationError{Name: "change_reason", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.change_reason": %w`, err)}
		}
	}
	if v, ok := acu.mutation.ChangeRecord(); ok {
		if err := activitychange.ChangeRecordValidator(v); err != nil {
			return &ValidationError{Name: "change_record", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.change_record": %w`, err)}
		}
	}
	return nil
}

func (acu *ActivityChangeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := acu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(activitychange.Table, activitychange.Columns, sqlgraph.NewFieldSpec(activitychange.FieldID, field.TypeUUID))
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.Operator(); ok {
		_spec.SetField(activitychange.FieldOperator, field.TypeString, value)
	}
	if value, ok := acu.mutation.Approver(); ok {
		_spec.SetField(activitychange.FieldApprover, field.TypeString, value)
	}
	if value, ok := acu.mutation.ApproveTime(); ok {
		_spec.SetField(activitychange.FieldApproveTime, field.TypeTime, value)
	}
	if value, ok := acu.mutation.Status(); ok {
		_spec.SetField(activitychange.FieldStatus, field.TypeInt, value)
	}
	if value, ok := acu.mutation.AddedStatus(); ok {
		_spec.AddField(activitychange.FieldStatus, field.TypeInt, value)
	}
	if value, ok := acu.mutation.ChangeSummary(); ok {
		_spec.SetField(activitychange.FieldChangeSummary, field.TypeString, value)
	}
	if value, ok := acu.mutation.ChangeReason(); ok {
		_spec.SetField(activitychange.FieldChangeReason, field.TypeString, value)
	}
	if value, ok := acu.mutation.ChangeRecord(); ok {
		_spec.SetField(activitychange.FieldChangeRecord, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activitychange.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acu.mutation.done = true
	return n, nil
}

// ActivityChangeUpdateOne is the builder for updating a single ActivityChange entity.
type ActivityChangeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActivityChangeMutation
}

// SetOperator sets the "operator" field.
func (acuo *ActivityChangeUpdateOne) SetOperator(s string) *ActivityChangeUpdateOne {
	acuo.mutation.SetOperator(s)
	return acuo
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (acuo *ActivityChangeUpdateOne) SetNillableOperator(s *string) *ActivityChangeUpdateOne {
	if s != nil {
		acuo.SetOperator(*s)
	}
	return acuo
}

// SetApprover sets the "approver" field.
func (acuo *ActivityChangeUpdateOne) SetApprover(s string) *ActivityChangeUpdateOne {
	acuo.mutation.SetApprover(s)
	return acuo
}

// SetNillableApprover sets the "approver" field if the given value is not nil.
func (acuo *ActivityChangeUpdateOne) SetNillableApprover(s *string) *ActivityChangeUpdateOne {
	if s != nil {
		acuo.SetApprover(*s)
	}
	return acuo
}

// SetApproveTime sets the "approve_time" field.
func (acuo *ActivityChangeUpdateOne) SetApproveTime(t time.Time) *ActivityChangeUpdateOne {
	acuo.mutation.SetApproveTime(t)
	return acuo
}

// SetNillableApproveTime sets the "approve_time" field if the given value is not nil.
func (acuo *ActivityChangeUpdateOne) SetNillableApproveTime(t *time.Time) *ActivityChangeUpdateOne {
	if t != nil {
		acuo.SetApproveTime(*t)
	}
	return acuo
}

// SetStatus sets the "status" field.
func (acuo *ActivityChangeUpdateOne) SetStatus(i int) *ActivityChangeUpdateOne {
	acuo.mutation.ResetStatus()
	acuo.mutation.SetStatus(i)
	return acuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (acuo *ActivityChangeUpdateOne) SetNillableStatus(i *int) *ActivityChangeUpdateOne {
	if i != nil {
		acuo.SetStatus(*i)
	}
	return acuo
}

// AddStatus adds i to the "status" field.
func (acuo *ActivityChangeUpdateOne) AddStatus(i int) *ActivityChangeUpdateOne {
	acuo.mutation.AddStatus(i)
	return acuo
}

// SetChangeSummary sets the "change_summary" field.
func (acuo *ActivityChangeUpdateOne) SetChangeSummary(s string) *ActivityChangeUpdateOne {
	acuo.mutation.SetChangeSummary(s)
	return acuo
}

// SetNillableChangeSummary sets the "change_summary" field if the given value is not nil.
func (acuo *ActivityChangeUpdateOne) SetNillableChangeSummary(s *string) *ActivityChangeUpdateOne {
	if s != nil {
		acuo.SetChangeSummary(*s)
	}
	return acuo
}

// SetChangeReason sets the "change_reason" field.
func (acuo *ActivityChangeUpdateOne) SetChangeReason(s string) *ActivityChangeUpdateOne {
	acuo.mutation.SetChangeReason(s)
	return acuo
}

// SetNillableChangeReason sets the "change_reason" field if the given value is not nil.
func (acuo *ActivityChangeUpdateOne) SetNillableChangeReason(s *string) *ActivityChangeUpdateOne {
	if s != nil {
		acuo.SetChangeReason(*s)
	}
	return acuo
}

// SetChangeRecord sets the "change_record" field.
func (acuo *ActivityChangeUpdateOne) SetChangeRecord(s string) *ActivityChangeUpdateOne {
	acuo.mutation.SetChangeRecord(s)
	return acuo
}

// SetNillableChangeRecord sets the "change_record" field if the given value is not nil.
func (acuo *ActivityChangeUpdateOne) SetNillableChangeRecord(s *string) *ActivityChangeUpdateOne {
	if s != nil {
		acuo.SetChangeRecord(*s)
	}
	return acuo
}

// Mutation returns the ActivityChangeMutation object of the builder.
func (acuo *ActivityChangeUpdateOne) Mutation() *ActivityChangeMutation {
	return acuo.mutation
}

// Where appends a list predicates to the ActivityChangeUpdate builder.
func (acuo *ActivityChangeUpdateOne) Where(ps ...predicate.ActivityChange) *ActivityChangeUpdateOne {
	acuo.mutation.Where(ps...)
	return acuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *ActivityChangeUpdateOne) Select(field string, fields ...string) *ActivityChangeUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated ActivityChange entity.
func (acuo *ActivityChangeUpdateOne) Save(ctx context.Context) (*ActivityChange, error) {
	return withHooks(ctx, acuo.sqlSave, acuo.mutation, acuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *ActivityChangeUpdateOne) SaveX(ctx context.Context) *ActivityChange {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *ActivityChangeUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *ActivityChangeUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acuo *ActivityChangeUpdateOne) check() error {
	if v, ok := acuo.mutation.Operator(); ok {
		if err := activitychange.OperatorValidator(v); err != nil {
			return &ValidationError{Name: "operator", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.operator": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.Approver(); ok {
		if err := activitychange.ApproverValidator(v); err != nil {
			return &ValidationError{Name: "approver", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.approver": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.ChangeSummary(); ok {
		if err := activitychange.ChangeSummaryValidator(v); err != nil {
			return &ValidationError{Name: "change_summary", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.change_summary": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.ChangeReason(); ok {
		if err := activitychange.ChangeReasonValidator(v); err != nil {
			return &ValidationError{Name: "change_reason", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.change_reason": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.ChangeRecord(); ok {
		if err := activitychange.ChangeRecordValidator(v); err != nil {
			return &ValidationError{Name: "change_record", err: fmt.Errorf(`ent: validator failed for field "ActivityChange.change_record": %w`, err)}
		}
	}
	return nil
}

func (acuo *ActivityChangeUpdateOne) sqlSave(ctx context.Context) (_node *ActivityChange, err error) {
	if err := acuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(activitychange.Table, activitychange.Columns, sqlgraph.NewFieldSpec(activitychange.FieldID, field.TypeUUID))
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ActivityChange.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activitychange.FieldID)
		for _, f := range fields {
			if !activitychange.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != activitychange.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.Operator(); ok {
		_spec.SetField(activitychange.FieldOperator, field.TypeString, value)
	}
	if value, ok := acuo.mutation.Approver(); ok {
		_spec.SetField(activitychange.FieldApprover, field.TypeString, value)
	}
	if value, ok := acuo.mutation.ApproveTime(); ok {
		_spec.SetField(activitychange.FieldApproveTime, field.TypeTime, value)
	}
	if value, ok := acuo.mutation.Status(); ok {
		_spec.SetField(activitychange.FieldStatus, field.TypeInt, value)
	}
	if value, ok := acuo.mutation.AddedStatus(); ok {
		_spec.AddField(activitychange.FieldStatus, field.TypeInt, value)
	}
	if value, ok := acuo.mutation.ChangeSummary(); ok {
		_spec.SetField(activitychange.FieldChangeSummary, field.TypeString, value)
	}
	if value, ok := acuo.mutation.ChangeReason(); ok {
		_spec.SetField(activitychange.FieldChangeReason, field.TypeString, value)
	}
	if value, ok := acuo.mutation.ChangeRecord(); ok {
		_spec.SetField(activitychange.FieldChangeRecord, field.TypeString, value)
	}
	_node = &ActivityChange{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activitychange.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acuo.mutation.done = true
	return _node, nil
}