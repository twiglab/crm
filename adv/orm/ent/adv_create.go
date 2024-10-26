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
	"github.com/twiglab/crm/adv/orm/ent/adv"
)

// AdvCreate is the builder for creating a Adv entity.
type AdvCreate struct {
	config
	mutation *AdvMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ac *AdvCreate) SetCreateTime(t time.Time) *AdvCreate {
	ac.mutation.SetCreateTime(t)
	return ac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ac *AdvCreate) SetNillableCreateTime(t *time.Time) *AdvCreate {
	if t != nil {
		ac.SetCreateTime(*t)
	}
	return ac
}

// SetUpdateTime sets the "update_time" field.
func (ac *AdvCreate) SetUpdateTime(t time.Time) *AdvCreate {
	ac.mutation.SetUpdateTime(t)
	return ac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ac *AdvCreate) SetNillableUpdateTime(t *time.Time) *AdvCreate {
	if t != nil {
		ac.SetUpdateTime(*t)
	}
	return ac
}

// SetCode sets the "code" field.
func (ac *AdvCreate) SetCode(s string) *AdvCreate {
	ac.mutation.SetCode(s)
	return ac
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ac *AdvCreate) SetNillableCode(s *string) *AdvCreate {
	if s != nil {
		ac.SetCode(*s)
	}
	return ac
}

// SetMallCode sets the "mall_code" field.
func (ac *AdvCreate) SetMallCode(s string) *AdvCreate {
	ac.mutation.SetMallCode(s)
	return ac
}

// SetMallName sets the "mall_name" field.
func (ac *AdvCreate) SetMallName(s string) *AdvCreate {
	ac.mutation.SetMallName(s)
	return ac
}

// SetH3Index6 sets the "h3_index_6" field.
func (ac *AdvCreate) SetH3Index6(s string) *AdvCreate {
	ac.mutation.SetH3Index6(s)
	return ac
}

// SetH3Index7 sets the "h3_index_7" field.
func (ac *AdvCreate) SetH3Index7(s string) *AdvCreate {
	ac.mutation.SetH3Index7(s)
	return ac
}

// SetH3Index8 sets the "h3_index_8" field.
func (ac *AdvCreate) SetH3Index8(s string) *AdvCreate {
	ac.mutation.SetH3Index8(s)
	return ac
}

// SetImgPath sets the "img_path" field.
func (ac *AdvCreate) SetImgPath(s string) *AdvCreate {
	ac.mutation.SetImgPath(s)
	return ac
}

// SetURL sets the "url" field.
func (ac *AdvCreate) SetURL(s string) *AdvCreate {
	ac.mutation.SetURL(s)
	return ac
}

// SetRuler sets the "ruler" field.
func (ac *AdvCreate) SetRuler(s string) *AdvCreate {
	ac.mutation.SetRuler(s)
	return ac
}

// SetOrd sets the "ord" field.
func (ac *AdvCreate) SetOrd(i int) *AdvCreate {
	ac.mutation.SetOrd(i)
	return ac
}

// SetNillableOrd sets the "ord" field if the given value is not nil.
func (ac *AdvCreate) SetNillableOrd(i *int) *AdvCreate {
	if i != nil {
		ac.SetOrd(*i)
	}
	return ac
}

// SetMemo sets the "memo" field.
func (ac *AdvCreate) SetMemo(s string) *AdvCreate {
	ac.mutation.SetMemo(s)
	return ac
}

