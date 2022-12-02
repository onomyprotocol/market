package types

import (
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
	return nil
}
