// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/twiglab/crm/card/orm/ent/chargerecord"
	"github.com/twiglab/crm/card/orm/ent/predicate"
)

// ChargeRecordDelete is the builder for deleting a ChargeRecord entity.
type ChargeRecordDelete struct {
	config
	hooks    []Hook
	mutation *ChargeRecordMutation
}

// Where appends a list predicates to the ChargeRecordDelete builder.
func (crd *ChargeRecordDelete) Where(ps ...predicate.ChargeRecord) *ChargeRecordDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *ChargeRecordDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crd.sqlExec, crd.mutation, crd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *ChargeRecordDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *ChargeRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(chargerecord.Table, sqlgraph.NewFieldSpec(chargerecord.FieldID, field.TypeInt))
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crd.mutation.done = true
	return affected, err
}

// ChargeRecordDeleteOne is the builder for deleting a single ChargeRecord entity.
type ChargeRecordDeleteOne struct {
	crd *ChargeRecordDelete
}

// Where appends a list predicates to the ChargeRecordDelete builder.
func (crdo *ChargeRecordDeleteOne) Where(ps ...predicate.ChargeRecord) *ChargeRecordDeleteOne {
	crdo.crd.mutation.Where(ps...)
	return crdo
}

// Exec executes the deletion query.
func (crdo *ChargeRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chargerecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *ChargeRecordDeleteOne) ExecX(ctx context.Context) {
	if err := crdo.Exec(ctx); err != nil {
		panic(err)
	}
}
