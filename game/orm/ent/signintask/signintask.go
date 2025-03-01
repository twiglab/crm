// Code generated by ent, DO NOT EDIT.

package signintask

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the signintask type in the database.
	Label = "sign_in_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldRewardInfo holds the string denoting the reward_info field in the database.
	FieldRewardInfo = "reward_info"
	// Table holds the table name of the signintask in the database.
	Table = "sign_in_tasks"
)

// Columns holds all SQL columns for signintask fields.
var Columns = []string{
	FieldID,
	FieldTaskID,
	FieldDescription,
	FieldStartTime,
	FieldEndTime,
	FieldRewardInfo,
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

// OrderOption defines the ordering options for the SignInTask queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTaskID orders the results by the task_id field.
func ByTaskID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaskID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByRewardInfo orders the results by the reward_info field.
func ByRewardInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRewardInfo, opts...).ToFunc()
}
