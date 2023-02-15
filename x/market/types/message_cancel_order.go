package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelOrder = "cancel_order"

var _ sdk.Msg = &MsgCancelOrder{}

func NewMsgCancelOrder(creator string, uid string) *MsgCancelOrder {
	return &MsgCancelOrder{
		Creator: creator,
		Uid:     uid,
	}
}

func (msg *MsgCancelOrder) Route() string {
	return RouterKey
}

func (msg *MsgCancelOrder) Type() string {
	return TypeMsgCancelOrder
}

func (msg *MsgCancelOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = strconv.ParseUint(msg.Uid, 10, 64)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "drop uid is not an integer")
	}

	return nil
}
