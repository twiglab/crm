// Code generated by ent, DO NOT EDIT.

package mall

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/twiglab/crm/adv/orm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldUpdateTime, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldCode, v))
}

// MallCode applies equality check predicate on the "mall_code" field. It's identical to MallCodeEQ.
func MallCode(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldMallCode, v))
}

// MallName applies equality check predicate on the "mall_name" field. It's identical to MallNameEQ.
func MallName(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldMallName, v))
}

// H3Index6 applies equality check predicate on the "h3_index_6" field. It's identical to H3Index6EQ.
func H3Index6(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldH3Index6, v))
}

// H3Index7 applies equality check predicate on the "h3_index_7" field. It's identical to H3Index7EQ.
func H3Index7(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldH3Index7, v))
}

// H3Index8 applies equality check predicate on the "h3_index_8" field. It's identical to H3Index8EQ.
func H3Index8(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldH3Index8, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldMemo, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldStartTime, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldStatus, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldUpdateTime, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContainsFold(FieldCode, v))
}

// MallCodeEQ applies the EQ predicate on the "mall_code" field.
func MallCodeEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldMallCode, v))
}

// MallCodeNEQ applies the NEQ predicate on the "mall_code" field.
func MallCodeNEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldMallCode, v))
}

// MallCodeIn applies the In predicate on the "mall_code" field.
func MallCodeIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldMallCode, vs...))
}

// MallCodeNotIn applies the NotIn predicate on the "mall_code" field.
func MallCodeNotIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldMallCode, vs...))
}

// MallCodeGT applies the GT predicate on the "mall_code" field.
func MallCodeGT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldMallCode, v))
}

// MallCodeGTE applies the GTE predicate on the "mall_code" field.
func MallCodeGTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldMallCode, v))
}

// MallCodeLT applies the LT predicate on the "mall_code" field.
func MallCodeLT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldMallCode, v))
}

// MallCodeLTE applies the LTE predicate on the "mall_code" field.
func MallCodeLTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldMallCode, v))
}

// MallCodeContains applies the Contains predicate on the "mall_code" field.
func MallCodeContains(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContains(FieldMallCode, v))
}

// MallCodeHasPrefix applies the HasPrefix predicate on the "mall_code" field.
func MallCodeHasPrefix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasPrefix(FieldMallCode, v))
}

// MallCodeHasSuffix applies the HasSuffix predicate on the "mall_code" field.
func MallCodeHasSuffix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasSuffix(FieldMallCode, v))
}

// MallCodeEqualFold applies the EqualFold predicate on the "mall_code" field.
func MallCodeEqualFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEqualFold(FieldMallCode, v))
}

// MallCodeContainsFold applies the ContainsFold predicate on the "mall_code" field.
func MallCodeContainsFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContainsFold(FieldMallCode, v))
}

// MallNameEQ applies the EQ predicate on the "mall_name" field.
func MallNameEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldMallName, v))
}

// MallNameNEQ applies the NEQ predicate on the "mall_name" field.
func MallNameNEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldMallName, v))
}

// MallNameIn applies the In predicate on the "mall_name" field.
func MallNameIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldMallName, vs...))
}

// MallNameNotIn applies the NotIn predicate on the "mall_name" field.
func MallNameNotIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldMallName, vs...))
}

// MallNameGT applies the GT predicate on the "mall_name" field.
func MallNameGT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldMallName, v))
}

// MallNameGTE applies the GTE predicate on the "mall_name" field.
func MallNameGTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldMallName, v))
}

// MallNameLT applies the LT predicate on the "mall_name" field.
func MallNameLT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldMallName, v))
}

// MallNameLTE applies the LTE predicate on the "mall_name" field.
func MallNameLTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldMallName, v))
}

// MallNameContains applies the Contains predicate on the "mall_name" field.
func MallNameContains(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContains(FieldMallName, v))
}

// MallNameHasPrefix applies the HasPrefix predicate on the "mall_name" field.
func MallNameHasPrefix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasPrefix(FieldMallName, v))
}

// MallNameHasSuffix applies the HasSuffix predicate on the "mall_name" field.
func MallNameHasSuffix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasSuffix(FieldMallName, v))
}

// MallNameEqualFold applies the EqualFold predicate on the "mall_name" field.
func MallNameEqualFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEqualFold(FieldMallName, v))
}

// MallNameContainsFold applies the ContainsFold predicate on the "mall_name" field.
func MallNameContainsFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContainsFold(FieldMallName, v))
}

// H3Index6EQ applies the EQ predicate on the "h3_index_6" field.
func H3Index6EQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldH3Index6, v))
}

// H3Index6NEQ applies the NEQ predicate on the "h3_index_6" field.
func H3Index6NEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldH3Index6, v))
}

// H3Index6In applies the In predicate on the "h3_index_6" field.
func H3Index6In(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldH3Index6, vs...))
}

// H3Index6NotIn applies the NotIn predicate on the "h3_index_6" field.
func H3Index6NotIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldH3Index6, vs...))
}

// H3Index6GT applies the GT predicate on the "h3_index_6" field.
func H3Index6GT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldH3Index6, v))
}

// H3Index6GTE applies the GTE predicate on the "h3_index_6" field.
func H3Index6GTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldH3Index6, v))
}

// H3Index6LT applies the LT predicate on the "h3_index_6" field.
func H3Index6LT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldH3Index6, v))
}

