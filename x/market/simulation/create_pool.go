package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
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
		simAccount, _ := simtypes.RandomAcc(r, accs)
		addr := simAccount.Address.String()
		msg := &types.MsgCreatePool{
			Creator: simAccount.Address.String(),
			CoinA:   sdk.NewCoin("CoinA", sdk.NewInt(150)).String(),
			CoinB:   sdk.NewCoin("CoinB", sdk.NewInt(150)).String(),
			RateA:   []string{},
			RateB:   []string{},
		}
		coinPair, _ := sample.SampleCoins(msg.CoinA, msg.CoinB)

		bk.MintCoins(ctx, types.ModuleName, coinPair)
		//SendCoinsFromModuleToAccount
		requestAddress, _ := sdk.AccAddressFromBech32(addr)

		bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, requestAddress, coinPair)

		keeper.NewMsgServerImpl(k).CreatePool(sdk.WrapSDKContext(ctx), msg)

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreatePool simulation not implemented"), nil, nil
	}
}
