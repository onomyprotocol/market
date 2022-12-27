package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "market/CreatePool", nil)
	cdc.RegisterConcrete(&MsgCreateDrop{}, "market/CreateDrop", nil)
	cdc.RegisterConcrete(&MsgRedeemDrop{}, "market/RedeemDrop", nil)
	cdc.RegisterConcrete(&MsgCreateOrder{}, "market/CreateOrder", nil)
	cdc.RegisterConcrete(&MsgCancelOrder{}, "market/CancelOrder", nil)
	cdc.RegisterConcrete(&MsgMarketOrder{}, "market/MarketOrder", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDrop{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRedeemDrop{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMarketOrder{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
