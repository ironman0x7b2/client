package config

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	CONFIG = "config"

	errFailedToCallUpdateHook = 12
	errFailedToSaveConfig     = 13

	errMsgFailedToCallUpdateHook = "failed to call update hook"
	errMsgFailedToSaveConfig     = "failed to save config"
)

func errorFailedToCallUpdateHook() *types.Error {
	return types.NewError(CONFIG, errFailedToCallUpdateHook, errMsgFailedToCallUpdateHook)
}
func errorFailedToSaveConfig() *types.Error {
	return types.NewError(CONFIG, errFailedToSaveConfig, errMsgFailedToSaveConfig)
}
