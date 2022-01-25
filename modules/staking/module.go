package staking

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/pylons-tech/juno/modules"

	"github.com/pylons-tech/bdjuno/database"
	stakingsource "github.com/pylons-tech/bdjuno/modules/staking/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.BlockModule   = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/staking module
type Module struct {
	cdc            codec.Marshaler
	db             *database.Db
	source         stakingsource.Source
	bankModule     BankModule
	distrModule    DistrModule
	historyModule  HistoryModule
	slashingModule SlashingModule
}

// NewModule returns a new Module instance
func NewModule(
	source stakingsource.Source,
	bankModule BankModule, distrModule DistrModule, historyModule HistoryModule, slashingModule SlashingModule,
	cdc codec.Marshaler, db *database.Db,
) *Module {
	return &Module{
		cdc:            cdc,
		db:             db,
		source:         source,
		bankModule:     bankModule,
		distrModule:    distrModule,
		historyModule:  historyModule,
		slashingModule: slashingModule,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "staking"
}
