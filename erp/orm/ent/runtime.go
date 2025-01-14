// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/twiglab/crm/erp/orm/ent/shop"
	"github.com/twiglab/crm/erp/orm/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	shopMixin := schema.Shop{}.Mixin()
	shopMixinFields0 := shopMixin[0].Fields()
	_ = shopMixinFields0
	shopFields := schema.Shop{}.Fields()
	_ = shopFields
	// shopDescCreateTime is the schema descriptor for create_time field.
	shopDescCreateTime := shopMixinFields0[0].Descriptor()
	// shop.DefaultCreateTime holds the default value on creation for the create_time field.
	shop.DefaultCreateTime = shopDescCreateTime.Default.(func() time.Time)
	// shopDescUpdateTime is the schema descriptor for update_time field.
	shopDescUpdateTime := shopMixinFields0[1].Descriptor()
	// shop.DefaultUpdateTime holds the default value on creation for the update_time field.
	shop.DefaultUpdateTime = shopDescUpdateTime.Default.(func() time.Time)
	// shop.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	shop.UpdateDefaultUpdateTime = shopDescUpdateTime.UpdateDefault.(func() time.Time)
	// shopDescCode is the schema descriptor for code field.
	shopDescCode := shopFields[0].Descriptor()
	// shop.DefaultCode holds the default value on creation for the code field.
	shop.DefaultCode = shopDescCode.Default.(func() string)
	// shop.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	shop.CodeValidator = func() func(string) error {
		validators := shopDescCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(code string) error {
			for _, fn := range fns {
				if err := fn(code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shopDescMallCode is the schema descriptor for mall_code field.
	shopDescMallCode := shopFields[1].Descriptor()
	// shop.MallCodeValidator is a validator for the "mall_code" field. It is called by the builders before save.
	shop.MallCodeValidator = shopDescMallCode.Validators[0].(func(string) error)
	// shopDescMallName is the schema descriptor for mall_name field.
	shopDescMallName := shopFields[2].Descriptor()
	// shop.MallNameValidator is a validator for the "mall_name" field. It is called by the builders before save.
	shop.MallNameValidator = shopDescMallName.Validators[0].(func(string) error)
	// shopDescContractCode is the schema descriptor for contract_code field.
	shopDescContractCode := shopFields[3].Descriptor()
	// shop.ContractCodeValidator is a validator for the "contract_code" field. It is called by the builders before save.
	shop.ContractCodeValidator = shopDescContractCode.Validators[0].(func(string) error)
	// shopDescFloor is the schema descriptor for floor field.
	shopDescFloor := shopFields[4].Descriptor()
	// shop.FloorValidator is a validator for the "floor" field. It is called by the builders before save.
	shop.FloorValidator = shopDescFloor.Validators[0].(func(string) error)
	// shopDescPos is the schema descriptor for pos field.
	shopDescPos := shopFields[5].Descriptor()
	// shop.PosValidator is a validator for the "pos" field. It is called by the builders before save.
	shop.PosValidator = shopDescPos.Validators[0].(func(string) error)
	// shopDescShopCode is the schema descriptor for shop_code field.
	shopDescShopCode := shopFields[6].Descriptor()
	// shop.ShopCodeValidator is a validator for the "shop_code" field. It is called by the builders before save.
	shop.ShopCodeValidator = shopDescShopCode.Validators[0].(func(string) error)
	// shopDescShopName is the schema descriptor for shop_name field.
	shopDescShopName := shopFields[7].Descriptor()
	// shop.ShopNameValidator is a validator for the "shop_name" field. It is called by the builders before save.
	shop.ShopNameValidator = shopDescShopName.Validators[0].(func(string) error)
	// shopDescBizClass1 is the schema descriptor for biz_class_1 field.
	shopDescBizClass1 := shopFields[8].Descriptor()
	// shop.BizClass1Validator is a validator for the "biz_class_1" field. It is called by the builders before save.
	shop.BizClass1Validator = shopDescBizClass1.Validators[0].(func(string) error)
	// shopDescBizClassName1 is the schema descriptor for biz_class_name_1 field.
	shopDescBizClassName1 := shopFields[9].Descriptor()
	// shop.BizClassName1Validator is a validator for the "biz_class_name_1" field. It is called by the builders before save.
	shop.BizClassName1Validator = shopDescBizClassName1.Validators[0].(func(string) error)
	// shopDescBizClass2 is the schema descriptor for biz_class_2 field.
	shopDescBizClass2 := shopFields[10].Descriptor()
	// shop.BizClass2Validator is a validator for the "biz_class_2" field. It is called by the builders before save.
	shop.BizClass2Validator = shopDescBizClass2.Validators[0].(func(string) error)
	// shopDescBizClassName2 is the schema descriptor for biz_class_name_2 field.
	shopDescBizClassName2 := shopFields[11].Descriptor()
	// shop.BizClassName2Validator is a validator for the "biz_class_name_2" field. It is called by the builders before save.
	shop.BizClassName2Validator = shopDescBizClassName2.Validators[0].(func(string) error)
}
