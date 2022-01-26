package source

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/pylons-tech/bdjuno/types"
)

type Source interface {
	GetBalances(addresses []string, height int64) ([]types.AccountBalance, error)
	GetSupply(height int64) (sdk.Coins, error)
}
