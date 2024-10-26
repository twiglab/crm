// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TAdvColumns holds the columns for the "t_adv" table.
	TAdvColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "mall_code", Type: field.TypeString, Unique: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "mall_name", Type: field.TypeString, Unique: true, Size: 64, SchemaType: map[string]string{"mysql": "char(64)", "postgres": "char(64)", "sqlite3": "char(64)"}},
		{Name: "h3_index_6", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "h3_index_7", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "h3_index_8", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "img_path", Type: field.TypeString, Unique: true, Size: 64, SchemaType: map[string]string{"mysql": "char(64)", "postgres": "char(64)", "sqlite3": "char(64)"}},
		{Name: "url", Type: field.TypeString, Unique: true, Size: 64, SchemaType: map[string]string{"mysql": "char(64)", "postgres": "char(64)", "sqlite3": "char(64)"}},
		{Name: "ruler", Type: field.TypeString, Size: 256, SchemaType: map[string]string{"mysql": "char(256)", "postgres": "char(256)", "sqlite3": "char(256)"}},
		{Name: "ord", Type: field.TypeInt, Default: 0},
		{Name: "memo", Type: field.TypeString, Size: 64, SchemaType: map[string]string{"mysql": "varchar(64)", "postgres": "varchar(64)", "sqlite3": "varchar(64)"}},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt, Default: 1},
	}
	// TAdvTable holds the schema information for the "t_adv" table.
	TAdvTable = &schema.Table{
		Name:       "t_adv",
		Columns:    TAdvColumns,
		PrimaryKey: []*schema.Column{TAdvColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "adv_code",
				Unique:  false,
				Columns: []*schema.Column{TAdvColumns[3]},
			},
			{
				Name:    "adv_status",
				Unique:  false,
				Columns: []*schema.Column{TAdvColumns[16]},
			},
			{
				Name:    "adv_ord",
				Unique:  false,
				Columns: []*schema.Column{TAdvColumns[12]},
			},
			{
				Name:    "adv_start_time",
				Unique:  false,
				Columns: []*schema.Column{TAdvColumns[14]},
			},
			{
				Name:    "adv_end_time",
				Unique:  false,
				Columns: []*schema.Column{TAdvColumns[15]},
			},
		},
	}
	// TMallColumns holds the columns for the "t_mall" table.
	TMallColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "mall_code", Type: field.TypeString, Unique: true, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "mall_name", Type: field.TypeString, Unique: true, Size: 64, SchemaType: map[string]string{"mysql": "char(64)", "postgres": "char(64)", "sqlite3": "char(64)"}},
		{Name: "h3_index_6", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "h3_index_7", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "h3_index_8", Type: field.TypeString, Size: 36, SchemaType: map[string]string{"mysql": "char(36)", "postgres": "char(36)", "sqlite3": "char(36)"}},
		{Name: "memo", Type: field.TypeString, Size: 64, SchemaType: map[string]string{"mysql": "varchar(64)", "postgres": "varchar(64)", "sqlite3": "varchar(64)"}},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt, Default: 1},
	}
	// TMallTable holds the schema information for the "t_mall" table.
	TMallTable = &schema.Table{
		Name:       "t_mall",
		Columns:    TMallColumns,
		PrimaryKey: []*schema.Column{TMallColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "mall_code",
				Unique:  false,
				Columns: []*schema.Column{TMallColumns[3]},
			},
			{
				Name:    "mall_status",
				Unique:  false,
				Columns: []*schema.Column{TMallColumns[11]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TAdvTable,
		TMallTable,
	}
)

func init() {
	TAdvTable.Annotation = &entsql.Annotation{
		Table: "t_adv",
	}
	TMallTable.Annotation = &entsql.Annotation{
		Table: "t_mall",
	}
}
