package torram_test

import (
	"testing"

	keepertest "github.com/balu6914/torram/testutil/keeper"
	"github.com/balu6914/torram/testutil/nullify"
	torram "github.com/balu6914/torram/x/torram/module"
	"github.com/balu6914/torram/x/torram/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TorramKeeper(t)
	torram.InitGenesis(ctx, k, genesisState)
	got := torram.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
