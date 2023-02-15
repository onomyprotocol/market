package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
