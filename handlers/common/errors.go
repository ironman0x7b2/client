package common

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	errParseRequestBody               = 1
	errParseQueryParams               = 2
	errValidateRequestBody            = 3
	errDecodeAddress                  = 4
	errFailedToPrepareMsg             = 5
	errFailedToBroadcastTransaction   = 6
	errFailedToGetValidator           = 7
	errFailedToGetValidators          = 8
	errFailedToGetDelegatorValidators = 9
	errFailedToGetTransaction         = 10
	errFailedToGetTransactions        = 11

	errMsgParseRequestBody               = "failed to parse the request body"
	errMsgParseQueryParams               = "failed to parse the query params"
	errMsgValidateRequestBody            = "failed to validate request body"
	errMsgDecodeAddress                  = "failed to decode the address"
	errMsgFailedToPrepareTxMsg           = "failed to prepare the message"
	errMsgFailedToBroadcastTransaction   = "failed to broadcast the transaction"
	errMsgFailedToGetValidator           = "failed to get validator"
	errMsgFailedToGetValidators          = "failed to get validators"
	errMsgFailedToGetDelegatorValidators = "failed to get delegator validators"
	errMsgFailedToGetTransaction         = "failed to get transaction"
	errMsgFailedToGetTransactions        = "failed to get transactions"
)

func ErrorParseRequestBody(module string) *types.Error {
	return types.NewError(module, errParseRequestBody, errMsgParseRequestBody)
}
func ErrorParseQueryParams(module string) *types.Error {
	return types.NewError(module, errParseQueryParams, errMsgParseQueryParams)
}
func ErrorValidateRequestBody(module string) *types.Error {
	return types.NewError(module, errValidateRequestBody, errMsgValidateRequestBody)
}
func ErrorDecodeAddress(module string) *types.Error {
	return types.NewError(module, errDecodeAddress, errMsgDecodeAddress)
}
func ErrorFailedToPrepareMsg(module string) *types.Error {
	return types.NewError(module, errFailedToPrepareMsg, errMsgFailedToPrepareTxMsg)
}
func ErrorFailedToBroadcastTransaction(module string) *types.Error {
	return types.NewError(module, errFailedToBroadcastTransaction, errMsgFailedToBroadcastTransaction)
}
func ErrorFailedToGetValidator(module string) *types.Error {
	return types.NewError(module, errFailedToGetValidator, errMsgFailedToGetValidator)
}
func ErrorFailedToGetValidators(module string) *types.Error {
	return types.NewError(module, errFailedToGetValidators, errMsgFailedToGetValidators)
}
func ErrorFailedToGetDelegatorValidators(module string) *types.Error {
	return types.NewError(module, errFailedToGetDelegatorValidators, errMsgFailedToGetDelegatorValidators)
}
func ErrorFailedToGetTransaction(module string) *types.Error {
	return types.NewError(module, errFailedToGetTransaction, errMsgFailedToGetTransaction)
}
func ErrorFailedToGetTransactions(module string) *types.Error {
	return types.NewError(module, errFailedToGetTransactions, errMsgFailedToGetTransactions)
}
