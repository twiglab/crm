package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/twiglab/crm/psdk/code"
)

type Poly struct {
	ent.Schema
}

func (Poly) Fields() []ent.Field {
	return []ent.Field{

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

		// 规则code
		field.String("rule").
			MaxLen(36).
			NotEmpty().
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

		// 活动描述
		field.String("desc").NotEmpty(),

		// 活动开始时间
		field.Time("start_time"),
		// 活动结束时间

		field.Time("end_time"),

		// 活动状态 （1:等待审批 2：未开启   3：已开启   4：已结束   5：已废弃）
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
