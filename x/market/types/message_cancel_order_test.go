package types

import (
	"testing"

	"market/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCancelOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelOrder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCancelOrder{
				Creator: "invalid_address",
				Uid:     "2",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: MsgCancelOrder{
				Creator: sample.AccAddress(),
				Uid:     "0",
			},
		},
		{
			name: "negative uid",
			msg: MsgCancelOrder{
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
