package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Create Post message handler
func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Extract message received and create a Post object
	var post = types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	// Call AppendPost method to add new post to block chain
	id := k.AppendPost(
		ctx,
		post,
	)

	// Return a response indicating the status of the operation
	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
