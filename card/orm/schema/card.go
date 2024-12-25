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
		field.String("card_bin").
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
		field.String("pic1").
			MaxLen(256).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(256)", // Override MySQL.
				dialect.Postgres: "char(256)", // Override Postgres.
				dialect.SQLite:   "char(256)", // Override Postgres.
			}),

		field.String("pic2").
			MaxLen(256).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(256)", // Override MySQL.
				dialect.Postgres: "char(256)", // Override Postgres.
				dialect.SQLite:   "char(256)", // Override Postgres.
			}),

		// 额度, 单位分
		field.Int64("amount").Immutable().Default(0),

		// Member code
		field.String("member_code").
			MaxLen(36).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(36)", // Override MySQL.
				dialect.Postgres: "char(36)", // Override Postgres.
				dialect.SQLite:   "char(36)", // Override Postgres.
			}),

		// 绑定时间
		field.Time("bind_time").
			Optional().
			Nillable(),

		// 清算后的余额，单位分，清算脚本设置
		field.Int64("last_clean_balance").Immutable().Default(0),
		// 最后一次清算时间戳,到1/1000秒
		field.Int16("last_clean_ts").Immutable().Default(0),

		// 开卡完毕（可绑定状态） = 100, 绑定（正常使用状态） = 120, 锁定 = 130
		// 小于100 内部管理，和用户无关，均为无效状态
		field.Int("status").Default(0),
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
		index.Fields("card_bin"),
		index.Fields("member_code"),
		index.Fields("status"),
	}
}

func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_card"},
	}
}
