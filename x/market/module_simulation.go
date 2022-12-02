package market

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/onomyprotocol/market/testutil/sample"
	marketsimulation "github.com/onomyprotocol/market/x/market/simulation"
	"github.com/onomyprotocol/market/x/market/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = marketsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreatePool = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePool int = 100

	opWeightMsgCreateDrop = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDrop int = 100

	opWeightMsgRedeemDrop = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRedeemDrop int = 100

	opWeightMsgCreateOrder = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOrder int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	marketGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&marketGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePool, &weightMsgCreatePool, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePool = defaultWeightMsgCreatePool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePool,
		marketsimulation.SimulateMsgCreatePool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateDrop int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDrop, &weightMsgCreateDrop, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDrop = defaultWeightMsgCreateDrop
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDrop,
		marketsimulation.SimulateMsgCreateDrop(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRedeemDrop int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRedeemDrop, &weightMsgRedeemDrop, nil,
		func(_ *rand.Rand) {
			weightMsgRedeemDrop = defaultWeightMsgRedeemDrop
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRedeemDrop,
		marketsimulation.SimulateMsgRedeemDrop(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateOrder, &weightMsgCreateOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOrder = defaultWeightMsgCreateOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOrder,
		marketsimulation.SimulateMsgCreateOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
