package errors

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	errParseRequestBody               = 1
	errParseQueryParams               = 1
	errValidateRequestBody            = 2
	errDecodeAddress                  = 3
	errFaildToPrepareMsg              = 4
	errFaildToMarshalParams           = 6
	errFaildToQueryValidators         = 7
	errUnmarshallDelegatorlValidators = 8
	errFailedToReadResponseBody       = 9

	errMsgParseRequestBody              = "failed to parse the request body"
	errMsgParseQueryParams              = "failed to parse the query params"
	errMsgValidateRequestBody           = "failed to validate request body"
	errMsgDecodeAddress                 = "failed to decode the address"
	errMsgFaildToPrepareTxMsg           = "failed to prepare the message"
	errMsgFaildToMarshalParams          = "failed to marshal params"
	errMsgFaildToQueryValidators        = "failed to query validators"
	errMsgUnmarshallDelegatorValidators = "failed to unmarshall delegator validators "
	errMsgFailedToReadResponseBody      = "failed to read response body"
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
	return types.NewError(module, errFaildToPrepareMsg, errMsgFaildToPrepareTxMsg)
}
func ErrorFailedToMarshalParams(module string) *types.Error {
	return types.NewError(module, errFaildToMarshalParams, errMsgFaildToMarshalParams)
}
func ErrorFailedToQueryValidators(module string) *types.Error {
	return types.NewError(module, errFaildToQueryValidators, errMsgFaildToQueryValidators)
}
func ErrorFailedToUnmarshallDelegatorValidators(module string) *types.Error {
	return types.NewError(module, errUnmarshallDelegatorlValidators, errMsgUnmarshallDelegatorValidators)
}
func ErrorFailedToReadResponseBody(module string) *types.Error {
	return types.NewError(module, errFailedToReadResponseBody, errMsgFailedToReadResponseBody)
}
