import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgWithdrawAllAvailable } from "./types/cfevesting/tx";
import { MsgSendToVestingAccount } from "./types/cfevesting/tx";
import { MsgCreateVestingPool } from "./types/cfevesting/tx";
import { MsgCreateVestingAccount } from "./types/cfevesting/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfevesting.MsgWithdrawAllAvailable", MsgWithdrawAllAvailable],
    ["/chain4energy.c4echain.cfevesting.MsgSendToVestingAccount", MsgSendToVestingAccount],
    ["/chain4energy.c4echain.cfevesting.MsgCreateVestingPool", MsgCreateVestingPool],
    ["/chain4energy.c4echain.cfevesting.MsgCreateVestingAccount", MsgCreateVestingAccount],
    
];

export { msgTypes }