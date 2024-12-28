package mix

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// The Query interface represents an operation that queries with WhereP.
type (
	Query interface {
		WhereP(...func(*sql.Selector))
	}
)

// SoftDeleted implements the soft delete pattern for schemas.
type SoftDeleted[T Query, Q ent.Mutator] struct {
	mixin.Schema
	QueryFunc func(ent.Query) (T, error)
}

func NewSoftDeleted[T Query, Q ent.Mutator](qf func(ent.Query) (T, error)) SoftDeleted[T, Q] {
	return SoftDeleted[T, Q]{
		QueryFunc: qf,
	}
}

// Fields of the SoftDeleteMixin.
func (d SoftDeleted[T, Q]) Fields() []ent.Field {
	return []ent.Field{
		field.Int("deleted").Default(0),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the SoftDeleteMixin.
func (d SoftDeleted[T, Q]) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.TraverseFunc(func(ctx context.Context, q ent.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}
			df, err := d.QueryFunc(q)
			if err != nil {
				return err
			}
			d.P(df)
			return nil
		}),
	}
}

type softDeleter[Q ent.Mutator] interface {
	Client() Q
	SetOp(ent.Op)
	SetDeleted(int)
}

// Hooks of the SoftDeleteMixin.
func (d SoftDeleted[T, Q]) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				if !m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
					return next.Mutate(ctx, m)
				}
				// Skip soft-delete, means delete the entity permanently.
				if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
					return next.Mutate(ctx, m)
				}

				if mx, ok := m.(softDeleter[Q]); ok {
					mx.SetOp(ent.OpUpdate)
					mx.SetDeleted(1)
					return mx.Client().Mutate(ctx, m)
				}
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			})
		},
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d SoftDeleted[T, Q]) P(w Query) {
	w.WhereP(
		sql.FieldEQ("deleted", 0),
	)
}
