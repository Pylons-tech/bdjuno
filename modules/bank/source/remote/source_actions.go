package remote

import (
	"fmt"

	"github.com/Pylons-tech/bdjuno/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// GetAccountBalances implements bankkeeper.Source
func (s Source) GetAccountBalance(address string, height int64) ([]sdk.Coin, error) {

	// Get account balance at certain height
	ctx := utils.GetHeightRequestContext(s.Ctx, height)
	balRes, err := s.bankClient.AllBalances(ctx, &banktypes.QueryAllBalancesRequest{Address: address})
	if err != nil {
		return nil, fmt.Errorf("error while getting all balances: %s", err)
	}

	return balRes.Balances, nil
}
