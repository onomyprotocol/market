package types

import (
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateDrop = "create_drop"

var _ sdk.Msg = &MsgCreateDrop{}

func NewMsgCreateDrop(creator string, pair string, drops string) *MsgCreateDrop {
	return &MsgCreateDrop{
		Creator: creator,
		Pair:    pair,
		Drops:   drops,
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

	if len(pairMsg) != 2 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "pair not a valid denom pair")
	}

	drops, ok := sdk.NewIntFromString(msg.Drops)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "drops not a valid integer")
	}

	if !drops.GT(sdk.NewInt(0)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "drops not >0")
	}

	return nil
}
