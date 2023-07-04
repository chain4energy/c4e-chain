import { Client, registry, MissingWalletError } from 'chain4energy-c4e-chain-client-ts'

import { Campaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { UserEntry } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { ClaimRecord } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { ClaimRecordEntry } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventNewCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventCloseCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventRemoveCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventEnableCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventAddMission } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventClaim } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventInitialClaim } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventAddClaimRecords } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventDeleteClaimRecord } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { EventCompleteMission } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { MissionCount } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"
import { Mission } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeclaim/types"


export { Campaign, UserEntry, ClaimRecord, ClaimRecordEntry, EventNewCampaign, EventCloseCampaign, EventRemoveCampaign, EventEnableCampaign, EventAddMission, EventClaim, EventInitialClaim, EventAddClaimRecords, EventDeleteClaimRecord, EventCompleteMission, MissionCount, Mission };

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
				UserEntry: {},
				UsersEntries: {},
				Mission: {},
				CampaignMissions: {},
				Missions: {},
				Campaign: {},
				Campaigns: {},
				
				_Structure: {
						Campaign: getStructure(Campaign.fromPartial({})),
						UserEntry: getStructure(UserEntry.fromPartial({})),
						ClaimRecord: getStructure(ClaimRecord.fromPartial({})),
						ClaimRecordEntry: getStructure(ClaimRecordEntry.fromPartial({})),
						EventNewCampaign: getStructure(EventNewCampaign.fromPartial({})),
						EventCloseCampaign: getStructure(EventCloseCampaign.fromPartial({})),
						EventRemoveCampaign: getStructure(EventRemoveCampaign.fromPartial({})),
						EventEnableCampaign: getStructure(EventEnableCampaign.fromPartial({})),
						EventAddMission: getStructure(EventAddMission.fromPartial({})),
						EventClaim: getStructure(EventClaim.fromPartial({})),
						EventInitialClaim: getStructure(EventInitialClaim.fromPartial({})),
						EventAddClaimRecords: getStructure(EventAddClaimRecords.fromPartial({})),
						EventDeleteClaimRecord: getStructure(EventDeleteClaimRecord.fromPartial({})),
						EventCompleteMission: getStructure(EventCompleteMission.fromPartial({})),
						MissionCount: getStructure(MissionCount.fromPartial({})),
						Mission: getStructure(Mission.fromPartial({})),
						
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
				getCampaignMissions: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CampaignMissions[JSON.stringify(params)] ?? {}
		},
				getMissions: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Missions[JSON.stringify(params)] ?? {}
		},
				getCampaign: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Campaign[JSON.stringify(params)] ?? {}
		},
				getCampaigns: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Campaigns[JSON.stringify(params)] ?? {}
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
		
		
		
		
		 		
		
		
		async QueryCampaignMissions({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryCampaignMissions( key.campaign_id)).data
				
					
				commit('QUERY', { query: 'CampaignMissions', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCampaignMissions', payload: { options: { all }, params: {...key},query }})
				return getters['getCampaignMissions']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCampaignMissions API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMissions({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeclaim.query.queryMissions(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.Chain4EnergyC4EchainCfeclaim.query.queryMissions({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Missions', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMissions', payload: { options: { all }, params: {...key},query }})
				return getters['getMissions']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMissions API Node Unavailable. Could not perform query: ' + e.message)
				
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
		
		
		async sendMsgInitialClaim({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgInitialClaim({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInitialClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInitialClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddMission({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgAddMission({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddMission:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddMission:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteClaimRecord({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgDeleteClaimRecord({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteClaimRecord:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteClaimRecord:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateCampaign({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgCreateCampaign({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgEnableCampaign({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgEnableCampaign({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEnableCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgEnableCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRemoveCampaign({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgRemoveCampaign({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRemoveCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgClaim({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgClaim({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddClaimRecords({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgAddClaimRecords({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddClaimRecords:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddClaimRecords:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCloseCampaign({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeclaim.tx.sendMsgCloseCampaign({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCloseCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCloseCampaign:Send Could not broadcast Tx: '+ e.message)
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
		async MsgAddMission({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgAddMission({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddMission:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddMission:Create Could not create message: ' + e.message)
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
		async MsgEnableCampaign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeclaim.tx.msgEnableCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEnableCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgEnableCampaign:Create Could not create message: ' + e.message)
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
		
	}
}