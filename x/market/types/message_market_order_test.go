package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pendulum-labs/market/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgMarketOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMarketOrder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMarketOrder{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "20CoinA",
				DenomBid:  "30CoinB",
				AmountBid: "40",
				Slippage:  "20",
			},
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
