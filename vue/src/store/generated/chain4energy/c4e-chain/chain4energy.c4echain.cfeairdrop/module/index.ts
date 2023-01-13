// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateAirdropEntry } from "./types/cfeairdrop/tx";
import { MsgCreateAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgDeleteAirdropEntry } from "./types/cfeairdrop/tx";
import { MsgAddMissionToAidropCampaign } from "./types/cfeairdrop/tx";
import { MsgClaim } from "./types/cfeairdrop/tx";
import { MsgUpdateAirdropEntry } from "./types/cfeairdrop/tx";


const types = [
  ["/chain4energy.c4echain.cfeairdrop.MsgCreateAirdropEntry", MsgCreateAirdropEntry],
  ["/chain4energy.c4echain.cfeairdrop.MsgCreateAirdropCampaign", MsgCreateAirdropCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgDeleteAirdropEntry", MsgDeleteAirdropEntry],
  ["/chain4energy.c4echain.cfeairdrop.MsgAddMissionToAidropCampaign", MsgAddMissionToAidropCampaign],
  ["/chain4energy.c4echain.cfeairdrop.MsgClaim", MsgClaim],
  ["/chain4energy.c4echain.cfeairdrop.MsgUpdateAirdropEntry", MsgUpdateAirdropEntry],
  
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
    msgCreateAirdropEntry: (data: MsgCreateAirdropEntry): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCreateAirdropEntry", value: MsgCreateAirdropEntry.fromPartial( data ) }),
    msgCreateAirdropCampaign: (data: MsgCreateAirdropCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCreateAirdropCampaign", value: MsgCreateAirdropCampaign.fromPartial( data ) }),
    msgDeleteAirdropEntry: (data: MsgDeleteAirdropEntry): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgDeleteAirdropEntry", value: MsgDeleteAirdropEntry.fromPartial( data ) }),
    msgAddMissionToAidropCampaign: (data: MsgAddMissionToAidropCampaign): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddMissionToAidropCampaign", value: MsgAddMissionToAidropCampaign.fromPartial( data ) }),
    msgClaim: (data: MsgClaim): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgClaim", value: MsgClaim.fromPartial( data ) }),
    msgUpdateAirdropEntry: (data: MsgUpdateAirdropEntry): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgUpdateAirdropEntry", value: MsgUpdateAirdropEntry.fromPartial( data ) }),
    
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
