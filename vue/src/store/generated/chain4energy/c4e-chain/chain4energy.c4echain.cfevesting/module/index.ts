// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendToVestingAccount } from "./types/cfevesting/tx";
import { MsgWithdrawAllAvailable } from "./types/cfevesting/tx";
import { MsgCreateVestingAccount } from "./types/cfevesting/tx";
import { MsgCreateVestingPool } from "./types/cfevesting/tx";


const types = [
  ["/chain4energy.c4echain.cfevesting.MsgSendToVestingAccount", MsgSendToVestingAccount],
  ["/chain4energy.c4echain.cfevesting.MsgWithdrawAllAvailable", MsgWithdrawAllAvailable],
  ["/chain4energy.c4echain.cfevesting.MsgCreateVestingAccount", MsgCreateVestingAccount],
  ["/chain4energy.c4echain.cfevesting.MsgCreateVestingPool", MsgCreateVestingPool],
  
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
    msgSendToVestingAccount: (data: MsgSendToVestingAccount): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgSendToVestingAccount", value: MsgSendToVestingAccount.fromPartial( data ) }),
    msgWithdrawAllAvailable: (data: MsgWithdrawAllAvailable): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgWithdrawAllAvailable", value: MsgWithdrawAllAvailable.fromPartial( data ) }),
    msgCreateVestingAccount: (data: MsgCreateVestingAccount): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgCreateVestingAccount", value: MsgCreateVestingAccount.fromPartial( data ) }),
    msgCreateVestingPool: (data: MsgCreateVestingPool): EncodeObject => ({ typeUrl: "/chain4energy.c4echain.cfevesting.MsgCreateVestingPool", value: MsgCreateVestingPool.fromPartial( data ) }),
    
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
