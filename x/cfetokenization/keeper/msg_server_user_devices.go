package keeper

//
//func (k msgServer) CreateUserDevices(goCtx context.Context,  msg *types.MsgCreateUserDevices) (*types.MsgCreateUserDevicesResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//    var userDevices = types.UserDevices{
//        Owner: msg.Owner,
//        Devices: msg.Devices,
//    }
//
//    id := k.AppendUserDevices(
//        ctx,
//        userDevices,
//    )
//
//	return &types.MsgCreateUserDevicesResponse{
//	    Id: id,
//	}, nil
//}
//
//func (k msgServer) UpdateUserDevices(goCtx context.Context,  msg *types.MsgUpdateUserDevices) (*types.MsgUpdateUserDevicesResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//    var userDevices = types.UserDevices{
//		Owner: msg.Owner,
//		Id:      msg.Id,
//    	Devices: msg.Devices,
//	}
//
//    // Checks that the element exists
//    val, found := k.GetUserDevices(ctx, msg.Id)
//    if !found {
//        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
//    }
//
//    // Checks if the msg owner is the same as the current owner
//    if msg.Owner != val.Creator {
//        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
//    }
//
//	k.SetUserDevices(ctx, userDevices)
//
//	return &types.MsgUpdateUserDevicesResponse{}, nil
//}
//
//func (k msgServer) DeleteUserDevices(goCtx context.Context,  msg *types.MsgDeleteUserDevices) (*types.MsgDeleteUserDevicesResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//    // Checks that the element exists
//    val, found := k.GetUserDevices(ctx, msg.Id)
//    if !found {
//        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
//    }
//
//    // Checks if the msg owner is the same as the current owner
//    if msg.Owner != val.Creator {
//        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
//    }
//
//	k.RemoveUserDevices(ctx, msg.Id)
//
//	return &types.MsgDeleteUserDevicesResponse{}, nil
//}
