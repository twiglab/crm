// Code generated by ent, DO NOT EDIT.

package bonusitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the bonusitem type in the database.
	Label = "bonus_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldMchID holds the string denoting the mch_id field in the database.
	FieldMchID = "mch_id"
	// FieldMallCode holds the string denoting the mall_code field in the database.
	FieldMallCode = "mall_code"
	// FieldMallName holds the string denoting the mall_name field in the database.
	FieldMallName = "mall_name"
	// FieldShopCode holds the string denoting the shop_code field in the database.
	FieldShopCode = "shop_code"
	// FieldShopName holds the string denoting the shop_name field in the database.
	FieldShopName = "shop_name"
	// FieldMemberCode holds the string denoting the member_code field in the database.
	FieldMemberCode = "member_code"
	// FieldWxOpenID holds the string denoting the wx_open_id field in the database.
	FieldWxOpenID = "wx_open_id"
	// FieldBcmbNotifyTime holds the string denoting the bcmb_notify_time field in the database.
	FieldBcmbNotifyTime = "bcmb_notify_time"
	// FieldBcmbNotifyID holds the string denoting the bcmb_notify_id field in the database.
	FieldBcmbNotifyID = "bcmb_notify_id"
	// FieldBcmbTransCode holds the string denoting the bcmb_trans_code field in the database.
	FieldBcmbTransCode = "bcmb_trans_code"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldBcmbTransTime holds the string denoting the bcmb_trans_time field in the database.
	FieldBcmbTransTime = "bcmb_trans_time"
	// FieldCreateTs holds the string denoting the create_ts field in the database.
	FieldCreateTs = "create_ts"
	// FieldBcmbTransPayCode holds the string denoting the bcmb_trans_pay_code field in the database.
	FieldBcmbTransPayCode = "bcmb_trans_pay_code"
	// FieldBcmbTransType holds the string denoting the bcmb_trans_type field in the database.
	FieldBcmbTransType = "bcmb_trans_type"
	// FieldBonus holds the string denoting the bonus field in the database.
	FieldBonus = "bonus"
	// FieldBonusRate holds the string denoting the bonus_rate field in the database.
	FieldBonusRate = "bonus_rate"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the bonusitem in the database.
	Table = "t_bonus_list"
)

// Columns holds all SQL columns for bonusitem fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCode,
	FieldMchID,
	FieldMallCode,
	FieldMallName,
	FieldShopCode,
	FieldShopName,
	FieldMemberCode,
	FieldWxOpenID,
	FieldBcmbNotifyTime,
	FieldBcmbNotifyID,
	FieldBcmbTransCode,
	FieldAmount,
	FieldBcmbTransTime,
	FieldCreateTs,
	FieldBcmbTransPayCode,
	FieldBcmbTransType,
	FieldBonus,
	FieldBonusRate,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultCode holds the default value on creation for the "code" field.
	DefaultCode func() string
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// MchIDValidator is a validator for the "mch_id" field. It is called by the builders before save.
	MchIDValidator func(string) error
	// MallCodeValidator is a validator for the "mall_code" field. It is called by the builders before save.
	MallCodeValidator func(string) error
	// MallNameValidator is a validator for the "mall_name" field. It is called by the builders before save.
	MallNameValidator func(string) error
	// ShopCodeValidator is a validator for the "shop_code" field. It is called by the builders before save.
	ShopCodeValidator func(string) error
	// ShopNameValidator is a validator for the "shop_name" field. It is called by the builders before save.
	ShopNameValidator func(string) error
	// MemberCodeValidator is a validator for the "member_code" field. It is called by the builders before save.
	MemberCodeValidator func(string) error
	// WxOpenIDValidator is a validator for the "wx_open_id" field. It is called by the builders before save.
	WxOpenIDValidator func(string) error
	// DefaultBcmbNotifyTime holds the default value on creation for the "bcmb_notify_time" field.
	DefaultBcmbNotifyTime func() time.Time
	// BcmbNotifyIDValidator is a validator for the "bcmb_notify_id" field. It is called by the builders before save.
	BcmbNotifyIDValidator func(string) error
	// BcmbTransCodeValidator is a validator for the "bcmb_trans_code" field. It is called by the builders before save.
	BcmbTransCodeValidator func(string) error
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int
	// DefaultBcmbTransTime holds the default value on creation for the "bcmb_trans_time" field.
	DefaultBcmbTransTime func() time.Time
	// DefaultCreateTs holds the default value on creation for the "create_ts" field.
	DefaultCreateTs int64
	// BcmbTransPayCodeValidator is a validator for the "bcmb_trans_pay_code" field. It is called by the builders before save.
	BcmbTransPayCodeValidator func(string) error
	// DefaultBcmbTransType holds the default value on creation for the "bcmb_trans_type" field.
	DefaultBcmbTransType int
	// DefaultBonus holds the default value on creation for the "bonus" field.
	DefaultBonus int
	// DefaultBonusRate holds the default value on creation for the "bonus_rate" field.
	DefaultBonusRate int
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
)

// OrderOption defines the ordering options for the BonusItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByMchID orders the results by the mch_id field.
func ByMchID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMchID, opts...).ToFunc()
}

// ByMallCode orders the results by the mall_code field.
func ByMallCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMallCode, opts...).ToFunc()
}

// ByMallName orders the results by the mall_name field.
func ByMallName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMallName, opts...).ToFunc()
}

// ByShopCode orders the results by the shop_code field.
func ByShopCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShopCode, opts...).ToFunc()
}

// ByShopName orders the results by the shop_name field.
func ByShopName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShopName, opts...).ToFunc()
}

// ByMemberCode orders the results by the member_code field.
func ByMemberCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberCode, opts...).ToFunc()
}

// ByWxOpenID orders the results by the wx_open_id field.
func ByWxOpenID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWxOpenID, opts...).ToFunc()
}

// ByBcmbNotifyTime orders the results by the bcmb_notify_time field.
func ByBcmbNotifyTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbNotifyTime, opts...).ToFunc()
}

// ByBcmbNotifyID orders the results by the bcmb_notify_id field.
func ByBcmbNotifyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbNotifyID, opts...).ToFunc()
}

// ByBcmbTransCode orders the results by the bcmb_trans_code field.
func ByBcmbTransCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbTransCode, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByBcmbTransTime orders the results by the bcmb_trans_time field.
func ByBcmbTransTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbTransTime, opts...).ToFunc()
}

// ByCreateTs orders the results by the create_ts field.
func ByCreateTs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTs, opts...).ToFunc()
}

// ByBcmbTransPayCode orders the results by the bcmb_trans_pay_code field.
func ByBcmbTransPayCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbTransPayCode, opts...).ToFunc()
}

// ByBcmbTransType orders the results by the bcmb_trans_type field.
func ByBcmbTransType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBcmbTransType, opts...).ToFunc()
}

// ByBonus orders the results by the bonus field.
func ByBonus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBonus, opts...).ToFunc()
}

// ByBonusRate orders the results by the bonus_rate field.
func ByBonusRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBonusRate, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}