package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "tfchain4/testutil/keeper"
	"tfchain4/x/tfchain4/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.Tfchain4Keeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
