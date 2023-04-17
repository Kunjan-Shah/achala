package keeper_test

import (
	"context"
	"testing"

	keepertest "achala/testutil/keeper"
	"achala/x/achala/keeper"
	"achala/x/achala/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AchalaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
