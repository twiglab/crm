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
	"github.com/twiglab/crm/member/orm/schema/internal/x"
)

type Member struct {
	ent.Schema
}

func (Member) Fields() []ent.Field {
	return []ent.Field{

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

		field.String("phone").
			MaxLen(64).
			Optional().
			Unique().
			Validate(x.CheckMobile).
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("nickname").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		field.String("wx_open_id").
			MaxLen(256).
			Unique().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(256)", // Override MySQL.
				dialect.Postgres: "varchar(256)", // Override Postgres.
				dialect.SQLite:   "varchar(256)", // Override Postgres.
			}),

		/*
			field.String("wx_unionid").
				MaxLen(256).
				Optional().
				Unique().
				SchemaType(map[string]string{
					dialect.MySQL:    "varchar(256)", // Override MySQL.
					dialect.Postgres: "varchar(256)", // Override Postgres.
					dialect.SQLite:   "varchar(256)", // Override Postgres.
				}),
		*/

		// 商圈会员卡号
		// 用户在商圈会员卡card_id下的唯一标志，用户领取会员卡后获得的code
		field.String("bcmb_code").
			MaxLen(64).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		// 商圈会员注册时间
		field.Time("bcmb_reg_time").
			Optional().
			Nillable(),

		// 商圈会员消息id，用于幂等
		field.String("bcmb_reg_msg_id").
			MaxLen(64).
			Unique().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "varchar(64)", // Override MySQL.
				dialect.Postgres: "varchar(64)", // Override Postgres.
				dialect.SQLite:   "varchar(64)", // Override Postgres.
			}),

		// 商圈会员授权类型
		// 0 未开通
		// 1 REGISTERED_MODE ：会员开卡(进卡包) + 未授权会员积分服务
		// 2 REGISTERED_AND_AUTHORIZATION_MODE：会员开卡(进卡包）+授权会员积分服务
		field.Int("bcmb_type").Default(0),

		// 会员等级
		field.Int("level").Default(0),

		field.Int("source").Default(0),

		// 最后一次登录时间
		field.Time("last_time").Default(time.Now),

		field.Int("status").Default(1),
	}
}

func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
		index.Fields("wx_open_id"),
	}
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_member"},
	}
}
