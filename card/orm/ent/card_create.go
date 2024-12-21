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
	"github.com/twiglab/crm/card/orm/ent/card"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation *CardMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (cc *CardCreate) SetCreateTime(t time.Time) *CardCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CardCreate) SetNillableCreateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CardCreate) SetUpdateTime(t time.Time) *CardCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetCode sets the "code" field.
func (cc *CardCreate) SetCode(s string) *CardCreate {
	cc.mutation.SetCode(s)
	return cc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (cc *CardCreate) SetNillableCode(s *string) *CardCreate {
	if s != nil {
		cc.SetCode(*s)
	}
	return cc
}

// SetCardCode sets the "card_code" field.
func (cc *CardCreate) SetCardCode(s string) *CardCreate {
	cc.mutation.SetCardCode(s)
	return cc
}

// SetMemberCode sets the "member_code" field.
func (cc *CardCreate) SetMemberCode(s string) *CardCreate {
	cc.mutation.SetMemberCode(s)
	return cc
}

// SetNillableMemberCode sets the "member_code" field if the given value is not nil.
func (cc *CardCreate) SetNillableMemberCode(s *string) *CardCreate {
	if s != nil {
		cc.SetMemberCode(*s)
	}
	return cc
}

// SetType sets the "type" field.
func (cc *CardCreate) SetType(i int) *CardCreate {
	cc.mutation.SetType(i)
	return cc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cc *CardCreate) SetNillableType(i *int) *CardCreate {
	if i != nil {
		cc.SetType(*i)
	}
	return cc
}

// SetBalance sets the "balance" field.
func (cc *CardCreate) SetBalance(i int) *CardCreate {
	cc.mutation.SetBalance(i)
	return cc
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (cc *CardCreate) SetNillableBalance(i *int) *CardCreate {
	if i != nil {
		cc.SetBalance(*i)
	}
	return cc
}

// SetAmount sets the "amount" field.
func (cc *CardCreate) SetAmount(i int) *CardCreate {
	cc.mutation.SetAmount(i)
	return cc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (cc *CardCreate) SetNillableAmount(i *int) *CardCreate {
	if i != nil {
		cc.SetAmount(*i)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CardCreate) SetStatus(i int) *CardCreate {
	cc.mutation.SetStatus(i)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CardCreate) SetNillableStatus(i *int) *CardCreate {
	if i != nil {
		cc.SetStatus(*i)
	}
	return cc
}

// Mutation returns the CardMutation object of the builder.
func (cc *CardCreate) Mutation() *CardMutation {
	return cc.mutation
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CardCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CardCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CardCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := card.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := card.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.Code(); !ok {
		v := card.DefaultCode()
		cc.mutation.SetCode(v)
	}
	if _, ok := cc.mutation.MemberCode(); !ok {
		v := card.DefaultMemberCode()
		cc.mutation.SetMemberCode(v)
	}
	if _, ok := cc.mutation.GetType(); !ok {
		v := card.DefaultType
		cc.mutation.SetType(v)
	}
	if _, ok := cc.mutation.Balance(); !ok {
		v := card.DefaultBalance
		cc.mutation.SetBalance(v)
	}
	if _, ok := cc.mutation.Amount(); !ok {
		v := card.DefaultAmount
		cc.mutation.SetAmount(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := card.DefaultStatus
		cc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CardCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Card.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Card.update_time"`)}
	}
	if _, ok := cc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Card.code"`)}
	}
	if v, ok := cc.mutation.Code(); ok {
		if err := card.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Card.code": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CardCode(); !ok {
		return &ValidationError{Name: "card_code", err: errors.New(`ent: missing required field "Card.card_code"`)}
	}
	if v, ok := cc.mutation.CardCode(); ok {
		if err := card.CardCodeValidator(v); err != nil {
			return &ValidationError{Name: "card_code", err: fmt.Errorf(`ent: validator failed for field "Card.card_code": %w`, err)}
		}
	}
	if _, ok := cc.mutation.MemberCode(); !ok {
		return &ValidationError{Name: "member_code", err: errors.New(`ent: missing required field "Card.member_code"`)}
	}
	if v, ok := cc.mutation.MemberCode(); ok {
		if err := card.MemberCodeValidator(v); err != nil {
			return &ValidationError{Name: "member_code", err: fmt.Errorf(`ent: validator failed for field "Card.member_code": %w`, err)}
		}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Card.type"`)}
	}
	if _, ok := cc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "Card.balance"`)}
	}
	if _, ok := cc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Card.amount"`)}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Card.status"`)}
	}
	return nil
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CardCreate) createSpec() (*Card, *sqlgraph.CreateSpec) {
	var (
		_node = &Card{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(card.Table, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.SetField(card.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.SetField(card.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Code(); ok {
		_spec.SetField(card.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := cc.mutation.CardCode(); ok {
		_spec.SetField(card.FieldCardCode, field.TypeString, value)
		_node.CardCode = value
	}
	if value, ok := cc.mutation.MemberCode(); ok {
		_spec.SetField(card.FieldMemberCode, field.TypeString, value)
		_node.MemberCode = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.SetField(card.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := cc.mutation.Balance(); ok {
		_spec.SetField(card.FieldBalance, field.TypeInt, value)
		_node.Balance = value
	}
	if value, ok := cc.mutation.Amount(); ok {
		_spec.SetField(card.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(card.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Card.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CardUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (cc *CardCreate) OnConflict(opts ...sql.ConflictOption) *CardUpsertOne {
	cc.conflict = opts
	return &CardUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CardCreate) OnConflictColumns(columns ...string) *CardUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CardUpsertOne{
		create: cc,
	}
}

type (
	// CardUpsertOne is the builder for "upsert"-ing
	//  one Card node.
	CardUpsertOne struct {
		create *CardCreate
	}

	// CardUpsert is the "OnConflict" setter.
	CardUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *CardUpsert) SetUpdateTime(v time.Time) *CardUpsert {
	u.Set(card.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CardUpsert) UpdateUpdateTime() *CardUpsert {
	u.SetExcluded(card.FieldUpdateTime)
	return u
}

// SetMemberCode sets the "member_code" field.
func (u *CardUpsert) SetMemberCode(v string) *CardUpsert {
	u.Set(card.FieldMemberCode, v)
	return u
}

// UpdateMemberCode sets the "member_code" field to the value that was provided on create.
func (u *CardUpsert) UpdateMemberCode() *CardUpsert {
	u.SetExcluded(card.FieldMemberCode)
	return u
}

// SetType sets the "type" field.
func (u *CardUpsert) SetType(v int) *CardUpsert {
	u.Set(card.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *CardUpsert) UpdateType() *CardUpsert {
	u.SetExcluded(card.FieldType)
	return u
}

// AddType adds v to the "type" field.
func (u *CardUpsert) AddType(v int) *CardUpsert {
	u.Add(card.FieldType, v)
	return u
}

// SetBalance sets the "balance" field.
func (u *CardUpsert) SetBalance(v int) *CardUpsert {
	u.Set(card.FieldBalance, v)
	return u
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CardUpsert) UpdateBalance() *CardUpsert {
	u.SetExcluded(card.FieldBalance)
	return u
}

// AddBalance adds v to the "balance" field.
func (u *CardUpsert) AddBalance(v int) *CardUpsert {
	u.Add(card.FieldBalance, v)
	return u
}

// SetAmount sets the "amount" field.
func (u *CardUpsert) SetAmount(v int) *CardUpsert {
	u.Set(card.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *CardUpsert) UpdateAmount() *CardUpsert {
	u.SetExcluded(card.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *CardUpsert) AddAmount(v int) *CardUpsert {
	u.Add(card.FieldAmount, v)
	return u
}

// SetStatus sets the "status" field.
func (u *CardUpsert) SetStatus(v int) *CardUpsert {
	u.Set(card.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CardUpsert) UpdateStatus() *CardUpsert {
	u.SetExcluded(card.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *CardUpsert) AddStatus(v int) *CardUpsert {
	u.Add(card.FieldStatus, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CardUpsertOne) UpdateNewValues() *CardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(card.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Code(); exists {
			s.SetIgnore(card.FieldCode)
		}
		if _, exists := u.create.mutation.CardCode(); exists {
			s.SetIgnore(card.FieldCardCode)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Card.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CardUpsertOne) Ignore() *CardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CardUpsertOne) DoNothing() *CardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CardCreate.OnConflict
// documentation for more info.
func (u *CardUpsertOne) Update(set func(*CardUpsert)) *CardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CardUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CardUpsertOne) SetUpdateTime(v time.Time) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateUpdateTime() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetMemberCode sets the "member_code" field.
func (u *CardUpsertOne) SetMemberCode(v string) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetMemberCode(v)
	})
}

// UpdateMemberCode sets the "member_code" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateMemberCode() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateMemberCode()
	})
}

// SetType sets the "type" field.
func (u *CardUpsertOne) SetType(v int) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *CardUpsertOne) AddType(v int) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateType() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateType()
	})
}

// SetBalance sets the "balance" field.
func (u *CardUpsertOne) SetBalance(v int) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetBalance(v)
	})
}

// AddBalance adds v to the "balance" field.
func (u *CardUpsertOne) AddBalance(v int) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.AddBalance(v)
	})
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateBalance() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateBalance()
	})
}

// SetAmount sets the "amount" field.
func (u *CardUpsertOne) SetAmount(v int) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *CardUpsertOne) AddAmount(v int) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateAmount() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateAmount()
	})
}

