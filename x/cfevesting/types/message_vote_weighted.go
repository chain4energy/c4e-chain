package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

)

const TypeMsgVoteWeighted = "vote_weighted"

var _ sdk.Msg = &MsgVoteWeighted{}

func NewMsgVoteWeighted(voter string, proposalId uint64, options govtypes.WeightedVoteOptions) *MsgVoteWeighted {
	return &MsgVoteWeighted{
		Voter:      voter,
		ProposalId: proposalId,
		Options:    options,
	}
}

func (msg *MsgVoteWeighted) Route() string {
	return RouterKey
}

func (msg *MsgVoteWeighted) Type() string {
	return TypeMsgVoteWeighted
}

func (msg *MsgVoteWeighted) GetSigners() []sdk.AccAddress {
	voter, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{voter}
}

func (msg *MsgVoteWeighted) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteWeighted) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Voter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid voter address (%s)", err)
	}
	return nil
}
