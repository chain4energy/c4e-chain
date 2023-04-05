import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgClaim } from "./types/c4echain/cfeairdrop/tx";
import { MsgCreateCampaign } from "./types/c4echain/cfeairdrop/tx";
import { MsgDeleteClaimRecord } from "./types/c4echain/cfeairdrop/tx";
import { MsgInitialClaim } from "./types/c4echain/cfeairdrop/tx";
import { MsgEditCampaign } from "./types/c4echain/cfeairdrop/tx";
import { MsgRemoveCampaign } from "./types/c4echain/cfeairdrop/tx";
import { MsgCloseCampaign } from "./types/c4echain/cfeairdrop/tx";
import { MsgAddClaimRecords } from "./types/c4echain/cfeairdrop/tx";
import { MsgAddMissionToCampaign } from "./types/c4echain/cfeairdrop/tx";
import { MsgStartCampaign } from "./types/c4echain/cfeairdrop/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfeairdrop.MsgClaim", MsgClaim],
    ["/chain4energy.c4echain.cfeairdrop.MsgCreateCampaign", MsgCreateCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgDeleteClaimRecord", MsgDeleteClaimRecord],
    ["/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", MsgInitialClaim],
    ["/chain4energy.c4echain.cfeairdrop.MsgEditCampaign", MsgEditCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgRemoveCampaign", MsgRemoveCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgCloseCampaign", MsgCloseCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgAddClaimRecords", MsgAddClaimRecords],
    ["/chain4energy.c4echain.cfeairdrop.MsgAddMissionToCampaign", MsgAddMissionToCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgStartCampaign", MsgStartCampaign],
    
];

export { msgTypes }