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
	"github.com/twiglab/crm/member/orm/ent/member"
	"github.com/twiglab/crm/member/orm/ent/predicate"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdateTime sets the "update_time" field.
func (mu *MemberUpdate) SetUpdateTime(t time.Time) *MemberUpdate {
	mu.mutation.SetUpdateTime(t)
	return mu
}

// SetCodeBin sets the "code_bin" field.
func (mu *MemberUpdate) SetCodeBin(s string) *MemberUpdate {
	mu.mutation.SetCodeBin(s)
	return mu
}

// SetNillableCodeBin sets the "code_bin" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableCodeBin(s *string) *MemberUpdate {
	if s != nil {
		mu.SetCodeBin(*s)
	}
	return mu
}

// ClearCodeBin clears the value of the "code_bin" field.
func (mu *MemberUpdate) ClearCodeBin() *MemberUpdate {
	mu.mutation.ClearCodeBin()
	return mu
}

// SetPhone sets the "phone" field.
func (mu *MemberUpdate) SetPhone(s string) *MemberUpdate {
	mu.mutation.SetPhone(s)
	return mu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (mu *MemberUpdate) SetNillablePhone(s *string) *MemberUpdate {
	if s != nil {
		mu.SetPhone(*s)
	}
	return mu
}

// ClearPhone clears the value of the "phone" field.
func (mu *MemberUpdate) ClearPhone() *MemberUpdate {
	mu.mutation.ClearPhone()
	return mu
}

// SetNickname sets the "nickname" field.
func (mu *MemberUpdate) SetNickname(s string) *MemberUpdate {
	mu.mutation.SetNickname(s)
	return mu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableNickname(s *string) *MemberUpdate {
	if s != nil {
		mu.SetNickname(*s)
	}
	return mu
}

// ClearNickname clears the value of the "nickname" field.
func (mu *MemberUpdate) ClearNickname() *MemberUpdate {
	mu.mutation.ClearNickname()
	return mu
}

// SetWxOpenID sets the "wx_open_id" field.
func (mu *MemberUpdate) SetWxOpenID(s string) *MemberUpdate {
	mu.mutation.SetWxOpenID(s)
	return mu
}

// SetNillableWxOpenID sets the "wx_open_id" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableWxOpenID(s *string) *MemberUpdate {
	if s != nil {
		mu.SetWxOpenID(*s)
	}
	return mu
}

// SetWxUnionID sets the "wx_union_id" field.
func (mu *MemberUpdate) SetWxUnionID(s string) *MemberUpdate {
	mu.mutation.SetWxUnionID(s)
	return mu
}

// SetNillableWxUnionID sets the "wx_union_id" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableWxUnionID(s *string) *MemberUpdate {
	if s != nil {
		mu.SetWxUnionID(*s)
	}
	return mu
}

// ClearWxUnionID clears the value of the "wx_union_id" field.
func (mu *MemberUpdate) ClearWxUnionID() *MemberUpdate {
	mu.mutation.ClearWxUnionID()
	return mu
}

// SetBcmbCode sets the "bcmb_code" field.
func (mu *MemberUpdate) SetBcmbCode(s string) *MemberUpdate {
	mu.mutation.SetBcmbCode(s)
	return mu
}

// SetNillableBcmbCode sets the "bcmb_code" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableBcmbCode(s *string) *MemberUpdate {
	if s != nil {
		mu.SetBcmbCode(*s)
	}
	return mu
}

// ClearBcmbCode clears the value of the "bcmb_code" field.
func (mu *MemberUpdate) ClearBcmbCode() *MemberUpdate {
	mu.mutation.ClearBcmbCode()
	return mu
}

// SetBcmbRegTime sets the "bcmb_reg_time" field.
func (mu *MemberUpdate) SetBcmbRegTime(t time.Time) *MemberUpdate {
	mu.mutation.SetBcmbRegTime(t)
	return mu
}

