package simulator

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (s *Simulator) RunQuery(requestData interface{}, path string, responseValue interface{}) abci.ResponseQuery {
	resp := s.app.Query(abci.RequestQuery{
		Data: codec.MustMarshalJSONIndent(s.cdc, requestData),
		Path: path,
	})

	if responseValue != nil && resp.IsOK() {
		require.NoError(s.t, s.cdc.UnmarshalJSON(resp.Value, responseValue))
	}

	return resp
}

// QueryAuthAccount queries account by address.
func (s *Simulator) QueryAuthAccount(addr sdk.AccAddress) (res auth.BaseAccount) {
	resp := s.RunQuery(
		auth.QueryAccountParams{
			Address: addr,
		},
		"/custom/"+auth.QuerierRoute+"/"+auth.QueryAccount,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryStakingValidators queries all validators with specified status and pagination.
func (s *Simulator) QueryStakingValidators(page, limit int, status string) []staking.Validator {
	res := make([]staking.Validator, 0)
	resp := s.RunQuery(
		staking.QueryValidatorsParams{
			Page:   page,
			Limit:  limit,
			Status: status,
		},
		"/custom/"+staking.QuerierRoute+"/"+staking.QueryValidators,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryStakingValidator queries validator by its operator address.
func (s *Simulator) QueryStakingValidator(valAddr sdk.ValAddress) (res staking.Validator) {
	resp := s.RunQuery(
		staking.QueryValidatorParams{
			ValidatorAddr: valAddr,
		},
		"/custom/"+staking.QuerierRoute+"/"+staking.QueryValidator,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryStakingPool queries bonded / unbonded pool state.
func (s *Simulator) QueryStakingPool() (res staking.Pool) {
	resp := s.RunQuery(
		nil,
		"/custom/"+staking.QuerierRoute+"/"+staking.QueryPool,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryStakingDelegation queries delegation for specified delegator and validator.
func (s *Simulator) QueryStakingDelegation(simAcc *SimAccount, val *staking.Validator) (res staking.DelegationResponse) {
	require.NotNil(s.t, simAcc)
	require.NotNil(s.t, val)

	resp := s.RunQuery(
		staking.QueryBondsParams{
			DelegatorAddr: simAcc.Address,
			ValidatorAddr: val.OperatorAddress,
		},
		"/custom/"+staking.QuerierRoute+"/"+staking.QueryDelegation,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryMintParams queries mint module parameters.
func (s *Simulator) QueryMintParams() (res mint.Params) {
	resp := s.RunQuery(
		nil,
		"/custom/"+mint.QuerierRoute+"/"+mint.QueryParameters,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryMintParams queries mint module annual provisions.
func (s *Simulator) QueryMintAnnualProvisions() (res sdk.Dec) {
	resp := s.RunQuery(
		nil,
		"/custom/"+mint.QuerierRoute+"/"+mint.QueryAnnualProvisions,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryMintParams queries mint module annual provisions.
func (s *Simulator) QueryMintBlocksPerYearEstimation() (res uint64) {
	resp := s.RunQuery(
		nil,
		"/custom/"+mint.QuerierRoute+"/"+mint.QueryBlocksPerYear,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryDistPool queries distribution module pool supply.
func (s *Simulator) QueryDistPool(poolName distribution.RewardPoolName) (res sdk.DecCoins) {
	resp := s.RunQuery(
		nil,
		"/custom/"+distribution.QuerierRoute+"/"+distribution.QueryPool+"/"+poolName.String(),
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}

// QueryDistPool queries distribution module pool supply.
func (s *Simulator) QuerySupplyTotal() (res sdk.Coins) {
	resp := s.RunQuery(
		supply.QueryTotalSupplyParams{
			Page:  1,
			Limit: 50,
		},
		"/custom/"+supply.QuerierRoute+"/"+supply.QueryTotalSupply,
		&res,
	)
	require.True(s.t, resp.IsOK())

	return res
}