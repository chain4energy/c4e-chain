import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgPublishReferencePayloadLink } from "./types/c4echain/cfesignature/tx";
import { MsgCreateAccount } from "./types/c4echain/cfesignature/tx";
import { MsgStoreSignature } from "./types/c4echain/cfesignature/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfesignature.MsgPublishReferencePayloadLink", MsgPublishReferencePayloadLink],
    ["/chain4energy.c4echain.cfesignature.MsgCreateAccount", MsgCreateAccount],
    ["/chain4energy.c4echain.cfesignature.MsgStoreSignature", MsgStoreSignature],
    
];

export { msgTypes }