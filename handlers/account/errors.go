package account

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	ACCOUNT = "account"

	errQueryAccount                    = 12
	errFailedToGetDelegatorDelegations = 13
	errFailedToGetDelegatorValidators  = 14

	errMsgQueryAccount                    = "failed to query the account"
	errMsgFailedToGetDelegatorDelegations = "failed to get delegator delegations"
	errMsgFailedToGetDelegatorValidators  = "failed to get delegator validators"
)

func errorQueryAccount() *types.Error {
	return types.NewError(ACCOUNT, errQueryAccount, errMsgQueryAccount)
}
func errorFailedToGetDelegatorDelegations() *types.Error {
	return types.NewError(ACCOUNT, errFailedToGetDelegatorDelegations, errMsgFailedToGetDelegatorDelegations)
}
func errorFailedToGetDelegatorValidators() *types.Error {
	return types.NewError(ACCOUNT, errFailedToGetDelegatorValidators, errMsgFailedToGetDelegatorValidators)
}
