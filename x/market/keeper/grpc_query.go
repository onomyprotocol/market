package keeper

import (
	"market/x/market/types"
)

var _ types.QueryServer = Keeper{}
