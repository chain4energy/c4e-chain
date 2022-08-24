// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgTransferTokensOptimally } from "./types/cfeenergybank/tx";
import { MsgCreateTokenParams } from "./types/cfeenergybank/tx";
import { MsgMintToken } from "./types/cfeenergybank/tx";
import { MsgTransferTokens } from "./types/cfeenergybank/tx";


const types = [
  ["/chain4energy.c4echain.cfeenergybank.MsgTransferTokensOptimally", MsgTransferTokensOptimally],
  ["/chain4energy.c4echain.cfeenergybank.MsgCreateTokenParams", MsgCreateTokenParams],
  ["/chain4energy.c4echain.cfeenergybank.MsgMintToken", MsgMintToken],
  ["/chain4energy.c4echain.cfeenergybank.MsgTransferTokens", MsgTransferTokens],
  
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
    msgTransferTokensOptimally: (data: MsgTransferTokensOptimally): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeenergybank.MsgTransferTokensOptimally", value: MsgTransferTokensOptimally.fromPartial( data ) }),
    msgCreateTokenParams: (data: MsgCreateTokenParams): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeenergybank.MsgCreateTokenParams", value: MsgCreateTokenParams.fromPartial( data ) }),
    msgMintToken: (data: MsgMintToken): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeenergybank.MsgMintToken", value: MsgMintToken.fromPartial( data ) }),
    msgTransferTokens: (data: MsgTransferTokens): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfeenergybank.MsgTransferTokens", value: MsgTransferTokens.fromPartial( data ) }),
    
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
