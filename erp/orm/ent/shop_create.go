// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/twiglab/crm/erp/orm/ent/shop"
)

// ShopCreate is the builder for creating a Shop entity.
type ShopCreate struct {
	config
	mutation *ShopMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (sc *ShopCreate) SetCreateTime(t time.Time) *ShopCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *ShopCreate) SetNillableCreateTime(t *time.Time) *ShopCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *ShopCreate) SetUpdateTime(t time.Time) *ShopCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *ShopCreate) SetNillableUpdateTime(t *time.Time) *ShopCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetCode sets the "code" field.
func (sc *ShopCreate) SetCode(s string) *ShopCreate {
	sc.mutation.SetCode(s)
	return sc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (sc *ShopCreate) SetNillableCode(s *string) *ShopCreate {
	if s != nil {
		sc.SetCode(*s)
	}
	return sc
}

// SetMallCode sets the "mall_code" field.
func (sc *ShopCreate) SetMallCode(s string) *ShopCreate {
	sc.mutation.SetMallCode(s)
	return sc
}

// SetMallName sets the "mall_name" field.
func (sc *ShopCreate) SetMallName(s string) *ShopCreate {
	sc.mutation.SetMallName(s)
	return sc
}

// SetContractCode sets the "contract_code" field.
func (sc *ShopCreate) SetContractCode(s string) *ShopCreate {
	sc.mutation.SetContractCode(s)
	return sc
}

// SetFloor sets the "floor" field.
func (sc *ShopCreate) SetFloor(s string) *ShopCreate {
	sc.mutation.SetFloor(s)
	return sc
}

// SetPos sets the "pos" field.
func (sc *ShopCreate) SetPos(s string) *ShopCreate {
	sc.mutation.SetPos(s)
	return sc
}

// SetShopCode sets the "shop_code" field.
func (sc *ShopCreate) SetShopCode(s string) *ShopCreate {
	sc.mutation.SetShopCode(s)
	return sc
}

// SetShopName sets the "shop_name" field.
func (sc *ShopCreate) SetShopName(s string) *ShopCreate {
	sc.mutation.SetShopName(s)
	return sc
}

// SetBizClass1 sets the "biz_class_1" field.
func (sc *ShopCreate) SetBizClass1(s string) *ShopCreate {
	sc.mutation.SetBizClass1(s)
	return sc
}

// SetNillableBizClass1 sets the "biz_class_1" field if the given value is not nil.
func (sc *ShopCreate) SetNillableBizClass1(s *string) *ShopCreate {
	if s != nil {
		sc.SetBizClass1(*s)
	}
	return sc
}

// SetBizClassName1 sets the "biz_class_name_1" field.
func (sc *ShopCreate) SetBizClassName1(s string) *ShopCreate {
	sc.mutation.SetBizClassName1(s)
	return sc
}

// SetNillableBizClassName1 sets the "biz_class_name_1" field if the given value is not nil.
func (sc *ShopCreate) SetNillableBizClassName1(s *string) *ShopCreate {
	if s != nil {
		sc.SetBizClassName1(*s)
	}
	return sc
}

// SetBizClass2 sets the "biz_class_2" field.
func (sc *ShopCreate) SetBizClass2(s string) *ShopCreate {
	sc.mutation.SetBizClass2(s)
	return sc
}

// SetNillableBizClass2 sets the "biz_class_2" field if the given value is not nil.
func (sc *ShopCreate) SetNillableBizClass2(s *string) *ShopCreate {
	if s != nil {
		sc.SetBizClass2(*s)
	}
	return sc
}

// SetBizClassName2 sets the "biz_class_name_2" field.
func (sc *ShopCreate) SetBizClassName2(s string) *ShopCreate {
	sc.mutation.SetBizClassName2(s)
	return sc
}

// SetNillableBizClassName2 sets the "biz_class_name_2" field if the given value is not nil.
func (sc *ShopCreate) SetNillableBizClassName2(s *string) *ShopCreate {
	if s != nil {
		sc.SetBizClassName2(*s)
	}
	return sc
}

// Mutation returns the ShopMutation object of the builder.
func (sc *ShopCreate) Mutation() *ShopMutation {
	return sc.mutation
}

// Save creates the Shop in the database.
func (sc *ShopCreate) Save(ctx context.Context) (*Shop, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShopCreate) SaveX(ctx context.Context) *Shop {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ShopCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ShopCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ShopCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := shop.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := shop.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.Code(); !ok {
		v := shop.DefaultCode()
		sc.mutation.SetCode(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShopCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Shop.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Shop.update_time"`)}
	}
	if _, ok := sc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Shop.code"`)}
	}
	if v, ok := sc.mutation.Code(); ok {
		if err := shop.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Shop.code": %w`, err)}
		}
	}
	if _, ok := sc.mutation.MallCode(); !ok {
		return &ValidationError{Name: "mall_code", err: errors.New(`ent: missing required field "Shop.mall_code"`)}
	}
	if v, ok := sc.mutation.MallCode(); ok {
		if err := shop.MallCodeValidator(v); err != nil {
			return &ValidationError{Name: "mall_code", err: fmt.Errorf(`ent: validator failed for field "Shop.mall_code": %w`, err)}
		}
	}
	if _, ok := sc.mutation.MallName(); !ok {
		return &ValidationError{Name: "mall_name", err: errors.New(`ent: missing required field "Shop.mall_name"`)}
	}
	if v, ok := sc.mutation.MallName(); ok {
		if err := shop.MallNameValidator(v); err != nil {
			return &ValidationError{Name: "mall_name", err: fmt.Errorf(`ent: validator failed for field "Shop.mall_name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.ContractCode(); !ok {
		return &ValidationError{Name: "contract_code", err: errors.New(`ent: missing required field "Shop.contract_code"`)}
	}
	if v, ok := sc.mutation.ContractCode(); ok {
		if err := shop.ContractCodeValidator(v); err != nil {
			return &ValidationError{Name: "contract_code", err: fmt.Errorf(`ent: validator failed for field "Shop.contract_code": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Floor(); !ok {
		return &ValidationError{Name: "floor", err: errors.New(`ent: missing required field "Shop.floor"`)}
	}
	if v, ok := sc.mutation.Floor(); ok {
		if err := shop.FloorValidator(v); err != nil {
			return &ValidationError{Name: "floor", err: fmt.Errorf(`ent: validator failed for field "Shop.floor": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Pos(); !ok {
		return &ValidationError{Name: "pos", err: errors.New(`ent: missing required field "Shop.pos"`)}
	}
	if v, ok := sc.mutation.Pos(); ok {
		if err := shop.PosValidator(v); err != nil {
			return &ValidationError{Name: "pos", err: fmt.Errorf(`ent: validator failed for field "Shop.pos": %w`, err)}
		}
	}
	if _, ok := sc.mutation.ShopCode(); !ok {
		return &ValidationError{Name: "shop_code", err: errors.New(`ent: missing required field "Shop.shop_code"`)}
	}
	if v, ok := sc.mutation.ShopCode(); ok {
		if err := shop.ShopCodeValidator(v); err != nil {
			return &ValidationError{Name: "shop_code", err: fmt.Errorf(`ent: validator failed for field "Shop.shop_code": %w`, err)}
		}
	}
	if _, ok := sc.mutation.ShopName(); !ok {
		return &ValidationError{Name: "shop_name", err: errors.New(`ent: missing required field "Shop.shop_name"`)}
	}
	if v, ok := sc.mutation.ShopName(); ok {
		if err := shop.ShopNameValidator(v); err != nil {
			return &ValidationError{Name: "shop_name", err: fmt.Errorf(`ent: validator failed for field "Shop.shop_name": %w`, err)}
		}
	}
	if v, ok := sc.mutation.BizClass1(); ok {
		if err := shop.BizClass1Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_1", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_1": %w`, err)}
		}
	}
	if v, ok := sc.mutation.BizClassName1(); ok {
		if err := shop.BizClassName1Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_name_1", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_name_1": %w`, err)}
		}
	}
	if v, ok := sc.mutation.BizClass2(); ok {
		if err := shop.BizClass2Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_2", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_2": %w`, err)}
		}
	}
	if v, ok := sc.mutation.BizClassName2(); ok {
		if err := shop.BizClassName2Validator(v); err != nil {
			return &ValidationError{Name: "biz_class_name_2", err: fmt.Errorf(`ent: validator failed for field "Shop.biz_class_name_2": %w`, err)}
		}
	}
	return nil
}

func (sc *ShopCreate) sqlSave(ctx context.Context) (*Shop, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ShopCreate) createSpec() (*Shop, *sqlgraph.CreateSpec) {
	var (
		_node = &Shop{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(shop.Table, sqlgraph.NewFieldSpec(shop.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(shop.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(shop.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.Code(); ok {
		_spec.SetField(shop.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := sc.mutation.MallCode(); ok {
		_spec.SetField(shop.FieldMallCode, field.TypeString, value)
		_node.MallCode = value
	}
	if value, ok := sc.mutation.MallName(); ok {
		_spec.SetField(shop.FieldMallName, field.TypeString, value)
		_node.MallName = value
	}
	if value, ok := sc.mutation.ContractCode(); ok {
		_spec.SetField(shop.FieldContractCode, field.TypeString, value)
		_node.ContractCode = value
	}
	if value, ok := sc.mutation.Floor(); ok {
		_spec.SetField(shop.FieldFloor, field.TypeString, value)
		_node.Floor = value
	}
	if value, ok := sc.mutation.Pos(); ok {
		_spec.SetField(shop.FieldPos, field.TypeString, value)
		_node.Pos = value
	}
	if value, ok := sc.mutation.ShopCode(); ok {
		_spec.SetField(shop.FieldShopCode, field.TypeString, value)
		_node.ShopCode = value
	}
	if value, ok := sc.mutation.ShopName(); ok {
		_spec.SetField(shop.FieldShopName, field.TypeString, value)
		_node.ShopName = value
	}
	if value, ok := sc.mutation.BizClass1(); ok {
		_spec.SetField(shop.FieldBizClass1, field.TypeString, value)
		_node.BizClass1 = value
	}
	if value, ok := sc.mutation.BizClassName1(); ok {
		_spec.SetField(shop.FieldBizClassName1, field.TypeString, value)
		_node.BizClassName1 = value
	}
	if value, ok := sc.mutation.BizClass2(); ok {
		_spec.SetField(shop.FieldBizClass2, field.TypeString, value)
		_node.BizClass2 = value
	}
	if value, ok := sc.mutation.BizClassName2(); ok {
		_spec.SetField(shop.FieldBizClassName2, field.TypeString, value)
		_node.BizClassName2 = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Shop.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShopUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (sc *ShopCreate) OnConflict(opts ...sql.ConflictOption) *ShopUpsertOne {
	sc.conflict = opts
	return &ShopUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *ShopCreate) OnConflictColumns(columns ...string) *ShopUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &ShopUpsertOne{
		create: sc,
	}
}

type (
	// ShopUpsertOne is the builder for "upsert"-ing
	//  one Shop node.
	ShopUpsertOne struct {
		create *ShopCreate
	}

	// ShopUpsert is the "OnConflict" setter.
	ShopUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *ShopUpsert) SetUpdateTime(v time.Time) *ShopUpsert {
	u.Set(shop.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ShopUpsert) UpdateUpdateTime() *ShopUpsert {
	u.SetExcluded(shop.FieldUpdateTime)
	return u
}

// SetBizClass1 sets the "biz_class_1" field.
func (u *ShopUpsert) SetBizClass1(v string) *ShopUpsert {
	u.Set(shop.FieldBizClass1, v)
	return u
}

// UpdateBizClass1 sets the "biz_class_1" field to the value that was provided on create.
func (u *ShopUpsert) UpdateBizClass1() *ShopUpsert {
	u.SetExcluded(shop.FieldBizClass1)
	return u
}

// ClearBizClass1 clears the value of the "biz_class_1" field.
func (u *ShopUpsert) ClearBizClass1() *ShopUpsert {
	u.SetNull(shop.FieldBizClass1)
	return u
}

// SetBizClassName1 sets the "biz_class_name_1" field.
func (u *ShopUpsert) SetBizClassName1(v string) *ShopUpsert {
	u.Set(shop.FieldBizClassName1, v)
	return u
}

// UpdateBizClassName1 sets the "biz_class_name_1" field to the value that was provided on create.
func (u *ShopUpsert) UpdateBizClassName1() *ShopUpsert {
	u.SetExcluded(shop.FieldBizClassName1)
	return u
}

// ClearBizClassName1 clears the value of the "biz_class_name_1" field.
func (u *ShopUpsert) ClearBizClassName1() *ShopUpsert {
	u.SetNull(shop.FieldBizClassName1)
	return u
}

// SetBizClass2 sets the "biz_class_2" field.
func (u *ShopUpsert) SetBizClass2(v string) *ShopUpsert {
	u.Set(shop.FieldBizClass2, v)
	return u
}

// UpdateBizClass2 sets the "biz_class_2" field to the value that was provided on create.
func (u *ShopUpsert) UpdateBizClass2() *ShopUpsert {
	u.SetExcluded(shop.FieldBizClass2)
	return u
}

// ClearBizClass2 clears the value of the "biz_class_2" field.
func (u *ShopUpsert) ClearBizClass2() *ShopUpsert {
	u.SetNull(shop.FieldBizClass2)
	return u
}

// SetBizClassName2 sets the "biz_class_name_2" field.
func (u *ShopUpsert) SetBizClassName2(v string) *ShopUpsert {
	u.Set(shop.FieldBizClassName2, v)
	return u
}

// UpdateBizClassName2 sets the "biz_class_name_2" field to the value that was provided on create.
func (u *ShopUpsert) UpdateBizClassName2() *ShopUpsert {
	u.SetExcluded(shop.FieldBizClassName2)
	return u
}

// ClearBizClassName2 clears the value of the "biz_class_name_2" field.
func (u *ShopUpsert) ClearBizClassName2() *ShopUpsert {
	u.SetNull(shop.FieldBizClassName2)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ShopUpsertOne) UpdateNewValues() *ShopUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(shop.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Code(); exists {
			s.SetIgnore(shop.FieldCode)
		}
		if _, exists := u.create.mutation.MallCode(); exists {
			s.SetIgnore(shop.FieldMallCode)
		}
		if _, exists := u.create.mutation.MallName(); exists {
			s.SetIgnore(shop.FieldMallName)
		}
		if _, exists := u.create.mutation.ContractCode(); exists {
			s.SetIgnore(shop.FieldContractCode)
		}
		if _, exists := u.create.mutation.Floor(); exists {
			s.SetIgnore(shop.FieldFloor)
		}
		if _, exists := u.create.mutation.Pos(); exists {
			s.SetIgnore(shop.FieldPos)
		}
		if _, exists := u.create.mutation.ShopCode(); exists {
			s.SetIgnore(shop.FieldShopCode)
		}
		if _, exists := u.create.mutation.ShopName(); exists {
			s.SetIgnore(shop.FieldShopName)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Shop.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ShopUpsertOne) Ignore() *ShopUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShopUpsertOne) DoNothing() *ShopUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShopCreate.OnConflict
// documentation for more info.
func (u *ShopUpsertOne) Update(set func(*ShopUpsert)) *ShopUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShopUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ShopUpsertOne) SetUpdateTime(v time.Time) *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ShopUpsertOne) UpdateUpdateTime() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetBizClass1 sets the "biz_class_1" field.
func (u *ShopUpsertOne) SetBizClass1(v string) *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.SetBizClass1(v)
	})
}

// UpdateBizClass1 sets the "biz_class_1" field to the value that was provided on create.
func (u *ShopUpsertOne) UpdateBizClass1() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateBizClass1()
	})
}

// ClearBizClass1 clears the value of the "biz_class_1" field.
func (u *ShopUpsertOne) ClearBizClass1() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.ClearBizClass1()
	})
}

