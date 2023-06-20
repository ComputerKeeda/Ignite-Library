package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "library/testutil/keeper"
	"library/x/library/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LibraryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
