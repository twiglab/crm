// Code generated by ent, DO NOT EDIT.

package bonusacc

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the bonusacc type in the database.
	Label = "bonus_acc"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldMemberCode holds the string denoting the member_code field in the database.
	FieldMemberCode = "member_code"
	// FieldLastCleanBalance holds the string denoting the last_clean_balance field in the database.
	FieldLastCleanBalance = "last_clean_balance"
	// FieldLastCleanTs holds the string denoting the last_clean_ts field in the database.
	FieldLastCleanTs = "last_clean_ts"
	// FieldLastCleanTime holds the string denoting the last_clean_time field in the database.
	FieldLastCleanTime = "last_clean_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the bonusacc in the database.
	Table = "t_bonus_acc"
)

// Columns holds all SQL columns for bonusacc fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldMemberCode,
	FieldLastCleanBalance,
	FieldLastCleanTs,
	FieldLastCleanTime,
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
	// MemberCodeValidator is a validator for the "member_code" field. It is called by the builders before save.
	MemberCodeValidator func(string) error
	// DefaultLastCleanBalance holds the default value on creation for the "last_clean_balance" field.
	DefaultLastCleanBalance int
	// DefaultLastCleanTs holds the default value on creation for the "last_clean_ts" field.
	DefaultLastCleanTs int64
	// DefaultLastCleanTime holds the default value on creation for the "last_clean_time" field.
	DefaultLastCleanTime func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
)

// OrderOption defines the ordering options for the BonusAcc queries.
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

// ByMemberCode orders the results by the member_code field.
func ByMemberCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberCode, opts...).ToFunc()
}

// ByLastCleanBalance orders the results by the last_clean_balance field.
func ByLastCleanBalance(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastCleanBalance, opts...).ToFunc()
}

// ByLastCleanTs orders the results by the last_clean_ts field.
func ByLastCleanTs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastCleanTs, opts...).ToFunc()
}

// ByLastCleanTime orders the results by the last_clean_time field.
func ByLastCleanTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastCleanTime, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
