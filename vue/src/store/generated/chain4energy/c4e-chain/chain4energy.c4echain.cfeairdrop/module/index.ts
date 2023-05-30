// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgInitialClaim } from "./types/cfeclaim/tx";
import { MsgAddMission } from "./types/cfeclaim/tx";
import { MsgEnableCampaign } from "./types/cfeclaim/tx";
import { MsgCloseCampaign } from "./types/cfeclaim/tx";
import { MsgEditCampaign } from "./types/cfeclaim/tx";
import { MsgRemoveCampaign } from "./types/cfeclaim/tx";
import { MsgAddClaimRecords } from "./types/cfeclaim/tx";
import { MsgClaim } from "./types/cfeclaim/tx";
import { MsgCreateCampaign } from "./types/cfeclaim/tx";
import { MsgDeleteClaimRecord } from "./types/cfeclaim/tx";


const types = [
  ["/chain4energy.c4echain.cfeclaim.MsgInitialClaim", MsgInitialClaim],
  ["/chain4energy.c4echain.cfeclaim.MsgAddMission", MsgAddMission],
  ["/chain4energy.c4echain.cfeclaim.MsgEnableCampaign", MsgEnableCampaign],
  ["/chain4energy.c4echain.cfeclaim.MsgCloseCampaign", MsgCloseCampaign],
  ["/chain4energy.c4echain.cfeclaim.MsgEditCampaign", MsgEditCampaign],
  ["/chain4energy.c4echain.cfeclaim.MsgRemoveCampaign", MsgRemoveCampaign],
  ["/chain4energy.c4echain.cfeclaim.MsgAddClaimRecords", MsgAddClaimRecords],
  ["/chain4energy.c4echain.cfeclaim.MsgClaim", MsgClaim],
  ["/chain4energy.c4echain.cfeclaim.MsgCreateCampaign", MsgCreateCampaign],
  ["/chain4energy.c4echain.cfeclaim.MsgDeleteClaimRecord", MsgDeleteClaimRecord],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgInitialClaim: (data: MsgInitialClaim): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgInitialClaim", value: MsgInitialClaim.fromPartial( data ) }),
    msgAddMission: (data: MsgAddMission): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgAddMission", value: MsgAddMission.fromPartial( data ) }),
    MsgEnableCampaign: (data: MsgEnableCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgEnableCampaign", value: MsgEnableCampaign.fromPartial( data ) }),
    msgCloseCampaign: (data: MsgCloseCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgCloseCampaign", value: MsgCloseCampaign.fromPartial( data ) }),
    msgEditCampaign: (data: MsgEditCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgEditCampaign", value: MsgEditCampaign.fromPartial( data ) }),
    msgRemoveCampaign: (data: MsgRemoveCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgRemoveCampaign", value: MsgRemoveCampaign.fromPartial( data ) }),
    msgAddClaimRecords: (data: MsgAddClaimRecords): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgAddClaimRecords", value: MsgAddClaimRecords.fromPartial( data ) }),
    msgClaim: (data: MsgClaim): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgClaim", value: MsgClaim.fromPartial( data ) }),
    msgCreateCampaign: (data: MsgCreateCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgCreateCampaign", value: MsgCreateCampaign.fromPartial( data ) }),
    msgDeleteClaimRecord: (data: MsgDeleteClaimRecord): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeclaim.MsgDeleteClaimRecord", value: MsgDeleteClaimRecord.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
