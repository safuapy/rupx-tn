package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "rupx-tn/testutil/keeper"
	"rupx-tn/x/rupxtn/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RupxtnKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
