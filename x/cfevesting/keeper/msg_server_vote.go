package keeper

import (
	"context"
	"fmt"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	keeper := k.Keeper

	accVestings, found := keeper.GetAccountVestings(ctx, msg.Voter)
	if !found {
		return nil, fmt.Errorf("no vestings for account: %q", msg.Voter)
	}
	if len(accVestings.DelegableAddress) == 0 {
		return nil, fmt.Errorf("no delegable vestings for account: %q", msg.Voter)
	}

	_, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
	if err != nil {
		return nil, err
	}

	_, err = sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		return nil, err
	}

	voteMsg := govtypes.MsgVote{ProposalId: msg.ProposalId,
		Voter: accVestings.DelegableAddress, Option: msg.Option}
	_, err = k.govMsgServer.Vote(goCtx, &voteMsg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Voter),
		),
	})

	return &types.MsgVoteResponse{}, nil
}
