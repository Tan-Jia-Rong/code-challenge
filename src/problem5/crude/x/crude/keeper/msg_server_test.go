package keeper_test

import (
	"context"
	"testing"

	keepertest "crude/testutil/keeper"
	"crude/x/crude/keeper"
	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CrudeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