// SetNillableBcmbRegTime sets the "bcmb_reg_time" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableBcmbRegTime(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetBcmbRegTime(*t)
	}
	return mu
}

// ClearBcmbRegTime clears the value of the "bcmb_reg_time" field.
func (mu *MemberUpdate) ClearBcmbRegTime() *MemberUpdate {
	mu.mutation.ClearBcmbRegTime()
	return mu
}

// SetBcmbRegMsgID sets the "bcmb_reg_msg_id" field.
func (mu *MemberUpdate) SetBcmbRegMsgID(s string) *MemberUpdate {
	mu.mutation.SetBcmbRegMsgID(s)
	return mu
}

// SetNillableBcmbRegMsgID sets the "bcmb_reg_msg_id" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableBcmbRegMsgID(s *string) *MemberUpdate {
	if s != nil {
		mu.SetBcmbRegMsgID(*s)
	}
	return mu
}

// ClearBcmbRegMsgID clears the value of the "bcmb_reg_msg_id" field.
func (mu *MemberUpdate) ClearBcmbRegMsgID() *MemberUpdate {
	mu.mutation.ClearBcmbRegMsgID()
	return mu
}

// SetBcmbType sets the "bcmb_type" field.
func (mu *MemberUpdate) SetBcmbType(i int32) *MemberUpdate {
	mu.mutation.ResetBcmbType()
	mu.mutation.SetBcmbType(i)
	return mu
}

// SetNillableBcmbType sets the "bcmb_type" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableBcmbType(i *int32) *MemberUpdate {
	if i != nil {
		mu.SetBcmbType(*i)
	}
	return mu
}

// AddBcmbType adds i to the "bcmb_type" field.
func (mu *MemberUpdate) AddBcmbType(i int32) *MemberUpdate {
	mu.mutation.AddBcmbType(i)
	return mu
}

// SetLevel sets the "level" field.
func (mu *MemberUpdate) SetLevel(i int32) *MemberUpdate {
	mu.mutation.ResetLevel()
	mu.mutation.SetLevel(i)
	return mu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableLevel(i *int32) *MemberUpdate {
	if i != nil {
		mu.SetLevel(*i)
	}
	return mu
}

// AddLevel adds i to the "level" field.
func (mu *MemberUpdate) AddLevel(i int32) *MemberUpdate {
	mu.mutation.AddLevel(i)
	return mu
}

// SetLastTime sets the "last_time" field.
func (mu *MemberUpdate) SetLastTime(t time.Time) *MemberUpdate {
	mu.mutation.SetLastTime(t)
	return mu
}

// SetNillableLastTime sets the "last_time" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableLastTime(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetLastTime(*t)
	}
	return mu
}

