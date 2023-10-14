package keeper

import (
	"context"

	"blog/x/blog/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPost(goCtx context.Context, req *types.QueryListPostRequest) (*types.QueryListPostResponse, error) {
	// Error handling: No request/value is provided
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Initialize variable to store posts
	var posts []types.Post

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Initialize Store with the Key-Value Store associated with the ctx and keyprefix
	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.KeyPrefix(types.PostKey))

	// Function to process posts
	processPost := func(key []byte, value []byte) error {
		var post types.Post

		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}

		posts = append(posts, post)
		return nil
	}

	// Retrieve posts from a data store one page at a time based on the pagination parameters specified by request
	pageRes, err := query.Paginate(postStore, req.Pagination, processPost)

	// Error Handling
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return posts if successful
	return &types.QueryListPostResponse{Post: posts, Pagination: pageRes}, nil
}
