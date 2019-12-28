package cli

import (
	"errors"
	"fmt"
	"log"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func (c *CLI) GetAccount(address sdk.AccAddress) (auth.Account, error) {
	bytes, err := c.Codec.MarshalJSON(auth.NewQueryAccountParams(address))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	
	res, _, err := c.QueryWithData(fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount), bytes)
	if err != nil {
		if err.Error() == errors.New("no RPC client defined").Error() {
			log.Println(err.Error())
			return nil, err
		}
		return nil, nil
	}
	
	var account auth.Account
	if err := c.Codec.UnmarshalJSON(res, &account); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	
	return account, nil
}
