package types

import (
	appparams "github.com/chain4energy/c4e-chain/v2/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const TypeMsgUpdateDenomParam = "update_denom_param"

var _ sdk.Msg = &MsgUpdateDenomParam{}

func (msg *MsgUpdateDenomParam) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDenomParam) Type() string {
	return TypeMsgUpdateDenomParam
}

func (msg *MsgUpdateDenomParam) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDenomParam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDenomParam) ValidateBasic() error {
	if msg.Authority != appparams.GetAuthority() {
		return govtypes.ErrInvalidSigner
	}
	return Params{Denom: msg.Denom}.Validate()
}
