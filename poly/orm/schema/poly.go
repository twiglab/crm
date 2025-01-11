package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/twiglab/crm/psdk/code"
)

type Poly struct {
	ent.Schema
}

func (Poly) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.Nil).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.String("code").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(code.Code36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		field.String("mall_code").
			MaxLen(36).
			NotEmpty().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		// 活动名称
		field.String("name").
			MaxLen(64).
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		// 用于标签或者列表显示
		field.String("title").
			MaxLen(64).
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		// 活动描述
		field.String("memo").
			MaxLen(512).
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(512)", // Override MySQL.
				dialect.Postgres: "varchar(512)", // Override Postgres.
				dialect.SQLite:   "varchar(512)", // Override Postgres.
			}),

		// 活动开始时间
		field.Time("start_time").Default(time.Now).Immutable(),
		// 活动结束时间

		field.Time("end_time").Default(time.Now).Immutable(),

		// 门槛规则
		field.String("menkan").
			MaxLen(512).
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(512)", // Override MySQL.
				dialect.Postgres: "char(512)", // Override Postgres.
				dialect.SQLite:   "char(512)", // Override Postgres.
			}),

		// 发放规则
		field.String("fafang").
			MaxLen(512).
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(512)", // Override MySQL.
				dialect.Postgres: "char(512)", // Override Postgres.
				dialect.SQLite:   "char(512)", // Override Postgres.
			}),

		// 有效期规则
		field.String("xiaoqi").
			MaxLen(512).
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(512)", // Override MySQL.
				dialect.Postgres: "char(512)", // Override Postgres.
				dialect.SQLite:   "char(512)", // Override Postgres.
			}),

		// 使用规则
		field.String("shiyong").
			MaxLen(512).
			NotEmpty().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(512)", // Override MySQL.
				dialect.Postgres: "char(512)", // Override Postgres.
				dialect.SQLite:   "char(512)", // Override Postgres.
			}),

		// 活动状态
		field.Int32("status").Default(1),

		// 活动类型 (预设字段，后续结合业务拓展)
		field.Int32("type").Default(1),
	}
}

func (Poly) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Poly) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}

func (Poly) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_poly"},
	}
}
