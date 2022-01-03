package auth

import (
	"github.com/pylons-tech/juno/cmd/parse"
	"github.com/spf13/cobra"
)

// NewAuthCmd returns the Cobra command that allows to fix all the things related to the x/auth module
func NewAuthCmd(parseCfg *parse.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Fix things related to the x/auth module",
	}

	cmd.AddCommand(
		vestingCmd(parseCfg),
	)

	return cmd
}
