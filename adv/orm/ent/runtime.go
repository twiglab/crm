// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/twiglab/crm/adv/orm/ent/adv"
	"github.com/twiglab/crm/adv/orm/ent/mall"
	"github.com/twiglab/crm/adv/orm/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	advMixin := schema.Adv{}.Mixin()
	advMixinFields0 := advMixin[0].Fields()
	_ = advMixinFields0
	advFields := schema.Adv{}.Fields()
	_ = advFields
	// advDescCreateTime is the schema descriptor for create_time field.
	advDescCreateTime := advMixinFields0[0].Descriptor()
	// adv.DefaultCreateTime holds the default value on creation for the create_time field.
	adv.DefaultCreateTime = advDescCreateTime.Default.(func() time.Time)
	// advDescUpdateTime is the schema descriptor for update_time field.
	advDescUpdateTime := advMixinFields0[1].Descriptor()
	// adv.DefaultUpdateTime holds the default value on creation for the update_time field.
	adv.DefaultUpdateTime = advDescUpdateTime.Default.(func() time.Time)
	// adv.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	adv.UpdateDefaultUpdateTime = advDescUpdateTime.UpdateDefault.(func() time.Time)
	// advDescCode is the schema descriptor for code field.
	advDescCode := advFields[1].Descriptor()
	// adv.DefaultCode holds the default value on creation for the code field.
	adv.DefaultCode = advDescCode.Default.(func() string)
	// adv.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	adv.CodeValidator = func() func(string) error {
		validators := advDescCode.Validators
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
	// advDescMallCode is the schema descriptor for mall_code field.
	advDescMallCode := advFields[2].Descriptor()
	// adv.MallCodeValidator is a validator for the "mall_code" field. It is called by the builders before save.
	adv.MallCodeValidator = func() func(string) error {
		validators := advDescMallCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(mall_code string) error {
			for _, fn := range fns {
				if err := fn(mall_code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// advDescMallName is the schema descriptor for mall_name field.
	advDescMallName := advFields[3].Descriptor()
	// adv.MallNameValidator is a validator for the "mall_name" field. It is called by the builders before save.
	adv.MallNameValidator = func() func(string) error {
		validators := advDescMallName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(mall_name string) error {
			for _, fn := range fns {
				if err := fn(mall_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// advDescH3Index6 is the schema descriptor for h3_index_6 field.
	advDescH3Index6 := advFields[4].Descriptor()
	// adv.H3Index6Validator is a validator for the "h3_index_6" field. It is called by the builders before save.
	adv.H3Index6Validator = advDescH3Index6.Validators[0].(func(string) error)
	// advDescH3Index7 is the schema descriptor for h3_index_7 field.
	advDescH3Index7 := advFields[5].Descriptor()
	// adv.H3Index7Validator is a validator for the "h3_index_7" field. It is called by the builders before save.
	adv.H3Index7Validator = advDescH3Index7.Validators[0].(func(string) error)
	// advDescH3Index8 is the schema descriptor for h3_index_8 field.
	advDescH3Index8 := advFields[6].Descriptor()
	// adv.H3Index8Validator is a validator for the "h3_index_8" field. It is called by the builders before save.
	adv.H3Index8Validator = advDescH3Index8.Validators[0].(func(string) error)
	// advDescImgPath is the schema descriptor for img_path field.
	advDescImgPath := advFields[7].Descriptor()
	// adv.ImgPathValidator is a validator for the "img_path" field. It is called by the builders before save.
	adv.ImgPathValidator = func() func(string) error {
		validators := advDescImgPath.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(img_path string) error {
			for _, fn := range fns {
				if err := fn(img_path); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// advDescURL is the schema descriptor for url field.
	advDescURL := advFields[8].Descriptor()
	// adv.URLValidator is a validator for the "url" field. It is called by the builders before save.
	adv.URLValidator = func() func(string) error {
		validators := advDescURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(url string) error {
			for _, fn := range fns {
				if err := fn(url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// advDescRuler is the schema descriptor for ruler field.
	advDescRuler := advFields[9].Descriptor()
	// adv.RulerValidator is a validator for the "ruler" field. It is called by the builders before save.
	adv.RulerValidator = advDescRuler.Validators[0].(func(string) error)
	// advDescOrd is the schema descriptor for ord field.
	advDescOrd := advFields[10].Descriptor()
	// adv.DefaultOrd holds the default value on creation for the ord field.
	adv.DefaultOrd = advDescOrd.Default.(int)
	// advDescMemo is the schema descriptor for memo field.
	advDescMemo := advFields[11].Descriptor()
	// adv.MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	adv.MemoValidator = advDescMemo.Validators[0].(func(string) error)
	// advDescStartTime is the schema descriptor for start_time field.
	advDescStartTime := advFields[12].Descriptor()
	// adv.DefaultStartTime holds the default value on creation for the start_time field.
	adv.DefaultStartTime = advDescStartTime.Default.(func() time.Time)
	// advDescEndTime is the schema descriptor for end_time field.
	advDescEndTime := advFields[13].Descriptor()
	// adv.DefaultEndTime holds the default value on creation for the end_time field.
	adv.DefaultEndTime = advDescEndTime.Default.(func() time.Time)
	// advDescStatus is the schema descriptor for status field.
	advDescStatus := advFields[14].Descriptor()
	// adv.DefaultStatus holds the default value on creation for the status field.
	adv.DefaultStatus = advDescStatus.Default.(int)
	// advDescID is the schema descriptor for id field.
	advDescID := advFields[0].Descriptor()
	// adv.DefaultID holds the default value on creation for the id field.
	adv.DefaultID = advDescID.Default.(func() uuid.UUID)
	mallMixin := schema.Mall{}.Mixin()
	mallMixinFields0 := mallMixin[0].Fields()
	_ = mallMixinFields0
	mallFields := schema.Mall{}.Fields()
	_ = mallFields
	// mallDescCreateTime is the schema descriptor for create_time field.
	mallDescCreateTime := mallMixinFields0[0].Descriptor()
	// mall.DefaultCreateTime holds the default value on creation for the create_time field.
	mall.DefaultCreateTime = mallDescCreateTime.Default.(func() time.Time)
	// mallDescUpdateTime is the schema descriptor for update_time field.
	mallDescUpdateTime := mallMixinFields0[1].Descriptor()
	// mall.DefaultUpdateTime holds the default value on creation for the update_time field.
	mall.DefaultUpdateTime = mallDescUpdateTime.Default.(func() time.Time)
	// mall.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	mall.UpdateDefaultUpdateTime = mallDescUpdateTime.UpdateDefault.(func() time.Time)
	// mallDescCode is the schema descriptor for code field.
	mallDescCode := mallFields[1].Descriptor()
	// mall.DefaultCode holds the default value on creation for the code field.
	mall.DefaultCode = mallDescCode.Default.(func() string)
	// mall.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	mall.CodeValidator = func() func(string) error {
		validators := mallDescCode.Validators
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
	// mallDescMallCode is the schema descriptor for mall_code field.
	mallDescMallCode := mallFields[2].Descriptor()
	// mall.MallCodeValidator is a validator for the "mall_code" field. It is called by the builders before save.
	mall.MallCodeValidator = func() func(string) error {
		validators := mallDescMallCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(mall_code string) error {
			for _, fn := range fns {
				if err := fn(mall_code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mallDescMallName is the schema descriptor for mall_name field.
	mallDescMallName := mallFields[3].Descriptor()
	// mall.MallNameValidator is a validator for the "mall_name" field. It is called by the builders before save.
	mall.MallNameValidator = func() func(string) error {
		validators := mallDescMallName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(mall_name string) error {
			for _, fn := range fns {
				if err := fn(mall_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// mallDescH3Index6 is the schema descriptor for h3_index_6 field.
	mallDescH3Index6 := mallFields[4].Descriptor()
	// mall.H3Index6Validator is a validator for the "h3_index_6" field. It is called by the builders before save.
	mall.H3Index6Validator = mallDescH3Index6.Validators[0].(func(string) error)
	// mallDescH3Index7 is the schema descriptor for h3_index_7 field.
	mallDescH3Index7 := mallFields[5].Descriptor()
	// mall.H3Index7Validator is a validator for the "h3_index_7" field. It is called by the builders before save.
	mall.H3Index7Validator = mallDescH3Index7.Validators[0].(func(string) error)
	// mallDescH3Index8 is the schema descriptor for h3_index_8 field.
	mallDescH3Index8 := mallFields[6].Descriptor()
	// mall.H3Index8Validator is a validator for the "h3_index_8" field. It is called by the builders before save.
	mall.H3Index8Validator = mallDescH3Index8.Validators[0].(func(string) error)
	// mallDescMemo is the schema descriptor for memo field.
	mallDescMemo := mallFields[7].Descriptor()
	// mall.MemoValidator is a validator for the "memo" field. It is called by the builders before save.
	mall.MemoValidator = mallDescMemo.Validators[0].(func(string) error)
	// mallDescStartTime is the schema descriptor for start_time field.
	mallDescStartTime := mallFields[8].Descriptor()
	// mall.DefaultStartTime holds the default value on creation for the start_time field.
	mall.DefaultStartTime = mallDescStartTime.Default.(func() time.Time)
	// mallDescStatus is the schema descriptor for status field.
	mallDescStatus := mallFields[9].Descriptor()
	// mall.DefaultStatus holds the default value on creation for the status field.
	mall.DefaultStatus = mallDescStatus.Default.(int)
	// mallDescID is the schema descriptor for id field.
	mallDescID := mallFields[0].Descriptor()
	// mall.DefaultID holds the default value on creation for the id field.
	mall.DefaultID = mallDescID.Default.(func() uuid.UUID)
}