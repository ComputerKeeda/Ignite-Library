package keeper

import (
    "encoding/binary"

    "library/x/library/types"

    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendAuthor(ctx sdk.Context, author types.Author) uint64 {
    count := k.GetAuthorCount(ctx)
    author.Id = count
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorKey))
    appendedValue := k.cdc.MustMarshal(&author)
    store.Set(GetAuthorIDBytes(author.Id), appendedValue)
    k.SetAuthorCount(ctx, count+1)
    return count
}

func (k Keeper) GetAuthorCount(ctx sdk.Context) uint64 {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
    byteKey := types.KeyPrefix(types.AuthorCountKey)
    bz := store.Get(byteKey)
    if bz == nil {
        return 0
    }
    return binary.BigEndian.Uint64(bz)
}

func GetAuthorIDBytes(id uint64) []byte {
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, id)
    return bz
}

func (k Keeper) SetAuthorCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AuthorCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}