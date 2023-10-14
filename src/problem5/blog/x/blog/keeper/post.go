package keeper

import (
	"encoding/binary"

	"blog/x/blog/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Appends new Post to back of list
func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)
	post.Id = count

	// Initialize Store with the Key-Value Store associated with the ctx and keyprefix
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))

	// Serialize Post to store in blockchain
	appendedValue := k.cdc.MustMarshal(&post)

	// Store serialized Post and associate it to its Id
	store.Set(GetPostIDBytes(post.Id), appendedValue)
	k.SetPostCount(ctx, count+1)
	return count
}

// Retrieve the current number of posts stored in database.
func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// Initialize Store with the Key-Value Store associated with the ctx
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})

	// Get key pointing to latest count
	byteKey := types.KeyPrefix(types.PostCountKey)

	// Retrieve Post Count associated with byteKey
	cnt := store.Get(byteKey)

	// Error checking
	if cnt == nil {
		return 0
	}
	return binary.BigEndian.Uint64(cnt)
}

// Convert post Id to byte array
func GetPostIDBytes(id uint64) []byte {
	temp := make([]byte, 8)

	// Put id into temp in BigEndian order
	binary.BigEndian.PutUint64(temp, id)
	return temp
}

// Update the post count stored in the database
func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {

	// Initialize Store with the Key-Value Store associated with the ctx
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})

	// Get key pointing to latest count
	byteKey := types.KeyPrefix(types.PostCountKey)

	// Convert count into binary form
	temp := make([]byte, 8)
	binary.BigEndian.PutUint64(temp, count)

	// Store count and associate it to bytekey
	store.Set(byteKey, temp)
}

// Get Post from store associated with the given id
func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
	// Initialize Store with the Key-Value Store associated with the ctx and keyprefix
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))

	// Obtain post from store
	post := store.Get(GetPostIDBytes(id))

	// Error handling: Post not found
	if post == nil {
		return val, false
	}

	// Convert byte slice obtained into Post object
	k.cdc.MustUnmarshal(post, &val)
	return val, true
}

// Update existing Post
func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	// Initialize Store with the Key-Value Store associated with the ctx and keyprefix
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))

	// Convert Post object into byte slice
	updatedPost := k.cdc.MustMarshal(&post)

	// Update Post associated with post Id
	store.Set(GetPostIDBytes(post.Id), updatedPost)
}

// Delete existing Post
func (k Keeper) RemovePost(ctx sdk.Context, id uint64) {
	// Initialize Store with the Key-Value Store associated with the ctx and keyprefix
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))

	// Delete Post associated with post Id
	store.Delete(GetPostIDBytes(id))
}
