import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateMintersParams } from "./types/c4echain/cfeminter/tx";
import { MsgBurn } from "./types/c4echain/cfeminter/tx";
import { MsgUpdateParams } from "./types/c4echain/cfeminter/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfeminter.MsgUpdateMintersParams", MsgUpdateMintersParams],
    ["/chain4energy.c4echain.cfeminter.MsgBurn", MsgBurn],
    ["/chain4energy.c4echain.cfeminter.MsgUpdateParams", MsgUpdateParams],
    
];

export { msgTypes }