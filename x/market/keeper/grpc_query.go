package keeper

import (
	"github.com/pendulum-labs/market/x/market/types"
)

var _ types.QueryServer = Keeper{}
