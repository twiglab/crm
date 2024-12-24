// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/twiglab/crm/card/orm/ent/card"
	"github.com/twiglab/crm/card/orm/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCard = "Card"
)

// CardMutation represents an operation that mutates the Card nodes in the graph.
type CardMutation struct {
	config
	op              Op
	typ             string
	id              *int
	create_time     *time.Time
	update_time     *time.Time
	code            *string
	code_bin        *string
	_type           *int
	add_type        *int
	pic1            *string
	pic2            *string
	balance         *int64
	addbalance      *int64
	amount          *int64
	addamount       *int64
	member_code     *string
	bind_time       *time.Time
	hit_time        *int64
	addhit_time     *int64
	last_clean_time *time.Time
	status          *int
	addstatus       *int
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*Card, error)
	predicates      []predicate.Card
}

var _ ent.Mutation = (*CardMutation)(nil)

// cardOption allows management of the mutation configuration using functional options.
type cardOption func(*CardMutation)

// newCardMutation creates new mutation for the Card entity.
func newCardMutation(c config, op Op, opts ...cardOption) *CardMutation {
	m := &CardMutation{
		config:        c,
		op:            op,
		typ:           TypeCard,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCardID sets the ID field of the mutation.
func withCardID(id int) cardOption {
	return func(m *CardMutation) {
		var (
			err   error
			once  sync.Once
			value *Card
		)
		m.oldValue = func(ctx context.Context) (*Card, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Card.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCard sets the old Card of the mutation.
func withCard(node *Card) cardOption {
	return func(m *CardMutation) {
		m.oldValue = func(context.Context) (*Card, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CardMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CardMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CardMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CardMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Card.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreateTime sets the "create_time" field.
func (m *CardMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *CardMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *CardMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *CardMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *CardMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *CardMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetCode sets the "code" field.
func (m *CardMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *CardMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *CardMutation) ResetCode() {
	m.code = nil
}

// SetCodeBin sets the "code_bin" field.
func (m *CardMutation) SetCodeBin(s string) {
	m.code_bin = &s
}

// CodeBin returns the value of the "code_bin" field in the mutation.
func (m *CardMutation) CodeBin() (r string, exists bool) {
	v := m.code_bin
	if v == nil {
		return
	}
	return *v, true
}

// OldCodeBin returns the old "code_bin" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldCodeBin(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCodeBin is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCodeBin requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCodeBin: %w", err)
	}
	return oldValue.CodeBin, nil
}

// ResetCodeBin resets all changes to the "code_bin" field.
func (m *CardMutation) ResetCodeBin() {
	m.code_bin = nil
}

// SetType sets the "type" field.
func (m *CardMutation) SetType(i int) {
	m._type = &i
	m.add_type = nil
}

// GetType returns the value of the "type" field in the mutation.
func (m *CardMutation) GetType() (r int, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldType(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// AddType adds i to the "type" field.
func (m *CardMutation) AddType(i int) {
	if m.add_type != nil {
		*m.add_type += i
	} else {
		m.add_type = &i
	}
}

// AddedType returns the value that was added to the "type" field in this mutation.
func (m *CardMutation) AddedType() (r int, exists bool) {
	v := m.add_type
	if v == nil {
		return
	}
	return *v, true
}

// ResetType resets all changes to the "type" field.
func (m *CardMutation) ResetType() {
	m._type = nil
	m.add_type = nil
}

// SetPic1 sets the "pic1" field.
func (m *CardMutation) SetPic1(s string) {
	m.pic1 = &s
}

// Pic1 returns the value of the "pic1" field in the mutation.
func (m *CardMutation) Pic1() (r string, exists bool) {
	v := m.pic1
	if v == nil {
		return
	}
	return *v, true
}

// OldPic1 returns the old "pic1" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldPic1(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPic1 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPic1 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPic1: %w", err)
	}
	return oldValue.Pic1, nil
}

// ResetPic1 resets all changes to the "pic1" field.
func (m *CardMutation) ResetPic1() {
	m.pic1 = nil
}

// SetPic2 sets the "pic2" field.
func (m *CardMutation) SetPic2(s string) {
	m.pic2 = &s
}

// Pic2 returns the value of the "pic2" field in the mutation.
func (m *CardMutation) Pic2() (r string, exists bool) {
	v := m.pic2
	if v == nil {
		return
	}
	return *v, true
}

// OldPic2 returns the old "pic2" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldPic2(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPic2 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPic2 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPic2: %w", err)
	}
	return oldValue.Pic2, nil
}

// ResetPic2 resets all changes to the "pic2" field.
func (m *CardMutation) ResetPic2() {
	m.pic2 = nil
}

// SetBalance sets the "balance" field.
func (m *CardMutation) SetBalance(i int64) {
	m.balance = &i
	m.addbalance = nil
}

// Balance returns the value of the "balance" field in the mutation.
func (m *CardMutation) Balance() (r int64, exists bool) {
	v := m.balance
	if v == nil {
		return
	}
	return *v, true
}

// OldBalance returns the old "balance" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldBalance(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBalance is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBalance requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBalance: %w", err)
	}
	return oldValue.Balance, nil
}

// AddBalance adds i to the "balance" field.
func (m *CardMutation) AddBalance(i int64) {
	if m.addbalance != nil {
		*m.addbalance += i
	} else {
		m.addbalance = &i
	}
}

// AddedBalance returns the value that was added to the "balance" field in this mutation.
func (m *CardMutation) AddedBalance() (r int64, exists bool) {
	v := m.addbalance
	if v == nil {
		return
	}
	return *v, true
}

// ResetBalance resets all changes to the "balance" field.
func (m *CardMutation) ResetBalance() {
	m.balance = nil
	m.addbalance = nil
}

// SetAmount sets the "amount" field.
func (m *CardMutation) SetAmount(i int64) {
	m.amount = &i
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *CardMutation) Amount() (r int64, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldAmount(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds i to the "amount" field.
func (m *CardMutation) AddAmount(i int64) {
	if m.addamount != nil {
		*m.addamount += i
	} else {
		m.addamount = &i
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *CardMutation) AddedAmount() (r int64, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *CardMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// SetMemberCode sets the "member_code" field.
func (m *CardMutation) SetMemberCode(s string) {
	m.member_code = &s
}

// MemberCode returns the value of the "member_code" field in the mutation.
func (m *CardMutation) MemberCode() (r string, exists bool) {
	v := m.member_code
	if v == nil {
		return
	}
	return *v, true
}

// OldMemberCode returns the old "member_code" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldMemberCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMemberCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMemberCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMemberCode: %w", err)
	}
	return oldValue.MemberCode, nil
}

// ResetMemberCode resets all changes to the "member_code" field.
func (m *CardMutation) ResetMemberCode() {
	m.member_code = nil
}

// SetBindTime sets the "bind_time" field.
func (m *CardMutation) SetBindTime(t time.Time) {
	m.bind_time = &t
}

// BindTime returns the value of the "bind_time" field in the mutation.
func (m *CardMutation) BindTime() (r time.Time, exists bool) {
	v := m.bind_time
	if v == nil {
		return
	}
	return *v, true
}

// OldBindTime returns the old "bind_time" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldBindTime(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBindTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBindTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBindTime: %w", err)
	}
	return oldValue.BindTime, nil
}

// ClearBindTime clears the value of the "bind_time" field.
func (m *CardMutation) ClearBindTime() {
	m.bind_time = nil
	m.clearedFields[card.FieldBindTime] = struct{}{}
}

// BindTimeCleared returns if the "bind_time" field was cleared in this mutation.
func (m *CardMutation) BindTimeCleared() bool {
	_, ok := m.clearedFields[card.FieldBindTime]
	return ok
}

// ResetBindTime resets all changes to the "bind_time" field.
func (m *CardMutation) ResetBindTime() {
	m.bind_time = nil
	delete(m.clearedFields, card.FieldBindTime)
}

// SetHitTime sets the "hit_time" field.
func (m *CardMutation) SetHitTime(i int64) {
	m.hit_time = &i
	m.addhit_time = nil
}

// HitTime returns the value of the "hit_time" field in the mutation.
func (m *CardMutation) HitTime() (r int64, exists bool) {
	v := m.hit_time
	if v == nil {
		return
	}
	return *v, true
}

// OldHitTime returns the old "hit_time" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldHitTime(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHitTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHitTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHitTime: %w", err)
	}
	return oldValue.HitTime, nil
}

// AddHitTime adds i to the "hit_time" field.
func (m *CardMutation) AddHitTime(i int64) {
	if m.addhit_time != nil {
		*m.addhit_time += i
	} else {
		m.addhit_time = &i
	}
}

// AddedHitTime returns the value that was added to the "hit_time" field in this mutation.
func (m *CardMutation) AddedHitTime() (r int64, exists bool) {
	v := m.addhit_time
	if v == nil {
		return
	}
	return *v, true
}

// ResetHitTime resets all changes to the "hit_time" field.
func (m *CardMutation) ResetHitTime() {
	m.hit_time = nil
	m.addhit_time = nil
}

// SetLastCleanTime sets the "last_clean_time" field.
func (m *CardMutation) SetLastCleanTime(t time.Time) {
	m.last_clean_time = &t
}

// LastCleanTime returns the value of the "last_clean_time" field in the mutation.
func (m *CardMutation) LastCleanTime() (r time.Time, exists bool) {
	v := m.last_clean_time
	if v == nil {
		return
	}
	return *v, true
}

// OldLastCleanTime returns the old "last_clean_time" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldLastCleanTime(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastCleanTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastCleanTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastCleanTime: %w", err)
	}
	return oldValue.LastCleanTime, nil
}

