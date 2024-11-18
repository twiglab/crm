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
	"github.com/twiglab/crm/member/orm/ent/member"
)

// MemberCreate is the builder for creating a Member entity.
type MemberCreate struct {
	config
	mutation *MemberMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (mc *MemberCreate) SetCreateTime(t time.Time) *MemberCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mc *MemberCreate) SetNillableCreateTime(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// SetUpdateTime sets the "update_time" field.
func (mc *MemberCreate) SetUpdateTime(t time.Time) *MemberCreate {
	mc.mutation.SetUpdateTime(t)
	return mc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mc *MemberCreate) SetNillableUpdateTime(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetUpdateTime(*t)
	}
	return mc
}

// SetCode sets the "code" field.
func (mc *MemberCreate) SetCode(s string) *MemberCreate {
	mc.mutation.SetCode(s)
	return mc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (mc *MemberCreate) SetNillableCode(s *string) *MemberCreate {
	if s != nil {
		mc.SetCode(*s)
	}
	return mc
}

// SetPhone sets the "phone" field.
func (mc *MemberCreate) SetPhone(s string) *MemberCreate {
	mc.mutation.SetPhone(s)
	return mc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (mc *MemberCreate) SetNillablePhone(s *string) *MemberCreate {
	if s != nil {
		mc.SetPhone(*s)
	}
	return mc
}

// SetNickname sets the "nickname" field.
func (mc *MemberCreate) SetNickname(s string) *MemberCreate {
	mc.mutation.SetNickname(s)
	return mc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mc *MemberCreate) SetNillableNickname(s *string) *MemberCreate {
	if s != nil {
		mc.SetNickname(*s)
	}
	return mc
}

// SetWxOpenID sets the "wx_open_id" field.
func (mc *MemberCreate) SetWxOpenID(s string) *MemberCreate {
	mc.mutation.SetWxOpenID(s)
	return mc
}

// SetWxUID sets the "wx_uid" field.
func (mc *MemberCreate) SetWxUID(s string) *MemberCreate {
	mc.mutation.SetWxUID(s)
	return mc
}

// SetWxMBCode sets the "wx_mb_code" field.
func (mc *MemberCreate) SetWxMBCode(s string) *MemberCreate {
	mc.mutation.SetWxMBCode(s)
	return mc
}

// SetNillableWxMBCode sets the "wx_mb_code" field if the given value is not nil.
func (mc *MemberCreate) SetNillableWxMBCode(s *string) *MemberCreate {
	if s != nil {
		mc.SetWxMBCode(*s)
	}
	return mc
}

// SetWxRegTime sets the "wx_reg_time" field.
func (mc *MemberCreate) SetWxRegTime(t time.Time) *MemberCreate {
	mc.mutation.SetWxRegTime(t)
	return mc
}

// SetNillableWxRegTime sets the "wx_reg_time" field if the given value is not nil.
func (mc *MemberCreate) SetNillableWxRegTime(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetWxRegTime(*t)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MemberCreate) SetStatus(i int) *MemberCreate {
	mc.mutation.SetStatus(i)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *MemberCreate) SetNillableStatus(i *int) *MemberCreate {
	if i != nil {
		mc.SetStatus(*i)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MemberCreate) SetID(u uuid.UUID) *MemberCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MemberCreate) SetNillableID(u *uuid.UUID) *MemberCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// Mutation returns the MemberMutation object of the builder.
func (mc *MemberCreate) Mutation() *MemberMutation {
	return mc.mutation
}

// Save creates the Member in the database.
func (mc *MemberCreate) Save(ctx context.Context) (*Member, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MemberCreate) SaveX(ctx context.Context) *Member {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MemberCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MemberCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MemberCreate) defaults() {
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := member.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		v := member.DefaultUpdateTime()
		mc.mutation.SetUpdateTime(v)
	}
	if _, ok := mc.mutation.Code(); !ok {
		v := member.DefaultCode()
		mc.mutation.SetCode(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := member.DefaultStatus
		mc.mutation.SetStatus(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := member.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MemberCreate) check() error {
	if _, ok := mc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Member.create_time"`)}
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Member.update_time"`)}
	}
	if _, ok := mc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Member.code"`)}
	}
	if v, ok := mc.mutation.Code(); ok {
		if err := member.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Member.code": %w`, err)}
		}
	}
	if v, ok := mc.mutation.Phone(); ok {
		if err := member.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Member.phone": %w`, err)}
		}
	}
	if v, ok := mc.mutation.Nickname(); ok {
		if err := member.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Member.nickname": %w`, err)}
		}
	}
	if _, ok := mc.mutation.WxOpenID(); !ok {
		return &ValidationError{Name: "wx_open_id", err: errors.New(`ent: missing required field "Member.wx_open_id"`)}
	}
	if v, ok := mc.mutation.WxOpenID(); ok {
		if err := member.WxOpenIDValidator(v); err != nil {
			return &ValidationError{Name: "wx_open_id", err: fmt.Errorf(`ent: validator failed for field "Member.wx_open_id": %w`, err)}
		}
	}
	if _, ok := mc.mutation.WxUID(); !ok {
		return &ValidationError{Name: "wx_uid", err: errors.New(`ent: missing required field "Member.wx_uid"`)}
	}
	if v, ok := mc.mutation.WxUID(); ok {
		if err := member.WxUIDValidator(v); err != nil {
			return &ValidationError{Name: "wx_uid", err: fmt.Errorf(`ent: validator failed for field "Member.wx_uid": %w`, err)}
		}
	}
	if v, ok := mc.mutation.WxMBCode(); ok {
		if err := member.WxMBCodeValidator(v); err != nil {
			return &ValidationError{Name: "wx_mb_code", err: fmt.Errorf(`ent: validator failed for field "Member.wx_mb_code": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Member.status"`)}
	}
	return nil
}