// SetStatus sets the "status" field.
func (u *CardUpsertOne) SetStatus(v int) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *CardUpsertOne) AddStatus(v int) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateStatus() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *CardUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CardCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CardUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CardUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CardCreateBulk is the builder for creating many Card entities in bulk.
type CardCreateBulk struct {
	config
	err      error
	builders []*CardCreate
	conflict []sql.ConflictOption
}

// Save creates the Card entities in the database.
func (ccb *CardCreateBulk) Save(ctx context.Context) ([]*Card, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Card, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CardCreateBulk) SaveX(ctx context.Context) []*Card {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CardCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CardCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Card.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CardUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (ccb *CardCreateBulk) OnConflict(opts ...sql.ConflictOption) *CardUpsertBulk {
	ccb.conflict = opts
	return &CardUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CardCreateBulk) OnConflictColumns(columns ...string) *CardUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CardUpsertBulk{
		create: ccb,
	}
}

// CardUpsertBulk is the builder for "upsert"-ing
// a bulk of Card nodes.
type CardUpsertBulk struct {
	create *CardCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CardUpsertBulk) UpdateNewValues() *CardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(card.FieldCreateTime)
			}
			if _, exists := b.mutation.Code(); exists {
				s.SetIgnore(card.FieldCode)
			}
			if _, exists := b.mutation.CardCode(); exists {
				s.SetIgnore(card.FieldCardCode)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CardUpsertBulk) Ignore() *CardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CardUpsertBulk) DoNothing() *CardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CardCreateBulk.OnConflict
