// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgUpdateSubDistributorDestinationShareParam } from "./types/c4echain/cfedistributor/tx";
import { MsgUpdateParams } from "./types/c4echain/cfedistributor/tx";
import { MsgUpdateSubDistributorBurnShareParam } from "./types/c4echain/cfedistributor/tx";
import { MsgUpdateSubDistributorParam } from "./types/c4echain/cfedistributor/tx";

import { EventDistribution as typeEventDistribution} from "./types"
import { EventDistributionBurn as typeEventDistributionBurn} from "./types"
import { Params as typeParams} from "./types"
import { State as typeState} from "./types"
import { SubDistributor as typeSubDistributor} from "./types"
import { Destinations as typeDestinations} from "./types"
import { DestinationShare as typeDestinationShare} from "./types"
import { Account as typeAccount} from "./types"

export { MsgUpdateSubDistributorDestinationShareParam, MsgUpdateParams, MsgUpdateSubDistributorBurnShareParam, MsgUpdateSubDistributorParam };

type sendMsgUpdateSubDistributorDestinationShareParamParams = {
  value: MsgUpdateSubDistributorDestinationShareParam,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateParamsParams = {
  value: MsgUpdateParams,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateSubDistributorBurnShareParamParams = {
  value: MsgUpdateSubDistributorBurnShareParam,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdateSubDistributorParamParams = {
  value: MsgUpdateSubDistributorParam,
  fee?: StdFee,
  memo?: string
};


type msgUpdateSubDistributorDestinationShareParamParams = {
  value: MsgUpdateSubDistributorDestinationShareParam,
};

type msgUpdateParamsParams = {
  value: MsgUpdateParams,
};

type msgUpdateSubDistributorBurnShareParamParams = {
  value: MsgUpdateSubDistributorBurnShareParam,
};

type msgUpdateSubDistributorParamParams = {
  value: MsgUpdateSubDistributorParam,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
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
		
		async sendMsgUpdateSubDistributorDestinationShareParam({ value, fee, memo }: sendMsgUpdateSubDistributorDestinationShareParamParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateSubDistributorDestinationShareParam: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateSubDistributorDestinationShareParam({ value: MsgUpdateSubDistributorDestinationShareParam.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateSubDistributorDestinationShareParam: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateParams({ value, fee, memo }: sendMsgUpdateParamsParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateParams: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateParams({ value: MsgUpdateParams.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateParams: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateSubDistributorBurnShareParam({ value, fee, memo }: sendMsgUpdateSubDistributorBurnShareParamParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateSubDistributorBurnShareParam: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateSubDistributorBurnShareParam({ value: MsgUpdateSubDistributorBurnShareParam.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateSubDistributorBurnShareParam: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdateSubDistributorParam({ value, fee, memo }: sendMsgUpdateSubDistributorParamParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateSubDistributorParam: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateSubDistributorParam({ value: MsgUpdateSubDistributorParam.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateSubDistributorParam: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgUpdateSubDistributorDestinationShareParam({ value }: msgUpdateSubDistributorDestinationShareParamParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfedistributor.MsgUpdateSubDistributorDestinationShareParam", value: MsgUpdateSubDistributorDestinationShareParam.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateSubDistributorDestinationShareParam: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateParams({ value }: msgUpdateParamsParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfedistributor.MsgUpdateParams", value: MsgUpdateParams.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateParams: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateSubDistributorBurnShareParam({ value }: msgUpdateSubDistributorBurnShareParamParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfedistributor.MsgUpdateSubDistributorBurnShareParam", value: MsgUpdateSubDistributorBurnShareParam.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateSubDistributorBurnShareParam: Could not create message: ' + e.message)
			}
		},
		
		msgUpdateSubDistributorParam({ value }: msgUpdateSubDistributorParamParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfedistributor.MsgUpdateSubDistributorParam", value: MsgUpdateSubDistributorParam.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateSubDistributorParam: Could not create message: ' + e.message)
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
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						EventDistribution: getStructure(typeEventDistribution.fromPartial({})),
						EventDistributionBurn: getStructure(typeEventDistributionBurn.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						State: getStructure(typeState.fromPartial({})),
						SubDistributor: getStructure(typeSubDistributor.fromPartial({})),
						Destinations: getStructure(typeDestinations.fromPartial({})),
						DestinationShare: getStructure(typeDestinationShare.fromPartial({})),
						Account: getStructure(typeAccount.fromPartial({})),
						
		};
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
			Chain4EnergyC4EchainCfedistributor: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;