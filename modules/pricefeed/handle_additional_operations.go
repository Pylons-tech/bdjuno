package pricefeed

import (
	"fmt"
	"time"

	"github.com/pylons-tech/bdjuno/types/config"

	historyutils "github.com/pylons-tech/bdjuno/modules/history/utils"
	"github.com/pylons-tech/bdjuno/modules/utils"

	"github.com/pylons-tech/bdjuno/database"

	"github.com/pylons-tech/bdjuno/types"

	"github.com/rs/zerolog/log"
)

// StoreTokens stores the tokens defined inside the given configuration into the database
func StoreTokens(cfg *config.Config, db *database.Db) error {
	log.Debug().Str("module", "pricefeed").Msg("storing tokens")

	var prices []types.TokenPrice
	for _, coin := range cfg.GetPricefeedConfig().GetTokens() {
		// Save the coin as a token with its units
		err := db.SaveToken(coin)
		if err != nil {
			return fmt.Errorf("error while saving token: %s", err)
		}

		// Create the price entry
		for _, unit := range coin.Units {
			// Skip units with empty price ids
			if unit.PriceID == "" {
				continue
			}

			prices = append(prices, types.NewTokenPrice(unit.Denom, 0, 0, time.Time{}))
		}
	}

	err := db.SaveTokensPrices(prices)
	if err != nil {
		return fmt.Errorf("error while storing token prices: %s", err)
	}

	if utils.IsModuleEnabled(cfg, types.HistoryModuleName) {
		return historyutils.UpdatePriceHistory(prices, db)
	}

	return nil
}
