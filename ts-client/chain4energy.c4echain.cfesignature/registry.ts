import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgStoreSignature } from "./types/cfesignature/tx";
import { MsgCreateAccount } from "./types/cfesignature/tx";
import { MsgPublishReferencePayloadLink } from "./types/cfesignature/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfesignature.MsgStoreSignature", MsgStoreSignature],
    ["/chain4energy.c4echain.cfesignature.MsgCreateAccount", MsgCreateAccount],
    ["/chain4energy.c4echain.cfesignature.MsgPublishReferencePayloadLink", MsgPublishReferencePayloadLink],
    
];

export { msgTypes }