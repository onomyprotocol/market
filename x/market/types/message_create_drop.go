package types

import (
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateDrop = "create_drop"

var _ sdk.Msg = &MsgCreateDrop{}

func NewMsgCreateDrop(creator string, pair string, drops string, rate1 []string, prev1 string, next1 string, rate2 []string, prev2 string, next2 string) *MsgCreateDrop {
	return &MsgCreateDrop{
		Creator: creator,
		Pair:    pair,
		Drops:   drops,
		Rate1:   rate1,
		Prev1:   prev1,
		Next1:   next1,
		Rate2:   rate2,
		Prev2:   prev1,
		Next2:   next2,
	}
}

func (msg *MsgCreateDrop) Route() string {
	return RouterKey
}

func (msg *MsgCreateDrop) Type() string {
	return TypeMsgCreateDrop
}

func (msg *MsgCreateDrop) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDrop) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDrop) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	pairMsg := strings.Split(msg.Pair, ",")
	sort.Strings(pairMsg)

	denom1 := pairMsg[1]

	coin1, _ := sdk.ParseCoinNormalized(`{1}{` + denom1 + `}`)
	if !coin1.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "pair not a valid denom pair")
	}

	denom2 := pairMsg[2]

	coin2, _ := sdk.ParseCoinNormalized(`{1}{` + denom2 + `}`)
	if !coin2.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "pair not a valid denom pair")
	}

	_, ok := sdk.NewIntFromString(msg.Drops)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "drops not a valid integer")
	}

	_, ok = sdk.NewIntFromString(msg.Rate1[0])
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	_, ok = sdk.NewIntFromString(msg.Rate1[1])
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	_, ok = sdk.NewIntFromString(msg.Prev1)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "prev1 uid is not an integer")
	}

	_, ok = sdk.NewIntFromString(msg.Next1)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "next1 uid is not an integer")
	}

	_, ok = sdk.NewIntFromString(msg.Rate2[0])
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	_, ok = sdk.NewIntFromString(msg.Rate2[1])
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	_, ok = sdk.NewIntFromString(msg.Prev2)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "prev2 uid is not an integer")
	}

	_, ok = sdk.NewIntFromString(msg.Next2)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "next2 uid is not an integer")
	}

	return nil
}
