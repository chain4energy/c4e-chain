// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateCampaign } from "./types/cfeairdrop/tx";
import { MsgCloseCampaign } from "./types/cfeairdrop/tx";
import { MsgAddClaimRecords } from "./types/cfeairdrop/tx";
import { MsgStartCampaign } from "./types/cfeairdrop/tx";
import { MsgRemoveCampaign } from "./types/cfeairdrop/tx";
import { MsgInitialClaim } from "./types/cfeairdrop/tx";
import { MsgAddMissionToCampaign } from "./types/cfeairdrop/tx";
import { MsgEditCampaign } from "./types/cfeairdrop/tx";
import { MsgClaim } from "./types/cfeairdrop/tx";
import { MsgDeleteClaimRecord } from "./types/cfeairdrop/tx";


const types = [
  ["/chain4energy.c4echain.cfeairdrop.MsgCreateCampaign", MsgCreateCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgCloseCampaign", MsgCloseCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgAddClaimRecords", MsgAddClaimRecords],
  ["/chain4energy.c4echain.cfeairdrop.MsgStartCampaign", MsgStartCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgRemoveCampaign", MsgRemoveCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", MsgInitialClaim],
  ["/chain4energy.c4echain.cfeairdrop.MsgAddMissionToCampaign", MsgAddMissionToCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgEditCampaign", MsgEditCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgClaim", MsgClaim],
  ["/chain4energy.c4echain.cfeairdrop.MsgDeleteClaimRecord", MsgDeleteClaimRecord],
  
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
    msgCreateCampaign: (data: MsgCreateCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCreateCampaign", value: MsgCreateCampaign.fromPartial( data ) }),
    msgCloseCampaign: (data: MsgCloseCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCloseCampaign", value: MsgCloseCampaign.fromPartial( data ) }),
    msgAddClaimRecords: (data: MsgAddClaimRecords): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddClaimRecords", value: MsgAddClaimRecords.fromPartial( data ) }),
    msgStartCampaign: (data: MsgStartCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgStartCampaign", value: MsgStartCampaign.fromPartial( data ) }),
    msgRemoveCampaign: (data: MsgRemoveCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgRemoveCampaign", value: MsgRemoveCampaign.fromPartial( data ) }),
    msgInitialClaim: (data: MsgInitialClaim): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", value: MsgInitialClaim.fromPartial( data ) }),
    msgAddMissionToCampaign: (data: MsgAddMissionToCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddMissionToCampaign", value: MsgAddMissionToCampaign.fromPartial( data ) }),
    msgEditCampaign: (data: MsgEditCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgEditCampaign", value: MsgEditCampaign.fromPartial( data ) }),
    msgClaim: (data: MsgClaim): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgClaim", value: MsgClaim.fromPartial( data ) }),
    msgDeleteClaimRecord: (data: MsgDeleteClaimRecord): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgDeleteClaimRecord", value: MsgDeleteClaimRecord.fromPartial( data ) }),
    
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
