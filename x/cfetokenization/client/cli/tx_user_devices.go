package cli

//func CmdCreateUserDevices() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "create-user-devices [devices]",
//		Short: "Create a new UserDevices",
//		Args:  cobra.ExactArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) (err error) {
//	  	 argDevices := new(types.Device)
//					err = json.Unmarshal([]byte(args[0]), argDevices)
//    				if err != nil {
//                		return err
//            		}
//
//			clientCtx, err := client.GetClientTxContext(cmd)
//			if err != nil {
//				return err
//			}
//
//			msg := types.NewMsgCreateUserDevices(clientCtx.GetFromAddress().String(), argDevices)
//			if err := msg.ValidateBasic(); err != nil {
//				return err
//			}
//			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
//		},
//	}
//
//	flags.AddTxFlagsToCmd(cmd)
//
//    return cmd
//}
//
//func CmdUpdateUserDevices() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "update-user-devices [id] [devices]",
//		Short: "Update a UserDevices",
//		Args:  cobra.ExactArgs(2),
//		RunE: func(cmd *cobra.Command, args []string) (err error) {
//            id, err := strconv.ParseUint(args[0], 10, 64)
//            if err != nil {
//                return err
//            }
//
//
//	  		argDevices := new(types.Device)
//					err = json.Unmarshal([]byte(args[1]), argDevices)
//    				if err != nil {
//                		return err
//            		}
//
//			clientCtx, err := client.GetClientTxContext(cmd)
//			if err != nil {
//				return err
//			}
//
//			msg := types.NewMsgUpdateUserDevices(clientCtx.GetFromAddress().String(), id, argDevices)
//			if err := msg.ValidateBasic(); err != nil {
//				return err
//			}
//			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
//		},
//	}
//
//	flags.AddTxFlagsToCmd(cmd)
//
//    return cmd
//}
//
//func CmdDeleteUserDevices() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "delete-user-devices [id]",
//		Short: "Delete a UserDevices by id",
//		Args:  cobra.ExactArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) error {
//            id, err := strconv.ParseUint(args[0], 10, 64)
//            if err != nil {
//                return err
//            }
//
//			clientCtx, err := client.GetClientTxContext(cmd)
//			if err != nil {
//				return err
//			}
//
//			msg := types.NewMsgDeleteUserDevices(clientCtx.GetFromAddress().String(), id)
//			if err := msg.ValidateBasic(); err != nil {
//				return err
//			}
//			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
//		},
//	}
//
//	flags.AddTxFlagsToCmd(cmd)
//
//    return cmd
//}
