// Code generated by ent, DO NOT EDIT.

package poly

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/twiglab/crm/poly/orm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldCode, v))
}

// MallCode applies equality check predicate on the "mall_code" field. It's identical to MallCodeEQ.
func MallCode(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldMallCode, v))
}

// Operator applies equality check predicate on the "operator" field. It's identical to OperatorEQ.
func Operator(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldOperator, v))
}

// AddTime applies equality check predicate on the "add_time" field. It's identical to AddTimeEQ.
func AddTime(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldAddTime, v))
}

// RuleCode applies equality check predicate on the "rule_code" field. It's identical to RuleCodeEQ.
func RuleCode(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldRuleCode, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldName, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldDesc, v))
}

// Budget applies equality check predicate on the "budget" field. It's identical to BudgetEQ.
func Budget(v int64) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldBudget, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldEndTime, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldStatus, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldType, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContainsFold(FieldCode, v))
}

// MallCodeEQ applies the EQ predicate on the "mall_code" field.
func MallCodeEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldMallCode, v))
}

// MallCodeNEQ applies the NEQ predicate on the "mall_code" field.
func MallCodeNEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldMallCode, v))
}

// MallCodeIn applies the In predicate on the "mall_code" field.
func MallCodeIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldMallCode, vs...))
}

// MallCodeNotIn applies the NotIn predicate on the "mall_code" field.
func MallCodeNotIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldMallCode, vs...))
}

// MallCodeGT applies the GT predicate on the "mall_code" field.
func MallCodeGT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldMallCode, v))
}

// MallCodeGTE applies the GTE predicate on the "mall_code" field.
func MallCodeGTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldMallCode, v))
}

// MallCodeLT applies the LT predicate on the "mall_code" field.
func MallCodeLT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldMallCode, v))
}

// MallCodeLTE applies the LTE predicate on the "mall_code" field.
func MallCodeLTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldMallCode, v))
}

// MallCodeContains applies the Contains predicate on the "mall_code" field.
func MallCodeContains(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContains(FieldMallCode, v))
}

// MallCodeHasPrefix applies the HasPrefix predicate on the "mall_code" field.
func MallCodeHasPrefix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasPrefix(FieldMallCode, v))
}

// MallCodeHasSuffix applies the HasSuffix predicate on the "mall_code" field.
func MallCodeHasSuffix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasSuffix(FieldMallCode, v))
}

// MallCodeEqualFold applies the EqualFold predicate on the "mall_code" field.
func MallCodeEqualFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEqualFold(FieldMallCode, v))
}

// MallCodeContainsFold applies the ContainsFold predicate on the "mall_code" field.
func MallCodeContainsFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContainsFold(FieldMallCode, v))
}

// OperatorEQ applies the EQ predicate on the "operator" field.
func OperatorEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldOperator, v))
}

// OperatorNEQ applies the NEQ predicate on the "operator" field.
func OperatorNEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldOperator, v))
}

// OperatorIn applies the In predicate on the "operator" field.
func OperatorIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldOperator, vs...))
}

// OperatorNotIn applies the NotIn predicate on the "operator" field.
func OperatorNotIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldOperator, vs...))
}

// OperatorGT applies the GT predicate on the "operator" field.
func OperatorGT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldOperator, v))
}

// OperatorGTE applies the GTE predicate on the "operator" field.
func OperatorGTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldOperator, v))
}

// OperatorLT applies the LT predicate on the "operator" field.
func OperatorLT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldOperator, v))
}

// OperatorLTE applies the LTE predicate on the "operator" field.
func OperatorLTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldOperator, v))
}

// OperatorContains applies the Contains predicate on the "operator" field.
func OperatorContains(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContains(FieldOperator, v))
}

// OperatorHasPrefix applies the HasPrefix predicate on the "operator" field.
func OperatorHasPrefix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasPrefix(FieldOperator, v))
}

// OperatorHasSuffix applies the HasSuffix predicate on the "operator" field.
func OperatorHasSuffix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasSuffix(FieldOperator, v))
}

// OperatorEqualFold applies the EqualFold predicate on the "operator" field.
func OperatorEqualFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEqualFold(FieldOperator, v))
}

// OperatorContainsFold applies the ContainsFold predicate on the "operator" field.
func OperatorContainsFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContainsFold(FieldOperator, v))
}

// AddTimeEQ applies the EQ predicate on the "add_time" field.
func AddTimeEQ(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldAddTime, v))
}

// AddTimeNEQ applies the NEQ predicate on the "add_time" field.
func AddTimeNEQ(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldAddTime, v))
}

