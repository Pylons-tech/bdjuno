package staking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Pylons-tech/bdjuno/types"
)

type SlashingModule interface {
	GetSigningInfo(height int64, consAddr sdk.ConsAddress) (types.ValidatorSigningInfo, error)
}
