package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/juno/modules/messages"
	"github.com/gogo/protobuf/proto"
	"github.com/rs/zerolog/log"

	"github.com/pylons-tech/bdjuno/database"
	authutils "github.com/pylons-tech/bdjuno/modules/auth/utils"
	"github.com/pylons-tech/bdjuno/modules/utils"
)

// HandleMsg handles any message updating the involved accounts
func HandleMsg(msg sdk.Msg, getAddresses messages.MessageAddressesParser, cdc codec.Codec, db *database.Db) error {
	addresses, err := getAddresses(cdc, msg)
	if err != nil {
		log.Error().Str("module", "auth").Err(err).
			Str("operation", "refresh account").
			Msgf("error while refreshing accounts after message of type %s", proto.MessageName(msg))
	}

	return authutils.UpdateAccounts(utils.FilterNonAccountAddresses(addresses), db)
}
