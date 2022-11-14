package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/market module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	// ErrPoolAlreadyExists - the pool is already exist.
	ErrPoolAlreadyExists = sdkerrors.Register(ModuleName, 3, "the pool already exists") // nolint: gomnd
)
