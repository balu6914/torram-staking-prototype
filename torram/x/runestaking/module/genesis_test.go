package runestaking_test

import (
	"testing"

	keepertest "github.com/balu6914/torram/testutil/keeper"
	"github.com/balu6914/torram/testutil/nullify"
	runestaking "github.com/balu6914/torram/x/runestaking/module"
	"github.com/balu6914/torram/x/runestaking/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RunestakingKeeper(t)
	runestaking.InitGenesis(ctx, k, genesisState)
	got := runestaking.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
