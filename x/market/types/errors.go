package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/market module sentinel errors
var (
	// ErrInvalidCoins - coin/coins are invalid.
	ErrInvalidCoins = sdkerrors.Register(ModuleName, 1, "coins are invalid")
	// ErrInsufficientBalance - the user balance is insufficient for the operation.
	ErrInsufficientBalance = sdkerrors.Register(ModuleName, 2, "insufficient balance") // nolint: gomnd
	// ErrPoolAlreadyExists - the pool is already exist.
	ErrPoolAlreadyExists = sdkerrors.Register(ModuleName, 3, "the pool already exists") // nolint: gomnd
	// ErrPoolNotFound - the pool not found.
	ErrPoolNotFound = sdkerrors.Register(ModuleName, 4, "the pool not found") // nolint: gomnd
	// ErrPoolNotFound - the drop not found.
	ErrDropNotFound = sdkerrors.Register(ModuleName, 5, "the pool not found") // nolint: gomnd
	// ErrPoolNotFound - the drop not found.
	ErrNotDropOwner = sdkerrors.Register(ModuleName, 6, "the pool not found") // nolint: gomnd
	// ErrMemberNotFound - the pool member not found.
	ErrMemberNotFound = sdkerrors.Register(ModuleName, 7, "the pool member not found") // nolint: gomnd
	// ErrInvalidDropAmount - the drop amount is invalid.
	ErrInvalidDropAmount = sdkerrors.Register(ModuleName, 8, "invalid drop amount") // nolint: gomnd
	// ErrInvalidDenomsPair - invalid demos pair.
	ErrInvalidDenomsPair = sdkerrors.Register(ModuleName, 9, "invalid demos pair") // nolint: gomnd
	// ErrInvalidOrder - invalid demos pair.
	ErrInvalidOrder = sdkerrors.Register(ModuleName, 10, "invalid order") // nolint: gomnd
)