// SetStartTime sets the "start_time" field.
func (ac *AdvCreate) SetStartTime(t time.Time) *AdvCreate {
	ac.mutation.SetStartTime(t)
	return ac
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (ac *AdvCreate) SetNillableStartTime(t *time.Time) *AdvCreate {
	if t != nil {
		ac.SetStartTime(*t)
	}
	return ac
}

// SetEndTime sets the "end_time" field.
func (ac *AdvCreate) SetEndTime(t time.Time) *AdvCreate {
	ac.mutation.SetEndTime(t)
	return ac
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (ac *AdvCreate) SetNillableEndTime(t *time.Time) *AdvCreate {
	if t != nil {
		ac.SetEndTime(*t)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *AdvCreate) SetStatus(i int) *AdvCreate {
	ac.mutation.SetStatus(i)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *AdvCreate) SetNillableStatus(i *int) *AdvCreate {
	if i != nil {
		ac.SetStatus(*i)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AdvCreate) SetID(u uuid.UUID) *AdvCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AdvCreate) SetNillableID(u *uuid.UUID) *AdvCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// Mutation returns the AdvMutation object of the builder.
func (ac *AdvCreate) Mutation() *AdvMutation {
	return ac.mutation
}

// Save creates the Adv in the database.
func (ac *AdvCreate) Save(ctx context.Context) (*Adv, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdvCreate) SaveX(ctx context.Context) *Adv {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdvCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdvCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdvCreate) defaults() {
	if _, ok := ac.mutation.CreateTime(); !ok {
		v := adv.DefaultCreateTime()
		ac.mutation.SetCreateTime(v)
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		v := adv.DefaultUpdateTime()
		ac.mutation.SetUpdateTime(v)
	}
	if _, ok := ac.mutation.Code(); !ok {
		v := adv.DefaultCode()
		ac.mutation.SetCode(v)
	}
	if _, ok := ac.mutation.Ord(); !ok {
		v := adv.DefaultOrd
		ac.mutation.SetOrd(v)
	}
	if _, ok := ac.mutation.StartTime(); !ok {
		v := adv.DefaultStartTime()
		ac.mutation.SetStartTime(v)
	}
	if _, ok := ac.mutation.EndTime(); !ok {
		v := adv.DefaultEndTime()
		ac.mutation.SetEndTime(v)
	}
	if _, ok := ac.mutation.Status(); !ok {
		v := adv.DefaultStatus
		ac.mutation.SetStatus(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := adv.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdvCreate) check() error {
	if _, ok := ac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Adv.create_time"`)}
	}
	if _, ok := ac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Adv.update_time"`)}
	}
	if _, ok := ac.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Adv.code"`)}
	}
	if v, ok := ac.mutation.Code(); ok {
		if err := adv.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Adv.code": %w`, err)}
		}
	}
	if _, ok := ac.mutation.MallCode(); !ok {
		return &ValidationError{Name: "mall_code", err: errors.New(`ent: missing required field "Adv.mall_code"`)}
	}
	if v, ok := ac.mutation.MallCode(); ok {
		if err := adv.MallCodeValidator(v); err != nil {
			return &ValidationError{Name: "mall_code", err: fmt.Errorf(`ent: validator failed for field "Adv.mall_code": %w`, err)}
		}
	}
	if _, ok := ac.mutation.MallName(); !ok {
		return &ValidationError{Name: "mall_name", err: errors.New(`ent: missing required field "Adv.mall_name"`)}
	}
	if v, ok := ac.mutation.MallName(); ok {
		if err := adv.MallNameValidator(v); err != nil {
			return &ValidationError{Name: "mall_name", err: fmt.Errorf(`ent: validator failed for field "Adv.mall_name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.H3Index6(); !ok {
		return &ValidationError{Name: "h3_index_6", err: errors.New(`ent: missing required field "Adv.h3_index_6"`)}
	}
	if v, ok := ac.mutation.H3Index6(); ok {
		if err := adv.H3Index6Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_6", err: fmt.Errorf(`ent: validator failed for field "Adv.h3_index_6": %w`, err)}
		}
	}
	if _, ok := ac.mutation.H3Index7(); !ok {
		return &ValidationError{Name: "h3_index_7", err: errors.New(`ent: missing required field "Adv.h3_index_7"`)}
	}
	if v, ok := ac.mutation.H3Index7(); ok {
		if err := adv.H3Index7Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_7", err: fmt.Errorf(`ent: validator failed for field "Adv.h3_index_7": %w`, err)}
		}
	}
	if _, ok := ac.mutation.H3Index8(); !ok {
		return &ValidationError{Name: "h3_index_8", err: errors.New(`ent: missing required field "Adv.h3_index_8"`)}
	}
	if v, ok := ac.mutation.H3Index8(); ok {
		if err := adv.H3Index8Validator(v); err != nil {
			return &ValidationError{Name: "h3_index_8", err: fmt.Errorf(`ent: validator failed for field "Adv.h3_index_8": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ImgPath(); !ok {
		return &ValidationError{Name: "img_path", err: errors.New(`ent: missing required field "Adv.img_path"`)}
	}
	if v, ok := ac.mutation.ImgPath(); ok {
		if err := adv.ImgPathValidator(v); err != nil {
			return &ValidationError{Name: "img_path", err: fmt.Errorf(`ent: validator failed for field "Adv.img_path": %w`, err)}
		}
	}
	if _, ok := ac.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Adv.url"`)}
	}
	if v, ok := ac.mutation.URL(); ok {
		if err := adv.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Adv.url": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Ruler(); !ok {
		return &ValidationError{Name: "ruler", err: errors.New(`ent: missing required field "Adv.ruler"`)}
	}
	if v, ok := ac.mutation.Ruler(); ok {
		if err := adv.RulerValidator(v); err != nil {
			return &ValidationError{Name: "ruler", err: fmt.Errorf(`ent: validator failed for field "Adv.ruler": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Ord(); !ok {
		return &ValidationError{Name: "ord", err: errors.New(`ent: missing required field "Adv.ord"`)}
	}
	if _, ok := ac.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New(`ent: missing required field "Adv.memo"`)}
	}
	if v, ok := ac.mutation.Memo(); ok {
		if err := adv.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf(`ent: validator failed for field "Adv.memo": %w`, err)}
		}
	}
	if _, ok := ac.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "Adv.start_time"`)}
	}
	if _, ok := ac.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "Adv.end_time"`)}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Adv.status"`)}
	}
	return nil
}

func (ac *AdvCreate) sqlSave(ctx context.Context) (*Adv, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AdvCreate) createSpec() (*Adv, *sqlgraph.CreateSpec) {
	var (
		_node = &Adv{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(adv.Table, sqlgraph.NewFieldSpec(adv.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreateTime(); ok {
		_spec.SetField(adv.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := ac.mutation.UpdateTime(); ok {
		_spec.SetField(adv.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := ac.mutation.Code(); ok {
		_spec.SetField(adv.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := ac.mutation.MallCode(); ok {
		_spec.SetField(adv.FieldMallCode, field.TypeString, value)
		_node.MallCode = value
	}
	if value, ok := ac.mutation.MallName(); ok {
		_spec.SetField(adv.FieldMallName, field.TypeString, value)
		_node.MallName = value
	}
	if value, ok := ac.mutation.H3Index6(); ok {
		_spec.SetField(adv.FieldH3Index6, field.TypeString, value)
		_node.H3Index6 = value
	}
	if value, ok := ac.mutation.H3Index7(); ok {
		_spec.SetField(adv.FieldH3Index7, field.TypeString, value)
		_node.H3Index7 = value
	}
	if value, ok := ac.mutation.H3Index8(); ok {
		_spec.SetField(adv.FieldH3Index8, field.TypeString, value)
		_node.H3Index8 = value
	}
	if value, ok := ac.mutation.ImgPath(); ok {
		_spec.SetField(adv.FieldImgPath, field.TypeString, value)
		_node.ImgPath = value
	}
	if value, ok := ac.mutation.URL(); ok {
		_spec.SetField(adv.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ac.mutation.Ruler(); ok {
		_spec.SetField(adv.FieldRuler, field.TypeString, value)
		_node.Ruler = value
	}
	if value, ok := ac.mutation.Ord(); ok {
		_spec.SetField(adv.FieldOrd, field.TypeInt, value)
		_node.Ord = value
	}
	if value, ok := ac.mutation.Memo(); ok {
		_spec.SetField(adv.FieldMemo, field.TypeString, value)
		_node.Memo = value
	}
	if value, ok := ac.mutation.StartTime(); ok {
		_spec.SetField(adv.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := ac.mutation.EndTime(); ok {
		_spec.SetField(adv.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(adv.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Adv.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AdvUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ac *AdvCreate) OnConflict(opts ...sql.ConflictOption) *AdvUpsertOne {
	ac.conflict = opts
	return &AdvUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Adv.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AdvCreate) OnConflictColumns(columns ...string) *AdvUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AdvUpsertOne{
		create: ac,
	}
}

type (
	// AdvUpsertOne is the builder for "upsert"-ing
	//  one Adv node.
	AdvUpsertOne struct {
		create *AdvCreate
	}

	// AdvUpsert is the "OnConflict" setter.
	AdvUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *AdvUpsert) SetUpdateTime(v time.Time) *AdvUpsert {
	u.Set(adv.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AdvUpsert) UpdateUpdateTime() *AdvUpsert {
	u.SetExcluded(adv.FieldUpdateTime)
	return u
}

// SetMallName sets the "mall_name" field.
func (u *AdvUpsert) SetMallName(v string) *AdvUpsert {
	u.Set(adv.FieldMallName, v)
	return u
}

// UpdateMallName sets the "mall_name" field to the value that was provided on create.
func (u *AdvUpsert) UpdateMallName() *AdvUpsert {
	u.SetExcluded(adv.FieldMallName)
	return u
}

// SetH3Index6 sets the "h3_index_6" field.
func (u *AdvUpsert) SetH3Index6(v string) *AdvUpsert {
	u.Set(adv.FieldH3Index6, v)
	return u
}

// UpdateH3Index6 sets the "h3_index_6" field to the value that was provided on create.
func (u *AdvUpsert) UpdateH3Index6() *AdvUpsert {
	u.SetExcluded(adv.FieldH3Index6)
	return u
}

// SetH3Index7 sets the "h3_index_7" field.
func (u *AdvUpsert) SetH3Index7(v string) *AdvUpsert {
	u.Set(adv.FieldH3Index7, v)
	return u
}

// UpdateH3Index7 sets the "h3_index_7" field to the value that was provided on create.
func (u *AdvUpsert) UpdateH3Index7() *AdvUpsert {
	u.SetExcluded(adv.FieldH3Index7)
	return u
}

// SetH3Index8 sets the "h3_index_8" field.
func (u *AdvUpsert) SetH3Index8(v string) *AdvUpsert {
	u.Set(adv.FieldH3Index8, v)
	return u
}

// UpdateH3Index8 sets the "h3_index_8" field to the value that was provided on create.
func (u *AdvUpsert) UpdateH3Index8() *AdvUpsert {
	u.SetExcluded(adv.FieldH3Index8)
	return u
}

// SetRuler sets the "ruler" field.
func (u *AdvUpsert) SetRuler(v string) *AdvUpsert {
	u.Set(adv.FieldRuler, v)
	return u
}

// UpdateRuler sets the "ruler" field to the value that was provided on create.
func (u *AdvUpsert) UpdateRuler() *AdvUpsert {
	u.SetExcluded(adv.FieldRuler)
	return u
}

// SetOrd sets the "ord" field.
func (u *AdvUpsert) SetOrd(v int) *AdvUpsert {
	u.Set(adv.FieldOrd, v)
	return u
}

// UpdateOrd sets the "ord" field to the value that was provided on create.
func (u *AdvUpsert) UpdateOrd() *AdvUpsert {
	u.SetExcluded(adv.FieldOrd)
	return u
}

// AddOrd adds v to the "ord" field.
func (u *AdvUpsert) AddOrd(v int) *AdvUpsert {
	u.Add(adv.FieldOrd, v)
	return u
}

// SetMemo sets the "memo" field.
func (u *AdvUpsert) SetMemo(v string) *AdvUpsert {
	u.Set(adv.FieldMemo, v)
	return u
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *AdvUpsert) UpdateMemo() *AdvUpsert {
	u.SetExcluded(adv.FieldMemo)
	return u
}

// SetStatus sets the "status" field.
func (u *AdvUpsert) SetStatus(v int) *AdvUpsert {
	u.Set(adv.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AdvUpsert) UpdateStatus() *AdvUpsert {
	u.SetExcluded(adv.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *AdvUpsert) AddStatus(v int) *AdvUpsert {
	u.Add(adv.FieldStatus, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Adv.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(adv.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AdvUpsertOne) UpdateNewValues() *AdvUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(adv.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(adv.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Code(); exists {
			s.SetIgnore(adv.FieldCode)
		}
		if _, exists := u.create.mutation.MallCode(); exists {
			s.SetIgnore(adv.FieldMallCode)
		}
		if _, exists := u.create.mutation.ImgPath(); exists {
			s.SetIgnore(adv.FieldImgPath)
		}
		if _, exists := u.create.mutation.URL(); exists {
			s.SetIgnore(adv.FieldURL)
		}
		if _, exists := u.create.mutation.StartTime(); exists {
			s.SetIgnore(adv.FieldStartTime)
		}
		if _, exists := u.create.mutation.EndTime(); exists {
			s.SetIgnore(adv.FieldEndTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Adv.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AdvUpsertOne) Ignore() *AdvUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AdvUpsertOne) DoNothing() *AdvUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AdvCreate.OnConflict
// documentation for more info.
func (u *AdvUpsertOne) Update(set func(*AdvUpsert)) *AdvUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AdvUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *AdvUpsertOne) SetUpdateTime(v time.Time) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateUpdateTime() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetMallName sets the "mall_name" field.
func (u *AdvUpsertOne) SetMallName(v string) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetMallName(v)
	})
}

// UpdateMallName sets the "mall_name" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateMallName() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateMallName()
	})
}

// SetH3Index6 sets the "h3_index_6" field.
func (u *AdvUpsertOne) SetH3Index6(v string) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetH3Index6(v)
	})
}

// UpdateH3Index6 sets the "h3_index_6" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateH3Index6() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateH3Index6()
	})
}

// SetH3Index7 sets the "h3_index_7" field.
func (u *AdvUpsertOne) SetH3Index7(v string) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetH3Index7(v)
	})
}

// UpdateH3Index7 sets the "h3_index_7" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateH3Index7() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateH3Index7()
	})
}

// SetH3Index8 sets the "h3_index_8" field.
func (u *AdvUpsertOne) SetH3Index8(v string) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetH3Index8(v)
	})
}

// UpdateH3Index8 sets the "h3_index_8" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateH3Index8() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateH3Index8()
	})
}