// SetBizClassName1 sets the "biz_class_name_1" field.
func (u *ShopUpsertOne) SetBizClassName1(v string) *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.SetBizClassName1(v)
	})
}

// UpdateBizClassName1 sets the "biz_class_name_1" field to the value that was provided on create.
func (u *ShopUpsertOne) UpdateBizClassName1() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateBizClassName1()
	})
}

// ClearBizClassName1 clears the value of the "biz_class_name_1" field.
func (u *ShopUpsertOne) ClearBizClassName1() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.ClearBizClassName1()
	})
}

// SetBizClass2 sets the "biz_class_2" field.
func (u *ShopUpsertOne) SetBizClass2(v string) *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.SetBizClass2(v)
	})
}

// UpdateBizClass2 sets the "biz_class_2" field to the value that was provided on create.
func (u *ShopUpsertOne) UpdateBizClass2() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateBizClass2()
	})
}

// ClearBizClass2 clears the value of the "biz_class_2" field.
func (u *ShopUpsertOne) ClearBizClass2() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.ClearBizClass2()
	})
}

// SetBizClassName2 sets the "biz_class_name_2" field.
func (u *ShopUpsertOne) SetBizClassName2(v string) *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.SetBizClassName2(v)
	})
}

// UpdateBizClassName2 sets the "biz_class_name_2" field to the value that was provided on create.
func (u *ShopUpsertOne) UpdateBizClassName2() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateBizClassName2()
	})
}