func (mc *MemberCreate) sqlSave(ctx context.Context) (*Member, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MemberCreate) createSpec() (*Member, *sqlgraph.CreateSpec) {
	var (
		_node = &Member{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(member.Table, sqlgraph.NewFieldSpec(member.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.SetField(member.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := mc.mutation.UpdateTime(); ok {
		_spec.SetField(member.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := mc.mutation.Code(); ok {
		_spec.SetField(member.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := mc.mutation.Phone(); ok {
		_spec.SetField(member.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := mc.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := mc.mutation.WxOpenID(); ok {
		_spec.SetField(member.FieldWxOpenID, field.TypeString, value)
		_node.WxOpenID = value
	}
	if value, ok := mc.mutation.WxUID(); ok {
		_spec.SetField(member.FieldWxUID, field.TypeString, value)
		_node.WxUID = value
	}
	if value, ok := mc.mutation.WxMBCode(); ok {
		_spec.SetField(member.FieldWxMBCode, field.TypeString, value)
		_node.WxMBCode = value
	}
	if value, ok := mc.mutation.WxRegTime(); ok {
		_spec.SetField(member.FieldWxRegTime, field.TypeTime, value)
		_node.WxRegTime = &value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(member.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Member.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemberUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (mc *MemberCreate) OnConflict(opts ...sql.ConflictOption) *MemberUpsertOne {
	mc.conflict = opts
	return &MemberUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MemberCreate) OnConflictColumns(columns ...string) *MemberUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MemberUpsertOne{
		create: mc,
	}
}

type (
	// MemberUpsertOne is the builder for "upsert"-ing
	//  one Member node.
	MemberUpsertOne struct {
		create *MemberCreate
	}

	// MemberUpsert is the "OnConflict" setter.
	MemberUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *MemberUpsert) SetUpdateTime(v time.Time) *MemberUpsert {
	u.Set(member.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MemberUpsert) UpdateUpdateTime() *MemberUpsert {
	u.SetExcluded(member.FieldUpdateTime)
	return u
}

// SetPhone sets the "phone" field.
func (u *MemberUpsert) SetPhone(v string) *MemberUpsert {
	u.Set(member.FieldPhone, v)
	return u
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *MemberUpsert) UpdatePhone() *MemberUpsert {
	u.SetExcluded(member.FieldPhone)
	return u
}

// ClearPhone clears the value of the "phone" field.
func (u *MemberUpsert) ClearPhone() *MemberUpsert {
	u.SetNull(member.FieldPhone)
	return u
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsert) SetNickname(v string) *MemberUpsert {
	u.Set(member.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsert) UpdateNickname() *MemberUpsert {
	u.SetExcluded(member.FieldNickname)
	return u
}

// ClearNickname clears the value of the "nickname" field.
func (u *MemberUpsert) ClearNickname() *MemberUpsert {
	u.SetNull(member.FieldNickname)
	return u
}

// SetWxOpenID sets the "wx_open_id" field.
func (u *MemberUpsert) SetWxOpenID(v string) *MemberUpsert {
	u.Set(member.FieldWxOpenID, v)
	return u
}

// UpdateWxOpenID sets the "wx_open_id" field to the value that was provided on create.
func (u *MemberUpsert) UpdateWxOpenID() *MemberUpsert {
	u.SetExcluded(member.FieldWxOpenID)
	return u
}

// SetWxUID sets the "wx_uid" field.
func (u *MemberUpsert) SetWxUID(v string) *MemberUpsert {
	u.Set(member.FieldWxUID, v)
	return u
}

// UpdateWxUID sets the "wx_uid" field to the value that was provided on create.
func (u *MemberUpsert) UpdateWxUID() *MemberUpsert {
	u.SetExcluded(member.FieldWxUID)
	return u
}

// SetWxMBCode sets the "wx_mb_code" field.
func (u *MemberUpsert) SetWxMBCode(v string) *MemberUpsert {
	u.Set(member.FieldWxMBCode, v)
	return u
}

// UpdateWxMBCode sets the "wx_mb_code" field to the value that was provided on create.
func (u *MemberUpsert) UpdateWxMBCode() *MemberUpsert {
	u.SetExcluded(member.FieldWxMBCode)
	return u
}

// ClearWxMBCode clears the value of the "wx_mb_code" field.
func (u *MemberUpsert) ClearWxMBCode() *MemberUpsert {
	u.SetNull(member.FieldWxMBCode)
	return u
}

// SetStatus sets the "status" field.
func (u *MemberUpsert) SetStatus(v int) *MemberUpsert {
	u.Set(member.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MemberUpsert) UpdateStatus() *MemberUpsert {
	u.SetExcluded(member.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *MemberUpsert) AddStatus(v int) *MemberUpsert {
	u.Add(member.FieldStatus, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(member.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MemberUpsertOne) UpdateNewValues() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(member.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(member.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Code(); exists {
			s.SetIgnore(member.FieldCode)
		}
		if _, exists := u.create.mutation.WxRegTime(); exists {
			s.SetIgnore(member.FieldWxRegTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MemberUpsertOne) Ignore() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemberUpsertOne) DoNothing() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemberCreate.OnConflict
// documentation for more info.
func (u *MemberUpsertOne) Update(set func(*MemberUpsert)) *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *MemberUpsertOne) SetUpdateTime(v time.Time) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateUpdateTime() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetPhone sets the "phone" field.
func (u *MemberUpsertOne) SetPhone(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdatePhone() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdatePhone()
	})
}

// ClearPhone clears the value of the "phone" field.
func (u *MemberUpsertOne) ClearPhone() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearPhone()
	})
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsertOne) SetNickname(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateNickname() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *MemberUpsertOne) ClearNickname() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearNickname()
	})
}

// SetWxOpenID sets the "wx_open_id" field.
func (u *MemberUpsertOne) SetWxOpenID(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetWxOpenID(v)
	})
}

// UpdateWxOpenID sets the "wx_open_id" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateWxOpenID() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateWxOpenID()
	})
}

// SetWxUID sets the "wx_uid" field.
func (u *MemberUpsertOne) SetWxUID(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetWxUID(v)
	})
}

// UpdateWxUID sets the "wx_uid" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateWxUID() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateWxUID()
	})
}

// SetWxMBCode sets the "wx_mb_code" field.
func (u *MemberUpsertOne) SetWxMBCode(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetWxMBCode(v)
	})
}

