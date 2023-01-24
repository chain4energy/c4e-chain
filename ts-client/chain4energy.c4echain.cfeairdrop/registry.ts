import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgStartAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgAddMissionToAidropCampaign } from "./types/cfeairdrop/tx";
import { MsgCreateAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgInitialClaim } from "./types/cfeairdrop/tx";
import { MsgEditAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgAddAirdropEntries } from "./types/cfeairdrop/tx";
import { MsgDeleteAirdropEntry } from "./types/cfeairdrop/tx";
import { MsgCloseAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgClaim } from "./types/cfeairdrop/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfeairdrop.MsgStartAirdropCampaign", MsgStartAirdropCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgAddMissionToAidropCampaign", MsgAddMissionToAidropCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgCreateAirdropCampaign", MsgCreateAirdropCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", MsgInitialClaim],
    ["/chain4energy.c4echain.cfeairdrop.MsgEditAirdropCampaign", MsgEditAirdropCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgAddAirdropEntries", MsgAddAirdropEntries],
    ["/chain4energy.c4echain.cfeairdrop.MsgDeleteAirdropEntry", MsgDeleteAirdropEntry],
    ["/chain4energy.c4echain.cfeairdrop.MsgCloseAirdropCampaign", MsgCloseAirdropCampaign],
    ["/chain4energy.c4echain.cfeairdrop.MsgClaim", MsgClaim],
    
];

export { msgTypes }