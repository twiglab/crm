// Code generated by ent, DO NOT EDIT.

package usersignin

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the usersignin type in the database.
	Label = "user_sign_in"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldSignInTime holds the string denoting the sign_in_time field in the database.
	FieldSignInTime = "sign_in_time"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// FieldExtraPoints holds the string denoting the extra_points field in the database.
	FieldExtraPoints = "extra_points"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// Table holds the table name of the usersignin in the database.
	Table = "user_sign_ins"
)

// Columns holds all SQL columns for usersignin fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldSignInTime,
	FieldPoints,
	FieldExtraPoints,
	FieldRemark,
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
	// DefaultSignInTime holds the default value on creation for the "sign_in_time" field.
	DefaultSignInTime func() time.Time
	// DefaultPoints holds the default value on creation for the "points" field.
	DefaultPoints int
	// DefaultExtraPoints holds the default value on creation for the "extra_points" field.
	DefaultExtraPoints int
	// DefaultRemark holds the default value on creation for the "remark" field.
	DefaultRemark string
)

// OrderOption defines the ordering options for the UserSignIn queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// BySignInTime orders the results by the sign_in_time field.
func BySignInTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignInTime, opts...).ToFunc()
}

// ByPoints orders the results by the points field.
func ByPoints(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoints, opts...).ToFunc()
}

// ByExtraPoints orders the results by the extra_points field.
func ByExtraPoints(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtraPoints, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}
