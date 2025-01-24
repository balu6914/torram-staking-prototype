package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/balu6914/torram/testutil/keeper"
	"github.com/balu6914/torram/x/runestaking/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RunestakingKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
