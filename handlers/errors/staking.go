package errors

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	STAKING = "staking"

	errInvalidExplorer              = 8
	errFailedToGetValidator         = 8
	errFailedToUnmarshallValidator  = 8
	errFailedToUnmarshallValidators = 8

	errMsgInvalidExplorer              = "invalid explorer address"
	errMsgFailedToGetValidator         = "failed to get validator"
	errMsgFailedToUnmarshallValidator  = "failed to unmarshall validator"
	errMsgFailedToUnmarshallValidators = "failed to unmarshall validators"
)

func ErrorInvalidExplorerAddress() *types.Error {
	return types.NewError(STAKING, errInvalidExplorer, errMsgInvalidExplorer)
}
func ErrorFailedToGetValidator() *types.Error {
	return types.NewError(STAKING, errFailedToGetValidator, errMsgFailedToGetValidator)
}
func ErrorFailedToUnmarshallValidator() *types.Error {
	return types.NewError(STAKING, errFailedToUnmarshallValidator, errMsgFailedToUnmarshallValidator)
}
func ErrorFailedToUnmarshallValidators() *types.Error {
	return types.NewError(STAKING, errFailedToUnmarshallValidators, errMsgFailedToUnmarshallValidators)
}
