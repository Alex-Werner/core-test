package newchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "newchain/testutil/keeper"
	"newchain/testutil/nullify"
	"newchain/x/newchain"
	"newchain/x/newchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NewchainKeeper(t)
	newchain.InitGenesis(ctx, *k, genesisState)
	got := newchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
