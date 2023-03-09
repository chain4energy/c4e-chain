package types

import (
	"cosmossdk.io/errors"
	appparams "github.com/chain4energy/c4e-chain/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const TypeMsgUpdateAllSubDistributorsParams = "update_all_subdistributors_params"

var _ sdk.Msg = &MsgUpdateAllSubDistributorsParams{}

func (msg *MsgUpdateAllSubDistributorsParams) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAllSubDistributorsParams) Type() string {
	return TypeMsgUpdateAllSubDistributorsParams
}

func (msg *MsgUpdateAllSubDistributorsParams) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAllSubDistributorsParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAllSubDistributorsParams) ValidateBasic() error {
	if msg.Authority != appparams.GetAuthority() {
		return govtypes.ErrInvalidSigner
	}
	err := Params{SubDistributors: msg.SubDistributors}.Validate()
	if err != nil {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}
	return nil
}

const TypeMsgUpdateSubDistributorParam = "update_single_subdistributor_param"

var _ sdk.Msg = &MsgUpdateSubDistributorParam{}

func (msg *MsgUpdateSubDistributorParam) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSubDistributorParam) Type() string {
	return TypeMsgUpdateSubDistributorParam
}

func (msg *MsgUpdateSubDistributorParam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSubDistributorParam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSubDistributorParam) ValidateBasic() error {
	if msg.Authority != appparams.GetAuthority() {
		return govtypes.ErrInvalidSigner
	}

	if err := msg.SubDistributor.Validate(); err != nil {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "validation error: %s", err)
	}
	return nil
}

const TypeMsgUpdateSubDistributorBurnShareParam = "update_sub_distributor_burn_share_param"

var _ sdk.Msg = &MsgUpdateSubDistributorBurnShareParam{}

func (msg *MsgUpdateSubDistributorBurnShareParam) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSubDistributorBurnShareParam) Type() string {
	return TypeMsgUpdateSubDistributorBurnShareParam
}

func (msg *MsgUpdateSubDistributorBurnShareParam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSubDistributorBurnShareParam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSubDistributorBurnShareParam) ValidateBasic() error {
	if msg.Authority != appparams.GetAuthority() {
		return govtypes.ErrInvalidSigner
	}
	if msg.SubDistributorName == "" {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "empty sub distributor name")
	}
	if msg.BurnShare.IsNil() {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "burn share cannot be nil")
	}
	if msg.BurnShare.GTE(sdk.NewDec(maxShare)) || msg.BurnShare.IsNegative() {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "burn share must be between 0 and 1")
	}
	return nil
}

const TypeMsgUpdateSubDistributorDestinationShareParam = "update_sub_distributor_destination_share_param"

var _ sdk.Msg = &MsgUpdateSubDistributorBurnShareParam{}

func (msg *MsgUpdateSubDistributorDestinationShareParam) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSubDistributorDestinationShareParam) Type() string {
	return TypeMsgUpdateSubDistributorDestinationShareParam
}

func (msg *MsgUpdateSubDistributorDestinationShareParam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSubDistributorDestinationShareParam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSubDistributorDestinationShareParam) ValidateBasic() error {
	if msg.Authority != appparams.GetAuthority() {
		return govtypes.ErrInvalidSigner
	}
	if msg.SubDistributorName == "" {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "empty sub distributor name")
	}
	if msg.DestinationName == "" {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "empty destination name")
	}
	if msg.Share.IsNil() {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "share cannot be nil")
	}
	if msg.Share.GTE(sdk.NewDec(maxShare)) || msg.Share.IsNegative() {
		return errors.Wrapf(govtypes.ErrInvalidProposalContent, "share must be between 0 and 1")
	}
	return nil
}
