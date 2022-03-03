package pricefeed

import "github.com/Pylons-tech/bdjuno/types"

type HistoryModule interface {
	UpdatePricesHistory([]types.TokenPrice) error
}
