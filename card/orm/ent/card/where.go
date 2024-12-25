// Code generated by ent, DO NOT EDIT.

package card

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/twiglab/crm/card/orm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldUpdateTime, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCode, v))
}

// CodeBin applies equality check predicate on the "code_bin" field. It's identical to CodeBinEQ.
func CodeBin(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCodeBin, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldType, v))
}

// Pic1 applies equality check predicate on the "pic1" field. It's identical to Pic1EQ.
func Pic1(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldPic1, v))
}

// Pic2 applies equality check predicate on the "pic2" field. It's identical to Pic2EQ.
func Pic2(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldPic2, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v int64) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldAmount, v))
}

// MemberCode applies equality check predicate on the "member_code" field. It's identical to MemberCodeEQ.
func MemberCode(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldMemberCode, v))
}

// BindTime applies equality check predicate on the "bind_time" field. It's identical to BindTimeEQ.
func BindTime(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldBindTime, v))
}

// LastCleanBalance applies equality check predicate on the "last_clean_balance" field. It's identical to LastCleanBalanceEQ.
func LastCleanBalance(v int64) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldLastCleanBalance, v))
}

// LastCleanTs applies equality check predicate on the "last_clean_ts" field. It's identical to LastCleanTsEQ.
func LastCleanTs(v int16) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldLastCleanTs, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldStatus, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldUpdateTime, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldCode, v))
}

// CodeBinEQ applies the EQ predicate on the "code_bin" field.
func CodeBinEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldCodeBin, v))
}

// CodeBinNEQ applies the NEQ predicate on the "code_bin" field.
func CodeBinNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldCodeBin, v))
}

// CodeBinIn applies the In predicate on the "code_bin" field.
func CodeBinIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldCodeBin, vs...))
}

// CodeBinNotIn applies the NotIn predicate on the "code_bin" field.
func CodeBinNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldCodeBin, vs...))
}

// CodeBinGT applies the GT predicate on the "code_bin" field.
func CodeBinGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldCodeBin, v))
}

// CodeBinGTE applies the GTE predicate on the "code_bin" field.
func CodeBinGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldCodeBin, v))
}

// CodeBinLT applies the LT predicate on the "code_bin" field.
func CodeBinLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldCodeBin, v))
}

// CodeBinLTE applies the LTE predicate on the "code_bin" field.
func CodeBinLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldCodeBin, v))
}

// CodeBinContains applies the Contains predicate on the "code_bin" field.
func CodeBinContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldCodeBin, v))
}

// CodeBinHasPrefix applies the HasPrefix predicate on the "code_bin" field.
func CodeBinHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldCodeBin, v))
}

// CodeBinHasSuffix applies the HasSuffix predicate on the "code_bin" field.
func CodeBinHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldCodeBin, v))
}

// CodeBinEqualFold applies the EqualFold predicate on the "code_bin" field.
func CodeBinEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldCodeBin, v))
}

// CodeBinContainsFold applies the ContainsFold predicate on the "code_bin" field.
func CodeBinContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldCodeBin, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldType, v))
}

// Pic1EQ applies the EQ predicate on the "pic1" field.
func Pic1EQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldPic1, v))
}

// Pic1NEQ applies the NEQ predicate on the "pic1" field.
func Pic1NEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldPic1, v))
}

// Pic1In applies the In predicate on the "pic1" field.
func Pic1In(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldPic1, vs...))
}

// Pic1NotIn applies the NotIn predicate on the "pic1" field.
func Pic1NotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldPic1, vs...))
}

// Pic1GT applies the GT predicate on the "pic1" field.
func Pic1GT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldPic1, v))
}

// Pic1GTE applies the GTE predicate on the "pic1" field.
func Pic1GTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldPic1, v))
}

// Pic1LT applies the LT predicate on the "pic1" field.
func Pic1LT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldPic1, v))
}

// Pic1LTE applies the LTE predicate on the "pic1" field.
func Pic1LTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldPic1, v))
}

// Pic1Contains applies the Contains predicate on the "pic1" field.
func Pic1Contains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldPic1, v))
}

// Pic1HasPrefix applies the HasPrefix predicate on the "pic1" field.
func Pic1HasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldPic1, v))
}

// Pic1HasSuffix applies the HasSuffix predicate on the "pic1" field.
func Pic1HasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldPic1, v))
}

