package errors

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	VPN = "vpn"

	errFailedToQuerySubscriptionOfClient      = 6
	errFailedToUnmarshallSubscriptionOfClient = 6
	errFailedToMarshallNodeReqMsg             = 6

	errMsgFailedToQuerySubscriptionOfClient      = "failed to query the subscription of client"
	errMsgFailedToUnmarshallSubscriptionOfClient = "failed to unmarshall the subscription of client"
	errMsgFailedToMarshallNodeReqMsg             = "failed to marshall node req msg"
)

func ErrorFailedToQuerySubscriptionOfClient() *types.Error {
	return types.NewError(VPN, errFailedToQuerySubscriptionOfClient, errMsgFailedToQuerySubscriptionOfClient)
}
func ErrorFailedToUnmarshallSubscriptionOfClient() *types.Error {
	return types.NewError(VPN, errFailedToUnmarshallSubscriptionOfClient, errMsgFailedToUnmarshallSubscriptionOfClient)
}
func ErrorFailedToMarshallNodeReqMsg() *types.Error {
	return types.NewError(VPN, errFailedToMarshallNodeReqMsg, errMsgFailedToMarshallNodeReqMsg)
}
