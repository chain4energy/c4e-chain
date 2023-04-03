package types

import (
	"cosmossdk.io/errors"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const TypeMsgUpdateMintersParams = "update_minters_params"

var _ sdk.Msg = &MsgUpdateMintersParams{}

func (msg *MsgUpdateMintersParams) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMintersParams) Type() string {
	return TypeMsgUpdateMintersParams
}

func (msg *MsgUpdateMintersParams) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMintersParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMintersParams) ValidateBasic() error {
	if msg.Authority != appparams.GetAuthority() {
		return govtypes.ErrInvalidSigner
	}

	params := Params{
		StartTime: msg.StartTime,
		Minters:   msg.Minters,
	}
	if err := params.ValidateParamsMinters(); err != nil {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}
	return nil
}

const TypeMsgUpdateParams = "update_params"

var _ sdk.Msg = &MsgUpdateParams{}

func (msg *MsgUpdateParams) Route() string {
	return RouterKey
}

func (msg *MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}

func (msg *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateParams) ValidateBasic() error {
	if msg.Authority != appparams.GetAuthority() {
		return govtypes.ErrInvalidSigner
	}

	params := Params{
		MintDenom: msg.MintDenom,
		StartTime: msg.StartTime,
		Minters:   msg.Minters,
	}
	if err := params.Validate(); err != nil {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}
	return nil
}