// UpdateWxMBCode sets the "wx_mb_code" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateWxMBCode() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateWxMBCode()
	})
}

// ClearWxMBCode clears the value of the "wx_mb_code" field.
func (u *MemberUpsertOne) ClearWxMBCode() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.ClearWxMBCode()
	})
}

// SetStatus sets the "status" field.
func (u *MemberUpsertOne) SetStatus(v int) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *MemberUpsertOne) AddStatus(v int) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateStatus() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *MemberUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemberCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemberUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MemberUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MemberUpsertOne.ID is not supported by MySQL driver. Use MemberUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MemberUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MemberCreateBulk is the builder for creating many Member entities in bulk.
type MemberCreateBulk struct {
	config
	err      error
	builders []*MemberCreate
	conflict []sql.ConflictOption
}

// Save creates the Member entities in the database.
func (mcb *MemberCreateBulk) Save(ctx context.Context) ([]*Member, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Member, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MemberCreateBulk) SaveX(ctx context.Context) []*Member {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MemberCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MemberCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Member.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemberUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (mcb *MemberCreateBulk) OnConflict(opts ...sql.ConflictOption) *MemberUpsertBulk {
	mcb.conflict = opts
	return &MemberUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MemberCreateBulk) OnConflictColumns(columns ...string) *MemberUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MemberUpsertBulk{
		create: mcb,
	}
}

