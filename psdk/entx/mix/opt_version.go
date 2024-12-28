package mix

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

const lockVerField = "version"

type ver interface {
	WhereP(...func(*sql.Selector))
	ResetVersion()
	Version() (int64, bool)
	AddVersion(int64)
}

type optimisticLockingKey struct{}

// IgnoreOptimisticLocking ingore Optionistic Locking
func IgnoreOptimisticLocking(parent context.Context) context.Context {
	return context.WithValue(parent, optimisticLockingKey{}, true)
}

// Version for ent. db field is lock_ver int64
// only OpUpdateOne do update table set lock_ver = lock_ver +1 , set ... where lock_ver = n -- n is SetLockVer(n)
// when opUpdate Version was ResetLockVer
type Version struct {
	mixin.Schema
}

func (v Version) Fields() []ent.Field {
	return []ent.Field{
		field.Int64(lockVerField).Default(0),
	}
}

func (v Version) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(lockVerField),
	}
}

func (v Version) Hooks() []ent.Hook {
	return []ent.Hook{
		versionHook(),
	}
}

func versionHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(ent.OpUpdateOne | ent.OpUpdate) {
				if skip, _ := ctx.Value(optimisticLockingKey{}).(bool); skip {
					return next.Mutate(ctx, m)
				}

				mx, ok := m.(ver)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}

				if m.Op().Is(ent.OpUpdateOne) {
					if val, exists := mx.Version(); exists {
						// update table set lock_ver = lock_ver + 1 where lock_ver = val
						mx.ResetVersion()
						mx.AddVersion(1)
						addEq(mx, val)
					} else {
						log.Println("no call SetVersion(), nothing to do")
					}

				} else {
					// Op = OpUpdate
					mx.ResetVersion()
				}
			}

			return next.Mutate(ctx, m)
		})
	}
}

func addEq(w ver, v int64) {
	w.WhereP(
		sql.FieldEQ(lockVerField, v),
	)
}
