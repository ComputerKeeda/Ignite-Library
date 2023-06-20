package library_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "library/testutil/keeper"
	"library/testutil/nullify"
	"library/x/library"
	"library/x/library/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LibraryKeeper(t)
	library.InitGenesis(ctx, *k, genesisState)
	got := library.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
