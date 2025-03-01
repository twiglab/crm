// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/twiglab/crm/game/orm/schema"
	"time"

	"github.com/twiglab/crm/game/orm/ent/usersignin"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	usersigninFields := schema.UserSignIn{}.Fields()
	_ = usersigninFields
	// usersigninDescSignInTime is the schema descriptor for sign_in_time field.
	usersigninDescSignInTime := usersigninFields[1].Descriptor()
	// usersignin.DefaultSignInTime holds the default value on creation for the sign_in_time field.
	usersignin.DefaultSignInTime = usersigninDescSignInTime.Default.(func() time.Time)
	// usersigninDescPoints is the schema descriptor for points field.
	usersigninDescPoints := usersigninFields[2].Descriptor()
	// usersignin.DefaultPoints holds the default value on creation for the points field.
	usersignin.DefaultPoints = usersigninDescPoints.Default.(int)
	// usersigninDescExtraPoints is the schema descriptor for extra_points field.
	usersigninDescExtraPoints := usersigninFields[3].Descriptor()
	// usersignin.DefaultExtraPoints holds the default value on creation for the extra_points field.
	usersignin.DefaultExtraPoints = usersigninDescExtraPoints.Default.(int)
	// usersigninDescRemark is the schema descriptor for remark field.
	usersigninDescRemark := usersigninFields[4].Descriptor()
	// usersignin.DefaultRemark holds the default value on creation for the remark field.
	usersignin.DefaultRemark = usersigninDescRemark.Default.(string)
}
