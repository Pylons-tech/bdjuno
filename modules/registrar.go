package modules

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/pylons-tech/juno/modules/pruning"
	"github.com/pylons-tech/juno/modules/telemetry"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/pylons-tech/juno/node/remote"

	"github.com/pylons-tech/bdjuno/modules/history"
	"github.com/pylons-tech/bdjuno/modules/slashing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pylons-tech/juno/node/local"

	jmodules "github.com/pylons-tech/juno/modules"
	"github.com/pylons-tech/juno/modules/messages"
	"github.com/pylons-tech/juno/modules/registrar"

	"github.com/pylons-tech/bdjuno/utils"

	nodeconfig "github.com/pylons-tech/juno/node/config"

	"github.com/pylons-tech/bdjuno/database"
	"github.com/pylons-tech/bdjuno/modules/auth"
	"github.com/pylons-tech/bdjuno/modules/bank"
	banksource "github.com/pylons-tech/bdjuno/modules/bank/source"
	localbanksource "github.com/pylons-tech/bdjuno/modules/bank/source/local"
	remotebanksource "github.com/pylons-tech/bdjuno/modules/bank/source/remote"
	"github.com/pylons-tech/bdjuno/modules/consensus"
	"github.com/pylons-tech/bdjuno/modules/distribution"
	distrsource "github.com/pylons-tech/bdjuno/modules/distribution/source"
	localdistrsource "github.com/pylons-tech/bdjuno/modules/distribution/source/local"
	remotedistrsource "github.com/pylons-tech/bdjuno/modules/distribution/source/remote"
	"github.com/pylons-tech/bdjuno/modules/gov"
	govsource "github.com/pylons-tech/bdjuno/modules/gov/source"
	localgovsource "github.com/pylons-tech/bdjuno/modules/gov/source/local"
	remotegovsource "github.com/pylons-tech/bdjuno/modules/gov/source/remote"
	"github.com/pylons-tech/bdjuno/modules/mint"
	mintsource "github.com/pylons-tech/bdjuno/modules/mint/source"
	localmintsource "github.com/pylons-tech/bdjuno/modules/mint/source/local"
	remotemintsource "github.com/pylons-tech/bdjuno/modules/mint/source/remote"
	"github.com/pylons-tech/bdjuno/modules/modules"
	"github.com/pylons-tech/bdjuno/modules/pricefeed"
	slashingsource "github.com/pylons-tech/bdjuno/modules/slashing/source"
	localslashingsource "github.com/pylons-tech/bdjuno/modules/slashing/source/local"
	remoteslashingsource "github.com/pylons-tech/bdjuno/modules/slashing/source/remote"
	"github.com/pylons-tech/bdjuno/modules/staking"
	stakingsource "github.com/pylons-tech/bdjuno/modules/staking/source"
	localstakingsource "github.com/pylons-tech/bdjuno/modules/staking/source/local"
	remotestakingsource "github.com/pylons-tech/bdjuno/modules/staking/source/remote"
)

// UniqueAddressesParser returns a wrapper around the given parser that removes all duplicated addresses
func UniqueAddressesParser(parser messages.MessageAddressesParser) messages.MessageAddressesParser {
	return func(cdc codec.Marshaler, msg sdk.Msg) ([]string, error) {
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
	cdc := ctx.EncodingConfig.Marshaler
	db := database.Cast(ctx.Database)

	sources, err := BuildSources(ctx.JunoConfig.Node, ctx.EncodingConfig)
	if err != nil {
		panic(err)
	}

	authModule := auth.NewModule(r.parser, cdc, db)
	bankModule := bank.NewModule(r.parser, sources.BankSource, cdc, db)
	consensusModule := consensus.NewModule(db)
	distrModule := distribution.NewModule(ctx.JunoConfig, sources.DistrSource, bankModule, cdc, db)
	historyModule := history.NewModule(ctx.JunoConfig.Chain, r.parser, cdc, db)
	mintModule := mint.NewModule(sources.MintSource, cdc, db)
	slashingModule := slashing.NewModule(sources.SlashingSource, nil, cdc, db)
	stakingModule := staking.NewModule(sources.StakingSource, bankModule, distrModule, historyModule, slashingModule, cdc, db)
	govModule := gov.NewModule(sources.GovSource, authModule, bankModule, distrModule, mintModule, slashingModule, stakingModule, cdc, db)

	return []jmodules.Module{
		messages.NewModule(r.parser, cdc, ctx.Database),
		telemetry.NewModule(ctx.JunoConfig),
		pruning.NewModule(ctx.JunoConfig, db, ctx.Logger),

		authModule,
		bankModule,
		consensusModule,
		distrModule,
		govModule,
		historyModule,
		mint.NewModule(sources.MintSource, cdc, db),
		modules.NewModule(ctx.JunoConfig.Chain, db),
		pricefeed.NewModule(ctx.JunoConfig, historyModule, cdc, db),
		slashing.NewModule(sources.SlashingSource, stakingModule, cdc, db),
		stakingModule,
	}
}

type Sources struct {
	BankSource     banksource.Source
	DistrSource    distrsource.Source
	GovSource      govsource.Source
	MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	StakingSource  stakingsource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, encodingConfig)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, encodingConfig *params.EncodingConfig) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, encodingConfig)
	if err != nil {
		return nil, err
	}

	app := simapp.NewSimApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, map[int64]bool{},
		cfg.Home, 0, simapp.MakeTestEncodingConfig(), simapp.EmptyAppOptions{},
	)

	sources := &Sources{
		BankSource:     localbanksource.NewSource(source, banktypes.QueryServer(app.BankKeeper)),
		DistrSource:    localdistrsource.NewSource(source, distrtypes.QueryServer(app.DistrKeeper)),
		GovSource:      localgovsource.NewSource(source, govtypes.QueryServer(app.GovKeeper)),
		MintSource:     localmintsource.NewSource(source, minttypes.QueryServer(app.MintKeeper)),
		SlashingSource: localslashingsource.NewSource(source, slashingtypes.QueryServer(app.SlashingKeeper)),
		StakingSource:  localstakingsource.NewSource(source, stakingkeeper.Querier{Keeper: app.StakingKeeper}),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(app, "keys")
	if err != nil {
		return nil, err
	}

	err = source.MountTransientStores(app, "tkeys")
	if err != nil {
		return nil, err
	}

	err = source.MountMemoryStores(app, "memKeys")
	if err != nil {
		return nil, err
	}

	err = source.InitStores()
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource:     remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
		DistrSource:    remotedistrsource.NewSource(source, distrtypes.NewQueryClient(source.GrpcConn)),
		GovSource:      remotegovsource.NewSource(source, govtypes.NewQueryClient(source.GrpcConn)),
		MintSource:     remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource: remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		StakingSource:  remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
