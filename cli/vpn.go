package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub/x/vpn"

	"github.com/ironman0x7b2/client/types"
)

func (cli *CLI) GetSubscriptonsOfClientFromRPC(address sdk.AccAddress) ([]vpn.Subscription, error) {
	params := vpn.NewQuerySubscriptionsOfAddressParams(address)

	bz, err := cli.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, _, err := cli.QueryWithData(fmt.Sprintf("custom/%s/%s", vpn.QuerierRoute, vpn.QuerySubscriptionsOfAddress), bz)
	if err != nil {
		return nil, err
	}

	var subscriptions []vpn.Subscription
	err = cli.Codec.UnmarshalJSON(res, &subscriptions)
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (cli *CLI) GetNodes(cfg *types.Config, client *http.Client) (interface{}, error) {
	for _, resolver := range cfg.Resolvers {
		port := strconv.FormatUint(resolver.Port, 10)
		url := "http://" + resolver.IP + ":" + port + "/nodes"

		res, err := client.Get(url)
		if err != nil {
			return nil, err
		}

		var _resp types.Response
		_body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("Error while reading response body from node")
		}

		err = json.Unmarshal(_body, &_resp)
		if err != nil {
			log.Println("Error while unmarshal node response")
		}

		return _resp.Result.(*[]vpn.Node), nil
	}

	return nil, nil
}
