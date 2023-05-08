import { Client, registry, MissingWalletError } from 'chain4energy-c4e-chain-client-ts'

import { Campaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { CampaignTotalAmount } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { CampaignAmountLeft } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { UserEntry } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { ClaimRecord } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { NewCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { EditCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { CloseCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { RemoveCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { StartCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { AddMissionToCampaign } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { Claim } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { InitialClaim } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { AddClaimRecords } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { DeleteClaimRecord } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { CompleteMissionFromHook } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { Mission } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"
import { Params } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeairdrop/types"


export { Campaign, CampaignTotalAmount, CampaignAmountLeft, UserEntry, ClaimRecord, NewCampaign, EditCampaign, CloseCampaign, RemoveCampaign, StartCampaign, AddMissionToCampaign, Claim, InitialClaim, AddClaimRecords, DeleteClaimRecord, CompleteMissionFromHook, Mission, Params };

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
						StartCampaign: getStructure(StartCampaign.fromPartial({})),
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
			console.log('Vuex module: chain4energy.c4echain.cfeairdrop initialized!')
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryParams()).data
				
					
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryUserEntry( key.address)).data
				
					
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryUsersEntries(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.Chain4EnergyC4EchainCfeairdrop.query.queryUsersEntries({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryMission( key.campaign_id,  key.mission_id)).data
				
					
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryMissionAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.Chain4EnergyC4EchainCfeairdrop.query.queryMissionAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryCampaigns(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.Chain4EnergyC4EchainCfeairdrop.query.queryCampaigns({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryCampaign( key.campaign_id)).data
				
					
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryCampaignTotalAmount( key.campaign_id)).data
				
					
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
				let value= (await client.Chain4EnergyC4EchainCfeairdrop.query.queryCampaignAmountLeft( key.campaign_id)).data
				
					
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgClaim({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgCreateCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgDeleteClaimRecord({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgInitialClaim({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgEditCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgRemoveCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgCloseCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgAddClaimRecords({ value, fee: {amount: fee, gas: "200000"}, memo })
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
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgAddMissionToCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddMissionToCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddMissionToCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgStartCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfeairdrop.tx.sendMsgStartCampaign({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgStartCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgStartCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgClaim({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgClaim({value})
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
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgCreateCampaign({value})
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
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgDeleteClaimRecord({value})
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
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgInitialClaim({value})
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
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgEditCampaign({value})
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
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgRemoveCampaign({value})
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
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgCloseCampaign({value})
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
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgAddClaimRecords({value})
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
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgAddMissionToCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddMissionToCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddMissionToCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgStartCampaign({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeairdrop.tx.msgStartCampaign({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgStartCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgStartCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
