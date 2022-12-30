package rupxtn_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "rupx-tn/testutil/keeper"
	"rupx-tn/testutil/nullify"
	"rupx-tn/x/rupxtn"
	"rupx-tn/x/rupxtn/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RupxtnKeeper(t)
	rupxtn.InitGenesis(ctx, *k, genesisState)
	got := rupxtn.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
