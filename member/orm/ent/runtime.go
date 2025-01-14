// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/twiglab/crm/member/orm/ent/member"
	"github.com/twiglab/crm/member/orm/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	memberMixin := schema.Member{}.Mixin()
	memberMixinFields0 := memberMixin[0].Fields()
	_ = memberMixinFields0
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescCreateTime is the schema descriptor for create_time field.
	memberDescCreateTime := memberMixinFields0[0].Descriptor()
	// member.DefaultCreateTime holds the default value on creation for the create_time field.
	member.DefaultCreateTime = memberDescCreateTime.Default.(func() time.Time)
	// memberDescUpdateTime is the schema descriptor for update_time field.
	memberDescUpdateTime := memberMixinFields0[1].Descriptor()
	// member.DefaultUpdateTime holds the default value on creation for the update_time field.
	member.DefaultUpdateTime = memberDescUpdateTime.Default.(func() time.Time)
	// member.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	member.UpdateDefaultUpdateTime = memberDescUpdateTime.UpdateDefault.(func() time.Time)
	// memberDescCode is the schema descriptor for code field.
	memberDescCode := memberFields[1].Descriptor()
	// member.DefaultCode holds the default value on creation for the code field.
	member.DefaultCode = memberDescCode.Default.(func() string)
	// member.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	member.CodeValidator = func() func(string) error {
		validators := memberDescCode.Validators
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
	// memberDescPhone is the schema descriptor for phone field.
	memberDescPhone := memberFields[2].Descriptor()
	// member.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	member.PhoneValidator = func() func(string) error {
		validators := memberDescPhone.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(phone string) error {
			for _, fn := range fns {
				if err := fn(phone); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// memberDescNickname is the schema descriptor for nickname field.
	memberDescNickname := memberFields[3].Descriptor()
	// member.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	member.NicknameValidator = memberDescNickname.Validators[0].(func(string) error)
	// memberDescWxOpenID is the schema descriptor for wx_open_id field.
	memberDescWxOpenID := memberFields[4].Descriptor()
	// member.WxOpenIDValidator is a validator for the "wx_open_id" field. It is called by the builders before save.
	member.WxOpenIDValidator = memberDescWxOpenID.Validators[0].(func(string) error)
	// memberDescWxUnionID is the schema descriptor for wx_union_id field.
	memberDescWxUnionID := memberFields[5].Descriptor()
	// member.WxUnionIDValidator is a validator for the "wx_union_id" field. It is called by the builders before save.
	member.WxUnionIDValidator = memberDescWxUnionID.Validators[0].(func(string) error)
	// memberDescBcmbCode is the schema descriptor for bcmb_code field.
	memberDescBcmbCode := memberFields[6].Descriptor()
	// member.BcmbCodeValidator is a validator for the "bcmb_code" field. It is called by the builders before save.
	member.BcmbCodeValidator = memberDescBcmbCode.Validators[0].(func(string) error)
	// memberDescBcmbRegMsgID is the schema descriptor for bcmb_reg_msg_id field.
	memberDescBcmbRegMsgID := memberFields[8].Descriptor()
	// member.BcmbRegMsgIDValidator is a validator for the "bcmb_reg_msg_id" field. It is called by the builders before save.
	member.BcmbRegMsgIDValidator = memberDescBcmbRegMsgID.Validators[0].(func(string) error)
	// memberDescBcmbType is the schema descriptor for bcmb_type field.
	memberDescBcmbType := memberFields[9].Descriptor()
	// member.DefaultBcmbType holds the default value on creation for the bcmb_type field.
	member.DefaultBcmbType = memberDescBcmbType.Default.(int32)
	// memberDescLevel is the schema descriptor for level field.
	memberDescLevel := memberFields[10].Descriptor()
	// member.DefaultLevel holds the default value on creation for the level field.
	member.DefaultLevel = memberDescLevel.Default.(int32)
	// memberDescLastTime is the schema descriptor for last_time field.
	memberDescLastTime := memberFields[11].Descriptor()
	// member.DefaultLastTime holds the default value on creation for the last_time field.
	member.DefaultLastTime = memberDescLastTime.Default.(func() time.Time)
	// memberDescSource is the schema descriptor for source field.
	memberDescSource := memberFields[12].Descriptor()
	// member.DefaultSource holds the default value on creation for the source field.
	member.DefaultSource = memberDescSource.Default.(int32)
	// memberDescStatus is the schema descriptor for status field.
	memberDescStatus := memberFields[13].Descriptor()
	// member.DefaultStatus holds the default value on creation for the status field.
	member.DefaultStatus = memberDescStatus.Default.(int32)
	// memberDescID is the schema descriptor for id field.
	memberDescID := memberFields[0].Descriptor()
	// member.DefaultID holds the default value on creation for the id field.
	member.DefaultID = memberDescID.Default.(func() uuid.UUID)
}
