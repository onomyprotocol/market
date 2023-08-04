package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRedeemDrop = "redeem_drop"

var _ sdk.Msg = &MsgRedeemDrop{}

func NewMsgRedeemDrop(creator string, uid string) *MsgRedeemDrop {
	return &MsgRedeemDrop{
		Creator: creator,
		Uid:     uid,
	}
}

func (msg *MsgRedeemDrop) Route() string {
	return RouterKey
}

func (msg *MsgRedeemDrop) Type() string {
	return TypeMsgRedeemDrop
}

func (msg *MsgRedeemDrop) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRedeemDrop) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRedeemDrop) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = strconv.ParseUint(msg.Uid, 10, 64)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "drop uid is not an integer or is negative")
	}

	return nil
}
