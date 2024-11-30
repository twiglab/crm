// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/twiglab/crm/member/orm/ent/member"
)

// Member is the model entity for the Member schema.
type Member struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// WxOpenID holds the value of the "wx_open_id" field.
	WxOpenID string `json:"wx_open_id,omitempty"`
	// WxUID holds the value of the "wx_uid" field.
	WxUID string `json:"wx_uid,omitempty"`
	// WxBcmbCode holds the value of the "wx_bcmb_code" field.
	WxBcmbCode string `json:"wx_bcmb_code,omitempty"`
	// WxBcmbRegTime holds the value of the "wx_bcmb_reg_time" field.
	WxBcmbRegTime *time.Time `json:"wx_bcmb_reg_time,omitempty"`
	// WxBcmbMsgID holds the value of the "wx_bcmb_msg_id" field.
	WxBcmbMsgID string `json:"wx_bcmb_msg_id,omitempty"`
	// WxBcmbAuthType holds the value of the "wx_bcmb_auth_type" field.
	WxBcmbAuthType int `json:"wx_bcmb_auth_type,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Source holds the value of the "source" field.
	Source       int `json:"source,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Member) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case member.FieldWxBcmbAuthType, member.FieldStatus, member.FieldSource:
			values[i] = new(sql.NullInt64)
		case member.FieldCode, member.FieldPhone, member.FieldNickname, member.FieldWxOpenID, member.FieldWxUID, member.FieldWxBcmbCode, member.FieldWxBcmbMsgID:
			values[i] = new(sql.NullString)
		case member.FieldCreateTime, member.FieldUpdateTime, member.FieldWxBcmbRegTime:
			values[i] = new(sql.NullTime)
		case member.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Member fields.
func (m *Member) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case member.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case member.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				m.CreateTime = value.Time
			}
		case member.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				m.UpdateTime = value.Time
			}
		case member.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				m.Code = value.String
			}
		case member.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				m.Phone = value.String
			}
		case member.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				m.Nickname = value.String
			}
		case member.FieldWxOpenID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wx_open_id", values[i])
			} else if value.Valid {
				m.WxOpenID = value.String
			}
		case member.FieldWxUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wx_uid", values[i])
			} else if value.Valid {
				m.WxUID = value.String
			}
		case member.FieldWxBcmbCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wx_bcmb_code", values[i])
			} else if value.Valid {
				m.WxBcmbCode = value.String
			}
		case member.FieldWxBcmbRegTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field wx_bcmb_reg_time", values[i])
			} else if value.Valid {
				m.WxBcmbRegTime = new(time.Time)
				*m.WxBcmbRegTime = value.Time
			}
		case member.FieldWxBcmbMsgID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wx_bcmb_msg_id", values[i])
			} else if value.Valid {
				m.WxBcmbMsgID = value.String
			}
		case member.FieldWxBcmbAuthType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wx_bcmb_auth_type", values[i])
			} else if value.Valid {
				m.WxBcmbAuthType = int(value.Int64)
			}
		case member.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = int(value.Int64)
			}
		case member.FieldSource:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				m.Source = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Member.
// This includes values selected through modifiers, order, etc.
func (m *Member) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Member.
// Note that you need to call Member.Unwrap() before calling this method if this Member
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Member) Update() *MemberUpdateOne {
	return NewMemberClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Member entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Member) Unwrap() *Member {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Member is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Member) String() string {
	var builder strings.Builder
	builder.WriteString("Member(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("create_time=")
	builder.WriteString(m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(m.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(m.Code)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(m.Phone)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(m.Nickname)
	builder.WriteString(", ")
	builder.WriteString("wx_open_id=")
	builder.WriteString(m.WxOpenID)
	builder.WriteString(", ")
	builder.WriteString("wx_uid=")
	builder.WriteString(m.WxUID)
	builder.WriteString(", ")
	builder.WriteString("wx_bcmb_code=")
	builder.WriteString(m.WxBcmbCode)
	builder.WriteString(", ")
	if v := m.WxBcmbRegTime; v != nil {
		builder.WriteString("wx_bcmb_reg_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("wx_bcmb_msg_id=")
	builder.WriteString(m.WxBcmbMsgID)
	builder.WriteString(", ")
	builder.WriteString("wx_bcmb_auth_type=")
	builder.WriteString(fmt.Sprintf("%v", m.WxBcmbAuthType))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", m.Source))
	builder.WriteByte(')')
	return builder.String()
}

// Members is a parsable slice of Member.
type Members []*Member
