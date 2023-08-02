import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgMoveAvailableVesting } from "./types/c4echain/cfevesting/tx";
import { MsgSplitVesting } from "./types/c4echain/cfevesting/tx";
import { MsgWithdrawAllAvailable } from "./types/c4echain/cfevesting/tx";
import { MsgUpdateDenomParam } from "./types/c4echain/cfevesting/tx";
import { MsgCreateVestingPool } from "./types/c4echain/cfevesting/tx";
import { MsgMoveAvailableVestingByDenoms } from "./types/c4echain/cfevesting/tx";
import { MsgSendToVestingAccount } from "./types/c4echain/cfevesting/tx";
import { MsgCreateVestingAccount } from "./types/c4echain/cfevesting/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfevesting.MsgMoveAvailableVesting", MsgMoveAvailableVesting],
    ["/chain4energy.c4echain.cfevesting.MsgSplitVesting", MsgSplitVesting],
    ["/chain4energy.c4echain.cfevesting.MsgWithdrawAllAvailable", MsgWithdrawAllAvailable],
    ["/chain4energy.c4echain.cfevesting.MsgUpdateDenomParam", MsgUpdateDenomParam],
    ["/chain4energy.c4echain.cfevesting.MsgCreateVestingPool", MsgCreateVestingPool],
    ["/chain4energy.c4echain.cfevesting.MsgMoveAvailableVestingByDenoms", MsgMoveAvailableVestingByDenoms],
    ["/chain4energy.c4echain.cfevesting.MsgSendToVestingAccount", MsgSendToVestingAccount],
    ["/chain4energy.c4echain.cfevesting.MsgCreateVestingAccount", MsgCreateVestingAccount],
    
];

export { msgTypes }