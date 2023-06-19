package types

import (
	"strings"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/testutil/sample"
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
		}, {
			name: "valid address",
			msg:  MsgCreateDrop{Creator: sample.AccAddress(), Pair: strings.Join([]string{"10CoinA", "20CoinB"}, ","), Drops: "70"},
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
