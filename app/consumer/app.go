package app

import (
	consumertypes "github.com/cosmos/interchain-security/x/ccv/consumer/types"
)

// dummy import
const (
	dummy = consumertypes.ModuleName
)

const (
	AppName              = "market"
	upgradeName          = "v0.1.0"
	AccountAddressPrefix = "onomy"
)

// note: find `"."+AppName` and change to `".onomy_"+AppName`
