import { Client, registry, MissingWalletError } from 'chain4energy-c4e-chain-client-ts'

import { Campaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { CampaignTotalAmount } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { CampaignAmountLeft } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { UserEntry } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { ClaimRecord } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { NewCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EditCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { CloseCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { RemoveCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EnableCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { AddMissionToCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { Claim } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { InitialClaim } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { AddClaimRecords } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { DeleteClaimRecord } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { CompleteMissionFromHook } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { Mission } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { Params } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"


export { Campaign, CampaignTotalAmount, CampaignAmountLeft, UserEntry, ClaimRecord, NewCampaign, EditCampaign, CloseCampaign, RemoveCampaign, EnableCampaign, AddMissionToCampaign, Claim, InitialClaim, AddClaimRecords, DeleteClaimRecord, CompleteMissionFromHook, Mission, Params };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				UserEntry: {},
				UsersEntries: {},
				Mission: {},
				MissionAll: {},
				Campaigns: {},
				Campaign: {},
				CampaignTotalAmount: {},
				CampaignAmountLeft: {},
				
				_Structure: {
						Campaign: getStructure(Campaign.fromPartial({})),
						CampaignTotalAmount: getStructure(CampaignTotalAmount.fromPartial({})),
						CampaignAmountLeft: getStructure(CampaignAmountLeft.fromPartial({})),
						UserEntry: getStructure(UserEntry.fromPartial({})),
						ClaimRecord: getStructure(ClaimRecord.fromPartial({})),
						NewCampaign: getStructure(NewCampaign.fromPartial({})),
						EditCampaign: getStructure(EditCampaign.fromPartial({})),
						CloseCampaign: getStructure(CloseCampaign.fromPartial({})),
						RemoveCampaign: getStructure(RemoveCampaign.fromPartial({})),
						EnableCampaign: getStructure(EnableCampaign.fromPartial({})),
						AddMissionToCampaign: getStructure(AddMissionToCampaign.fromPartial({})),
						Claim: getStructure(Claim.fromPartial({})),
						InitialClaim: getStructure(InitialClaim.fromPartial({})),
						AddClaimRecords: getStructure(AddClaimRecords.fromPartial({})),
						DeleteClaimRecord: getStructure(DeleteClaimRecord.fromPartial({})),
						CompleteMissionFromHook: getStructure(CompleteMissionFromHook.fromPartial({})),
						Mission: getStructure(Mission.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getUserEntry: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UserEntry[JSON.stringify(params)] ?? {}
		},
				getUsersEntries: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UsersEntries[JSON.stringify(params)] ?? {}
		},
				getMission: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Mission[JSON.stringify(params)] ?? {}
		},
				getMissionAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MissionAll[JSON.stringify(params)] ?? {}
		},
				getCampaigns: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Campaigns[JSON.stringify(params)] ?? {}
		},
				getCampaign: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Campaign[JSON.stringify(params)] ?? {}
		},
				getCampaignTotalAmount: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CampaignTotalAmount[JSON.stringify(params)] ?? {}
		},
				getCampaignAmountLeft: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CampaignAmountLeft[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: chain4energy.c4echain.cfeclaim initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUserEntry({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryUserEntry( key.address)).data
				
					
				commit('QUERY', { query: 'UserEntry', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUserEntry', payload: { options: { all }, params: {...key},query }})
				return getters['getUserEntry']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUserEntry API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUsersEntries({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryUsersEntries(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.Chain4EnergyC4EchainCfeclaim.query.queryUsersEntries({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'UsersEntries', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUsersEntries', payload: { options: { all }, params: {...key},query }})
				return getters['getUsersEntries']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUsersEntries API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMission({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryMission( key.campaign_id,  key.mission_id)).data
				
					
				commit('QUERY', { query: 'Mission', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMission', payload: { options: { all }, params: {...key},query }})
				return getters['getMission']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMission API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMissionAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryMissionAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.Chain4EnergyC4EchainCfeclaim.query.queryMissionAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'MissionAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMissionAll', payload: { options: { all }, params: {...key},query }})
				return getters['getMissionAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMissionAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCampaigns({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryCampaigns(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.Chain4EnergyC4EchainCfeclaim.query.queryCampaigns({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Campaigns', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCampaigns', payload: { options: { all }, params: {...key},query }})
				return getters['getCampaigns']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCampaigns API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCampaign({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryCampaign( key.campaign_id)).data
				
					
				commit('QUERY', { query: 'Campaign', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCampaign', payload: { options: { all }, params: {...key},query }})
				return getters['getCampaign']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCampaign API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCampaignTotalAmount({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryCampaignTotalAmount( key.campaign_id)).data
				
					
				commit('QUERY', { query: 'CampaignTotalAmount', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCampaignTotalAmount', payload: { options: { all }, params: {...key},query }})
				return getters['getCampaignTotalAmount']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCampaignTotalAmount API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCampaignAmountLeft({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryCampaignAmountLeft( key.campaign_id)).data
				
					
				commit('QUERY', { query: 'CampaignAmountLeft', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCampaignAmountLeft', payload: { options: { all }, params: {...key},query }})
				return getters['getCampaignAmountLeft']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCampaignAmountLeft API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgClaim({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgClaim({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgCreateCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteClaimRecord({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgDeleteClaimRecord({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteClaimRecord:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteClaimRecord:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgInitialClaim({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgInitialClaim({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInitialClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInitialClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgEditCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgEditCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEditCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgEditCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRemoveCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgRemoveCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRemoveCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCloseCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgCloseCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCloseCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCloseCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddClaimRecords({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgAddClaimRecords({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddClaimRecords:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddClaimRecords:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddMissionToCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgAddMissionToCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddMissionToCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddMissionToCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgEnableCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgEnableCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEnableCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgEnableCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgClaim({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgClaim({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaim:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgClaim:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateCampaign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgCreateCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteClaimRecord({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgDeleteClaimRecord({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteClaimRecord:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteClaimRecord:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgInitialClaim({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgInitialClaim({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInitialClaim:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInitialClaim:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgEditCampaign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgEditCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEditCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgEditCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRemoveCampaign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgRemoveCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRemoveCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCloseCampaign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgCloseCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCloseCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCloseCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddClaimRecords({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgAddClaimRecords({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddClaimRecords:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddClaimRecords:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddMissionToCampaign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgAddMissionToCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddMissionToCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddMissionToCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgEnableCampaign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.MsgEnableCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEnableCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgEnableCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
