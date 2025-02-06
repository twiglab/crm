package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type UserSignIn struct {
	ent.Schema
}

func (UserSignIn) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Comment("用户ID"),
		field.Time("sign_in_time").
			Default(time.Now).
			Comment("签到时间"),
		field.Int("points").
			Default(0).
			Comment("签到获得积分"),
		field.Int("extra_points").
			Default(0).
			Comment("额外获得积分"),
		field.String("remark").
			Default("").
			Comment("备注"),
	}
}

func (UserSignIn) Edges() []ent.Edge {
	return nil
}

func (UserSignIn) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "sign_in_time"),
	}
}
