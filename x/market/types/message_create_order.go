package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateOrder = "create_order"

var _ sdk.Msg = &MsgCreateOrder{}

func NewMsgCreateOrder(creator string, denomAsk string, denomBid string, orderType string, amount string, rate []string, prev string, next string) *MsgCreateOrder {
	return &MsgCreateOrder{
		Creator:   creator,
		DenomAsk:  denomAsk,
		DenomBid:  denomBid,
		OrderType: orderType,
		Amount:    amount,
		Rate:      rate,
		Prev:      prev,
		Next:      next,
	}
}

func (msg *MsgCreateOrder) Route() string {
	return RouterKey
}

func (msg *MsgCreateOrder) Type() string {
	return TypeMsgCreateOrder
}

func (msg *MsgCreateOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	coinAsk, _ := sdk.ParseCoinNormalized(msg.DenomAsk)
	if !coinAsk.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid ask denom")
	}

	coinBid, _ := sdk.ParseCoinNormalized(msg.DenomBid)
	if !coinBid.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid bid denom")
	}

	if msg.OrderType != "stop" && msg.OrderType != "limit" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid order type")
	}

	_, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount integer")
	}

	_, ok = sdk.NewIntFromString(msg.Rate[0])
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	_, ok = sdk.NewIntFromString(msg.Rate[1])
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	_, err = strconv.ParseUint(msg.Prev, 10, 64)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "prev uid is not an integer")
	}

	_, err = strconv.ParseUint(msg.Next, 10, 64)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "next uid is not an integer")
	}

	return nil
}
