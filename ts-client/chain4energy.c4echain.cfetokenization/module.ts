// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgBuyCertificate } from "./types/c4echain/cfetokenization/tx";
import { MsgAuthorizeCertificate } from "./types/c4echain/cfetokenization/tx";
import { MsgAcceptDevice } from "./types/c4echain/cfetokenization/tx";
import { MsgBurnCertificate } from "./types/c4echain/cfetokenization/tx";
import { MsgCreateUserCertificates } from "./types/c4echain/cfetokenization/tx";
import { MsgAssignDeviceToUser } from "./types/c4echain/cfetokenization/tx";
import { MsgAddCertificateToMarketplace } from "./types/c4echain/cfetokenization/tx";
import { MsgAddMeasurement } from "./types/c4echain/cfetokenization/tx";

import { CertificateType as typeCertificateType} from "./types"
import { Params as typeParams} from "./types"
import { UserCertificates as typeUserCertificates} from "./types"
import { Certificate as typeCertificate} from "./types"
import { CertificateOffer as typeCertificateOffer} from "./types"
import { UserDevices as typeUserDevices} from "./types"
import { UserDevice as typeUserDevice} from "./types"
import { PendingDevice as typePendingDevice} from "./types"
import { Device as typeDevice} from "./types"
import { Measurement as typeMeasurement} from "./types"

export { MsgBuyCertificate, MsgAuthorizeCertificate, MsgAcceptDevice, MsgBurnCertificate, MsgCreateUserCertificates, MsgAssignDeviceToUser, MsgAddCertificateToMarketplace, MsgAddMeasurement };

type sendMsgBuyCertificateParams = {
  value: MsgBuyCertificate,
  fee?: StdFee,
  memo?: string
};

type sendMsgAuthorizeCertificateParams = {
  value: MsgAuthorizeCertificate,
  fee?: StdFee,
  memo?: string
};

type sendMsgAcceptDeviceParams = {
  value: MsgAcceptDevice,
  fee?: StdFee,
  memo?: string
};

type sendMsgBurnCertificateParams = {
  value: MsgBurnCertificate,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateUserCertificatesParams = {
  value: MsgCreateUserCertificates,
  fee?: StdFee,
  memo?: string
};

type sendMsgAssignDeviceToUserParams = {
  value: MsgAssignDeviceToUser,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddCertificateToMarketplaceParams = {
  value: MsgAddCertificateToMarketplace,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddMeasurementParams = {
  value: MsgAddMeasurement,
  fee?: StdFee,
  memo?: string
};


type msgBuyCertificateParams = {
  value: MsgBuyCertificate,
};

type msgAuthorizeCertificateParams = {
  value: MsgAuthorizeCertificate,
};

type msgAcceptDeviceParams = {
  value: MsgAcceptDevice,
};

type msgBurnCertificateParams = {
  value: MsgBurnCertificate,
};

type msgCreateUserCertificatesParams = {
  value: MsgCreateUserCertificates,
};

type msgAssignDeviceToUserParams = {
  value: MsgAssignDeviceToUser,
};

type msgAddCertificateToMarketplaceParams = {
  value: MsgAddCertificateToMarketplace,
};

type msgAddMeasurementParams = {
  value: MsgAddMeasurement,
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
		
		async sendMsgBuyCertificate({ value, fee, memo }: sendMsgBuyCertificateParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgBuyCertificate: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgBuyCertificate({ value: MsgBuyCertificate.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgBuyCertificate: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAuthorizeCertificate({ value, fee, memo }: sendMsgAuthorizeCertificateParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAuthorizeCertificate: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAuthorizeCertificate({ value: MsgAuthorizeCertificate.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAuthorizeCertificate: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAcceptDevice({ value, fee, memo }: sendMsgAcceptDeviceParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAcceptDevice: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAcceptDevice({ value: MsgAcceptDevice.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAcceptDevice: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgBurnCertificate({ value, fee, memo }: sendMsgBurnCertificateParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgBurnCertificate: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgBurnCertificate({ value: MsgBurnCertificate.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgBurnCertificate: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateUserCertificates({ value, fee, memo }: sendMsgCreateUserCertificatesParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateUserCertificates: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateUserCertificates({ value: MsgCreateUserCertificates.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateUserCertificates: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAssignDeviceToUser({ value, fee, memo }: sendMsgAssignDeviceToUserParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAssignDeviceToUser: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAssignDeviceToUser({ value: MsgAssignDeviceToUser.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAssignDeviceToUser: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddCertificateToMarketplace({ value, fee, memo }: sendMsgAddCertificateToMarketplaceParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddCertificateToMarketplace: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddCertificateToMarketplace({ value: MsgAddCertificateToMarketplace.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddCertificateToMarketplace: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddMeasurement({ value, fee, memo }: sendMsgAddMeasurementParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddMeasurement: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddMeasurement({ value: MsgAddMeasurement.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddMeasurement: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgBuyCertificate({ value }: msgBuyCertificateParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfetokenization.MsgBuyCertificate", value: MsgBuyCertificate.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBuyCertificate: Could not create message: ' + e.message)
			}
		},
		
		msgAuthorizeCertificate({ value }: msgAuthorizeCertificateParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfetokenization.MsgAuthorizeCertificate", value: MsgAuthorizeCertificate.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAuthorizeCertificate: Could not create message: ' + e.message)
			}
		},
		
		msgAcceptDevice({ value }: msgAcceptDeviceParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfetokenization.MsgAcceptDevice", value: MsgAcceptDevice.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAcceptDevice: Could not create message: ' + e.message)
			}
		},
		
		msgBurnCertificate({ value }: msgBurnCertificateParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfetokenization.MsgBurnCertificate", value: MsgBurnCertificate.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgBurnCertificate: Could not create message: ' + e.message)
			}
		},
		
		msgCreateUserCertificates({ value }: msgCreateUserCertificatesParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfetokenization.MsgCreateUserCertificates", value: MsgCreateUserCertificates.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateUserCertificates: Could not create message: ' + e.message)
			}
		},
		
		msgAssignDeviceToUser({ value }: msgAssignDeviceToUserParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfetokenization.MsgAssignDeviceToUser", value: MsgAssignDeviceToUser.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAssignDeviceToUser: Could not create message: ' + e.message)
			}
		},
		
		msgAddCertificateToMarketplace({ value }: msgAddCertificateToMarketplaceParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfetokenization.MsgAddCertificateToMarketplace", value: MsgAddCertificateToMarketplace.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddCertificateToMarketplace: Could not create message: ' + e.message)
			}
		},
		
		msgAddMeasurement({ value }: msgAddMeasurementParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfetokenization.MsgAddMeasurement", value: MsgAddMeasurement.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddMeasurement: Could not create message: ' + e.message)
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
						CertificateType: getStructure(typeCertificateType.fromPartial({})),
						Params: getStructure(typeParams.fromPartial({})),
						UserCertificates: getStructure(typeUserCertificates.fromPartial({})),
						Certificate: getStructure(typeCertificate.fromPartial({})),
						CertificateOffer: getStructure(typeCertificateOffer.fromPartial({})),
						UserDevices: getStructure(typeUserDevices.fromPartial({})),
						UserDevice: getStructure(typeUserDevice.fromPartial({})),
						PendingDevice: getStructure(typePendingDevice.fromPartial({})),
						Device: getStructure(typeDevice.fromPartial({})),
						Measurement: getStructure(typeMeasurement.fromPartial({})),
						
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
			Chain4EnergyC4EchainCfetokenization: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;