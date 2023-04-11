package simulation_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	SimApp "github.com/pendulum-labs/market/app"
	marketsimulation "github.com/pendulum-labs/market/x/market/simulation"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"

	//SimulationState "github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	//keepertest "github.com/pendulum-labs/market/testutil/keeper"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type SimTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *SimApp.App
}

func (suite *SimTestSuite) SetupTest() {
	//TestInput := keepertest.CreateTestEnvironment(suite.T())

	simapp.FlagEnabledValue = true
	simapp.FlagCommitValue = true
	encoding := cosmoscmd.MakeEncodingConfig(SimApp.ModuleBasics)
	config, db, _, logger, _, _ := simapp.SetupSimulation("goleveldb-app-sim", "Simulation")

	checkTx := false
	app := SimApp.New(
		logger,
		db,
		nil,
		false,
		map[int64]bool{},
		SimApp.DefaultNodeHome,
		0,
		encoding,
		simapp.EmptyAppOptions{},
	)
	simApp, _ := app.(*SimApp.App)
	suite.app = simApp
	suite.ctx = simApp.BaseApp.NewContext(checkTx, tmproto.Header{})
	_, simParams, _ := simulation.SimulateFromSeed(
		suite.T(),
		os.Stdout,
		simApp.GetBaseApp(),
		simapp.AppStateFn(simApp.AppCodec(), simApp.SimulationManager()),
		simtypes.RandomAccounts,
		simapp.SimulationOperations(simApp, simApp.AppCodec(), config),
		simApp.ModuleAccountAddrs(),
		config,
		simApp.AppCodec(),
	)
	simapp.CheckExportSimulation(simApp, config, simParams)

}

func (suite *SimTestSuite) TestSimulateMsgCreatePool() {
	const (
		accCount       = 1
		moduleAccCount = 1
	)

	s := rand.NewSource(1)
	r := rand.New(s)
	accounts := suite.getTestingAccounts(r, accCount)

	// begin a new block
	suite.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: suite.app.LastBlockHeight() + 1, AppHash: suite.app.LastCommitID().Hash}})

	// execute operation
	op := marketsimulation.SimulateMsgCreatePool(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.MarketKeeper)

	s = rand.NewSource(1)
	r = rand.New(s)

	operationMsg, futureOperations, err := op(r, suite.app.BaseApp, suite.ctx, accounts, "")
	suite.Require().Error(err)

	var msg types.MsgSend
	types.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	suite.Require().False(operationMsg.OK)
	suite.Require().Equal(operationMsg.Comment, "invalid transfers")
	suite.Require().Equal(types.TypeMsgSend, msg.Type())
	suite.Require().Equal(types.ModuleName, msg.Route())
	suite.Require().Len(futureOperations, 0)
}

/*
	func (suite *SimTestSuite) TestWeightedOperations() {
		cdc := suite.app.AppCodec()
		appParams := make(simtypes.AppParams)

		weightesOps := simulation.WeightedOperations(appParams, cdc, suite.app.AccountKeeper, suite.app.BankKeeper)

}*/

func TestWeightedOperations(t *testing.T) {
	simapp.FlagEnabledValue = true
	simapp.FlagCommitValue = true

	config, db, dir, logger, _, err := simapp.SetupSimulation("goleveldb-app-sim", "Simulation")
	require.NoError(t, err, "simulation setup failed")

	t.Cleanup(func() {
		db.Close()
		err = os.RemoveAll(dir)
		require.NoError(t, err)
	})

	encoding := cosmoscmd.MakeEncodingConfig(SimApp.ModuleBasics)

	app := SimApp.New(
		logger,
		db,
		nil,
		false,
		map[int64]bool{},
		SimApp.DefaultNodeHome,
		0,
		encoding,
		simapp.EmptyAppOptions{},
	)
	simApp, ok := app.(*SimApp.App)
	require.True(t, ok, "can't use simapp")
	// Run randomized simulations
	_, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		simApp.GetBaseApp(),
		simapp.AppStateFn(simApp.AppCodec(), simApp.SimulationManager()),
		simtypes.RandomAccounts,
		simapp.SimulationOperations(simApp, simApp.AppCodec(), config),
		simApp.ModuleAccountAddrs(),
		config,
		simApp.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = simapp.CheckExportSimulation(simApp, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	if config.Commit {
		simapp.PrintStats(db)
	}

}

func (suite *SimTestSuite) getTestingAccounts(r *rand.Rand, n int) []simtypes.Account {
	accounts := simtypes.RandomAccounts(r, n)

	initAmt := suite.app.StakingKeeper.TokensFromConsensusPower(suite.ctx, 200)
	initCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, initAmt))

	// add coins to the accounts
	for _, account := range accounts {
		acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, account.Address)
		suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
		suite.Require().NoError(simapp.FundAccount(suite.app.BankKeeper, suite.ctx, account.Address, initCoins))
	}

	return accounts
}

func TestSimTestSuite(t *testing.T) {
	suite.Run(t, new(SimTestSuite))
}
