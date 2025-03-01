// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TCardColumns holds the columns for the "t_card" table.
	TCardColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "card_bin", Type: field.TypeString, Unique: true, Size: 16, SchemaType: map[string]string{"mysql": "char(16)", "postgres": "char(16)", "sqlite3": "char(16)"}},
		{Name: "type", Type: field.TypeInt, Default: 0},
		{Name: "pic1", Type: field.TypeString, Size: 256, SchemaType: map[string]string{"mysql": "char(256)", "postgres": "char(256)", "sqlite3": "char(256)"}},
		{Name: "pic2", Type: field.TypeString, Size: 256, SchemaType: map[string]string{"mysql": "char(256)", "postgres": "char(256)", "sqlite3": "char(256)"}},
		{Name: "amount", Type: field.TypeInt64, Default: 0},
		{Name: "member_code", Type: field.TypeString, Nullable: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "bind_time", Type: field.TypeTime, Nullable: true},
		{Name: "last_use_ts", Type: field.TypeInt64, Default: 0},
		{Name: "last_clean_balance", Type: field.TypeInt64, Default: 0},
		{Name: "last_clean_ts", Type: field.TypeInt16, Default: 0},
		{Name: "status", Type: field.TypeInt, Default: 0},
	}
	// TCardTable holds the schema information for the "t_card" table.
	TCardTable = &schema.Table{
		Name:       "t_card",
		Columns:    TCardColumns,
		PrimaryKey: []*schema.Column{TCardColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "card_code",
				Unique:  false,
				Columns: []*schema.Column{TCardColumns[3]},
			},
			{
				Name:    "card_card_bin",
				Unique:  false,
				Columns: []*schema.Column{TCardColumns[4]},
			},
			{
				Name:    "card_member_code",
				Unique:  false,
				Columns: []*schema.Column{TCardColumns[9]},
			},
			{
				Name:    "card_status",
				Unique:  false,
				Columns: []*schema.Column{TCardColumns[14]},
			},
		},
	}
	// TChargeRecordColumns holds the columns for the "t_charge_record" table.
	TChargeRecordColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "pay_code", Type: field.TypeString, Size: 128, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "pay_ts", Type: field.TypeInt64, Default: 0},
		{Name: "deduct", Type: field.TypeInt64, Default: 0},
		{Name: "card_code", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "status", Type: field.TypeInt, Default: 0},
	}
	// TChargeRecordTable holds the schema information for the "t_charge_record" table.
	TChargeRecordTable = &schema.Table{
		Name:       "t_charge_record",
		Columns:    TChargeRecordColumns,
		PrimaryKey: []*schema.Column{TChargeRecordColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "chargerecord_code",
				Unique:  false,
				Columns: []*schema.Column{TChargeRecordColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TCardTable,
		TChargeRecordTable,
	}
)

func init() {
	TCardTable.Annotation = &entsql.Annotation{
		Table: "t_card",
	}
	TChargeRecordTable.Annotation = &entsql.Annotation{
		Table: "t_charge_record",
	}
}
