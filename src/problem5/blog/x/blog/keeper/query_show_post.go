package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowPost(goCtx context.Context, req *types.QueryShowPostRequest) (*types.QueryShowPostResponse, error) {
	// Error handling: No request/value is provided
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Post from Store
	post, found := k.GetPost(ctx, req.Id)

	// Error Handling: Post not found
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	// Return a pointer to the associated Post
	return &types.QueryShowPostResponse{Post: post}, nil
}
