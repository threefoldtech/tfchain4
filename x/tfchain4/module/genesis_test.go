package tfchain4_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "tfchain4/testutil/keeper"
	"tfchain4/testutil/nullify"
	"tfchain4/x/tfchain4/module"
	"tfchain4/x/tfchain4/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Tfchain4Keeper(t)
	tfchain4.InitGenesis(ctx, k, genesisState)
	got := tfchain4.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
