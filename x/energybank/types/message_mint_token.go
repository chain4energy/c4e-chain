package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintToken = "mint_token"

var _ sdk.Msg = &MsgMintToken{}

func NewMsgMintToken(creator string, name string, amount uint64, userAddress string) *MsgMintToken {
	return &MsgMintToken{
		Creator:     creator,
		Name:        name,
		Amount:      amount,
		UserAddress: userAddress,
	}
}

func (msg *MsgMintToken) Route() string {
	return RouterKey
}

func (msg *MsgMintToken) Type() string {
	return TypeMsgMintToken
}

func (msg *MsgMintToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