// Pic1EqualFold applies the EqualFold predicate on the "pic1" field.
func Pic1EqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldPic1, v))
}

// Pic1ContainsFold applies the ContainsFold predicate on the "pic1" field.
func Pic1ContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldPic1, v))
}

// Pic2EQ applies the EQ predicate on the "pic2" field.
func Pic2EQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldPic2, v))
}

// Pic2NEQ applies the NEQ predicate on the "pic2" field.
func Pic2NEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldPic2, v))
}

// Pic2In applies the In predicate on the "pic2" field.
func Pic2In(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldPic2, vs...))
}

// Pic2NotIn applies the NotIn predicate on the "pic2" field.
func Pic2NotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldPic2, vs...))
}

// Pic2GT applies the GT predicate on the "pic2" field.
func Pic2GT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldPic2, v))
}

// Pic2GTE applies the GTE predicate on the "pic2" field.
func Pic2GTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldPic2, v))
}

// Pic2LT applies the LT predicate on the "pic2" field.
func Pic2LT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldPic2, v))
}

// Pic2LTE applies the LTE predicate on the "pic2" field.
func Pic2LTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldPic2, v))
}

// Pic2Contains applies the Contains predicate on the "pic2" field.
func Pic2Contains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldPic2, v))
}

// Pic2HasPrefix applies the HasPrefix predicate on the "pic2" field.
func Pic2HasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldPic2, v))
}

// Pic2HasSuffix applies the HasSuffix predicate on the "pic2" field.
func Pic2HasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldPic2, v))
}

// Pic2EqualFold applies the EqualFold predicate on the "pic2" field.
func Pic2EqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldPic2, v))
}

// Pic2ContainsFold applies the ContainsFold predicate on the "pic2" field.
func Pic2ContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldPic2, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v int64) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v int64) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...int64) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...int64) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v int64) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v int64) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v int64) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v int64) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldAmount, v))
}

// MemberCodeEQ applies the EQ predicate on the "member_code" field.
func MemberCodeEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldMemberCode, v))
}

// MemberCodeNEQ applies the NEQ predicate on the "member_code" field.
func MemberCodeNEQ(v string) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldMemberCode, v))
}

// MemberCodeIn applies the In predicate on the "member_code" field.
func MemberCodeIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldMemberCode, vs...))
}

// MemberCodeNotIn applies the NotIn predicate on the "member_code" field.
func MemberCodeNotIn(vs ...string) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldMemberCode, vs...))
}

// MemberCodeGT applies the GT predicate on the "member_code" field.
func MemberCodeGT(v string) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldMemberCode, v))
}

// MemberCodeGTE applies the GTE predicate on the "member_code" field.
func MemberCodeGTE(v string) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldMemberCode, v))
}

// MemberCodeLT applies the LT predicate on the "member_code" field.
func MemberCodeLT(v string) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldMemberCode, v))
}

// MemberCodeLTE applies the LTE predicate on the "member_code" field.
func MemberCodeLTE(v string) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldMemberCode, v))
}

// MemberCodeContains applies the Contains predicate on the "member_code" field.
func MemberCodeContains(v string) predicate.Card {
	return predicate.Card(sql.FieldContains(FieldMemberCode, v))
}

// MemberCodeHasPrefix applies the HasPrefix predicate on the "member_code" field.
func MemberCodeHasPrefix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasPrefix(FieldMemberCode, v))
}

// MemberCodeHasSuffix applies the HasSuffix predicate on the "member_code" field.
func MemberCodeHasSuffix(v string) predicate.Card {
	return predicate.Card(sql.FieldHasSuffix(FieldMemberCode, v))
}

// MemberCodeIsNil applies the IsNil predicate on the "member_code" field.
func MemberCodeIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldMemberCode))
}

// MemberCodeNotNil applies the NotNil predicate on the "member_code" field.
func MemberCodeNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldMemberCode))
}

// MemberCodeEqualFold applies the EqualFold predicate on the "member_code" field.
func MemberCodeEqualFold(v string) predicate.Card {
	return predicate.Card(sql.FieldEqualFold(FieldMemberCode, v))
}

// MemberCodeContainsFold applies the ContainsFold predicate on the "member_code" field.
func MemberCodeContainsFold(v string) predicate.Card {
	return predicate.Card(sql.FieldContainsFold(FieldMemberCode, v))
}

