package errors

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	TXS = "key"

	errFailedToGetKeyInfo            = 8
	errFailedToGetAccount            = 8
	errAccountDoesNotExist           = 8
	errSignTransaction               = 8
	errGetRPCNode                    = 8
	errFaildToBroadcastTransaction   = 5
	errFaildToGetTransaction         = 5
	errFaildToUnmarshallTransaction  = 5
	errFaildToGetTransactions        = 5
	errFaildToUnmarshallTransactions = 5

	errMsgFailedToGetKeyInfo            = "failed to get key info"
	errMsgFailedToGetAccount            = "failed to get account"
	errMsgAccountDoesNotExist           = "account does not exist"
	errMsgSignTransaction               = "failed to sign transaction"
	errMsgGetRPCNode                    = "failed to get RPC node"
	errMsgFaildToBroadcastTransaction   = "failed to broadcast the transaction"
	errMsgFaildToGetTransaction         = "failed to get the transaction"
	errMsgFaildToUnmarshallTransaction  = "failed to get the transaction"
	errMsgFaildToGetTransactions        = "failed to get the transactions"
	errMsgFaildToUnmarshallTransactions = "failed to get the transactions"
)

func ErrorFailedToGetKeyInfo() *types.Error {
	return types.NewError(TXS, errFailedToGetKeyInfo, errMsgFailedToGetKeyInfo)
}
func ErrorFailedToGetAccount() *types.Error {
	return types.NewError(TXS, errFailedToGetAccount, errMsgFailedToGetAccount)
}
func ErrorAccountDoesNotExist() *types.Error {
	return types.NewError(TXS, errAccountDoesNotExist, errMsgAccountDoesNotExist)
}
func ErrorSignTransactions() *types.Error {
	return types.NewError(TXS, errSignTransaction, errMsgSignTransaction)
}
func ErrorGetRPCNode() *types.Error {
	return types.NewError(TXS, errGetRPCNode, errMsgGetRPCNode)
}
func ErrorFailedToBroadcastTransaction() *types.Error {
	return types.NewError(TXS, errFaildToBroadcastTransaction, errMsgFaildToBroadcastTransaction)
}
func ErrorFailedToGetTransaction() *types.Error {
	return types.NewError(TXS, errFaildToGetTransaction, errMsgFaildToGetTransaction)
}
func ErrorFailedToUnmarshalTransaction() *types.Error {
	return types.NewError(TXS, errFaildToUnmarshallTransaction, errMsgFaildToUnmarshallTransaction)
}
func ErrorFailedToGetTransactions() *types.Error {
	return types.NewError(TXS, errFaildToGetTransactions, errMsgFaildToGetTransactions)
}
func ErrorFailedToUnmarshalTransactions() *types.Error {
	return types.NewError(TXS, errFaildToUnmarshallTransactions, errMsgFaildToUnmarshallTransactions)
}
