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

	if msg.OrderType != "stop" && msg.OrderType != "limit" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid order type")
	}

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount integer")
	}

	if !amount.GT(sdk.NewInt(0)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount integer")
	}

	// Rate[0] needs to fit into uint64 to avoid numerical errors
	// Rate[0] will be converted to sdk.Int type in execution
	_, err = strconv.ParseUint(msg.Rate[0], 10, 64)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	// Rate[1] needs to fit into uint64 to avoid numerical errors
	// Rate[1] will be converted to sdk.Int type in execution
	_, err = strconv.ParseUint(msg.Rate[1], 10, 64)
	if err != nil {
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