// ClearBizClassName2 clears the value of the "biz_class_name_2" field.
func (u *ShopUpsertOne) ClearBizClassName2() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.ClearBizClassName2()
	})
}

// Exec executes the query.
func (u *ShopUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ShopCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShopUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ShopUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ShopUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ShopCreateBulk is the builder for creating many Shop entities in bulk.
type ShopCreateBulk struct {
	config
	err      error
	builders []*ShopCreate
	conflict []sql.ConflictOption
}

// Save creates the Shop entities in the database.
func (scb *ShopCreateBulk) Save(ctx context.Context) ([]*Shop, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Shop, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShopMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShopCreateBulk) SaveX(ctx context.Context) []*Shop {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ShopCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ShopCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Shop.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShopUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (scb *ShopCreateBulk) OnConflict(opts ...sql.ConflictOption) *ShopUpsertBulk {
	scb.conflict = opts
	return &ShopUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *ShopCreateBulk) OnConflictColumns(columns ...string) *ShopUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &ShopUpsertBulk{
		create: scb,
	}
}

// ShopUpsertBulk is the builder for "upsert"-ing
// a bulk of Shop nodes.
type ShopUpsertBulk struct {
	create *ShopCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ShopUpsertBulk) UpdateNewValues() *ShopUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(shop.FieldCreateTime)
			}
			if _, exists := b.mutation.Code(); exists {
				s.SetIgnore(shop.FieldCode)
			}
			if _, exists := b.mutation.MallCode(); exists {
				s.SetIgnore(shop.FieldMallCode)
			}
			if _, exists := b.mutation.MallName(); exists {
				s.SetIgnore(shop.FieldMallName)
			}
			if _, exists := b.mutation.ContractCode(); exists {
				s.SetIgnore(shop.FieldContractCode)
			}
			if _, exists := b.mutation.Floor(); exists {
				s.SetIgnore(shop.FieldFloor)
			}
			if _, exists := b.mutation.Pos(); exists {
				s.SetIgnore(shop.FieldPos)
			}
			if _, exists := b.mutation.ShopCode(); exists {
				s.SetIgnore(shop.FieldShopCode)
			}
			if _, exists := b.mutation.ShopName(); exists {
				s.SetIgnore(shop.FieldShopName)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ShopUpsertBulk) Ignore() *ShopUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShopUpsertBulk) DoNothing() *ShopUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShopCreateBulk.OnConflict
