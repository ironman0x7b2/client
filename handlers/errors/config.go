package errors

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	CONFIG = "config"

	errFailedToCallUpdateHook = 6
	errFailedToSaveConfig     = 7

	errMsgFailedToCallUpdateHook = "failed to call update hook"
	errMsgFailedToSaveConfig     = "failed to to save config"
)

func ErrorFailedToCallUpdateHook() *types.Error {
	return types.NewError(CONFIG, errFailedToCallUpdateHook, errMsgFailedToCallUpdateHook)
}
func ErrorFailedToSaveConfig() *types.Error {
	return types.NewError(CONFIG, errFailedToSaveConfig, errMsgFailedToSaveConfig)
}
