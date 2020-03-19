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

func (cli *CLI) GetResolversNodes(cfg *types.Config, client *http.Client) (interface{}, error) {
	var nodes []types.Nodes
	var ns types.Nodes
	var _nodes []types.Node

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

		bz, err := json.Marshal(_resp.Result)
		err = json.Unmarshal(bz, &_nodes)

		ns.Nodes = _nodes
		ns.Resolver = resolver.ID

		nodes = append(nodes, ns)
	}

	return nodes, nil
}

func (cli *CLI) GetResolverNodes(cfg *types.Config, client *http.Client, id string) (interface{}, error) {
	for _, resolver := range cfg.Resolvers {
		if resolver.ID == id {
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

			return _resp.Result, nil
		}
	}

	return nil, nil
}