// BindTimeEQ applies the EQ predicate on the "bind_time" field.
func BindTimeEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldBindTime, v))
}

// BindTimeNEQ applies the NEQ predicate on the "bind_time" field.
func BindTimeNEQ(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldBindTime, v))
}

// BindTimeIn applies the In predicate on the "bind_time" field.
func BindTimeIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldBindTime, vs...))
}

// BindTimeNotIn applies the NotIn predicate on the "bind_time" field.
func BindTimeNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldBindTime, vs...))
}

// BindTimeGT applies the GT predicate on the "bind_time" field.
func BindTimeGT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldBindTime, v))
}

// BindTimeGTE applies the GTE predicate on the "bind_time" field.
func BindTimeGTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldBindTime, v))
}

// BindTimeLT applies the LT predicate on the "bind_time" field.
func BindTimeLT(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldBindTime, v))
}

// BindTimeLTE applies the LTE predicate on the "bind_time" field.
func BindTimeLTE(v time.Time) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldBindTime, v))
}

// BindTimeIsNil applies the IsNil predicate on the "bind_time" field.
func BindTimeIsNil() predicate.Card {
	return predicate.Card(sql.FieldIsNull(FieldBindTime))
}

// BindTimeNotNil applies the NotNil predicate on the "bind_time" field.
func BindTimeNotNil() predicate.Card {
	return predicate.Card(sql.FieldNotNull(FieldBindTime))
}

// LastCleanBalanceEQ applies the EQ predicate on the "last_clean_balance" field.
func LastCleanBalanceEQ(v int64) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldLastCleanBalance, v))
}

// LastCleanBalanceNEQ applies the NEQ predicate on the "last_clean_balance" field.
func LastCleanBalanceNEQ(v int64) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldLastCleanBalance, v))
}

// LastCleanBalanceIn applies the In predicate on the "last_clean_balance" field.
func LastCleanBalanceIn(vs ...int64) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldLastCleanBalance, vs...))
}

// LastCleanBalanceNotIn applies the NotIn predicate on the "last_clean_balance" field.
func LastCleanBalanceNotIn(vs ...int64) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldLastCleanBalance, vs...))
}

// LastCleanBalanceGT applies the GT predicate on the "last_clean_balance" field.
func LastCleanBalanceGT(v int64) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldLastCleanBalance, v))
}

// LastCleanBalanceGTE applies the GTE predicate on the "last_clean_balance" field.
func LastCleanBalanceGTE(v int64) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldLastCleanBalance, v))
}

// LastCleanBalanceLT applies the LT predicate on the "last_clean_balance" field.
func LastCleanBalanceLT(v int64) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldLastCleanBalance, v))
}

// LastCleanBalanceLTE applies the LTE predicate on the "last_clean_balance" field.
func LastCleanBalanceLTE(v int64) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldLastCleanBalance, v))
}

// LastCleanTsEQ applies the EQ predicate on the "last_clean_ts" field.
func LastCleanTsEQ(v int16) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldLastCleanTs, v))
}

// LastCleanTsNEQ applies the NEQ predicate on the "last_clean_ts" field.
func LastCleanTsNEQ(v int16) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldLastCleanTs, v))
}

// LastCleanTsIn applies the In predicate on the "last_clean_ts" field.
func LastCleanTsIn(vs ...int16) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldLastCleanTs, vs...))
}

// LastCleanTsNotIn applies the NotIn predicate on the "last_clean_ts" field.
func LastCleanTsNotIn(vs ...int16) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldLastCleanTs, vs...))
}

// LastCleanTsGT applies the GT predicate on the "last_clean_ts" field.
func LastCleanTsGT(v int16) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldLastCleanTs, v))
}

// LastCleanTsGTE applies the GTE predicate on the "last_clean_ts" field.
func LastCleanTsGTE(v int16) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldLastCleanTs, v))
}

// LastCleanTsLT applies the LT predicate on the "last_clean_ts" field.
func LastCleanTsLT(v int16) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldLastCleanTs, v))
}

// LastCleanTsLTE applies the LTE predicate on the "last_clean_ts" field.
func LastCleanTsLTE(v int16) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldLastCleanTs, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Card {
	return predicate.Card(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Card {
	return predicate.Card(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Card {
	return predicate.Card(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Card {
	return predicate.Card(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Card {
	return predicate.Card(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Card {
	return predicate.Card(sql.FieldLTE(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(sql.NotPredicates(p))
}
