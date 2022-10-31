// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgPublishReferencePayloadLink } from "./types/cfesignature/tx";
import { MsgCreateAccount } from "./types/cfesignature/tx";
import { MsgStoreSignature } from "./types/cfesignature/tx";


const types = [
  ["/chain4energy.c4echain.cfesignature.MsgPublishReferencePayloadLink", MsgPublishReferencePayloadLink],
  ["/chain4energy.c4echain.cfesignature.MsgCreateAccount", MsgCreateAccount],
  ["/chain4energy.c4echain.cfesignature.MsgStoreSignature", MsgStoreSignature],
  
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
    msgPublishReferencePayloadLink: (data: MsgPublishReferencePayloadLink): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfesignature.MsgPublishReferencePayloadLink", value: MsgPublishReferencePayloadLink.fromPartial( data ) }),
    msgCreateAccount: (data: MsgCreateAccount): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfesignature.MsgCreateAccount", value: MsgCreateAccount.fromPartial( data ) }),
    msgStoreSignature: (data: MsgStoreSignature): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfesignature.MsgStoreSignature", value: MsgStoreSignature.fromPartial( data ) }),
    
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
