// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TBonusAccColumns holds the columns for the "t_bonus_acc" table.
	TBonusAccColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "member_code", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "varchar(64)", "postgres": "varchar(64)", "sqlite3": "varchar(64)"}},
		{Name: "last_clean_balance", Type: field.TypeInt, Default: 0},
		{Name: "last_clean_ts", Type: field.TypeInt64, Default: 0},
		{Name: "last_clean_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt, Default: 1},
	}
	// TBonusAccTable holds the schema information for the "t_bonus_acc" table.
	TBonusAccTable = &schema.Table{
		Name:       "t_bonus_acc",
		Columns:    TBonusAccColumns,
		PrimaryKey: []*schema.Column{TBonusAccColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "bonusacc_member_code",
				Unique:  false,
				Columns: []*schema.Column{TBonusAccColumns[3]},
			},
		},
	}
	// TBonusItemColumns holds the columns for the "t_bonus_item" table.
	TBonusItemColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "mch_id", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "varchar(36)", "postgres": "varchar(36)", "sqlite3": "varchar(36)"}},
		{Name: "mall_code", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "varchar(36)", "postgres": "varchar(36)", "sqlite3": "varchar(36)"}},
		{Name: "mall_name", Type: field.TypeString, Size: 256, SchemaType: map[string]string{"mysql": "varchar(256)", "postgres": "varchar(256)", "sqlite3": "varchar(256)"}},
		{Name: "shop_code", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "varchar(36)", "postgres": "varchar(36)", "sqlite3": "varchar(36)"}},
		{Name: "shop_name", Type: field.TypeString, Size: 256, SchemaType: map[string]string{"mysql": "varchar(256)", "postgres": "varchar(256)", "sqlite3": "varchar(256)"}},
		{Name: "member_code", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "varchar(36)", "postgres": "varchar(36)", "sqlite3": "varchar(36)"}},
		{Name: "wx_open_id", Type: field.TypeString, Size: 256, SchemaType: map[string]string{"mysql": "varchar(256)", "postgres": "varchar(256)", "sqlite3": "varchar(256)"}},
		{Name: "bcmb_notify_time", Type: field.TypeTime},
		{Name: "bcmb_notify_id", Type: field.TypeString, Unique: true, Nullable: true, Size: 64, SchemaType: map[string]string{"mysql": "varchar(64)", "postgres": "varchar(64)", "sqlite3": "varchar(64)"}},
		{Name: "bcmb_trans_code", Type: field.TypeString, Unique: true, Size: 64, SchemaType: map[string]string{"mysql": "varchar(64)", "postgres": "varchar(64)", "sqlite3": "varchar(64)"}},
		{Name: "amount", Type: field.TypeInt, Default: 0},
		{Name: "bcmb_trans_time", Type: field.TypeTime},
		{Name: "create_ts", Type: field.TypeInt64, Default: 0},
		{Name: "bcmb_trans_pay_code", Type: field.TypeString, Nullable: true, Size: 64, SchemaType: map[string]string{"mysql": "varchar(64)", "postgres": "varchar(64)", "sqlite3": "varchar(64)"}},
		{Name: "bcmb_trans_type", Type: field.TypeInt, Default: 0},
		{Name: "bonus", Type: field.TypeInt, Default: 0},
		{Name: "bonus_rate", Type: field.TypeInt32, Default: 100},
		{Name: "status", Type: field.TypeInt, Default: 1},
	}
	// TBonusItemTable holds the schema information for the "t_bonus_item" table.
	TBonusItemTable = &schema.Table{
		Name:       "t_bonus_item",
		Columns:    TBonusItemColumns,
		PrimaryKey: []*schema.Column{TBonusItemColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "bonusitem_code",
				Unique:  false,
				Columns: []*schema.Column{TBonusItemColumns[3]},
			},
			{
				Name:    "bonusitem_wx_open_id",
				Unique:  false,
				Columns: []*schema.Column{TBonusItemColumns[10]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TBonusAccTable,
		TBonusItemTable,
	}
)

func init() {
	TBonusAccTable.Annotation = &entsql.Annotation{
		Table: "t_bonus_acc",
	}
	TBonusItemTable.Annotation = &entsql.Annotation{
		Table: "t_bonus_item",
	}
}
