package cli

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub/x/vpn"

	"github.com/ironman0x7b2/client/handlers/errors"
	"github.com/ironman0x7b2/client/types"
)

func (cli *CLI) GetSubscriptonsOfClientFromRPC(address sdk.AccAddress, module string) ([]vpn.Subscription, *types.Error) {
	params := vpn.NewQuerySubscriptionsOfAddressParams(address)

	bz, err := cli.Codec.MarshalJSON(params)
	if err != nil {
		return nil, errors.ErrorFailedToMarshalParams(module)
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscriptionsOfAddress), bz)
	if err != nil {
		return nil, errors.ErrorFailedToQuerySubscriptionOfClient()
	}

	var subscriptions []vpn.Subscription
	err = cli.Codec.UnmarshalJSON(res, &subscriptions)
	if err != nil {
		return nil, errors.ErrorFailedToUnmarshallSubscriptionOfClient()
	}

	return subscriptions, nil
}
