package carpchaintestnet_test

import (
	"testing"

	keepertest "github.com/carpchain/carpchain-testnet/testutil/keeper"
	"github.com/carpchain/carpchain-testnet/testutil/nullify"
	carpchaintestnet "github.com/carpchain/carpchain-testnet/x/carpchaintestnet/module"
	"github.com/carpchain/carpchain-testnet/x/carpchaintestnet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CarpchaintestnetKeeper(t)
	carpchaintestnet.InitGenesis(ctx, k, genesisState)
	got := carpchaintestnet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
