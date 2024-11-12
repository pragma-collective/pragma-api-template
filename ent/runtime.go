// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/pragma-collective/0xStarter-api/ent/schema"
	"github.com/pragma-collective/0xStarter-api/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescWalletAddress is the schema descriptor for wallet_address field.
	userDescWalletAddress := userFields[0].Descriptor()
	// user.WalletAddressValidator is a validator for the "wallet_address" field. It is called by the builders before save.
	user.WalletAddressValidator = userDescWalletAddress.Validators[0].(func(string) error)
}
