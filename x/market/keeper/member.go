package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pendulum-labs/market/x/market/types"
)

// TODO(aaron) show Charles all places where `SetMember` and `GetMember` are called, it seems like
// they should come in pairs but sometimes aren't

// SetMember set a specific member in the store from its index
func (k Keeper) SetMember(ctx sdk.Context, member types.Member) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKeyPrefix))
	b := k.cdc.MustMarshal(&member)
	store.Set(types.MemberSetKey(
		member.DenomA,
		member.DenomB,
		// TODO(aaron) what about this?
		//member.Balance,
		//member.Previous,
		//member.Limit,
		//member.Stop,
		//member.Protect,
	), b)
}

// GetMember returns a member from its index
func (k Keeper) GetMember(
	ctx sdk.Context,
	denomA string,
	denomB string,

) (val types.Member, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKeyPrefix))

	b := store.Get(types.MemberKey(
		denomA,
		denomB,
	))
	if b == nil {
		// TODO(aaron) `val` is not initialized here, should use some invalid or empty constant?
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetMemberWithPair(
	ctx sdk.Context,
	pair string,
	denomA string,
	denomB string,

) (val types.Member, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKeyPrefix))

	b := store.Get(types.MemberKeyPair(
		pair,
		denomA,
		denomB,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMember removes a member from the store
func (k Keeper) RemoveMember(
	ctx sdk.Context,
	denomA string,
	denomB string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKeyPrefix))
	store.Delete(types.MemberKey(
		denomA,
		denomB,
	))
}

// GetAllMember returns all member
func (k Keeper) GetAllMember(ctx sdk.Context) (list []types.Member) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MemberKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Member
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
