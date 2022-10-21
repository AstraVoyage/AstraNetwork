package keeper_test

import (
	"context"
	"testing"

	keepertest "AstraNetwork/testutil/keeper"
	"AstraNetwork/x/airdrop/keeper"
	"AstraNetwork/x/airdrop/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AirdropKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
