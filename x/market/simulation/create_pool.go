package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/pendulum-labs/market/testutil/sample"

	//"github.com/pendulum-labs/market/testutil/sample"

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

		msg := &types.MsgCreatePool{
			Creator: addr.String(),
			CoinA:   sdk.NewCoin(sample.RandomString(5), sdk.NewInt(int64(r.Intn(200)))).String(),
			CoinB:   sdk.NewCoin(sample.RandomString(5), sdk.NewInt(int64(r.Intn(200)))).String(),
			RateA:   []string{sdk.NewInt(int64(r.Intn(200))).String(), sdk.NewInt(int64(r.Intn(200))).String()},
			RateB:   []string{sdk.NewInt(int64(r.Intn(200))).String(), sdk.NewInt(int64(r.Intn(200))).String()},
		}
		coins, _ := sample.SampleCoins(msg.CoinA, msg.CoinB)

		bk.MintCoins(ctx, types.ModuleName, coins)

		bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, coins)

		_, err := keeper.NewMsgServerImpl(k).CreatePool(sdk.WrapSDKContext(ctx), msg)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, err
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreatePool simulation"), nil, nil
	}
}
