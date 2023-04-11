package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

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
		creatorAccount := accs[0]
		msg := &types.MsgCreatePool{
			Creator: creatorAccount.Address.String(),
			CoinA:   sdk.NewCoin("CoinA", sdk.NewInt(150)).String(),
			CoinB:   sdk.NewCoin("CoinB", sdk.NewInt(150)).String(),
			RateA:   []string{sdk.NewInt(150).String(), sdk.NewInt(150).String()},
			RateB:   []string{sdk.NewInt(150).String(), sdk.NewInt(150).String()},
		}
		err := sendMsgSend(r, app, bk, ak, msg, ctx, chainID, []cryptotypes.PrivKey{creatorAccount.PrivKey})
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "invalid transfers"), nil, err
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreatePool simulation"), nil, nil
	}
}

func sendMsgSend(
	r *rand.Rand, app *baseapp.BaseApp, bk types.BankKeeper, ak types.AccountKeeper,
	msg sdk.Msg, ctx sdk.Context, chainID string, privkeys []cryptotypes.PrivKey,
) error {
	addr := msg.GetSigners()[0]
	account := ak.GetAccount(ctx, addr)
	coins := bk.SpendableCoins(ctx, account.GetAddress())
	fees, err := simtypes.RandomFees(r, ctx, coins)
	if err != nil {
		return err
	}
	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{msg},
		fees,
		helpers.DefaultGenTxGas,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
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
}
