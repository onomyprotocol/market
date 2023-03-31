package keeper

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	ccodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/pendulum-labs/market/x/market/keeper"
	markettypes "github.com/pendulum-labs/market/x/market/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	tmdb "github.com/tendermint/tm-db"
)

var (
	// ModuleBasics is a mock module basic manager for testing
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		mint.AppModuleBasic{},
		params.AppModuleBasic{},
	)
)

// TestInput stores the various keepers required to test the exchange
type TestInput struct {
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.BaseKeeper
	Context       sdk.Context
	Marshaler     codec.Codec
	MarketKeeper  *keeper.Keeper
	LegacyAmino   *codec.LegacyAmino
}

// MakeTestLegacyCodec creates a legacy codec for use in testing
func MakeTestLegacyCodec() *codec.LegacyAmino {
	var cdc = codec.NewLegacyAmino()
	auth.AppModuleBasic{}.RegisterLegacyAminoCodec(cdc)
	bank.AppModuleBasic{}.RegisterLegacyAminoCodec(cdc)

	sdk.RegisterLegacyAminoCodec(cdc)
	ccodec.RegisterCrypto(cdc)
	params.AppModuleBasic{}.RegisterLegacyAminoCodec(cdc)
	markettypes.RegisterCodec(cdc)
	return cdc
}

// MakeTestCodec creates a proto codec for use in testing
func MakeTestCodec() codec.Codec {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	ModuleBasics.RegisterInterfaces(interfaceRegistry)
	markettypes.RegisterInterfaces(interfaceRegistry)
	return codec.NewProtoCodec(interfaceRegistry)
}

func CreateTestEnvironment(t testing.TB) TestInput {
	//poolKey := sdk.NewKVStoreKey(markettypes.PoolKeyPrefix)
	storeKey := sdk.NewKVStoreKey(markettypes.StoreKey)
	keyAuth := sdk.NewKVStoreKey(authtypes.StoreKey)
	keyBank := sdk.NewKVStoreKey(banktypes.StoreKey)
	keyParams := sdk.NewKVStoreKey(paramstypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(markettypes.MemStoreKey)
	tkeyParams := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	db := tmdb.NewMemDB()

	stateStore := store.NewCommitMultiStore(db)

	//stateStore.MountStoreWithDB(poolKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyAuth, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyBank, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	//ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	ctx := sdk.NewContext(stateStore, tmproto.Header{
		Version: tmversion.Consensus{
			Block: 0,
			App:   0,
		},
		ChainID: "",
		Height:  1234567,
		Time:    time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
		LastBlockId: tmproto.BlockID{
			Hash: []byte{},
			PartSetHeader: tmproto.PartSetHeader{
				Total: 0,
				Hash:  []byte{},
			},
		},
		LastCommitHash:     []byte{},
		DataHash:           []byte{},
		ValidatorsHash:     []byte{},
		NextValidatorsHash: []byte{},
		ConsensusHash:      []byte{},
		AppHash:            []byte{},
		LastResultsHash:    []byte{},
		EvidenceHash:       []byte{},
		ProposerAddress:    []byte{},
	}, false, log.TestingLogger())

	cdc := MakeTestCodec()
	legacyCodec := MakeTestLegacyCodec()

	paramsKeeper := paramskeeper.NewKeeper(cdc, legacyCodec, keyParams, tkeyParams)
	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(markettypes.ModuleName)

	paramsSubspace := paramstypes.NewSubspace(cdc,
		markettypes.Amino,
		storeKey,
		memStoreKey,
		"MarketParams",
	)
	// this is also used to initialize module accounts for all the map keys
	maccPerms := map[string][]string{
		markettypes.ModuleName:     {authtypes.Minter},
		authtypes.FeeCollectorName: nil,
	}

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		keyAuth, // target store
		getSubspace(paramsKeeper, authtypes.ModuleName),
		authtypes.ProtoBaseAccount, // prototype
		maccPerms,
	)

	blockedAddr := make(map[string]bool, len(maccPerms))
	for acc := range maccPerms {
		blockedAddr[authtypes.NewModuleAddress(acc).String()] = true
	}
	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		keyBank,
		accountKeeper,
		getSubspace(paramsKeeper, banktypes.ModuleName),
		blockedAddr,
	)
	bankKeeper.SetParams(ctx, banktypes.Params{
		SendEnabled:        []*banktypes.SendEnabled{},
		DefaultSendEnabled: true,
	})
	marketKeeper := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		bankKeeper,
	)
	// Initialize params
	//marketKeeper.setID
	marketKeeper.SetParams(ctx, markettypes.DefaultParams())

	return TestInput{
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
		Context:       ctx,
		Marshaler:     cdc,
		LegacyAmino:   legacyCodec,
		MarketKeeper:  marketKeeper,
	}
}

// getSubspace returns a param subspace for a given module name.
func getSubspace(k paramskeeper.Keeper, moduleName string) paramstypes.Subspace {
	subspace, _ := k.GetSubspace(moduleName)
	return subspace
}
