package pricefeed

import "github.com/pylons-tech/bdjuno/types"

type HistoryModule interface {
	UpdatePricesHistory([]types.TokenPrice) error
}
