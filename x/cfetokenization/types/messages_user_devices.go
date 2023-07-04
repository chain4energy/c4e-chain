package types

//
//import (
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//)
//
//const (
//	TypeMsgCreateUserDevices = "create_user_devices"
//	TypeMsgUpdateUserDevices = "update_user_devices"
//	TypeMsgDeleteUserDevices = "delete_user_devices"
//)
//
//var _ sdk.Msg = &MsgCreateUserDevices{}
//
//func NewMsgCreateUserDevices(owner string, devices *Device) *MsgCreateUserDevices {
//  return &MsgCreateUserDevices{
//		Owner: owner,
//    Devices: devices,
//	}
//}
//
//func (msg *MsgCreateUserDevices) Route() string {
//  return RouterKey
//}
//
//func (msg *MsgCreateUserDevices) Type() string {
//  return TypeMsgCreateUserDevices
//}
//
//func (msg *MsgCreateUserDevices) GetSigners() []sdk.AccAddress {
//  owner, err := sdk.AccAddressFromBech32(msg.Owner)
//  if err != nil {
//    panic(err)
//  }
//  return []sdk.AccAddress{owner}
//}
//
//func (msg *MsgCreateUserDevices) GetSignBytes() []byte {
//  bz := ModuleCdc.MustMarshalJSON(msg)
//  return sdk.MustSortJSON(bz)
//}
//
//func (msg *MsgCreateUserDevices) ValidateBasic() error {
//  _, err := sdk.AccAddressFromBech32(msg.Owner)
//  	if err != nil {
//  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
//  	}
//  return nil
//}
//
//var _ sdk.Msg = &MsgUpdateUserDevices{}
//
//func NewMsgUpdateUserDevices(owner string, id uint64, devices *Device) *MsgUpdateUserDevices {
//  return &MsgUpdateUserDevices{
//        Id: id,
//		Owner: owner,
//    Devices: devices,
//	}
//}
//
//func (msg *MsgUpdateUserDevices) Route() string {
//  return RouterKey
//}
//
//func (msg *MsgUpdateUserDevices) Type() string {
//  return TypeMsgUpdateUserDevices
//}
//
//func (msg *MsgUpdateUserDevices) GetSigners() []sdk.AccAddress {
//  owner, err := sdk.AccAddressFromBech32(msg.Owner)
//  if err != nil {
//    panic(err)
//  }
//  return []sdk.AccAddress{owner}
//}
//
//func (msg *MsgUpdateUserDevices) GetSignBytes() []byte {
//  bz := ModuleCdc.MustMarshalJSON(msg)
//  return sdk.MustSortJSON(bz)
//}
//
//func (msg *MsgUpdateUserDevices) ValidateBasic() error {
//  _, err := sdk.AccAddressFromBech32(msg.Owner)
//  if err != nil {
//    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
//  }
//   return nil
//}
//
//var _ sdk.Msg = &MsgDeleteUserDevices{}
//
//func NewMsgDeleteUserDevices(owner string, id uint64) *MsgDeleteUserDevices {
//  return &MsgDeleteUserDevices{
//        Id: id,
//		Owner: owner,
//	}
//}
//func (msg *MsgDeleteUserDevices) Route() string {
//  return RouterKey
//}
//
//func (msg *MsgDeleteUserDevices) Type() string {
//  return TypeMsgDeleteUserDevices
//}
//
//func (msg *MsgDeleteUserDevices) GetSigners() []sdk.AccAddress {
//  owner, err := sdk.AccAddressFromBech32(msg.Owner)
//  if err != nil {
//    panic(err)
//  }
//  return []sdk.AccAddress{owner}
//}
//
//func (msg *MsgDeleteUserDevices) GetSignBytes() []byte {
//  bz := ModuleCdc.MustMarshalJSON(msg)
//  return sdk.MustSortJSON(bz)
//}
//
//func (msg *MsgDeleteUserDevices) ValidateBasic() error {
//  _, err := sdk.AccAddressFromBech32(msg.Owner)
//  if err != nil {
//    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
//  }
//  return nil
//}
