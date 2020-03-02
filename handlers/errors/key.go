package errors

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	KEY = "key"

	errFailedToListKeys         = 8
	errDuplicateName            = 9
	errFailedToCreatingMnemonic = 10
	errInvalidMnemonic          = 11
	errFailedToCreateKey        = 12
	errFailedToDeleteKey        = 13

	errMsgFailedToListKeys         = "failed to list the keys"
	errMsgDuplicateName            = "duplicate key name"
	errMsgFailedToCreatingMnemonic = "failed creating mnemonic"
	errMsgInvalidMnemonic          = "failed to invalid mnemonic"
	errMsgFailedToCreateKey        = "failed to creating key"
	errMsgFailedToDeleteKey        = "failed to deleting key"
)

func ErrorFailedToListKeys() *types.Error {
	return types.NewError(KEY, errFailedToListKeys, errMsgFailedToListKeys)
}
func ErrorDuplicateKeyName() *types.Error {
	return types.NewError(KEY, errDuplicateName, errMsgDuplicateName)
}
func ErrorFailedToCreateMnemonic() *types.Error {
	return types.NewError(KEY, errFailedToCreatingMnemonic, errMsgFailedToCreatingMnemonic)
}
func ErrorInvalidMnemonic() *types.Error {
	return types.NewError(KEY, errInvalidMnemonic, errMsgInvalidMnemonic)
}
func ErrorFailedToCreateKey() *types.Error {
	return types.NewError(KEY, errFailedToCreateKey, errMsgFailedToCreateKey)
}
func ErrorFailedToDeleteKey() *types.Error {
	return types.NewError(KEY, errFailedToDeleteKey, errMsgFailedToDeleteKey)
}
