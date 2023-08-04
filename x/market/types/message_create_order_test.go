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
		},
		{
			name: "valid address",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
		},
		{
			name: "limit order",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "limit",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
		},
		{
			name: "invalid denom",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "10CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid denom",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "10CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid type",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "invalid",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "zero amount",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "stop",
				Amount:    "0",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "negative amount",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "stop",
				Amount:    "-1",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate is not a 2-tuple",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate is not a 2-tuple",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate is not a 2-tuple",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20", "30"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate component is negative",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"-10", "20"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate component is negative",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "-20"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate component is zero",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"0", "20"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate component is zero",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "0"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate component fits in uint64",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"18446744073709551615", "18446744073709551615"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
		},
		{
			name: "rate component does not fit in uint64",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"18446744073709551616", "18446744073709551615"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "rate component does not fit in uint64",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"18446744073709551615", "18446744073709551616"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "Prev Uid is invalid",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "-1",
				Next:      "0",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "Prev Uid is invalid",
			msg: MsgCreateOrder{
				Creator:   sample.AccAddress(),
				DenomAsk:  "CoinA",
				DenomBid:  "CoinB",
				Rate:      []string{"10", "20"},
				OrderType: "stop",
				Amount:    "10",
				Prev:      "0",
				Next:      "-1",
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
