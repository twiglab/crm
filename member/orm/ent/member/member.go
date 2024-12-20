// Code generated by ent, DO NOT EDIT.

package member

import (
	"time"

	"entgo.io/ent/dialect/sql"
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
	// FieldWxUID holds the string denoting the wx_uid field in the database.
	FieldWxUID = "wx_uid"
	// FieldBcmbCode holds the string denoting the bcmb_code field in the database.
	FieldBcmbCode = "bcmb_code"
	// FieldBcmbRegTime holds the string denoting the bcmb_reg_time field in the database.
	FieldBcmbRegTime = "bcmb_reg_time"
	// FieldBcmbRegMsgID holds the string denoting the bcmb_reg_msg_id field in the database.
	FieldBcmbRegMsgID = "bcmb_reg_msg_id"
	// FieldBcmbType holds the string denoting the bcmb_type field in the database.
	FieldBcmbType = "bcmb_type"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldLastTime holds the string denoting the last_time field in the database.
	FieldLastTime = "last_time"
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
	FieldWxUID,
	FieldBcmbCode,
	FieldBcmbRegTime,
	FieldBcmbRegMsgID,
	FieldBcmbType,
	FieldLevel,
	FieldSource,
	FieldLastTime,
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
	// WxUIDValidator is a validator for the "wx_uid" field. It is called by the builders before save.
	WxUIDValidator func(string) error
	// BcmbCodeValidator is a validator for the "bcmb_code" field. It is called by the builders before save.
	BcmbCodeValidator func(string) error
	// BcmbRegMsgIDValidator is a validator for the "bcmb_reg_msg_id" field. It is called by the builders before save.
	BcmbRegMsgIDValidator func(string) error
	// DefaultBcmbType holds the default value on creation for the "bcmb_type" field.
	DefaultBcmbType int
	// DefaultLevel holds the default value on creation for the "level" field.
	DefaultLevel int
	// DefaultSource holds the default value on creation for the "source" field.
	DefaultSource int
	// DefaultLastTime holds the default value on creation for the "last_time" field.
	DefaultLastTime func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
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

// ByWxUID orders the results by the wx_uid field.
func ByWxUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWxUID, opts...).ToFunc()
}

// ByBcmbCode orders the results by the bcmb_code field.
func ByBcmbCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbCode, opts...).ToFunc()
}

// ByBcmbRegTime orders the results by the bcmb_reg_time field.
func ByBcmbRegTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbRegTime, opts...).ToFunc()
}

// ByBcmbRegMsgID orders the results by the bcmb_reg_msg_id field.
func ByBcmbRegMsgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbRegMsgID, opts...).ToFunc()
}

// ByBcmbType orders the results by the bcmb_type field.
func ByBcmbType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbType, opts...).ToFunc()
}

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// ByLastTime orders the results by the last_time field.
func ByLastTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastTime, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
