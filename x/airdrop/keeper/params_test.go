package keeper_test

import (
	"testing"

	testkeeper "AstraNetwork/testutil/keeper"
	"AstraNetwork/x/airdrop/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AirdropKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
