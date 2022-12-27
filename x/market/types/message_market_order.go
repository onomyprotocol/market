package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMarketOrder = "market_order"

var _ sdk.Msg = &MsgMarketOrder{}

func NewMsgMarketOrder(creator string, denomAsk string, denomBid string, amountBid string, quoteAsk string, slippage string) *MsgMarketOrder {
	return &MsgMarketOrder{
		Creator:   creator,
		DenomAsk:  denomAsk,
		DenomBid:  denomBid,
		AmountBid: amountBid,
		QuoteAsk:  quoteAsk,
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
	return nil
}
