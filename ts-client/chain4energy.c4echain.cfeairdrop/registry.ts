import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgStartCampaign } from "./types/cfeairdrop/tx";
import { MsgAddMissionToAidropCampaign } from "./types/cfeairdrop/tx";
import { MsgCreateCampaign } from "./types/cfeairdrop/tx";
import { MsgInitialClaim } from "./types/cfeairdrop/tx";
import { MsgEditCampaign } from "./types/cfeairdrop/tx";
import { MsgAddClaimRecords } from "./types/cfeairdrop/tx";
import { MsgDeleteClaimRecord } from "./types/cfeairdrop/tx";
import { MsgCloseCampaign } from "./types/cfeairdrop/tx";
import { MsgClaim } from "./types/cfeairdrop/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfeairdrop.MsgStartCampaign", MsgStartCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgAddMissionToAidropCampaign", MsgAddMissionToAidropCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgCreateCampaign", MsgCreateCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", MsgInitialClaim],
    ["/chain4energy.c4echain.cfeairdrop.MsgEditCampaign", MsgEditCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgAddClaimRecords", MsgAddClaimRecords],
    ["/chain4energy.c4echain.cfeairdrop.MsgDeleteClaimRecord", MsgDeleteClaimRecord],
    ["/chain4energy.c4echain.cfeairdrop.MsgCloseCampaign", MsgCloseCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgClaim", MsgClaim],
    
];

export { msgTypes }