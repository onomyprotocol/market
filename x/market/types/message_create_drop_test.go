package types

import (
	"testing"

	"market/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateDrop_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDrop
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDrop{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: MsgCreateDrop{
				Creator: sample.AccAddress(),
				Pair:    "CoinA,CoinB",
				Drops:   "70",
			},
		},
		{
			name: "not a pair",
			msg: MsgCreateDrop{
				Creator: sample.AccAddress(),
				Pair:    "CoinACoinB",
				Drops:   "70",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "not a pair",
			msg: MsgCreateDrop{
				Creator: sample.AccAddress(),
				Pair:    "CoinA,CoinB,CoinC",
				Drops:   "70",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "negative drops",
			msg: MsgCreateDrop{
				Creator: sample.AccAddress(),
				Pair:    "CoinA,CoinB",
				Drops:   "-1",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "zero drops",
			msg: MsgCreateDrop{
				Creator: sample.AccAddress(),
				Pair:    "CoinA,CoinB",
				Drops:   "0",
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