// AddTimeIn applies the In predicate on the "add_time" field.
func AddTimeIn(vs ...time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldAddTime, vs...))
}

// AddTimeNotIn applies the NotIn predicate on the "add_time" field.
func AddTimeNotIn(vs ...time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldAddTime, vs...))
}

// AddTimeGT applies the GT predicate on the "add_time" field.
func AddTimeGT(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldAddTime, v))
}

// AddTimeGTE applies the GTE predicate on the "add_time" field.
func AddTimeGTE(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldAddTime, v))
}

// AddTimeLT applies the LT predicate on the "add_time" field.
func AddTimeLT(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldAddTime, v))
}

// AddTimeLTE applies the LTE predicate on the "add_time" field.
func AddTimeLTE(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldAddTime, v))
}

// RuleCodeEQ applies the EQ predicate on the "rule_code" field.
func RuleCodeEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldRuleCode, v))
}

// RuleCodeNEQ applies the NEQ predicate on the "rule_code" field.
func RuleCodeNEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldRuleCode, v))
}

// RuleCodeIn applies the In predicate on the "rule_code" field.
func RuleCodeIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldRuleCode, vs...))
}

// RuleCodeNotIn applies the NotIn predicate on the "rule_code" field.
func RuleCodeNotIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldRuleCode, vs...))
}

// RuleCodeGT applies the GT predicate on the "rule_code" field.
func RuleCodeGT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldRuleCode, v))
}

// RuleCodeGTE applies the GTE predicate on the "rule_code" field.
func RuleCodeGTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldRuleCode, v))
}

// RuleCodeLT applies the LT predicate on the "rule_code" field.
func RuleCodeLT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldRuleCode, v))
}

// RuleCodeLTE applies the LTE predicate on the "rule_code" field.
func RuleCodeLTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldRuleCode, v))
}

// RuleCodeContains applies the Contains predicate on the "rule_code" field.
func RuleCodeContains(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContains(FieldRuleCode, v))
}

// RuleCodeHasPrefix applies the HasPrefix predicate on the "rule_code" field.
func RuleCodeHasPrefix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasPrefix(FieldRuleCode, v))
}

// RuleCodeHasSuffix applies the HasSuffix predicate on the "rule_code" field.
func RuleCodeHasSuffix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasSuffix(FieldRuleCode, v))
}

// RuleCodeEqualFold applies the EqualFold predicate on the "rule_code" field.
func RuleCodeEqualFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEqualFold(FieldRuleCode, v))
}

// RuleCodeContainsFold applies the ContainsFold predicate on the "rule_code" field.
func RuleCodeContainsFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContainsFold(FieldRuleCode, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContainsFold(FieldName, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Poly {
	return predicate.Poly(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Poly {
	return predicate.Poly(sql.FieldContainsFold(FieldDesc, v))
}

// BudgetEQ applies the EQ predicate on the "budget" field.
func BudgetEQ(v int64) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldBudget, v))
}

// BudgetNEQ applies the NEQ predicate on the "budget" field.
func BudgetNEQ(v int64) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldBudget, v))
}

// BudgetIn applies the In predicate on the "budget" field.
func BudgetIn(vs ...int64) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldBudget, vs...))
}

// BudgetNotIn applies the NotIn predicate on the "budget" field.
func BudgetNotIn(vs ...int64) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldBudget, vs...))
}

// BudgetGT applies the GT predicate on the "budget" field.
func BudgetGT(v int64) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldBudget, v))
}

// BudgetGTE applies the GTE predicate on the "budget" field.
func BudgetGTE(v int64) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldBudget, v))
}

// BudgetLT applies the LT predicate on the "budget" field.
func BudgetLT(v int64) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldBudget, v))
}

// BudgetLTE applies the LTE predicate on the "budget" field.
func BudgetLTE(v int64) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldBudget, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldStartTime, v))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldEndTime, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldStatus, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.Poly {
	return predicate.Poly(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.Poly {
	return predicate.Poly(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.Poly {
	return predicate.Poly(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.Poly {
	return predicate.Poly(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.Poly {
	return predicate.Poly(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.Poly {
	return predicate.Poly(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.Poly {
	return predicate.Poly(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.Poly {
	return predicate.Poly(sql.FieldLTE(FieldType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Poly) predicate.Poly {
	return predicate.Poly(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Poly) predicate.Poly {
	return predicate.Poly(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Poly) predicate.Poly {
	return predicate.Poly(sql.NotPredicates(p))
}