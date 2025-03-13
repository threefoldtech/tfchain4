package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "tfchain4/testutil/keeper"
	"tfchain4/x/tfchain4/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.Tfchain4Keeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