// SetSource sets the "source" field.
func (mu *MemberUpdate) SetSource(i int32) *MemberUpdate {
	mu.mutation.ResetSource()
	mu.mutation.SetSource(i)
	return mu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableSource(i *int32) *MemberUpdate {
	if i != nil {
		mu.SetSource(*i)
	}
	return mu
}

// AddSource adds i to the "source" field.
func (mu *MemberUpdate) AddSource(i int32) *MemberUpdate {
	mu.mutation.AddSource(i)
	return mu
}

// SetStatus sets the "status" field.
func (mu *MemberUpdate) SetStatus(i int32) *MemberUpdate {
	mu.mutation.ResetStatus()
	mu.mutation.SetStatus(i)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableStatus(i *int32) *MemberUpdate {
	if i != nil {
		mu.SetStatus(*i)
	}
	return mu
}

// AddStatus adds i to the "status" field.
func (mu *MemberUpdate) AddStatus(i int32) *MemberUpdate {
	mu.mutation.AddStatus(i)
	return mu
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MemberUpdate) defaults() {
	if _, ok := mu.mutation.UpdateTime(); !ok {
		v := member.UpdateDefaultUpdateTime()
		mu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.CodeBin(); ok {
		if err := member.CodeBinValidator(v); err != nil {
			return &ValidationError{Name: "code_bin", err: fmt.Errorf(`ent: validator failed for field "Member.code_bin": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Phone(); ok {
		if err := member.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Member.phone": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Member.nickname": %w`, err)}
		}
	}
	if v, ok := mu.mutation.WxOpenID(); ok {
		if err := member.WxOpenIDValidator(v); err != nil {
			return &ValidationError{Name: "wx_open_id", err: fmt.Errorf(`ent: validator failed for field "Member.wx_open_id": %w`, err)}
		}
	}
	if v, ok := mu.mutation.WxUnionID(); ok {
		if err := member.WxUnionIDValidator(v); err != nil {
			return &ValidationError{Name: "wx_union_id", err: fmt.Errorf(`ent: validator failed for field "Member.wx_union_id": %w`, err)}
		}
	}
	if v, ok := mu.mutation.BcmbCode(); ok {
		if err := member.BcmbCodeValidator(v); err != nil {
			return &ValidationError{Name: "bcmb_code", err: fmt.Errorf(`ent: validator failed for field "Member.bcmb_code": %w`, err)}
		}
	}
	if v, ok := mu.mutation.BcmbRegMsgID(); ok {
		if err := member.BcmbRegMsgIDValidator(v); err != nil {
			return &ValidationError{Name: "bcmb_reg_msg_id", err: fmt.Errorf(`ent: validator failed for field "Member.bcmb_reg_msg_id": %w`, err)}
		}
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdateTime(); ok {
		_spec.SetField(member.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := mu.mutation.CodeBin(); ok {
		_spec.SetField(member.FieldCodeBin, field.TypeString, value)
	}
	if mu.mutation.CodeBinCleared() {
		_spec.ClearField(member.FieldCodeBin, field.TypeString)
	}
	if value, ok := mu.mutation.Phone(); ok {
		_spec.SetField(member.FieldPhone, field.TypeString, value)
	}
	if mu.mutation.PhoneCleared() {
		_spec.ClearField(member.FieldPhone, field.TypeString)
	}
	if value, ok := mu.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
	}
	if mu.mutation.NicknameCleared() {
		_spec.ClearField(member.FieldNickname, field.TypeString)
	}
	if value, ok := mu.mutation.WxOpenID(); ok {
		_spec.SetField(member.FieldWxOpenID, field.TypeString, value)
	}
	if value, ok := mu.mutation.WxUnionID(); ok {
		_spec.SetField(member.FieldWxUnionID, field.TypeString, value)
	}
	if mu.mutation.WxUnionIDCleared() {
		_spec.ClearField(member.FieldWxUnionID, field.TypeString)
	}
	if value, ok := mu.mutation.BcmbCode(); ok {
		_spec.SetField(member.FieldBcmbCode, field.TypeString, value)
	}
	if mu.mutation.BcmbCodeCleared() {
		_spec.ClearField(member.FieldBcmbCode, field.TypeString)
	}
	if value, ok := mu.mutation.BcmbRegTime(); ok {
		_spec.SetField(member.FieldBcmbRegTime, field.TypeTime, value)
	}
	if mu.mutation.BcmbRegTimeCleared() {
		_spec.ClearField(member.FieldBcmbRegTime, field.TypeTime)
	}
	if value, ok := mu.mutation.BcmbRegMsgID(); ok {
		_spec.SetField(member.FieldBcmbRegMsgID, field.TypeString, value)
	}
	if mu.mutation.BcmbRegMsgIDCleared() {
		_spec.ClearField(member.FieldBcmbRegMsgID, field.TypeString)
	}
	if value, ok := mu.mutation.BcmbType(); ok {
		_spec.SetField(member.FieldBcmbType, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedBcmbType(); ok {
		_spec.AddField(member.FieldBcmbType, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.Level(); ok {
		_spec.SetField(member.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedLevel(); ok {
		_spec.AddField(member.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.LastTime(); ok {
		_spec.SetField(member.FieldLastTime, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Source(); ok {
		_spec.SetField(member.FieldSource, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedSource(); ok {
		_spec.AddField(member.FieldSource, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(member.FieldStatus, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedStatus(); ok {
		_spec.AddField(member.FieldStatus, field.TypeInt32, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetUpdateTime sets the "update_time" field.
func (muo *MemberUpdateOne) SetUpdateTime(t time.Time) *MemberUpdateOne {
	muo.mutation.SetUpdateTime(t)
	return muo
}

// SetCodeBin sets the "code_bin" field.
func (muo *MemberUpdateOne) SetCodeBin(s string) *MemberUpdateOne {
	muo.mutation.SetCodeBin(s)
	return muo
}

// SetNillableCodeBin sets the "code_bin" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableCodeBin(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetCodeBin(*s)
	}
	return muo
}

// ClearCodeBin clears the value of the "code_bin" field.
func (muo *MemberUpdateOne) ClearCodeBin() *MemberUpdateOne {
	muo.mutation.ClearCodeBin()
	return muo
}

// SetPhone sets the "phone" field.
func (muo *MemberUpdateOne) SetPhone(s string) *MemberUpdateOne {
	muo.mutation.SetPhone(s)
	return muo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillablePhone(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetPhone(*s)
	}
	return muo
}

// ClearPhone clears the value of the "phone" field.
func (muo *MemberUpdateOne) ClearPhone() *MemberUpdateOne {
	muo.mutation.ClearPhone()
	return muo
}

// SetNickname sets the "nickname" field.
func (muo *MemberUpdateOne) SetNickname(s string) *MemberUpdateOne {
	muo.mutation.SetNickname(s)
	return muo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableNickname(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetNickname(*s)
	}
	return muo
}

// ClearNickname clears the value of the "nickname" field.
func (muo *MemberUpdateOne) ClearNickname() *MemberUpdateOne {
	muo.mutation.ClearNickname()
	return muo
}

// SetWxOpenID sets the "wx_open_id" field.
func (muo *MemberUpdateOne) SetWxOpenID(s string) *MemberUpdateOne {
	muo.mutation.SetWxOpenID(s)
	return muo
}

// SetNillableWxOpenID sets the "wx_open_id" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableWxOpenID(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetWxOpenID(*s)
	}
	return muo
}

// SetWxUnionID sets the "wx_union_id" field.
func (muo *MemberUpdateOne) SetWxUnionID(s string) *MemberUpdateOne {
	muo.mutation.SetWxUnionID(s)
	return muo
}

// SetNillableWxUnionID sets the "wx_union_id" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableWxUnionID(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetWxUnionID(*s)
	}
	return muo
}

// ClearWxUnionID clears the value of the "wx_union_id" field.
func (muo *MemberUpdateOne) ClearWxUnionID() *MemberUpdateOne {
	muo.mutation.ClearWxUnionID()
	return muo
}

// SetBcmbCode sets the "bcmb_code" field.
func (muo *MemberUpdateOne) SetBcmbCode(s string) *MemberUpdateOne {
	muo.mutation.SetBcmbCode(s)
	return muo
}

// SetNillableBcmbCode sets the "bcmb_code" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableBcmbCode(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetBcmbCode(*s)
	}
	return muo
}

// ClearBcmbCode clears the value of the "bcmb_code" field.
func (muo *MemberUpdateOne) ClearBcmbCode() *MemberUpdateOne {
	muo.mutation.ClearBcmbCode()
	return muo
}

// SetBcmbRegTime sets the "bcmb_reg_time" field.
func (muo *MemberUpdateOne) SetBcmbRegTime(t time.Time) *MemberUpdateOne {
	muo.mutation.SetBcmbRegTime(t)
	return muo
}

// SetNillableBcmbRegTime sets the "bcmb_reg_time" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableBcmbRegTime(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetBcmbRegTime(*t)
	}
	return muo
}

// ClearBcmbRegTime clears the value of the "bcmb_reg_time" field.
func (muo *MemberUpdateOne) ClearBcmbRegTime() *MemberUpdateOne {
	muo.mutation.ClearBcmbRegTime()
	return muo
}

// SetBcmbRegMsgID sets the "bcmb_reg_msg_id" field.
func (muo *MemberUpdateOne) SetBcmbRegMsgID(s string) *MemberUpdateOne {
	muo.mutation.SetBcmbRegMsgID(s)
	return muo
}

// SetNillableBcmbRegMsgID sets the "bcmb_reg_msg_id" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableBcmbRegMsgID(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetBcmbRegMsgID(*s)
	}
	return muo
}

// ClearBcmbRegMsgID clears the value of the "bcmb_reg_msg_id" field.
func (muo *MemberUpdateOne) ClearBcmbRegMsgID() *MemberUpdateOne {
	muo.mutation.ClearBcmbRegMsgID()
	return muo
}

// SetBcmbType sets the "bcmb_type" field.
func (muo *MemberUpdateOne) SetBcmbType(i int32) *MemberUpdateOne {
	muo.mutation.ResetBcmbType()
	muo.mutation.SetBcmbType(i)
	return muo
}

// SetNillableBcmbType sets the "bcmb_type" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableBcmbType(i *int32) *MemberUpdateOne {
	if i != nil {
		muo.SetBcmbType(*i)
	}
	return muo
}

// AddBcmbType adds i to the "bcmb_type" field.
func (muo *MemberUpdateOne) AddBcmbType(i int32) *MemberUpdateOne {
	muo.mutation.AddBcmbType(i)
	return muo
}

// SetLevel sets the "level" field.
func (muo *MemberUpdateOne) SetLevel(i int32) *MemberUpdateOne {
	muo.mutation.ResetLevel()
	muo.mutation.SetLevel(i)
	return muo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableLevel(i *int32) *MemberUpdateOne {
	if i != nil {
		muo.SetLevel(*i)
	}
	return muo
}

// AddLevel adds i to the "level" field.
func (muo *MemberUpdateOne) AddLevel(i int32) *MemberUpdateOne {
	muo.mutation.AddLevel(i)
	return muo
}

// SetLastTime sets the "last_time" field.
func (muo *MemberUpdateOne) SetLastTime(t time.Time) *MemberUpdateOne {
	muo.mutation.SetLastTime(t)
	return muo
}

// SetNillableLastTime sets the "last_time" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableLastTime(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetLastTime(*t)
	}
	return muo
}

// SetSource sets the "source" field.
func (muo *MemberUpdateOne) SetSource(i int32) *MemberUpdateOne {
	muo.mutation.ResetSource()
	muo.mutation.SetSource(i)
	return muo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableSource(i *int32) *MemberUpdateOne {
	if i != nil {
		muo.SetSource(*i)
	}
	return muo
}

// AddSource adds i to the "source" field.
func (muo *MemberUpdateOne) AddSource(i int32) *MemberUpdateOne {
	muo.mutation.AddSource(i)
	return muo
}

// SetStatus sets the "status" field.
func (muo *MemberUpdateOne) SetStatus(i int32) *MemberUpdateOne {
	muo.mutation.ResetStatus()
	muo.mutation.SetStatus(i)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableStatus(i *int32) *MemberUpdateOne {
	if i != nil {
		muo.SetStatus(*i)
	}
	return muo
}

// AddStatus adds i to the "status" field.
func (muo *MemberUpdateOne) AddStatus(i int32) *MemberUpdateOne {
	muo.mutation.AddStatus(i)
	return muo
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (muo *MemberUpdateOne) Where(ps ...predicate.Member) *MemberUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MemberUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdateTime(); !ok {
		v := member.UpdateDefaultUpdateTime()
		muo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.CodeBin(); ok {
		if err := member.CodeBinValidator(v); err != nil {
			return &ValidationError{Name: "code_bin", err: fmt.Errorf(`ent: validator failed for field "Member.code_bin": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Phone(); ok {
		if err := member.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Member.phone": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Member.nickname": %w`, err)}
		}
	}
	if v, ok := muo.mutation.WxOpenID(); ok {
		if err := member.WxOpenIDValidator(v); err != nil {
			return &ValidationError{Name: "wx_open_id", err: fmt.Errorf(`ent: validator failed for field "Member.wx_open_id": %w`, err)}
		}
	}
	if v, ok := muo.mutation.WxUnionID(); ok {
		if err := member.WxUnionIDValidator(v); err != nil {
			return &ValidationError{Name: "wx_union_id", err: fmt.Errorf(`ent: validator failed for field "Member.wx_union_id": %w`, err)}
		}
	}
	if v, ok := muo.mutation.BcmbCode(); ok {
		if err := member.BcmbCodeValidator(v); err != nil {
			return &ValidationError{Name: "bcmb_code", err: fmt.Errorf(`ent: validator failed for field "Member.bcmb_code": %w`, err)}
		}
	}
	if v, ok := muo.mutation.BcmbRegMsgID(); ok {
		if err := member.BcmbRegMsgIDValidator(v); err != nil {
			return &ValidationError{Name: "bcmb_reg_msg_id", err: fmt.Errorf(`ent: validator failed for field "Member.bcmb_reg_msg_id": %w`, err)}
		}
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
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
		_spec.SetField(member.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := muo.mutation.CodeBin(); ok {
		_spec.SetField(member.FieldCodeBin, field.TypeString, value)
	}
	if muo.mutation.CodeBinCleared() {
		_spec.ClearField(member.FieldCodeBin, field.TypeString)
	}
	if value, ok := muo.mutation.Phone(); ok {
		_spec.SetField(member.FieldPhone, field.TypeString, value)
	}
	if muo.mutation.PhoneCleared() {
		_spec.ClearField(member.FieldPhone, field.TypeString)
	}
	if value, ok := muo.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
	}
	if muo.mutation.NicknameCleared() {
		_spec.ClearField(member.FieldNickname, field.TypeString)
	}
	if value, ok := muo.mutation.WxOpenID(); ok {
		_spec.SetField(member.FieldWxOpenID, field.TypeString, value)
	}
	if value, ok := muo.mutation.WxUnionID(); ok {
		_spec.SetField(member.FieldWxUnionID, field.TypeString, value)
	}
	if muo.mutation.WxUnionIDCleared() {
		_spec.ClearField(member.FieldWxUnionID, field.TypeString)
	}
	if value, ok := muo.mutation.BcmbCode(); ok {
		_spec.SetField(member.FieldBcmbCode, field.TypeString, value)
	}
	if muo.mutation.BcmbCodeCleared() {
		_spec.ClearField(member.FieldBcmbCode, field.TypeString)
	}
	if value, ok := muo.mutation.BcmbRegTime(); ok {
		_spec.SetField(member.FieldBcmbRegTime, field.TypeTime, value)
	}
	if muo.mutation.BcmbRegTimeCleared() {
		_spec.ClearField(member.FieldBcmbRegTime, field.TypeTime)
	}
	if value, ok := muo.mutation.BcmbRegMsgID(); ok {
		_spec.SetField(member.FieldBcmbRegMsgID, field.TypeString, value)
	}
	if muo.mutation.BcmbRegMsgIDCleared() {
		_spec.ClearField(member.FieldBcmbRegMsgID, field.TypeString)
	}
	if value, ok := muo.mutation.BcmbType(); ok {
		_spec.SetField(member.FieldBcmbType, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedBcmbType(); ok {
		_spec.AddField(member.FieldBcmbType, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.Level(); ok {
		_spec.SetField(member.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedLevel(); ok {
		_spec.AddField(member.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.LastTime(); ok {
		_spec.SetField(member.FieldLastTime, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Source(); ok {
		_spec.SetField(member.FieldSource, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedSource(); ok {
		_spec.AddField(member.FieldSource, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(member.FieldStatus, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedStatus(); ok {
		_spec.AddField(member.FieldStatus, field.TypeInt32, value)
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
