// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/twiglab/crm/bonus/orm/ent/bonusacc"
)

// BonusAcc is the model entity for the BonusAcc schema.
type BonusAcc struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// MemberCode holds the value of the "member_code" field.
	MemberCode string `json:"member_code,omitempty"`
	// LastCleanBalance holds the value of the "last_clean_balance" field.
	LastCleanBalance int `json:"last_clean_balance,omitempty"`
	// LastCleanTs holds the value of the "last_clean_ts" field.
	LastCleanTs int64 `json:"last_clean_ts,omitempty"`
	// LastCleanTime holds the value of the "last_clean_time" field.
	LastCleanTime time.Time `json:"last_clean_time,omitempty"`
	// Status holds the value of the "status" field.
	Status       int `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BonusAcc) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bonusacc.FieldID, bonusacc.FieldLastCleanBalance, bonusacc.FieldLastCleanTs, bonusacc.FieldStatus:
			values[i] = new(sql.NullInt64)
		case bonusacc.FieldMemberCode:
			values[i] = new(sql.NullString)
		case bonusacc.FieldCreateTime, bonusacc.FieldUpdateTime, bonusacc.FieldLastCleanTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BonusAcc fields.
func (ba *BonusAcc) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bonusacc.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ba.ID = int(value.Int64)
		case bonusacc.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ba.CreateTime = value.Time
			}
		case bonusacc.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ba.UpdateTime = value.Time
			}
		case bonusacc.FieldMemberCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field member_code", values[i])
			} else if value.Valid {
				ba.MemberCode = value.String
			}
		case bonusacc.FieldLastCleanBalance:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_clean_balance", values[i])
			} else if value.Valid {
				ba.LastCleanBalance = int(value.Int64)
			}
		case bonusacc.FieldLastCleanTs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_clean_ts", values[i])
			} else if value.Valid {
				ba.LastCleanTs = value.Int64
			}
		case bonusacc.FieldLastCleanTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_clean_time", values[i])
			} else if value.Valid {
				ba.LastCleanTime = value.Time
			}
		case bonusacc.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ba.Status = int(value.Int64)
			}
		default:
			ba.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BonusAcc.
// This includes values selected through modifiers, order, etc.
func (ba *BonusAcc) Value(name string) (ent.Value, error) {
	return ba.selectValues.Get(name)
}

// Update returns a builder for updating this BonusAcc.
// Note that you need to call BonusAcc.Unwrap() before calling this method if this BonusAcc
// was returned from a transaction, and the transaction was committed or rolled back.
func (ba *BonusAcc) Update() *BonusAccUpdateOne {
	return NewBonusAccClient(ba.config).UpdateOne(ba)
}

// Unwrap unwraps the BonusAcc entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ba *BonusAcc) Unwrap() *BonusAcc {
	_tx, ok := ba.config.driver.(*txDriver)
	if !ok {
		panic("ent: BonusAcc is not a transactional entity")
	}
	ba.config.driver = _tx.drv
	return ba
}

// String implements the fmt.Stringer.
func (ba *BonusAcc) String() string {
	var builder strings.Builder
	builder.WriteString("BonusAcc(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ba.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ba.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ba.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("member_code=")
	builder.WriteString(ba.MemberCode)
	builder.WriteString(", ")
	builder.WriteString("last_clean_balance=")
	builder.WriteString(fmt.Sprintf("%v", ba.LastCleanBalance))
	builder.WriteString(", ")
	builder.WriteString("last_clean_ts=")
	builder.WriteString(fmt.Sprintf("%v", ba.LastCleanTs))
	builder.WriteString(", ")
	builder.WriteString("last_clean_time=")
	builder.WriteString(ba.LastCleanTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ba.Status))
	builder.WriteByte(')')
	return builder.String()
}

// BonusAccs is a parsable slice of BonusAcc.
type BonusAccs []*BonusAcc
