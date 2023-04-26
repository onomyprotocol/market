package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateOrder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateOrder{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  MsgCreateOrder{Creator: sample.AccAddress(), DenomAsk: "20CoinA", DenomBid: "20CoinB", Rate: []string{"10", "20"}, OrderType: "stop", Amount: "0", Prev: "0", Next: "0"},
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
