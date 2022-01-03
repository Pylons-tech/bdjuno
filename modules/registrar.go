package modules

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authttypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pylons-tech/juno/client"
	jmodules "github.com/pylons-tech/juno/modules"
	"github.com/pylons-tech/juno/modules/messages"
	"github.com/pylons-tech/juno/modules/registrar"

	"github.com/pylons-tech/bdjuno/types/config"

	"github.com/pylons-tech/bdjuno/utils"

	"github.com/pylons-tech/bdjuno/modules/history"

	"github.com/pylons-tech/bdjuno/database"
	"github.com/pylons-tech/bdjuno/modules/auth"
	"github.com/pylons-tech/bdjuno/modules/bank"
	"github.com/pylons-tech/bdjuno/modules/consensus"
	"github.com/pylons-tech/bdjuno/modules/distribution"
	"github.com/pylons-tech/bdjuno/modules/gov"
	"github.com/pylons-tech/bdjuno/modules/mint"
	"github.com/pylons-tech/bdjuno/modules/modules"
	"github.com/pylons-tech/bdjuno/modules/pricefeed"
	"github.com/pylons-tech/bdjuno/modules/slashing"
	"github.com/pylons-tech/bdjuno/modules/staking"
)

// UniqueAddressesParser returns a wrapper around the given parser that removes all duplicated addresses
func UniqueAddressesParser(parser messages.MessageAddressesParser) messages.MessageAddressesParser {
	return func(cdc codec.Codec, msg sdk.Msg) ([]string, error) {
		addresses, err := parser(cdc, msg)
		if err != nil {
			return nil, err
		}

		return utils.RemoveDuplicateValues(addresses), nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ registrar.Registrar = &Registrar{}
)

// Registrar represents the modules.Registrar that allows to register all modules that are supported by BigDipper
type Registrar struct {
	parser messages.MessageAddressesParser
}

// NewRegistrar allows to build a new Registrar instance
func NewRegistrar(parser messages.MessageAddressesParser) *Registrar {
	return &Registrar{
		parser: UniqueAddressesParser(parser),
	}
}

// BuildModules implements modules.Registrar
func (r *Registrar) BuildModules(ctx registrar.Context) jmodules.Modules {
	bdjunoCfg, ok := ctx.ParsingConfig.(*config.Config)
	if !ok {
		panic(fmt.Errorf("invalid configuration type: %T", ctx.ParsingConfig))
	}

	bigDipperBd := database.Cast(ctx.Database)
	grpcConnection := client.MustCreateGrpcConnection(ctx.ParsingConfig)
	encodingConfig := ctx.EncodingConfig

	authClient := authttypes.NewQueryClient(grpcConnection)
	bankClient := banktypes.NewQueryClient(grpcConnection)
	distrClient := distrtypes.NewQueryClient(grpcConnection)
	govClient := govtypes.NewQueryClient(grpcConnection)
	mintClient := minttypes.NewQueryClient(grpcConnection)
	slashingClient := slashingtypes.NewQueryClient(grpcConnection)
	stakingClient := stakingtypes.NewQueryClient(grpcConnection)

	return []jmodules.Module{
		messages.NewModule(r.parser, encodingConfig.Marshaler, ctx.Database),
		auth.NewModule(r.parser, authClient, encodingConfig, bigDipperBd),
		bank.NewModule(r.parser, authClient, bankClient, encodingConfig, bigDipperBd),
		consensus.NewModule(ctx.Proxy, bigDipperBd),
		distribution.NewModule(bdjunoCfg, bankClient, distrClient, bigDipperBd),
		gov.NewModule(bankClient, govClient, stakingClient, encodingConfig, bigDipperBd),
		mint.NewModule(mintClient, bigDipperBd),
		modules.NewModule(ctx.ParsingConfig, bigDipperBd),
		pricefeed.NewModule(bdjunoCfg, encodingConfig, bigDipperBd),
		slashing.NewModule(slashingClient, bigDipperBd),
		staking.NewModule(ctx.ParsingConfig, bankClient, stakingClient, distrClient, encodingConfig, bigDipperBd),
		history.NewModule(r.parser, encodingConfig, bigDipperBd),
	}
}
