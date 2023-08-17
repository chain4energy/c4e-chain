import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateParams } from "./types/c4echain/cfedistributor/tx";
import { MsgUpdateSubDistributorDestinationShareParam } from "./types/c4echain/cfedistributor/tx";
import { MsgUpdateSubDistributorBurnShareParam } from "./types/c4echain/cfedistributor/tx";
import { MsgUpdateSubDistributorParam } from "./types/c4echain/cfedistributor/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfedistributor.MsgUpdateParams", MsgUpdateParams],
    ["/chain4energy.c4echain.cfedistributor.MsgUpdateSubDistributorDestinationShareParam", MsgUpdateSubDistributorDestinationShareParam],
    ["/chain4energy.c4echain.cfedistributor.MsgUpdateSubDistributorBurnShareParam", MsgUpdateSubDistributorBurnShareParam],
    ["/chain4energy.c4echain.cfedistributor.MsgUpdateSubDistributorParam", MsgUpdateSubDistributorParam],
    
];

export { msgTypes }