// H3Index6LTE applies the LTE predicate on the "h3_index_6" field.
func H3Index6LTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldH3Index6, v))
}

// H3Index6Contains applies the Contains predicate on the "h3_index_6" field.
func H3Index6Contains(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContains(FieldH3Index6, v))
}

// H3Index6HasPrefix applies the HasPrefix predicate on the "h3_index_6" field.
func H3Index6HasPrefix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasPrefix(FieldH3Index6, v))
}

// H3Index6HasSuffix applies the HasSuffix predicate on the "h3_index_6" field.
func H3Index6HasSuffix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasSuffix(FieldH3Index6, v))
}

// H3Index6EqualFold applies the EqualFold predicate on the "h3_index_6" field.
func H3Index6EqualFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEqualFold(FieldH3Index6, v))
}

// H3Index6ContainsFold applies the ContainsFold predicate on the "h3_index_6" field.
func H3Index6ContainsFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContainsFold(FieldH3Index6, v))
}

// H3Index7EQ applies the EQ predicate on the "h3_index_7" field.
func H3Index7EQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldH3Index7, v))
}

// H3Index7NEQ applies the NEQ predicate on the "h3_index_7" field.
func H3Index7NEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldH3Index7, v))
}

// H3Index7In applies the In predicate on the "h3_index_7" field.
func H3Index7In(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldH3Index7, vs...))
}

// H3Index7NotIn applies the NotIn predicate on the "h3_index_7" field.
func H3Index7NotIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldH3Index7, vs...))
}

// H3Index7GT applies the GT predicate on the "h3_index_7" field.
func H3Index7GT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldH3Index7, v))
}

// H3Index7GTE applies the GTE predicate on the "h3_index_7" field.
func H3Index7GTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldH3Index7, v))
}

// H3Index7LT applies the LT predicate on the "h3_index_7" field.
func H3Index7LT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldH3Index7, v))
}

// H3Index7LTE applies the LTE predicate on the "h3_index_7" field.
func H3Index7LTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldH3Index7, v))
}

// H3Index7Contains applies the Contains predicate on the "h3_index_7" field.
func H3Index7Contains(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContains(FieldH3Index7, v))
}

// H3Index7HasPrefix applies the HasPrefix predicate on the "h3_index_7" field.
func H3Index7HasPrefix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasPrefix(FieldH3Index7, v))
}

// H3Index7HasSuffix applies the HasSuffix predicate on the "h3_index_7" field.
func H3Index7HasSuffix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasSuffix(FieldH3Index7, v))
}

// H3Index7EqualFold applies the EqualFold predicate on the "h3_index_7" field.
func H3Index7EqualFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEqualFold(FieldH3Index7, v))
}

// H3Index7ContainsFold applies the ContainsFold predicate on the "h3_index_7" field.
func H3Index7ContainsFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContainsFold(FieldH3Index7, v))
}

// H3Index8EQ applies the EQ predicate on the "h3_index_8" field.
func H3Index8EQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldH3Index8, v))
}

// H3Index8NEQ applies the NEQ predicate on the "h3_index_8" field.
func H3Index8NEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldH3Index8, v))
}

// H3Index8In applies the In predicate on the "h3_index_8" field.
func H3Index8In(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldH3Index8, vs...))
}

// H3Index8NotIn applies the NotIn predicate on the "h3_index_8" field.
func H3Index8NotIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldH3Index8, vs...))
}

// H3Index8GT applies the GT predicate on the "h3_index_8" field.
func H3Index8GT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldH3Index8, v))
}

// H3Index8GTE applies the GTE predicate on the "h3_index_8" field.
func H3Index8GTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldH3Index8, v))
}

// H3Index8LT applies the LT predicate on the "h3_index_8" field.
func H3Index8LT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldH3Index8, v))
}

// H3Index8LTE applies the LTE predicate on the "h3_index_8" field.
func H3Index8LTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldH3Index8, v))
}

// H3Index8Contains applies the Contains predicate on the "h3_index_8" field.
func H3Index8Contains(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContains(FieldH3Index8, v))
}

// H3Index8HasPrefix applies the HasPrefix predicate on the "h3_index_8" field.
func H3Index8HasPrefix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasPrefix(FieldH3Index8, v))
}

// H3Index8HasSuffix applies the HasSuffix predicate on the "h3_index_8" field.
func H3Index8HasSuffix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasSuffix(FieldH3Index8, v))
}

// H3Index8EqualFold applies the EqualFold predicate on the "h3_index_8" field.
func H3Index8EqualFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEqualFold(FieldH3Index8, v))
}

// H3Index8ContainsFold applies the ContainsFold predicate on the "h3_index_8" field.
func H3Index8ContainsFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContainsFold(FieldH3Index8, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.Mall {
	return predicate.Mall(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.Mall {
	return predicate.Mall(sql.FieldContainsFold(FieldMemo, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldStartTime, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Mall {
	return predicate.Mall(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Mall {
	return predicate.Mall(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Mall {
	return predicate.Mall(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Mall {
	return predicate.Mall(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Mall {
	return predicate.Mall(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Mall {
	return predicate.Mall(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Mall {
	return predicate.Mall(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Mall {
	return predicate.Mall(sql.FieldLTE(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Mall) predicate.Mall {
	return predicate.Mall(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Mall) predicate.Mall {
	return predicate.Mall(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Mall) predicate.Mall {
	return predicate.Mall(sql.NotPredicates(p))
}
