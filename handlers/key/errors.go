package key

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	KEY = "key"

	errFailedToListKeys         = 12
	errDuplicateName            = 13
	errFailedToCreatingMnemonic = 14
	errInvalidMnemonic          = 15
	errFailedToCreateKey        = 16
	errFailedToDeleteKey        = 17

	errMsgFailedToListKeys         = "failed to list the keys"
	errMsgDuplicateName            = "duplicate key name"
	errMsgFailedToCreatingMnemonic = "failed creating mnemonic"
	errMsgInvalidMnemonic          = "failed to invalid mnemonic"
	errMsgFailedToCreateKey        = "failed to creating key"
	errMsgFailedToDeleteKey        = "failed to deleting key"
)

func errorFailedToListKeys() *types.Error {
	return types.NewError(KEY, errFailedToListKeys, errMsgFailedToListKeys)
}
func errorDuplicateKeyName() *types.Error {
	return types.NewError(KEY, errDuplicateName, errMsgDuplicateName)
}
func errorFailedToCreateMnemonic() *types.Error {
	return types.NewError(KEY, errFailedToCreatingMnemonic, errMsgFailedToCreatingMnemonic)
}
func errorInvalidMnemonic() *types.Error {
	return types.NewError(KEY, errInvalidMnemonic, errMsgInvalidMnemonic)
}
func errorFailedToCreateKey() *types.Error {
	return types.NewError(KEY, errFailedToCreateKey, errMsgFailedToCreateKey)
}
func errorFailedToDeleteKey() *types.Error {
	return types.NewError(KEY, errFailedToDeleteKey, errMsgFailedToDeleteKey)
}
