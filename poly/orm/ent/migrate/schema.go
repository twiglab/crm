// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PoliesColumns holds the columns for the "polies" table.
	PoliesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "mall_code", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "rule_code", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "name", Type: field.TypeString, Size: 64, SchemaType: map[string]string{"mysql": "varchar(64)", "postgres": "varchar(64)", "sqlite3": "varchar(64)"}},
		{Name: "desc", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "type", Type: field.TypeInt, Default: 1},
	}
	// PoliesTable holds the schema information for the "polies" table.
	PoliesTable = &schema.Table{
		Name:       "polies",
		Columns:    PoliesColumns,
		PrimaryKey: []*schema.Column{PoliesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PoliesTable,
	}
)

func init() {
}
