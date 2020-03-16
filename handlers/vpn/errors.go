package vpn

import (
	"github.com/ironman0x7b2/client/types"
)

const (
	VPN = "vpn"

	errFailedToGetNodes             = 12
	errFailedToGetSubscriptions     = 13
	errNoSubscriptions              = 14
	errMarshalNodeRequest           = 15
	errQuerySubscriptionTransaction = 16

	errMsgFailedToGetNodes                  = "failed to get the nodes"
	errMsgFailedToGetSubscriptions          = "failed to get the subscriptions"
	errMsgNoSubscriptions                   = "no subscriptions found"
	errMsgMarshalNodeRequest                = "failed to marshal the node request"
	errMsgErrorQuerySubscriptionTransaction = "error query subscription transaction"
)

func errorFailedToGetNodes() *types.Error {
	return types.NewError(VPN, errFailedToGetNodes, errMsgFailedToGetNodes)
}
func errorFailedToGetSubscriptions() *types.Error {
	return types.NewError(VPN, errFailedToGetSubscriptions, errMsgFailedToGetSubscriptions)
}
func errorNoSubscriptions() *types.Error {
	return types.NewError(VPN, errNoSubscriptions, errMsgNoSubscriptions)
}
func errorMarshallNodeRequestBody() *types.Error {
	return types.NewError(VPN, errMarshalNodeRequest, errMsgMarshalNodeRequest)
}
func errorQuerySubscriptionTransaction() *types.Error {
	return types.NewError(VPN, errQuerySubscriptionTransaction, errMsgErrorQuerySubscriptionTransaction)
}