// MemberUpsertBulk is the builder for "upsert"-ing
// a bulk of Member nodes.
type MemberUpsertBulk struct {
	create *MemberCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(member.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MemberUpsertBulk) UpdateNewValues() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(member.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(member.FieldCreateTime)
			}
			if _, exists := b.mutation.Code(); exists {
				s.SetIgnore(member.FieldCode)
			}
			if _, exists := b.mutation.WxRegTime(); exists {
				s.SetIgnore(member.FieldWxRegTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MemberUpsertBulk) Ignore() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemberUpsertBulk) DoNothing() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemberCreateBulk.OnConflict
// documentation for more info.
func (u *MemberUpsertBulk) Update(set func(*MemberUpsert)) *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *MemberUpsertBulk) SetUpdateTime(v time.Time) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateUpdateTime() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetPhone sets the "phone" field.
func (u *MemberUpsertBulk) SetPhone(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdatePhone() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdatePhone()
	})
}

// ClearPhone clears the value of the "phone" field.
func (u *MemberUpsertBulk) ClearPhone() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearPhone()
	})
}

// SetNickname sets the "nickname" field.
func (u *MemberUpsertBulk) SetNickname(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateNickname() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateNickname()
	})
}

// ClearNickname clears the value of the "nickname" field.
func (u *MemberUpsertBulk) ClearNickname() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearNickname()
	})
}

// SetWxOpenID sets the "wx_open_id" field.
func (u *MemberUpsertBulk) SetWxOpenID(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetWxOpenID(v)
	})
}

// UpdateWxOpenID sets the "wx_open_id" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateWxOpenID() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateWxOpenID()
	})
}

// SetWxUID sets the "wx_uid" field.
func (u *MemberUpsertBulk) SetWxUID(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetWxUID(v)
	})
}

// UpdateWxUID sets the "wx_uid" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateWxUID() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateWxUID()
	})
}

// SetWxMBCode sets the "wx_mb_code" field.
func (u *MemberUpsertBulk) SetWxMBCode(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetWxMBCode(v)
	})
}

// UpdateWxMBCode sets the "wx_mb_code" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateWxMBCode() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateWxMBCode()
	})
}

// ClearWxMBCode clears the value of the "wx_mb_code" field.
func (u *MemberUpsertBulk) ClearWxMBCode() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.ClearWxMBCode()
	})
}

// SetStatus sets the "status" field.
func (u *MemberUpsertBulk) SetStatus(v int) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *MemberUpsertBulk) AddStatus(v int) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateStatus() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *MemberUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MemberCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemberCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemberUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
