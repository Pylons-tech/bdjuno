package main

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/forbole/juno/v2/cmd"
	initcmd "github.com/forbole/juno/v2/cmd/init"
	parsecmd "github.com/forbole/juno/v2/cmd/parse"
	"github.com/forbole/juno/v2/modules/messages"

	fixcmd "github.com/pylons-tech/bdjuno/cmd/fix"
	migratecmd "github.com/pylons-tech/bdjuno/cmd/migrate"
	parsegenesiscmd "github.com/pylons-tech/bdjuno/cmd/parse-genesis"

	"github.com/pylons-tech/bdjuno/types/config"

	"github.com/pylons-tech/bdjuno/database"
	"github.com/pylons-tech/bdjuno/modules"

	gaiaapp "github.com/cosmos/gaia/v6/app"
)

func main() {
	parseCfg := parsecmd.NewConfig().
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(config.MakeEncodingConfig(getBasicManagers())).
		WithRegistrar(modules.NewRegistrar(getAddressesParser()))

	cfg := cmd.NewConfig("bdjuno").
		WithParseConfig(parseCfg)

	// Run the command
	rootCmd := cmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		cmd.VersionCmd(),
		initcmd.InitCmd(cfg.GetInitConfig()),
		parsecmd.ParseCmd(cfg.GetParseConfig()),
		migratecmd.NewMigrateCmd(),
		fixcmd.NewFixCmd(cfg.GetParseConfig()),
		parsegenesiscmd.NewParseGenesisCmd(cfg.GetParseConfig()),
	)

	executor := cmd.PrepareRootCmd(cfg.GetName(), rootCmd)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

// getBasicManagers returns the various basic managers that are used to register the encoding to
// support custom messages.
// This should be edited by custom implementations if needed.
func getBasicManagers() []module.BasicManager {
	return []module.BasicManager{
		gaiaapp.ModuleBasics,
	}
}

// getAddressesParser returns the messages parser that should be used to get the users involved in
// a specific message.
// This should be edited by custom implementations if needed.
func getAddressesParser() messages.MessageAddressesParser {
	return messages.JoinMessageParsers(
		messages.CosmosMessageAddressesParser,
	)
}