// ClearLastCleanTime clears the value of the "last_clean_time" field.
func (m *CardMutation) ClearLastCleanTime() {
	m.last_clean_time = nil
	m.clearedFields[card.FieldLastCleanTime] = struct{}{}
}

// LastCleanTimeCleared returns if the "last_clean_time" field was cleared in this mutation.
func (m *CardMutation) LastCleanTimeCleared() bool {
	_, ok := m.clearedFields[card.FieldLastCleanTime]
	return ok
}

// ResetLastCleanTime resets all changes to the "last_clean_time" field.
func (m *CardMutation) ResetLastCleanTime() {
	m.last_clean_time = nil
	delete(m.clearedFields, card.FieldLastCleanTime)
}

// SetStatus sets the "status" field.
func (m *CardMutation) SetStatus(i int) {
	m.status = &i
	m.addstatus = nil
}

// Status returns the value of the "status" field in the mutation.
func (m *CardMutation) Status() (r int, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldStatus(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// AddStatus adds i to the "status" field.
func (m *CardMutation) AddStatus(i int) {
	if m.addstatus != nil {
		*m.addstatus += i
	} else {
		m.addstatus = &i
	}
}

// AddedStatus returns the value that was added to the "status" field in this mutation.
func (m *CardMutation) AddedStatus() (r int, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ResetStatus resets all changes to the "status" field.
func (m *CardMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
}

// Where appends a list predicates to the CardMutation builder.
func (m *CardMutation) Where(ps ...predicate.Card) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CardMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CardMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Card, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CardMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CardMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Card).
func (m *CardMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CardMutation) Fields() []string {
	fields := make([]string, 0, 14)
	if m.create_time != nil {
		fields = append(fields, card.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, card.FieldUpdateTime)
	}
	if m.code != nil {
		fields = append(fields, card.FieldCode)
	}
	if m.code_bin != nil {
		fields = append(fields, card.FieldCodeBin)
	}
	if m._type != nil {
		fields = append(fields, card.FieldType)
	}
	if m.pic1 != nil {
		fields = append(fields, card.FieldPic1)
	}
	if m.pic2 != nil {
		fields = append(fields, card.FieldPic2)
	}
	if m.balance != nil {
		fields = append(fields, card.FieldBalance)
	}
	if m.amount != nil {
		fields = append(fields, card.FieldAmount)
	}
	if m.member_code != nil {
		fields = append(fields, card.FieldMemberCode)
	}
	if m.bind_time != nil {
		fields = append(fields, card.FieldBindTime)
	}
	if m.hit_time != nil {
		fields = append(fields, card.FieldHitTime)
	}
	if m.last_clean_time != nil {
		fields = append(fields, card.FieldLastCleanTime)
	}
	if m.status != nil {
		fields = append(fields, card.FieldStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CardMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case card.FieldCreateTime:
		return m.CreateTime()
	case card.FieldUpdateTime:
		return m.UpdateTime()
	case card.FieldCode:
		return m.Code()
	case card.FieldCodeBin:
		return m.CodeBin()
	case card.FieldType:
		return m.GetType()
	case card.FieldPic1:
		return m.Pic1()
	case card.FieldPic2:
		return m.Pic2()
	case card.FieldBalance:
		return m.Balance()
	case card.FieldAmount:
		return m.Amount()
	case card.FieldMemberCode:
		return m.MemberCode()
	case card.FieldBindTime:
		return m.BindTime()
	case card.FieldHitTime:
		return m.HitTime()
	case card.FieldLastCleanTime:
		return m.LastCleanTime()
	case card.FieldStatus:
		return m.Status()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CardMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case card.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case card.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case card.FieldCode:
		return m.OldCode(ctx)
	case card.FieldCodeBin:
		return m.OldCodeBin(ctx)
	case card.FieldType:
		return m.OldType(ctx)
	case card.FieldPic1:
		return m.OldPic1(ctx)
	case card.FieldPic2:
		return m.OldPic2(ctx)
	case card.FieldBalance:
		return m.OldBalance(ctx)
	case card.FieldAmount:
		return m.OldAmount(ctx)
	case card.FieldMemberCode:
		return m.OldMemberCode(ctx)
	case card.FieldBindTime:
		return m.OldBindTime(ctx)
	case card.FieldHitTime:
		return m.OldHitTime(ctx)
	case card.FieldLastCleanTime:
		return m.OldLastCleanTime(ctx)
	case card.FieldStatus:
		return m.OldStatus(ctx)
	}
	return nil, fmt.Errorf("unknown Card field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CardMutation) SetField(name string, value ent.Value) error {
	switch name {
	case card.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case card.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case card.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case card.FieldCodeBin:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCodeBin(v)
		return nil
	case card.FieldType:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case card.FieldPic1:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPic1(v)
		return nil
	case card.FieldPic2:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPic2(v)
		return nil
	case card.FieldBalance:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBalance(v)
		return nil
	case card.FieldAmount:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case card.FieldMemberCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMemberCode(v)
		return nil
	case card.FieldBindTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBindTime(v)
		return nil
	case card.FieldHitTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHitTime(v)
		return nil
	case card.FieldLastCleanTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastCleanTime(v)
		return nil
	case card.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CardMutation) AddedFields() []string {
	var fields []string
	if m.add_type != nil {
		fields = append(fields, card.FieldType)
	}
	if m.addbalance != nil {
		fields = append(fields, card.FieldBalance)
	}
	if m.addamount != nil {
		fields = append(fields, card.FieldAmount)
	}
	if m.addhit_time != nil {
		fields = append(fields, card.FieldHitTime)
	}
	if m.addstatus != nil {
		fields = append(fields, card.FieldStatus)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CardMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case card.FieldType:
		return m.AddedType()
	case card.FieldBalance:
		return m.AddedBalance()
	case card.FieldAmount:
		return m.AddedAmount()
	case card.FieldHitTime:
		return m.AddedHitTime()
	case card.FieldStatus:
		return m.AddedStatus()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CardMutation) AddField(name string, value ent.Value) error {
	switch name {
	case card.FieldType:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddType(v)
		return nil
	case card.FieldBalance:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBalance(v)
		return nil
	case card.FieldAmount:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	case card.FieldHitTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddHitTime(v)
		return nil
	case card.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	}
	return fmt.Errorf("unknown Card numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CardMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(card.FieldBindTime) {
		fields = append(fields, card.FieldBindTime)
	}
	if m.FieldCleared(card.FieldLastCleanTime) {
		fields = append(fields, card.FieldLastCleanTime)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CardMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CardMutation) ClearField(name string) error {
	switch name {
	case card.FieldBindTime:
		m.ClearBindTime()
		return nil
	case card.FieldLastCleanTime:
		m.ClearLastCleanTime()
		return nil
	}
	return fmt.Errorf("unknown Card nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CardMutation) ResetField(name string) error {
	switch name {
	case card.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case card.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case card.FieldCode:
		m.ResetCode()
		return nil
	case card.FieldCodeBin:
		m.ResetCodeBin()
		return nil
	case card.FieldType:
		m.ResetType()
		return nil
	case card.FieldPic1:
		m.ResetPic1()
		return nil
	case card.FieldPic2:
		m.ResetPic2()
		return nil
	case card.FieldBalance:
		m.ResetBalance()
		return nil
	case card.FieldAmount:
		m.ResetAmount()
		return nil
	case card.FieldMemberCode:
		m.ResetMemberCode()
		return nil
	case card.FieldBindTime:
		m.ResetBindTime()
		return nil
	case card.FieldHitTime:
		m.ResetHitTime()
		return nil
	case card.FieldLastCleanTime:
		m.ResetLastCleanTime()
		return nil
	case card.FieldStatus:
		m.ResetStatus()
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CardMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CardMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CardMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CardMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CardMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CardMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CardMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Card unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CardMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Card edge %s", name)
}
