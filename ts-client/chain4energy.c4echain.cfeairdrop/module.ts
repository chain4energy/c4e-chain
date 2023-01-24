// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgStartAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgAddMissionToAidropCampaign } from "./types/cfeairdrop/tx";
import { MsgCreateAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgInitialClaim } from "./types/cfeairdrop/tx";
import { MsgEditAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgAddAirdropEntries } from "./types/cfeairdrop/tx";
import { MsgDeleteAirdropEntry } from "./types/cfeairdrop/tx";
import { MsgCloseAirdropCampaign } from "./types/cfeairdrop/tx";
import { MsgClaim } from "./types/cfeairdrop/tx";


export { MsgStartAirdropCampaign, MsgAddMissionToAidropCampaign, MsgCreateAirdropCampaign, MsgInitialClaim, MsgEditAirdropCampaign, MsgAddAirdropEntries, MsgDeleteAirdropEntry, MsgCloseAirdropCampaign, MsgClaim };

type sendMsgStartAirdropCampaignParams = {
  value: MsgStartAirdropCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddMissionToAidropCampaignParams = {
  value: MsgAddMissionToAidropCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateAirdropCampaignParams = {
  value: MsgCreateAirdropCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgInitialClaimParams = {
  value: MsgInitialClaim,
  fee?: StdFee,
  memo?: string
};

type sendMsgEditAirdropCampaignParams = {
  value: MsgEditAirdropCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddAirdropEntriesParams = {
  value: MsgAddAirdropEntries,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeleteAirdropEntryParams = {
  value: MsgDeleteAirdropEntry,
  fee?: StdFee,
  memo?: string
};

type sendMsgCloseAirdropCampaignParams = {
  value: MsgCloseAirdropCampaign,
  fee?: StdFee,
  memo?: string
};

type sendMsgClaimParams = {
  value: MsgClaim,
  fee?: StdFee,
  memo?: string
};


type msgStartAirdropCampaignParams = {
  value: MsgStartAirdropCampaign,
};

type msgAddMissionToAidropCampaignParams = {
  value: MsgAddMissionToAidropCampaign,
};

type msgCreateAirdropCampaignParams = {
  value: MsgCreateAirdropCampaign,
};

type msgInitialClaimParams = {
  value: MsgInitialClaim,
};

type msgEditAirdropCampaignParams = {
  value: MsgEditAirdropCampaign,
};

type msgAddAirdropEntriesParams = {
  value: MsgAddAirdropEntries,
};

type msgDeleteAirdropEntryParams = {
  value: MsgDeleteAirdropEntry,
};

type msgCloseAirdropCampaignParams = {
  value: MsgCloseAirdropCampaign,
};

type msgClaimParams = {
  value: MsgClaim,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgStartAirdropCampaign({ value, fee, memo }: sendMsgStartAirdropCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgStartAirdropCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgStartAirdropCampaign({ value: MsgStartAirdropCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgStartAirdropCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddMissionToAidropCampaign({ value, fee, memo }: sendMsgAddMissionToAidropCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddMissionToAidropCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddMissionToAidropCampaign({ value: MsgAddMissionToAidropCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddMissionToAidropCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateAirdropCampaign({ value, fee, memo }: sendMsgCreateAirdropCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateAirdropCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateAirdropCampaign({ value: MsgCreateAirdropCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateAirdropCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgInitialClaim({ value, fee, memo }: sendMsgInitialClaimParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgInitialClaim: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgInitialClaim({ value: MsgInitialClaim.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgInitialClaim: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgEditAirdropCampaign({ value, fee, memo }: sendMsgEditAirdropCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgEditAirdropCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgEditAirdropCampaign({ value: MsgEditAirdropCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgEditAirdropCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddAirdropEntries({ value, fee, memo }: sendMsgAddAirdropEntriesParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddAirdropEntries: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddAirdropEntries({ value: MsgAddAirdropEntries.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddAirdropEntries: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeleteAirdropEntry({ value, fee, memo }: sendMsgDeleteAirdropEntryParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteAirdropEntry: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteAirdropEntry({ value: MsgDeleteAirdropEntry.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteAirdropEntry: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCloseAirdropCampaign({ value, fee, memo }: sendMsgCloseAirdropCampaignParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCloseAirdropCampaign: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCloseAirdropCampaign({ value: MsgCloseAirdropCampaign.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCloseAirdropCampaign: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgClaim({ value, fee, memo }: sendMsgClaimParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgClaim: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgClaim({ value: MsgClaim.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgClaim: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgStartAirdropCampaign({ value }: msgStartAirdropCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgStartAirdropCampaign", value: MsgStartAirdropCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgStartAirdropCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgAddMissionToAidropCampaign({ value }: msgAddMissionToAidropCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddMissionToAidropCampaign", value: MsgAddMissionToAidropCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddMissionToAidropCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgCreateAirdropCampaign({ value }: msgCreateAirdropCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCreateAirdropCampaign", value: MsgCreateAirdropCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateAirdropCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgInitialClaim({ value }: msgInitialClaimParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgInitialClaim", value: MsgInitialClaim.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgInitialClaim: Could not create message: ' + e.message)
			}
		},
		
		msgEditAirdropCampaign({ value }: msgEditAirdropCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgEditAirdropCampaign", value: MsgEditAirdropCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgEditAirdropCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgAddAirdropEntries({ value }: msgAddAirdropEntriesParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgAddAirdropEntries", value: MsgAddAirdropEntries.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddAirdropEntries: Could not create message: ' + e.message)
			}
		},
		
		msgDeleteAirdropEntry({ value }: msgDeleteAirdropEntryParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgDeleteAirdropEntry", value: MsgDeleteAirdropEntry.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteAirdropEntry: Could not create message: ' + e.message)
			}
		},
		
		msgCloseAirdropCampaign({ value }: msgCloseAirdropCampaignParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgCloseAirdropCampaign", value: MsgCloseAirdropCampaign.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCloseAirdropCampaign: Could not create message: ' + e.message)
			}
		},
		
		msgClaim({ value }: msgClaimParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfeairdrop.MsgClaim", value: MsgClaim.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgClaim: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			Chain4EnergyC4EchainCfeairdrop: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;