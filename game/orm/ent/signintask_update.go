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
	"github.com/twiglab/crm/game/orm/ent/predicate"
	"github.com/twiglab/crm/game/orm/ent/signintask"
)

// SignInTaskUpdate is the builder for updating SignInTask entities.
type SignInTaskUpdate struct {
	config
	hooks    []Hook
	mutation *SignInTaskMutation
}

// Where appends a list predicates to the SignInTaskUpdate builder.
func (situ *SignInTaskUpdate) Where(ps ...predicate.SignInTask) *SignInTaskUpdate {
	situ.mutation.Where(ps...)
	return situ
}

// SetTaskID sets the "task_id" field.
func (situ *SignInTaskUpdate) SetTaskID(i int) *SignInTaskUpdate {
	situ.mutation.ResetTaskID()
	situ.mutation.SetTaskID(i)
	return situ
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (situ *SignInTaskUpdate) SetNillableTaskID(i *int) *SignInTaskUpdate {
	if i != nil {
		situ.SetTaskID(*i)
	}
	return situ
}

// AddTaskID adds i to the "task_id" field.
func (situ *SignInTaskUpdate) AddTaskID(i int) *SignInTaskUpdate {
	situ.mutation.AddTaskID(i)
	return situ
}

// SetDescription sets the "description" field.
func (situ *SignInTaskUpdate) SetDescription(s string) *SignInTaskUpdate {
	situ.mutation.SetDescription(s)
	return situ
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (situ *SignInTaskUpdate) SetNillableDescription(s *string) *SignInTaskUpdate {
	if s != nil {
		situ.SetDescription(*s)
	}
	return situ
}

// SetStartTime sets the "start_time" field.
func (situ *SignInTaskUpdate) SetStartTime(t time.Time) *SignInTaskUpdate {
	situ.mutation.SetStartTime(t)
	return situ
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (situ *SignInTaskUpdate) SetNillableStartTime(t *time.Time) *SignInTaskUpdate {
	if t != nil {
		situ.SetStartTime(*t)
	}
	return situ
}

// SetEndTime sets the "end_time" field.
func (situ *SignInTaskUpdate) SetEndTime(t time.Time) *SignInTaskUpdate {
	situ.mutation.SetEndTime(t)
	return situ
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (situ *SignInTaskUpdate) SetNillableEndTime(t *time.Time) *SignInTaskUpdate {
	if t != nil {
		situ.SetEndTime(*t)
	}
	return situ
}

// SetRewardInfo sets the "reward_info" field.
func (situ *SignInTaskUpdate) SetRewardInfo(i int) *SignInTaskUpdate {
	situ.mutation.ResetRewardInfo()
	situ.mutation.SetRewardInfo(i)
	return situ
}

// SetNillableRewardInfo sets the "reward_info" field if the given value is not nil.
func (situ *SignInTaskUpdate) SetNillableRewardInfo(i *int) *SignInTaskUpdate {
	if i != nil {
		situ.SetRewardInfo(*i)
	}
	return situ
}

// AddRewardInfo adds i to the "reward_info" field.
func (situ *SignInTaskUpdate) AddRewardInfo(i int) *SignInTaskUpdate {
	situ.mutation.AddRewardInfo(i)
	return situ
}

// Mutation returns the SignInTaskMutation object of the builder.
func (situ *SignInTaskUpdate) Mutation() *SignInTaskMutation {
	return situ.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (situ *SignInTaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, situ.sqlSave, situ.mutation, situ.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (situ *SignInTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := situ.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (situ *SignInTaskUpdate) Exec(ctx context.Context) error {
	_, err := situ.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (situ *SignInTaskUpdate) ExecX(ctx context.Context) {
	if err := situ.Exec(ctx); err != nil {
		panic(err)
	}
}

func (situ *SignInTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(signintask.Table, signintask.Columns, sqlgraph.NewFieldSpec(signintask.FieldID, field.TypeInt))
	if ps := situ.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := situ.mutation.TaskID(); ok {
		_spec.SetField(signintask.FieldTaskID, field.TypeInt, value)
	}
	if value, ok := situ.mutation.AddedTaskID(); ok {
		_spec.AddField(signintask.FieldTaskID, field.TypeInt, value)
	}
	if value, ok := situ.mutation.Description(); ok {
		_spec.SetField(signintask.FieldDescription, field.TypeString, value)
	}
	if value, ok := situ.mutation.StartTime(); ok {
		_spec.SetField(signintask.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := situ.mutation.EndTime(); ok {
		_spec.SetField(signintask.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := situ.mutation.RewardInfo(); ok {
		_spec.SetField(signintask.FieldRewardInfo, field.TypeInt, value)
	}
	if value, ok := situ.mutation.AddedRewardInfo(); ok {
		_spec.AddField(signintask.FieldRewardInfo, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, situ.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{signintask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	situ.mutation.done = true
	return n, nil
}

// SignInTaskUpdateOne is the builder for updating a single SignInTask entity.
type SignInTaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SignInTaskMutation
}

// SetTaskID sets the "task_id" field.
func (situo *SignInTaskUpdateOne) SetTaskID(i int) *SignInTaskUpdateOne {
	situo.mutation.ResetTaskID()
	situo.mutation.SetTaskID(i)
	return situo
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (situo *SignInTaskUpdateOne) SetNillableTaskID(i *int) *SignInTaskUpdateOne {
	if i != nil {
		situo.SetTaskID(*i)
	}
	return situo
}

// AddTaskID adds i to the "task_id" field.
func (situo *SignInTaskUpdateOne) AddTaskID(i int) *SignInTaskUpdateOne {
	situo.mutation.AddTaskID(i)
	return situo
}

// SetDescription sets the "description" field.
func (situo *SignInTaskUpdateOne) SetDescription(s string) *SignInTaskUpdateOne {
	situo.mutation.SetDescription(s)
	return situo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (situo *SignInTaskUpdateOne) SetNillableDescription(s *string) *SignInTaskUpdateOne {
	if s != nil {
		situo.SetDescription(*s)
	}
	return situo
}

// SetStartTime sets the "start_time" field.
func (situo *SignInTaskUpdateOne) SetStartTime(t time.Time) *SignInTaskUpdateOne {
	situo.mutation.SetStartTime(t)
	return situo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (situo *SignInTaskUpdateOne) SetNillableStartTime(t *time.Time) *SignInTaskUpdateOne {
	if t != nil {
		situo.SetStartTime(*t)
	}
	return situo
}

// SetEndTime sets the "end_time" field.
func (situo *SignInTaskUpdateOne) SetEndTime(t time.Time) *SignInTaskUpdateOne {
	situo.mutation.SetEndTime(t)
	return situo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (situo *SignInTaskUpdateOne) SetNillableEndTime(t *time.Time) *SignInTaskUpdateOne {
	if t != nil {
		situo.SetEndTime(*t)
	}
	return situo
}

// SetRewardInfo sets the "reward_info" field.
func (situo *SignInTaskUpdateOne) SetRewardInfo(i int) *SignInTaskUpdateOne {
	situo.mutation.ResetRewardInfo()
	situo.mutation.SetRewardInfo(i)
	return situo
}

// SetNillableRewardInfo sets the "reward_info" field if the given value is not nil.
func (situo *SignInTaskUpdateOne) SetNillableRewardInfo(i *int) *SignInTaskUpdateOne {
	if i != nil {
		situo.SetRewardInfo(*i)
	}
	return situo
}

// AddRewardInfo adds i to the "reward_info" field.
func (situo *SignInTaskUpdateOne) AddRewardInfo(i int) *SignInTaskUpdateOne {
	situo.mutation.AddRewardInfo(i)
	return situo
}

// Mutation returns the SignInTaskMutation object of the builder.
func (situo *SignInTaskUpdateOne) Mutation() *SignInTaskMutation {
	return situo.mutation
}

// Where appends a list predicates to the SignInTaskUpdate builder.
func (situo *SignInTaskUpdateOne) Where(ps ...predicate.SignInTask) *SignInTaskUpdateOne {
	situo.mutation.Where(ps...)
	return situo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (situo *SignInTaskUpdateOne) Select(field string, fields ...string) *SignInTaskUpdateOne {
	situo.fields = append([]string{field}, fields...)
	return situo
}

// Save executes the query and returns the updated SignInTask entity.
func (situo *SignInTaskUpdateOne) Save(ctx context.Context) (*SignInTask, error) {
	return withHooks(ctx, situo.sqlSave, situo.mutation, situo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (situo *SignInTaskUpdateOne) SaveX(ctx context.Context) *SignInTask {
	node, err := situo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (situo *SignInTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := situo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (situo *SignInTaskUpdateOne) ExecX(ctx context.Context) {
	if err := situo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (situo *SignInTaskUpdateOne) sqlSave(ctx context.Context) (_node *SignInTask, err error) {
	_spec := sqlgraph.NewUpdateSpec(signintask.Table, signintask.Columns, sqlgraph.NewFieldSpec(signintask.FieldID, field.TypeInt))
	id, ok := situo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SignInTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := situo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, signintask.FieldID)
		for _, f := range fields {
			if !signintask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != signintask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := situo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := situo.mutation.TaskID(); ok {
		_spec.SetField(signintask.FieldTaskID, field.TypeInt, value)
	}
	if value, ok := situo.mutation.AddedTaskID(); ok {
		_spec.AddField(signintask.FieldTaskID, field.TypeInt, value)
	}
	if value, ok := situo.mutation.Description(); ok {
		_spec.SetField(signintask.FieldDescription, field.TypeString, value)
	}
	if value, ok := situo.mutation.StartTime(); ok {
		_spec.SetField(signintask.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := situo.mutation.EndTime(); ok {
		_spec.SetField(signintask.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := situo.mutation.RewardInfo(); ok {
		_spec.SetField(signintask.FieldRewardInfo, field.TypeInt, value)
	}
	if value, ok := situo.mutation.AddedRewardInfo(); ok {
		_spec.AddField(signintask.FieldRewardInfo, field.TypeInt, value)
	}
	_node = &SignInTask{config: situo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, situo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{signintask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	situo.mutation.done = true
	return _node, nil
}
