package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreatePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreatePool
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreatePool{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: MsgCreatePool{
				CoinA:   "20CoinA",
				CoinB:   "20CoinB",
				Creator: sample.AccAddress(),
			},
		},
		{
			name: "equal denoms",
			msg: MsgCreatePool{
				CoinA:   "20CoinA",
				CoinB:   "20CoinA",
				Creator: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid coin A",
			msg: MsgCreatePool{
				CoinA:   "-20CoinA",
				CoinB:   "20CoinV",
				Creator: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid coin B",
			msg: MsgCreatePool{
				CoinA:   "20CoinA",
				CoinB:   "-20CoinV",
				Creator: sample.AccAddress(),
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
