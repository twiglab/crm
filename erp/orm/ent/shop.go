// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/twiglab/crm/erp/orm/ent/shop"
)

// Shop is the model entity for the Shop schema.
type Shop struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// MallCode holds the value of the "mall_code" field.
	MallCode string `json:"mall_code,omitempty"`
	// MallName holds the value of the "mall_name" field.
	MallName string `json:"mall_name,omitempty"`
	// ContractCode holds the value of the "contract_code" field.
	ContractCode string `json:"contract_code,omitempty"`
	// PosCode holds the value of the "pos_code" field.
	PosCode string `json:"pos_code,omitempty"`
	// ShopCode holds the value of the "shop_code" field.
	ShopCode string `json:"shop_code,omitempty"`
	// ShopName holds the value of the "shop_name" field.
	ShopName string `json:"shop_name,omitempty"`
	// BizClass1 holds the value of the "biz_class_1" field.
	BizClass1 string `json:"biz_class_1,omitempty"`
	// BizClassName1 holds the value of the "biz_class_name_1" field.
	BizClassName1 string `json:"biz_class_name_1,omitempty"`
	// BizClass2 holds the value of the "biz_class_2" field.
	BizClass2 string `json:"biz_class_2,omitempty"`
	// BizClassName2 holds the value of the "biz_class_name_2" field.
	BizClassName2 string `json:"biz_class_name_2,omitempty"`
	// BizBeginTime holds the value of the "biz_begin_time" field.
	BizBeginTime time.Time `json:"biz_begin_time,omitempty"`
	// BizEndTime holds the value of the "biz_end_time" field.
	BizEndTime time.Time `json:"biz_end_time,omitempty"`
	// Status holds the value of the "status" field.
	Status       int `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Shop) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shop.FieldID, shop.FieldStatus:
			values[i] = new(sql.NullInt64)
		case shop.FieldCode, shop.FieldMallCode, shop.FieldMallName, shop.FieldContractCode, shop.FieldPosCode, shop.FieldShopCode, shop.FieldShopName, shop.FieldBizClass1, shop.FieldBizClassName1, shop.FieldBizClass2, shop.FieldBizClassName2:
			values[i] = new(sql.NullString)
		case shop.FieldCreateTime, shop.FieldUpdateTime, shop.FieldBizBeginTime, shop.FieldBizEndTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Shop fields.
func (s *Shop) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shop.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case shop.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case shop.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case shop.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case shop.FieldMallCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mall_code", values[i])
			} else if value.Valid {
				s.MallCode = value.String
			}
		case shop.FieldMallName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mall_name", values[i])
			} else if value.Valid {
				s.MallName = value.String
			}
		case shop.FieldContractCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contract_code", values[i])
			} else if value.Valid {
				s.ContractCode = value.String
			}
		case shop.FieldPosCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pos_code", values[i])
			} else if value.Valid {
				s.PosCode = value.String
			}
		case shop.FieldShopCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shop_code", values[i])
			} else if value.Valid {
				s.ShopCode = value.String
			}
		case shop.FieldShopName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shop_name", values[i])
			} else if value.Valid {
				s.ShopName = value.String
			}
		case shop.FieldBizClass1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field biz_class_1", values[i])
			} else if value.Valid {
				s.BizClass1 = value.String
			}
		case shop.FieldBizClassName1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field biz_class_name_1", values[i])
			} else if value.Valid {
				s.BizClassName1 = value.String
			}
		case shop.FieldBizClass2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field biz_class_2", values[i])
			} else if value.Valid {
				s.BizClass2 = value.String
			}
		case shop.FieldBizClassName2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field biz_class_name_2", values[i])
			} else if value.Valid {
				s.BizClassName2 = value.String
			}
		case shop.FieldBizBeginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field biz_begin_time", values[i])
			} else if value.Valid {
				s.BizBeginTime = value.Time
			}
		case shop.FieldBizEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field biz_end_time", values[i])
			} else if value.Valid {
				s.BizEndTime = value.Time
			}
		case shop.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Shop.
// This includes values selected through modifiers, order, etc.
func (s *Shop) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Shop.
// Note that you need to call Shop.Unwrap() before calling this method if this Shop
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Shop) Update() *ShopUpdateOne {
	return NewShopClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Shop entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Shop) Unwrap() *Shop {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Shop is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Shop) String() string {
	var builder strings.Builder
	builder.WriteString("Shop(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(s.Code)
	builder.WriteString(", ")
	builder.WriteString("mall_code=")
	builder.WriteString(s.MallCode)
	builder.WriteString(", ")
	builder.WriteString("mall_name=")
	builder.WriteString(s.MallName)
	builder.WriteString(", ")
	builder.WriteString("contract_code=")
	builder.WriteString(s.ContractCode)
	builder.WriteString(", ")
	builder.WriteString("pos_code=")
	builder.WriteString(s.PosCode)
	builder.WriteString(", ")
	builder.WriteString("shop_code=")
	builder.WriteString(s.ShopCode)
	builder.WriteString(", ")
	builder.WriteString("shop_name=")
	builder.WriteString(s.ShopName)
	builder.WriteString(", ")
	builder.WriteString("biz_class_1=")
	builder.WriteString(s.BizClass1)
	builder.WriteString(", ")
	builder.WriteString("biz_class_name_1=")
	builder.WriteString(s.BizClassName1)
	builder.WriteString(", ")
	builder.WriteString("biz_class_2=")
	builder.WriteString(s.BizClass2)
	builder.WriteString(", ")
	builder.WriteString("biz_class_name_2=")
	builder.WriteString(s.BizClassName2)
	builder.WriteString(", ")
	builder.WriteString("biz_begin_time=")
	builder.WriteString(s.BizBeginTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("biz_end_time=")
	builder.WriteString(s.BizEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Shops is a parsable slice of Shop.
type Shops []*Shop
