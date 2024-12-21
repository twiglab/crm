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
	"github.com/twiglab/crm/member/orm/ent/member"
	"github.com/twiglab/crm/member/orm/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMember = "Member"
)

// MemberMutation represents an operation that mutates the Member nodes in the graph.
type MemberMutation struct {
	config
	op              Op
	typ             string
	id              *int
	create_time     *time.Time
	update_time     *time.Time
	code            *string
	phone           *string
	nickname        *string
	wx_open_id      *string
	bcmb_code       *string
	bcmb_reg_time   *time.Time
	bcmb_reg_msg_id *string
	bcmb_type       *int
	addbcmb_type    *int
	level           *int
	addlevel        *int
	source          *int
	addsource       *int
	last_time       *time.Time
	status          *int
	addstatus       *int
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*Member, error)
	predicates      []predicate.Member
}

var _ ent.Mutation = (*MemberMutation)(nil)

// memberOption allows management of the mutation configuration using functional options.
type memberOption func(*MemberMutation)

// newMemberMutation creates new mutation for the Member entity.
func newMemberMutation(c config, op Op, opts ...memberOption) *MemberMutation {
	m := &MemberMutation{
		config:        c,
		op:            op,
		typ:           TypeMember,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMemberID sets the ID field of the mutation.
func withMemberID(id int) memberOption {
	return func(m *MemberMutation) {
		var (
			err   error
			once  sync.Once
			value *Member
		)
		m.oldValue = func(ctx context.Context) (*Member, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Member.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMember sets the old Member of the mutation.
func withMember(node *Member) memberOption {
	return func(m *MemberMutation) {
		m.oldValue = func(context.Context) (*Member, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MemberMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MemberMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MemberMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MemberMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Member.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreateTime sets the "create_time" field.
func (m *MemberMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *MemberMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
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
func (m *MemberMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *MemberMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *MemberMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
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
func (m *MemberMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetCode sets the "code" field.
func (m *MemberMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *MemberMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldCode(ctx context.Context) (v string, err error) {
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
func (m *MemberMutation) ResetCode() {
	m.code = nil
}

// SetPhone sets the "phone" field.
func (m *MemberMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *MemberMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldPhone(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ClearPhone clears the value of the "phone" field.
func (m *MemberMutation) ClearPhone() {
	m.phone = nil
	m.clearedFields[member.FieldPhone] = struct{}{}
}

// PhoneCleared returns if the "phone" field was cleared in this mutation.
func (m *MemberMutation) PhoneCleared() bool {
	_, ok := m.clearedFields[member.FieldPhone]
	return ok
}

// ResetPhone resets all changes to the "phone" field.
func (m *MemberMutation) ResetPhone() {
	m.phone = nil
	delete(m.clearedFields, member.FieldPhone)
}

// SetNickname sets the "nickname" field.
func (m *MemberMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the value of the "nickname" field in the mutation.
func (m *MemberMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old "nickname" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ClearNickname clears the value of the "nickname" field.
func (m *MemberMutation) ClearNickname() {
	m.nickname = nil
	m.clearedFields[member.FieldNickname] = struct{}{}
}

// NicknameCleared returns if the "nickname" field was cleared in this mutation.
func (m *MemberMutation) NicknameCleared() bool {
	_, ok := m.clearedFields[member.FieldNickname]
	return ok
}

// ResetNickname resets all changes to the "nickname" field.
func (m *MemberMutation) ResetNickname() {
	m.nickname = nil
	delete(m.clearedFields, member.FieldNickname)
}

// SetWxOpenID sets the "wx_open_id" field.
func (m *MemberMutation) SetWxOpenID(s string) {
	m.wx_open_id = &s
}

// WxOpenID returns the value of the "wx_open_id" field in the mutation.
func (m *MemberMutation) WxOpenID() (r string, exists bool) {
	v := m.wx_open_id
	if v == nil {
		return
	}
	return *v, true
}

// OldWxOpenID returns the old "wx_open_id" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldWxOpenID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWxOpenID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWxOpenID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWxOpenID: %w", err)
	}
	return oldValue.WxOpenID, nil
}

// ResetWxOpenID resets all changes to the "wx_open_id" field.
func (m *MemberMutation) ResetWxOpenID() {
	m.wx_open_id = nil
}

// SetBcmbCode sets the "bcmb_code" field.
func (m *MemberMutation) SetBcmbCode(s string) {
	m.bcmb_code = &s
}

// BcmbCode returns the value of the "bcmb_code" field in the mutation.
func (m *MemberMutation) BcmbCode() (r string, exists bool) {
	v := m.bcmb_code
	if v == nil {
		return
	}
	return *v, true
}

// OldBcmbCode returns the old "bcmb_code" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldBcmbCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBcmbCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBcmbCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBcmbCode: %w", err)
	}
	return oldValue.BcmbCode, nil
}

// ClearBcmbCode clears the value of the "bcmb_code" field.
func (m *MemberMutation) ClearBcmbCode() {
	m.bcmb_code = nil
	m.clearedFields[member.FieldBcmbCode] = struct{}{}
}

// BcmbCodeCleared returns if the "bcmb_code" field was cleared in this mutation.
func (m *MemberMutation) BcmbCodeCleared() bool {
	_, ok := m.clearedFields[member.FieldBcmbCode]
	return ok
}

// ResetBcmbCode resets all changes to the "bcmb_code" field.
func (m *MemberMutation) ResetBcmbCode() {
	m.bcmb_code = nil
	delete(m.clearedFields, member.FieldBcmbCode)
}

// SetBcmbRegTime sets the "bcmb_reg_time" field.
func (m *MemberMutation) SetBcmbRegTime(t time.Time) {
	m.bcmb_reg_time = &t
}

// BcmbRegTime returns the value of the "bcmb_reg_time" field in the mutation.
func (m *MemberMutation) BcmbRegTime() (r time.Time, exists bool) {
	v := m.bcmb_reg_time
	if v == nil {
		return
	}
	return *v, true
}

// OldBcmbRegTime returns the old "bcmb_reg_time" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldBcmbRegTime(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBcmbRegTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBcmbRegTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBcmbRegTime: %w", err)
	}
	return oldValue.BcmbRegTime, nil
}

// ClearBcmbRegTime clears the value of the "bcmb_reg_time" field.
func (m *MemberMutation) ClearBcmbRegTime() {
	m.bcmb_reg_time = nil
	m.clearedFields[member.FieldBcmbRegTime] = struct{}{}
}

// BcmbRegTimeCleared returns if the "bcmb_reg_time" field was cleared in this mutation.
func (m *MemberMutation) BcmbRegTimeCleared() bool {
	_, ok := m.clearedFields[member.FieldBcmbRegTime]
	return ok
}

// ResetBcmbRegTime resets all changes to the "bcmb_reg_time" field.
func (m *MemberMutation) ResetBcmbRegTime() {
	m.bcmb_reg_time = nil
	delete(m.clearedFields, member.FieldBcmbRegTime)
}

// SetBcmbRegMsgID sets the "bcmb_reg_msg_id" field.
func (m *MemberMutation) SetBcmbRegMsgID(s string) {
	m.bcmb_reg_msg_id = &s
}

// BcmbRegMsgID returns the value of the "bcmb_reg_msg_id" field in the mutation.
func (m *MemberMutation) BcmbRegMsgID() (r string, exists bool) {
	v := m.bcmb_reg_msg_id
	if v == nil {
		return
	}
	return *v, true
}

// OldBcmbRegMsgID returns the old "bcmb_reg_msg_id" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldBcmbRegMsgID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBcmbRegMsgID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBcmbRegMsgID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBcmbRegMsgID: %w", err)
	}
	return oldValue.BcmbRegMsgID, nil
}

// ClearBcmbRegMsgID clears the value of the "bcmb_reg_msg_id" field.
func (m *MemberMutation) ClearBcmbRegMsgID() {
	m.bcmb_reg_msg_id = nil
	m.clearedFields[member.FieldBcmbRegMsgID] = struct{}{}
}

// BcmbRegMsgIDCleared returns if the "bcmb_reg_msg_id" field was cleared in this mutation.
func (m *MemberMutation) BcmbRegMsgIDCleared() bool {
	_, ok := m.clearedFields[member.FieldBcmbRegMsgID]
	return ok
}

// ResetBcmbRegMsgID resets all changes to the "bcmb_reg_msg_id" field.
func (m *MemberMutation) ResetBcmbRegMsgID() {
	m.bcmb_reg_msg_id = nil
	delete(m.clearedFields, member.FieldBcmbRegMsgID)
}

// SetBcmbType sets the "bcmb_type" field.
func (m *MemberMutation) SetBcmbType(i int) {
	m.bcmb_type = &i
	m.addbcmb_type = nil
}

// BcmbType returns the value of the "bcmb_type" field in the mutation.
func (m *MemberMutation) BcmbType() (r int, exists bool) {
	v := m.bcmb_type
	if v == nil {
		return
	}
	return *v, true
}

// OldBcmbType returns the old "bcmb_type" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldBcmbType(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBcmbType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBcmbType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBcmbType: %w", err)
	}
	return oldValue.BcmbType, nil
}

// AddBcmbType adds i to the "bcmb_type" field.
func (m *MemberMutation) AddBcmbType(i int) {
	if m.addbcmb_type != nil {
		*m.addbcmb_type += i
	} else {
		m.addbcmb_type = &i
	}
}

// AddedBcmbType returns the value that was added to the "bcmb_type" field in this mutation.
func (m *MemberMutation) AddedBcmbType() (r int, exists bool) {
	v := m.addbcmb_type
	if v == nil {
		return
	}
	return *v, true
}

// ResetBcmbType resets all changes to the "bcmb_type" field.
func (m *MemberMutation) ResetBcmbType() {
	m.bcmb_type = nil
	m.addbcmb_type = nil
}

// SetLevel sets the "level" field.
func (m *MemberMutation) SetLevel(i int) {
	m.level = &i
	m.addlevel = nil
}

// Level returns the value of the "level" field in the mutation.
func (m *MemberMutation) Level() (r int, exists bool) {
	v := m.level
	if v == nil {
		return
	}
	return *v, true
}

// OldLevel returns the old "level" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldLevel(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLevel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLevel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLevel: %w", err)
	}
	return oldValue.Level, nil
}

// AddLevel adds i to the "level" field.
func (m *MemberMutation) AddLevel(i int) {
	if m.addlevel != nil {
		*m.addlevel += i
	} else {
		m.addlevel = &i
	}
}

// AddedLevel returns the value that was added to the "level" field in this mutation.
func (m *MemberMutation) AddedLevel() (r int, exists bool) {
	v := m.addlevel
	if v == nil {
		return
	}
	return *v, true
}

// ResetLevel resets all changes to the "level" field.
func (m *MemberMutation) ResetLevel() {
	m.level = nil
	m.addlevel = nil
}

// SetSource sets the "source" field.
func (m *MemberMutation) SetSource(i int) {
	m.source = &i
	m.addsource = nil
}

// Source returns the value of the "source" field in the mutation.
func (m *MemberMutation) Source() (r int, exists bool) {
	v := m.source
	if v == nil {
		return
	}
	return *v, true
}

// OldSource returns the old "source" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldSource(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSource is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSource requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSource: %w", err)
	}
	return oldValue.Source, nil
}

// AddSource adds i to the "source" field.
func (m *MemberMutation) AddSource(i int) {
	if m.addsource != nil {
		*m.addsource += i
	} else {
		m.addsource = &i
	}
}

// AddedSource returns the value that was added to the "source" field in this mutation.
func (m *MemberMutation) AddedSource() (r int, exists bool) {
	v := m.addsource
	if v == nil {
		return
	}
	return *v, true
}

// ResetSource resets all changes to the "source" field.
func (m *MemberMutation) ResetSource() {
	m.source = nil
	m.addsource = nil
}

// SetLastTime sets the "last_time" field.
func (m *MemberMutation) SetLastTime(t time.Time) {
	m.last_time = &t
}

// LastTime returns the value of the "last_time" field in the mutation.
func (m *MemberMutation) LastTime() (r time.Time, exists bool) {
	v := m.last_time
	if v == nil {
		return
	}
	return *v, true
}

// OldLastTime returns the old "last_time" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldLastTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastTime: %w", err)
	}
	return oldValue.LastTime, nil
}

// ResetLastTime resets all changes to the "last_time" field.
func (m *MemberMutation) ResetLastTime() {
	m.last_time = nil
}

// SetStatus sets the "status" field.
func (m *MemberMutation) SetStatus(i int) {
	m.status = &i
	m.addstatus = nil
}

// Status returns the value of the "status" field in the mutation.
func (m *MemberMutation) Status() (r int, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldStatus(ctx context.Context) (v int, err error) {
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
func (m *MemberMutation) AddStatus(i int) {
	if m.addstatus != nil {
		*m.addstatus += i
	} else {
		m.addstatus = &i
	}
}

// AddedStatus returns the value that was added to the "status" field in this mutation.
func (m *MemberMutation) AddedStatus() (r int, exists bool) {
	v := m.addstatus
	if v == nil {
		return
	}
	return *v, true
}

// ResetStatus resets all changes to the "status" field.
func (m *MemberMutation) ResetStatus() {
	m.status = nil
	m.addstatus = nil
}

// Where appends a list predicates to the MemberMutation builder.
func (m *MemberMutation) Where(ps ...predicate.Member) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the MemberMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *MemberMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Member, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *MemberMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *MemberMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Member).
func (m *MemberMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MemberMutation) Fields() []string {
	fields := make([]string, 0, 14)
	if m.create_time != nil {
		fields = append(fields, member.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, member.FieldUpdateTime)
	}
	if m.code != nil {
		fields = append(fields, member.FieldCode)
	}
	if m.phone != nil {
		fields = append(fields, member.FieldPhone)
	}
	if m.nickname != nil {
		fields = append(fields, member.FieldNickname)
	}
	if m.wx_open_id != nil {
		fields = append(fields, member.FieldWxOpenID)
	}
	if m.bcmb_code != nil {
		fields = append(fields, member.FieldBcmbCode)
	}
	if m.bcmb_reg_time != nil {
		fields = append(fields, member.FieldBcmbRegTime)
	}
	if m.bcmb_reg_msg_id != nil {
		fields = append(fields, member.FieldBcmbRegMsgID)
	}
	if m.bcmb_type != nil {
		fields = append(fields, member.FieldBcmbType)
	}
	if m.level != nil {
		fields = append(fields, member.FieldLevel)
	}
	if m.source != nil {
		fields = append(fields, member.FieldSource)
	}
	if m.last_time != nil {
		fields = append(fields, member.FieldLastTime)
	}
	if m.status != nil {
		fields = append(fields, member.FieldStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MemberMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case member.FieldCreateTime:
		return m.CreateTime()
	case member.FieldUpdateTime:
		return m.UpdateTime()
	case member.FieldCode:
		return m.Code()
	case member.FieldPhone:
		return m.Phone()
	case member.FieldNickname:
		return m.Nickname()
	case member.FieldWxOpenID:
		return m.WxOpenID()
	case member.FieldBcmbCode:
		return m.BcmbCode()
	case member.FieldBcmbRegTime:
		return m.BcmbRegTime()
	case member.FieldBcmbRegMsgID:
		return m.BcmbRegMsgID()
	case member.FieldBcmbType:
		return m.BcmbType()
	case member.FieldLevel:
		return m.Level()
	case member.FieldSource:
		return m.Source()
	case member.FieldLastTime:
		return m.LastTime()
	case member.FieldStatus:
		return m.Status()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MemberMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case member.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case member.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case member.FieldCode:
		return m.OldCode(ctx)
	case member.FieldPhone:
		return m.OldPhone(ctx)
	case member.FieldNickname:
		return m.OldNickname(ctx)
	case member.FieldWxOpenID:
		return m.OldWxOpenID(ctx)
	case member.FieldBcmbCode:
		return m.OldBcmbCode(ctx)
	case member.FieldBcmbRegTime:
		return m.OldBcmbRegTime(ctx)
	case member.FieldBcmbRegMsgID:
		return m.OldBcmbRegMsgID(ctx)
	case member.FieldBcmbType:
		return m.OldBcmbType(ctx)
	case member.FieldLevel:
		return m.OldLevel(ctx)
	case member.FieldSource:
		return m.OldSource(ctx)
	case member.FieldLastTime:
		return m.OldLastTime(ctx)
	case member.FieldStatus:
		return m.OldStatus(ctx)
	}
	return nil, fmt.Errorf("unknown Member field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) SetField(name string, value ent.Value) error {
	switch name {
	case member.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case member.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case member.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case member.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case member.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case member.FieldWxOpenID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWxOpenID(v)
		return nil
	case member.FieldBcmbCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBcmbCode(v)
		return nil
	case member.FieldBcmbRegTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBcmbRegTime(v)
		return nil
	case member.FieldBcmbRegMsgID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBcmbRegMsgID(v)
		return nil
	case member.FieldBcmbType:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBcmbType(v)
		return nil
	case member.FieldLevel:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLevel(v)
		return nil
	case member.FieldSource:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSource(v)
		return nil
	case member.FieldLastTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastTime(v)
		return nil
	case member.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MemberMutation) AddedFields() []string {
	var fields []string
	if m.addbcmb_type != nil {
		fields = append(fields, member.FieldBcmbType)
	}
	if m.addlevel != nil {
		fields = append(fields, member.FieldLevel)
	}
	if m.addsource != nil {
		fields = append(fields, member.FieldSource)
	}
	if m.addstatus != nil {
		fields = append(fields, member.FieldStatus)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MemberMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case member.FieldBcmbType:
		return m.AddedBcmbType()
	case member.FieldLevel:
		return m.AddedLevel()
	case member.FieldSource:
		return m.AddedSource()
	case member.FieldStatus:
		return m.AddedStatus()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) AddField(name string, value ent.Value) error {
	switch name {
	case member.FieldBcmbType:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBcmbType(v)
		return nil
	case member.FieldLevel:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddLevel(v)
		return nil
	case member.FieldSource:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSource(v)
		return nil
	case member.FieldStatus:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStatus(v)
		return nil
	}
	return fmt.Errorf("unknown Member numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MemberMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(member.FieldPhone) {
		fields = append(fields, member.FieldPhone)
	}
	if m.FieldCleared(member.FieldNickname) {
		fields = append(fields, member.FieldNickname)
	}
	if m.FieldCleared(member.FieldBcmbCode) {
		fields = append(fields, member.FieldBcmbCode)
	}
	if m.FieldCleared(member.FieldBcmbRegTime) {
		fields = append(fields, member.FieldBcmbRegTime)
	}
	if m.FieldCleared(member.FieldBcmbRegMsgID) {
		fields = append(fields, member.FieldBcmbRegMsgID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MemberMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MemberMutation) ClearField(name string) error {
	switch name {
	case member.FieldPhone:
		m.ClearPhone()
		return nil
	case member.FieldNickname:
		m.ClearNickname()
		return nil
	case member.FieldBcmbCode:
		m.ClearBcmbCode()
		return nil
	case member.FieldBcmbRegTime:
		m.ClearBcmbRegTime()
		return nil
	case member.FieldBcmbRegMsgID:
		m.ClearBcmbRegMsgID()
		return nil
	}
	return fmt.Errorf("unknown Member nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MemberMutation) ResetField(name string) error {
	switch name {
	case member.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case member.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case member.FieldCode:
		m.ResetCode()
		return nil
	case member.FieldPhone:
		m.ResetPhone()
		return nil
	case member.FieldNickname:
		m.ResetNickname()
		return nil
	case member.FieldWxOpenID:
		m.ResetWxOpenID()
		return nil
	case member.FieldBcmbCode:
		m.ResetBcmbCode()
		return nil
	case member.FieldBcmbRegTime:
		m.ResetBcmbRegTime()
		return nil
	case member.FieldBcmbRegMsgID:
		m.ResetBcmbRegMsgID()
		return nil
	case member.FieldBcmbType:
		m.ResetBcmbType()
		return nil
	case member.FieldLevel:
		m.ResetLevel()
		return nil
	case member.FieldSource:
		m.ResetSource()
		return nil
	case member.FieldLastTime:
		m.ResetLastTime()
		return nil
	case member.FieldStatus:
		m.ResetStatus()
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MemberMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MemberMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MemberMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MemberMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MemberMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MemberMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MemberMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Member unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MemberMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Member edge %s", name)
}
