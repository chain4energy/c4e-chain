package types

import (
	"cosmossdk.io/errors"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitialClaim = "initial_claim"

var _ sdk.Msg = &MsgInitialClaim{}

func NewMsgInitialClaim(claimer string, campaignId uint64, destinationAddress string) *MsgInitialClaim {
	return &MsgInitialClaim{
		Claimer:            claimer,
		CampaignId:         campaignId,
		DestinationAddress: destinationAddress,
	}
}

func (msg *MsgInitialClaim) Route() string {
	return RouterKey
}

func (msg *MsgInitialClaim) Type() string {
	return TypeMsgInitialClaim
}

func (msg *MsgInitialClaim) GetSigners() []sdk.AccAddress {
	claimer, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{claimer}
}

func (msg *MsgInitialClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitialClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid claimer address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.DestinationAddress)
	if err != nil {
		return errors.Wrapf(c4eerrors.ErrParsing, "destAddress parsing error: %s", msg.DestinationAddress)
	}
	return nil
}
