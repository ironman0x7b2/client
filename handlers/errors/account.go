package errors

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	ACCOUNT = "account"

	errQueryAccount                   = 6
	errUnmarshallAccount              = 6
	errFaildToGetDelegatorDelegations = 7
	errFaildToGetDelegatorValidators  = 8

	errMsgQueryAccount                   = "failed to query the account"
	errMsgUnmarshallAccount              = "failed to unmarshall the account"
	errMsgFaildToGetDelegatorDelegations = "failed to get delegator delegations"
	errMsgFaildToGetDelegatorValidators  = "failed to get delegator validators"
)

func ErrorQueryAccount() *types.Error {
	return types.NewError(ACCOUNT, errQueryAccount, errMsgQueryAccount)
}
func ErrorUnmarshallAccount() *types.Error {
	return types.NewError(ACCOUNT, errUnmarshallAccount, errMsgUnmarshallAccount)
}
func ErrorFailedToGetDelegatorDelegations() *types.Error {
	return types.NewError(ACCOUNT, errFaildToGetDelegatorDelegations, errMsgFaildToGetDelegatorDelegations)
}
func ErrorFailedToGetDelegatorValidators() *types.Error {
	return types.NewError(ACCOUNT, errFaildToGetDelegatorValidators, errMsgFaildToGetDelegatorValidators)
}
