package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/balu6914/torram/testutil/keeper"
	"github.com/balu6914/torram/x/torram/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.TorramKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