// SetRuler sets the "ruler" field.
func (u *AdvUpsertOne) SetRuler(v string) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetRuler(v)
	})
}

// UpdateRuler sets the "ruler" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateRuler() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateRuler()
	})
}

// SetOrd sets the "ord" field.
func (u *AdvUpsertOne) SetOrd(v int) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetOrd(v)
	})
}

// AddOrd adds v to the "ord" field.
func (u *AdvUpsertOne) AddOrd(v int) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.AddOrd(v)
	})
}

// UpdateOrd sets the "ord" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateOrd() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateOrd()
	})
}

// SetMemo sets the "memo" field.
func (u *AdvUpsertOne) SetMemo(v string) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateMemo() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateMemo()
	})
}

// SetStatus sets the "status" field.
func (u *AdvUpsertOne) SetStatus(v int) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *AdvUpsertOne) AddStatus(v int) *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AdvUpsertOne) UpdateStatus() *AdvUpsertOne {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *AdvUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AdvCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AdvUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AdvUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AdvUpsertOne.ID is not supported by MySQL driver. Use AdvUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AdvUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AdvCreateBulk is the builder for creating many Adv entities in bulk.
type AdvCreateBulk struct {
	config
	err      error
	builders []*AdvCreate
	conflict []sql.ConflictOption
}

// Save creates the Adv entities in the database.
func (acb *AdvCreateBulk) Save(ctx context.Context) ([]*Adv, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Adv, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdvMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdvCreateBulk) SaveX(ctx context.Context) []*Adv {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdvCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdvCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Adv.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AdvUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (acb *AdvCreateBulk) OnConflict(opts ...sql.ConflictOption) *AdvUpsertBulk {
	acb.conflict = opts
	return &AdvUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Adv.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AdvCreateBulk) OnConflictColumns(columns ...string) *AdvUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AdvUpsertBulk{
		create: acb,
	}
}

// AdvUpsertBulk is the builder for "upsert"-ing
// a bulk of Adv nodes.
type AdvUpsertBulk struct {
	create *AdvCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Adv.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(adv.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AdvUpsertBulk) UpdateNewValues() *AdvUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(adv.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(adv.FieldCreateTime)
			}
			if _, exists := b.mutation.Code(); exists {
				s.SetIgnore(adv.FieldCode)
			}
			if _, exists := b.mutation.MallCode(); exists {
				s.SetIgnore(adv.FieldMallCode)
			}
			if _, exists := b.mutation.ImgPath(); exists {
				s.SetIgnore(adv.FieldImgPath)
			}
			if _, exists := b.mutation.URL(); exists {
				s.SetIgnore(adv.FieldURL)
			}
			if _, exists := b.mutation.StartTime(); exists {
				s.SetIgnore(adv.FieldStartTime)
			}
			if _, exists := b.mutation.EndTime(); exists {
				s.SetIgnore(adv.FieldEndTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Adv.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AdvUpsertBulk) Ignore() *AdvUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AdvUpsertBulk) DoNothing() *AdvUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AdvCreateBulk.OnConflict
// documentation for more info.
func (u *AdvUpsertBulk) Update(set func(*AdvUpsert)) *AdvUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AdvUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *AdvUpsertBulk) SetUpdateTime(v time.Time) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateUpdateTime() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetMallName sets the "mall_name" field.
func (u *AdvUpsertBulk) SetMallName(v string) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetMallName(v)
	})
}

