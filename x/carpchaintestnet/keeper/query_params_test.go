package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/carpchain/carpchain-testnet/testutil/keeper"
	"github.com/carpchain/carpchain-testnet/x/carpchaintestnet/keeper"
	"github.com/carpchain/carpchain-testnet/x/carpchaintestnet/types"
)

func TestParamsQuery(t *testing.T) {
	k, ctx := keepertest.CarpchaintestnetKeeper(t)

	qs := keeper.NewQueryServerImpl(k)
	params := types.DefaultParams()
	require.NoError(t, k.Params.Set(ctx, params))

	response, err := qs.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
