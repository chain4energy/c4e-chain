package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/chain4energy/c4e-chain/x/compatibility/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakingV3DelegatorDelegations(goCtx context.Context, req *types.QueryStakingV3DelegatorDelegationsRequest) (*types.QueryStakingV3DelegatorDelegationsResponse, error) {

	// sk := k.stakingKeeper
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.DelegatorAddr == "" {
		return nil, status.Error(codes.InvalidArgument, "delegator address cannot be empty")
	}
	var delegations stakingtypes.Delegations
	ctx := sdk.UnwrapSDKContext(goCtx)

	delAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddr)
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(k.stakingStoreKey)
	delStore := prefix.NewStore(store, stakingtypes.GetDelegationsKey(delAddr))
	pageRes, err := query.Paginate(delStore, req.Pagination, func(key []byte, value []byte) error {
		delegation, err := stakingtypes.UnmarshalDelegation(k.cdc, value)
		if err != nil {
			return err
		}
		delegations = append(delegations, delegation)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	delegationResps, err := DelegationsToDelegationResponses(ctx, k.stakingKeeper, delegations)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryStakingV3DelegatorDelegationsResponse{DelegationResponses: delegationResps, Pagination: pageRes}, nil

	// TODO: Process the query
	// _ = ctx

	// return &types.QueryStakingV3DelegatorDelegationsResponse{}, nil
}

func DelegationsToDelegationResponses(ctx sdk.Context, k types.StakingKeeper, delegations stakingtypes.Delegations) (stakingtypes.DelegationResponses, error) {
	resp := make(stakingtypes.DelegationResponses, len(delegations))

	for i, del := range delegations {
		delResp, err := DelegationToDelegationResponse(ctx, k, del)
		if err != nil {
			return nil, err
		}

		resp[i] = delResp
	}

	return resp, nil
}

func DelegationToDelegationResponse(ctx sdk.Context, k types.StakingKeeper, del stakingtypes.Delegation) (stakingtypes.DelegationResponse, error) {
	val, found := k.GetValidator(ctx, del.GetValidatorAddr())
	if !found {
		return stakingtypes.DelegationResponse{}, stakingtypes.ErrNoValidatorFound
	}

	delegatorAddress, err := sdk.AccAddressFromBech32(del.DelegatorAddress)
	if err != nil {
		return stakingtypes.DelegationResponse{}, err
	}

	denom := k.BondDenom(ctx)
	if denom == "" {
		denom = "uc4e" // TODO: make it configurable
	}

	return stakingtypes.NewDelegationResp(
		delegatorAddress,
		del.GetValidatorAddr(),
		del.Shares,
		sdk.NewCoin(denom, val.TokensFromShares(del.Shares).TruncateInt()),
	), nil
}
