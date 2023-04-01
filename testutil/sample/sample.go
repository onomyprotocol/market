package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// SampleCoins returns the required NewCoins
func SampleCoins(coina string, coinb string) sdk.Coins {

	coinA, err := sdk.ParseCoinNormalized(coina)
	if err != nil {
		panic(err)
	}

	coinB, err := sdk.ParseCoinNormalized(coinb)
	if err != nil {
		panic(err)
	}

	return sdk.NewCoins(coinA, coinB)
}

// SampleDenoms returns the required denoms values
func SampleDenoms(coins sdk.Coins) (denomA string, denomB string) {
	denom1 := coins.GetDenomByIndex(0)
	denom2 := coins.GetDenomByIndex(1)
	return denom1, denom2
}
