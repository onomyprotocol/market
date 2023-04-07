package simulation_test

import (
	"os"
	"testing"

	SimApp "github.com/pendulum-labs/market/app"
	//market "github.com/pendulum-labs/market/x/market"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"

	"github.com/cosmos/cosmos-sdk/simapp"
	//SimulationState "github.com/cosmos/cosmos-sdk/types/module"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	//keepertest "github.com/pendulum-labs/market/testutil/keeper"
)

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
		true,
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
		simulationtypes.RandomAccounts,
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
