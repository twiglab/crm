// Code generated by ent, DO NOT EDIT.

package member

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldWxOpenID holds the string denoting the wx_open_id field in the database.
	FieldWxOpenID = "wx_open_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the member in the database.
	Table = "t_member"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCode,
	FieldPhone,
	FieldNickname,
	FieldWxOpenID,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultCode holds the default value on creation for the "code" field.
	DefaultCode func() string
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// WxOpenIDValidator is a validator for the "wx_open_id" field. It is called by the builders before save.
	WxOpenIDValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Member queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByWxOpenID orders the results by the wx_open_id field.
func ByWxOpenID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWxOpenID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
