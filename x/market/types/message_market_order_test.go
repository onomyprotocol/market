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
		},
		{
			name: "valid address",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				AmountBid: "40",
				Slippage:  "20",
			},
		},
		{
			name: "max slippage",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				AmountBid: "40",
				Slippage:  "9999",
			},
		},
		{
			name: "too large slippage",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				AmountBid: "40",
				Slippage:  "10000",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "negative slippage",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				AmountBid: "40",
				Slippage:  "-1",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "negative bid",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				AmountBid: "-1",
				Slippage:  "20",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "zero bid",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				AmountBid: "0",
				Slippage:  "20",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid ask",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "20CoinA",
				DenomBid:  "CoinB",
				AmountBid: "40",
				Slippage:  "20",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid bid",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "20CoinB",
				AmountBid: "40",
				Slippage:  "20",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid slippage",
			msg: MsgMarketOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				AmountBid: "40",
				Slippage:  "0999",
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
