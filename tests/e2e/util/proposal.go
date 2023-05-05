package util

import (
	"encoding/json"
	cfeclaimmoduletypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
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

func NewClaimRecordsListJson(userEntries []cfeclaimmoduletypes.ClaimRecord) (string, error) {
	proposalJSON, err := json.Marshal(userEntries)
	if err != nil {
		return "", err
	}
	return string(proposalJSON), nil
}
