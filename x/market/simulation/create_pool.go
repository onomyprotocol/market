package simulation

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	//"github.com/pendulum-labs/market/testutil/sample"

	"github.com/pendulum-labs/market/testutil/sample"
	"github.com/pendulum-labs/market/x/market/keeper"
	"github.com/pendulum-labs/market/x/market/types"
)

func SimulateMsgCreatePool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		//addr := sample.AccAddress()
		simAccount, _ := simtypes.RandomAcc(r, accs)
		addr := simAccount.Address

		//requestAddress, _ := sdk.AccAddressFromBech32(addr)
		coinA := sdk.NewCoin("Coin"+sample.RandomString(1), sdk.NewInt(int64(r.Intn(200)))).String()
		coinB := sdk.NewCoin("Coin"+sample.RandomString(1), sdk.NewInt(int64(r.Intn(200)))).String()
		msg := &types.MsgCreatePool{
			Creator: addr.String(),
			CoinA:   coinA,
			CoinB:   coinB,
			RateA:   []string{sdk.NewInt(int64(r.Intn(200))).String(), sdk.NewInt(int64(r.Intn(200))).String()},
			RateB:   []string{sdk.NewInt(int64(r.Intn(200))).String(), sdk.NewInt(int64(r.Intn(200))).String()},
		}
		coins, _ := sample.SampleCoins(msg.CoinA, msg.CoinB)
		//coins = simtypes.RandSubsetCoins(r, coins)
		denomA, denomB := sample.SampleDenoms(coins)
		pair := strings.Join([]string{denomA, denomB}, ",")

		mintCoinsError := bk.MintCoins(ctx, types.ModuleName, coins)
		if mintCoinsError != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Mint Coins Error"), nil, mintCoinsError
		}

		sendCoinsFromModuleToAccountError := bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, coins)
		if sendCoinsFromModuleToAccountError != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SendCoinsFromModuleToAccount Error"), nil, sendCoinsFromModuleToAccountError
		}

		_, err := keeper.NewMsgServerImpl(k).CreatePool(sdk.WrapSDKContext(ctx), msg)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, err
		}

		_, found := keeper.Keeper.GetPool(k, ctx, pair)

		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Pool Not Found"), nil, err
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreatePool simulation"), nil, nil
	}
}
