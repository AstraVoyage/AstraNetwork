package keeper

import (
	"AstraNetwork/x/airdrop/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetClaimedCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ClaimedCountKey))

	// Convert the PostCountKey to bytes
	byteKey := []byte(types.ClaimedCountKey)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetClaimedCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "airdrop") and ClaimedCountKey (which is "Claimed/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ClaimedCountKey))

	// Convert the ClaimedCountKey to bytes
	byteKey := []byte(types.ClaimedCountKey)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of Claimed/count/ to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendClaimed(ctx sdk.Context, claimed types.Claimed) uint64 {
	// Get the current number of claimed in the store
	count := k.GetClaimedCount(ctx)

	// Assign an ID to the claimed based on the number of claimed in the store
	claimed.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ClaimedKey))

	// Convert the claimed ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, claimed.Id)

	// Marshal the claimed into bytes
	appendedValue := k.cdc.MustMarshal(&claimed)

	// Insert the claimed bytes using claimed ID as a key
	store.Set(byteKey, appendedValue)

	// Update the claimed count
	k.SetClaimedCount(ctx, count+1)
	return count
}
