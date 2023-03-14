package util

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

func NewProposalJSON(messages []sdk.Msg) (string, error) {
	var proposal v1.MsgSubmitProposal

	anys, err := sdktx.SetMsgs(messages)
	if err != nil {
		return "", err
	}

	proposal.Messages = anys
	proposalJSON, err := Cdc.MarshalJSON(&proposal)
	return string(proposalJSON), nil
}
