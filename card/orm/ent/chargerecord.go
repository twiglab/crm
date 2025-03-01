// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/twiglab/crm/card/orm/ent/chargerecord"
)

// ChargeRecord is the model entity for the ChargeRecord schema.
type ChargeRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// PayCode holds the value of the "pay_code" field.
	PayCode string `json:"pay_code,omitempty"`
	// PayTs holds the value of the "pay_ts" field.
	PayTs int64 `json:"pay_ts,omitempty"`
	// Deduct holds the value of the "deduct" field.
	Deduct int64 `json:"deduct,omitempty"`
	// CardCode holds the value of the "card_code" field.
	CardCode string `json:"card_code,omitempty"`
	// Status holds the value of the "status" field.
	Status       int `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChargeRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case chargerecord.FieldID, chargerecord.FieldPayTs, chargerecord.FieldDeduct, chargerecord.FieldStatus:
			values[i] = new(sql.NullInt64)
		case chargerecord.FieldCode, chargerecord.FieldPayCode, chargerecord.FieldCardCode:
			values[i] = new(sql.NullString)
		case chargerecord.FieldCreateTime, chargerecord.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChargeRecord fields.
func (cr *ChargeRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chargerecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cr.ID = int(value.Int64)
		case chargerecord.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				cr.CreateTime = value.Time
			}
		case chargerecord.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				cr.UpdateTime = value.Time
			}
		case chargerecord.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				cr.Code = value.String
			}
		case chargerecord.FieldPayCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pay_code", values[i])
			} else if value.Valid {
				cr.PayCode = value.String
			}
		case chargerecord.FieldPayTs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pay_ts", values[i])
			} else if value.Valid {
				cr.PayTs = value.Int64
			}
		case chargerecord.FieldDeduct:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deduct", values[i])
			} else if value.Valid {
				cr.Deduct = value.Int64
			}
		case chargerecord.FieldCardCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field card_code", values[i])
			} else if value.Valid {
				cr.CardCode = value.String
			}
		case chargerecord.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cr.Status = int(value.Int64)
			}
		default:
			cr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ChargeRecord.
// This includes values selected through modifiers, order, etc.
func (cr *ChargeRecord) Value(name string) (ent.Value, error) {
	return cr.selectValues.Get(name)
}

// Update returns a builder for updating this ChargeRecord.
// Note that you need to call ChargeRecord.Unwrap() before calling this method if this ChargeRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (cr *ChargeRecord) Update() *ChargeRecordUpdateOne {
	return NewChargeRecordClient(cr.config).UpdateOne(cr)
}

// Unwrap unwraps the ChargeRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cr *ChargeRecord) Unwrap() *ChargeRecord {
	_tx, ok := cr.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChargeRecord is not a transactional entity")
	}
	cr.config.driver = _tx.drv
	return cr
}

// String implements the fmt.Stringer.
func (cr *ChargeRecord) String() string {
	var builder strings.Builder
	builder.WriteString("ChargeRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cr.ID))
	builder.WriteString("create_time=")
	builder.WriteString(cr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(cr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(cr.Code)
	builder.WriteString(", ")
	builder.WriteString("pay_code=")
	builder.WriteString(cr.PayCode)
	builder.WriteString(", ")
	builder.WriteString("pay_ts=")
	builder.WriteString(fmt.Sprintf("%v", cr.PayTs))
	builder.WriteString(", ")
	builder.WriteString("deduct=")
	builder.WriteString(fmt.Sprintf("%v", cr.Deduct))
	builder.WriteString(", ")
	builder.WriteString("card_code=")
	builder.WriteString(cr.CardCode)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", cr.Status))
	builder.WriteByte(')')
	return builder.String()
}

// ChargeRecords is a parsable slice of ChargeRecord.
type ChargeRecords []*ChargeRecord