// documentation for more info.
func (u *CardUpsertBulk) Update(set func(*CardUpsert)) *CardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CardUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CardUpsertBulk) SetUpdateTime(v time.Time) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateUpdateTime() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetMemberCode sets the "member_code" field.
func (u *CardUpsertBulk) SetMemberCode(v string) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetMemberCode(v)
	})
}

// UpdateMemberCode sets the "member_code" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateMemberCode() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateMemberCode()
	})
}

// SetType sets the "type" field.
func (u *CardUpsertBulk) SetType(v int) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *CardUpsertBulk) AddType(v int) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateType() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateType()
	})
}

// SetBalance sets the "balance" field.
func (u *CardUpsertBulk) SetBalance(v int) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetBalance(v)
	})
}

// AddBalance adds v to the "balance" field.
func (u *CardUpsertBulk) AddBalance(v int) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.AddBalance(v)
	})
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateBalance() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateBalance()
	})
}

// SetAmount sets the "amount" field.
func (u *CardUpsertBulk) SetAmount(v int) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *CardUpsertBulk) AddAmount(v int) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateAmount() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateAmount()
	})
}

// SetStatus sets the "status" field.
func (u *CardUpsertBulk) SetStatus(v int) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *CardUpsertBulk) AddStatus(v int) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateStatus() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *CardUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CardCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CardCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}