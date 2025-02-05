package auth

import (
	"encoding/json"
	"fmt"

	"github.com/pylons-tech/juno/cmd/parse"
	"github.com/pylons-tech/juno/types/config"
	"github.com/spf13/cobra"

	"github.com/pylons-tech/bdjuno/database"
	authutils "github.com/pylons-tech/bdjuno/modules/auth"
	"github.com/pylons-tech/bdjuno/utils"
)

// vestingCmd returns a Cobra command that allows to fix the vesting data for the accounts
func vestingCmd(parseConfig *parse.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "vesting-accounts",
		Short: "Fix the vesting accounts stored by removing duplicated vesting periods",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parse.GetParsingContext(parseConfig)
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Get the genesis
			genesis, err := utils.ReadGenesis(config.Cfg, parseCtx.Node)
			if err != nil {
				return fmt.Errorf("error while reading the genesis: %s", err)
			}

			var appState map[string]json.RawMessage
			if err := json.Unmarshal(genesis.AppState, &appState); err != nil {
				return fmt.Errorf("error unmarshalling genesis doc: %s", err)
			}

			vestingAccounts, err := authutils.GetGenesisVestingAccounts(appState, parseCtx.EncodingConfig.Marshaler)
			if err != nil {
				return fmt.Errorf("error while gestting vesting accounts: %s", err)
			}

			err = db.SaveVestingAccounts(vestingAccounts)
			if err != nil {
				return fmt.Errorf("error while storing vesting accounts: %s", err)
			}

			return nil
		},
	}
}
