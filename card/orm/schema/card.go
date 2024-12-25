package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/twiglab/crm/card/orm/schema/internal/x"
)

type Card struct {
	ent.Schema
}

func (Card) Fields() []ent.Field {
	return []ent.Field{

		// Code
		field.String("code").
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(x.Code36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		// card code 就是一个普通的业务字段，卡上面显示的号码，全局唯一，不重复
		field.String("code_bin").
			MaxLen(16).
			NotEmpty().
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(16)", // Override MySQL.
				dialect.Postgres: "char(16)", // Override Postgres.
				dialect.SQLite:   "char(16)", // Override Postgres.
			}),

		// 种类
		field.Int("type").Default(0),

		// 图片
		field.String("pic1").MaxLen(255),
		field.String("pic2").MaxLen(255),

		// 余额 ，清算脚本设置
		field.Int64("balance").Immutable().Default(0),

		// 额度
		field.Int64("amount").Immutable().Default(0),

		// Member code
		field.String("member_code").
			MaxLen(36).
			DefaultFunc(x.Code36).
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		// 绑定时间
		field.Time("bind_time").
			Optional().
			Nillable(),

		// 激活时间
		field.Int64("hit_time").
			Optional().
			Nillable(),

		field.Time("last_clean_time").Optional().Nillable(),

		field.Int("status").Default(1),
	}
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Card) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
		index.Fields("member_code"),
	}
}

func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_card"},
	}
}
