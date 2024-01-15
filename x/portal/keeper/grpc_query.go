package keeper

import (
	"market/x/portal/types"
)

var _ types.QueryServer = Keeper{}
