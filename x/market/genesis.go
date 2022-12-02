package market

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/market/x/market/keeper"
	"github.com/onomyprotocol/market/x/market/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the pool
	for _, elem := range genState.PoolList {
		k.SetPool(ctx, elem)
	}
	// Set all the drop
	for _, elem := range genState.DropList {
		k.SetDrop(ctx, elem)
	}
	// Set all the member
	for _, elem := range genState.MemberList {
		k.SetMember(ctx, elem)
	}
	// Set all the burnings
	for _, elem := range genState.BurningsList {
		k.SetBurnings(ctx, elem)
	}
	// Set all the order
	for _, elem := range genState.OrderList {
		k.SetOrder(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PoolList = k.GetAllPool(ctx)
	genesis.DropList = k.GetAllDrop(ctx)
	genesis.MemberList = k.GetAllMember(ctx)
	genesis.BurningsList = k.GetAllBurnings(ctx)
	genesis.OrderList = k.GetAllOrder(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
