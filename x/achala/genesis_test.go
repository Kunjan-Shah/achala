package achala_test

import (
	"testing"

	keepertest "achala/testutil/keeper"
	"achala/testutil/nullify"
	"achala/x/achala"
	"achala/x/achala/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AchalaKeeper(t)
	achala.InitGenesis(ctx, *k, genesisState)
	got := achala.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