// UpdateMallName sets the "mall_name" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateMallName() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateMallName()
	})
}

// SetH3Index6 sets the "h3_index_6" field.
func (u *AdvUpsertBulk) SetH3Index6(v string) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetH3Index6(v)
	})
}

// UpdateH3Index6 sets the "h3_index_6" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateH3Index6() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateH3Index6()
	})
}

// SetH3Index7 sets the "h3_index_7" field.
func (u *AdvUpsertBulk) SetH3Index7(v string) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetH3Index7(v)
	})
}

// UpdateH3Index7 sets the "h3_index_7" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateH3Index7() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateH3Index7()
	})
}

// SetH3Index8 sets the "h3_index_8" field.
func (u *AdvUpsertBulk) SetH3Index8(v string) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetH3Index8(v)
	})
}

// UpdateH3Index8 sets the "h3_index_8" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateH3Index8() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateH3Index8()
	})
}

// SetRuler sets the "ruler" field.
func (u *AdvUpsertBulk) SetRuler(v string) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetRuler(v)
	})
}

// UpdateRuler sets the "ruler" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateRuler() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateRuler()
	})
}

// SetOrd sets the "ord" field.
func (u *AdvUpsertBulk) SetOrd(v int) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetOrd(v)
	})
}

// AddOrd adds v to the "ord" field.
func (u *AdvUpsertBulk) AddOrd(v int) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.AddOrd(v)
	})
}

// UpdateOrd sets the "ord" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateOrd() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateOrd()
	})
}

// SetMemo sets the "memo" field.
func (u *AdvUpsertBulk) SetMemo(v string) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetMemo(v)
	})
}

// UpdateMemo sets the "memo" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateMemo() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateMemo()
	})
}

// SetStatus sets the "status" field.
func (u *AdvUpsertBulk) SetStatus(v int) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *AdvUpsertBulk) AddStatus(v int) *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AdvUpsertBulk) UpdateStatus() *AdvUpsertBulk {
	return u.Update(func(s *AdvUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *AdvUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AdvCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AdvCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AdvUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
