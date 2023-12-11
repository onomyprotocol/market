package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMarketOrder = "market_order"

var _ sdk.Msg = &MsgMarketOrder{}

func NewMsgMarketOrder(creator string, denomAsk string, amountAsk string, denomBid string, amountBid string, slippage string) *MsgMarketOrder {
	return &MsgMarketOrder{
		Creator:   creator,
		DenomAsk:  denomAsk,
		AmountAsk: amountAsk,
		DenomBid:  denomBid,
		AmountBid: amountBid,
		Slippage:  slippage,
	}
}

func (msg *MsgMarketOrder) Route() string {
	return RouterKey
}

func (msg *MsgMarketOrder) Type() string {
	return TypeMsgMarketOrder
}

func (msg *MsgMarketOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMarketOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMarketOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = sdk.ValidateDenom(msg.DenomAsk)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid ask denom")
	}

	err = sdk.ValidateDenom(msg.DenomBid)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid bid denom")
	}

	amount, ok := sdk.NewIntFromString(msg.AmountBid)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount integer")
	}
	if amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount integer")
	}

	amount, ok = sdk.NewIntFromString(msg.AmountAsk)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount integer")
	}
	if amount.LTE(sdk.ZeroInt()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount integer")
	}

	slippage, ok := sdk.NewIntFromString(msg.Slippage)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid slippage integer")
	}
	if slippage.IsNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid slippage integer")
	}
	if slippage.GT(sdk.NewInt(9999)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid slippage integer")
	}

	return nil
}
