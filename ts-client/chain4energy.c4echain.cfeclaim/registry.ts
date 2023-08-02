import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddMission } from "./types/c4echain/cfeclaim/tx";
import { MsgEnableCampaign } from "./types/c4echain/cfeclaim/tx";
import { MsgCreateCampaign } from "./types/c4echain/cfeclaim/tx";
import { MsgCloseCampaign } from "./types/c4echain/cfeclaim/tx";
import { MsgInitialClaim } from "./types/c4echain/cfeclaim/tx";
import { MsgRemoveCampaign } from "./types/c4echain/cfeclaim/tx";
import { MsgDeleteClaimRecord } from "./types/c4echain/cfeclaim/tx";
import { MsgAddClaimRecords } from "./types/c4echain/cfeclaim/tx";
import { MsgClaim } from "./types/c4echain/cfeclaim/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfeclaim.MsgAddMission", MsgAddMission],
    ["/chain4energy.c4echain.cfeclaim.MsgEnableCampaign", MsgEnableCampaign],
    ["/chain4energy.c4echain.cfeclaim.MsgCreateCampaign", MsgCreateCampaign],
    ["/chain4energy.c4echain.cfeclaim.MsgCloseCampaign", MsgCloseCampaign],
    ["/chain4energy.c4echain.cfeclaim.MsgInitialClaim", MsgInitialClaim],
    ["/chain4energy.c4echain.cfeclaim.MsgRemoveCampaign", MsgRemoveCampaign],
    ["/chain4energy.c4echain.cfeclaim.MsgDeleteClaimRecord", MsgDeleteClaimRecord],
    ["/chain4energy.c4echain.cfeclaim.MsgAddClaimRecords", MsgAddClaimRecords],
    ["/chain4energy.c4echain.cfeclaim.MsgClaim", MsgClaim],
    
];

export { msgTypes }