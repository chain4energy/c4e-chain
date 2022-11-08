import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateAccount } from "./types/c4e-chain/cfesignature/tx";
import { MsgStoreSignature } from "./types/c4e-chain/cfesignature/tx";
import { MsgPublishReferencePayloadLink } from "./types/c4e-chain/cfesignature/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfesignature.MsgCreateAccount", MsgCreateAccount],
    ["/chain4energy.c4echain.cfesignature.MsgStoreSignature", MsgStoreSignature],
    ["/chain4energy.c4echain.cfesignature.MsgPublishReferencePayloadLink", MsgPublishReferencePayloadLink],
    
];

export { msgTypes }