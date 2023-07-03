package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func EQ(a []sdk.Int, b []sdk.Int) bool {
	return a[0].Mul(b[1]).Equal(a[1].Mul(b[0]))
}

func GT(a []sdk.Int, b []sdk.Int) bool {
	return a[0].Mul(b[1]).GT(a[1].Mul(b[0]))
}

func GTE(a []sdk.Int, b []sdk.Int) bool {
	return a[0].Mul(b[1]).GTE(a[1].Mul(b[0]))
}

func LT(a []sdk.Int, b []sdk.Int) bool {
	return a[0].Mul(b[1]).LT(a[1].Mul(b[0]))
}

func LTE(a []sdk.Int, b []sdk.Int) bool {
	return a[0].Mul(b[1]).LTE(a[1].Mul(b[0]))
}

func RateStringToInt(rateString []string) ([]sdk.Int, error) {
	var rateUint64 [2]uint64
	var err error

	// Rate[0] needs to fit into uint64 to avoid numerical errors
	rateUint64[0], err = strconv.ParseUint(rateString[0], 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate format")
	}

	// Rate[1] needs to fit into uint64 to avoid numerical errors
	rateUint64[1], err = strconv.ParseUint(rateString[1], 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate format")
	}

	var rate []sdk.Int

	rate = append(rate, sdk.NewIntFromUint64(rateUint64[0]))
	rate = append(rate, sdk.NewIntFromUint64(rateUint64[1]))

	return rate, nil
}
