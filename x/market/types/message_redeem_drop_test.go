package types

import (
	"testing"

	"market/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRedeemDrop_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRedeemDrop
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRedeemDrop{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: MsgRedeemDrop{
				Creator: sample.AccAddress(),
				Uid:     "0",
			},
		},
		{
			name: "negative uid",
			msg: MsgRedeemDrop{
				Creator: sample.AccAddress(),
				Uid:     "-1",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
