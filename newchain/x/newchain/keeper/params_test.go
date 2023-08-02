package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "newchain/testutil/keeper"
	"newchain/x/newchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NewchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
