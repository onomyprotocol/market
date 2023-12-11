package simulation

import (
	"math/rand"

	"market/testutil/sample"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	//"market/testutil/sample"

	"market/x/market/keeper"
	"market/x/market/types"
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
			CoinA:   sdk.NewCoin("CoinA", sdk.NewInt(160)).String(),
			CoinB:   sdk.NewCoin("CoinB", sdk.NewInt(170)).String(),
		}
		coins, _ := sample.SampleCoins(msg.CoinA, msg.CoinB)

		bk.MintCoins(ctx, types.ModuleName, coins)

		bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, coins)

		_, err := keeper.NewMsgServerImpl(k).CreatePool(sdk.WrapSDKContext(ctx), msg)

		//simtypes.
		//err := sendMsgSend(r, app, bk, ak, k, *msg, ctx, chainID, []cryptotypes.PrivKey{creatorAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, err
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreatePool simulation"), nil, nil
	}
}

/*
func sendMsgSend(
	r *rand.Rand, app *baseapp.BaseApp, bk types.BankKeeper, ak types.AccountKeeper, k keeper.Keeper,
	msg types.MsgCreatePool, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr := sample.AccAddress()
	requestAddress, err := sdk.AccAddressFromBech32(addr)
	//coina := msg.GetCoinA()
	//coinb  := msg.GetCoinB()
	coins,_ := sample.SampleCoins(msg.GetCoinA(), msg.GetCoinB())
	//fees, err := simtypes.RandomFees(r, ctx, coins)

	bk.MintCoins(ctx, types.ModuleName, coins)

	err = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, requestAddress, coins)
	if err != nil {
		return err
	}

	_, err = keeper.NewMsgServerImpl(k).CreatePool(sdk.WrapSDKContext(ctx), &msg)
	if err != nil {
		return err
	}
	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{&msg},
		coins,
		helpers.DefaultGenTxGas,
		chainID,
		[]uint64{requestAddress.GetAccountNumber()},
		[]uint64{requestAddress.GetSequence()},
		privkeys...,
	)
	if err != nil {
		return err
	}
	_, _, err = app.Deliver(txGen.TxEncoder(), tx)
	if err != nil {
		return err
	}
	return nil
}*/