// documentation for more info.
func (u *ShopUpsertBulk) Update(set func(*ShopUpsert)) *ShopUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShopUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ShopUpsertBulk) SetUpdateTime(v time.Time) *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ShopUpsertBulk) UpdateUpdateTime() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetBizClass1 sets the "biz_class_1" field.
func (u *ShopUpsertBulk) SetBizClass1(v string) *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.SetBizClass1(v)
	})
}

// UpdateBizClass1 sets the "biz_class_1" field to the value that was provided on create.
func (u *ShopUpsertBulk) UpdateBizClass1() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateBizClass1()
	})
}

// ClearBizClass1 clears the value of the "biz_class_1" field.
func (u *ShopUpsertBulk) ClearBizClass1() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.ClearBizClass1()
	})
}

// SetBizClassName1 sets the "biz_class_name_1" field.
func (u *ShopUpsertBulk) SetBizClassName1(v string) *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.SetBizClassName1(v)
	})
}

// UpdateBizClassName1 sets the "biz_class_name_1" field to the value that was provided on create.
func (u *ShopUpsertBulk) UpdateBizClassName1() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateBizClassName1()
	})
}

// ClearBizClassName1 clears the value of the "biz_class_name_1" field.
func (u *ShopUpsertBulk) ClearBizClassName1() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.ClearBizClassName1()
	})
}

// SetBizClass2 sets the "biz_class_2" field.
func (u *ShopUpsertBulk) SetBizClass2(v string) *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.SetBizClass2(v)
	})
}

// UpdateBizClass2 sets the "biz_class_2" field to the value that was provided on create.
func (u *ShopUpsertBulk) UpdateBizClass2() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateBizClass2()
	})
}

// ClearBizClass2 clears the value of the "biz_class_2" field.
func (u *ShopUpsertBulk) ClearBizClass2() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.ClearBizClass2()
	})
}

// SetBizClassName2 sets the "biz_class_name_2" field.
func (u *ShopUpsertBulk) SetBizClassName2(v string) *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.SetBizClassName2(v)
	})
}

// UpdateBizClassName2 sets the "biz_class_name_2" field to the value that was provided on create.
func (u *ShopUpsertBulk) UpdateBizClassName2() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateBizClassName2()
	})
}

// ClearBizClassName2 clears the value of the "biz_class_name_2" field.
func (u *ShopUpsertBulk) ClearBizClassName2() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.ClearBizClassName2()
	})
}

// Exec executes the query.
func (u *ShopUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ShopCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ShopCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShopUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
