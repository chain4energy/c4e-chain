import { txClient, queryClient, MissingWalletError , registry} from './module'

import { ContinuousVestingPeriod } from "./module/types/cfeairdrop/account"
import { AirdropVestingAccount } from "./module/types/cfeairdrop/account"
import { CampaignRecord } from "./module/types/cfeairdrop/airdrop"
import { ClaimRecord } from "./module/types/cfeairdrop/airdrop"
import { Campaign } from "./module/types/cfeairdrop/airdrop"
import { InitialClaimff } from "./module/types/cfeairdrop/airdrop"
import { Mission } from "./module/types/cfeairdrop/airdrop"
import { InitialClaim } from "./module/types/cfeairdrop/initial_claim"
import { Params } from "./module/types/cfeairdrop/params"


export { ContinuousVestingPeriod, AirdropVestingAccount, CampaignRecord, ClaimRecord, Campaign, InitialClaimff, Mission, InitialClaim, Params };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
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

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},
				ClaimRecord: {},
				ClaimRecordAll: {},
				InitialClaim: {},
				InitialClaimAll: {},
				
				_Structure: {
						ContinuousVestingPeriod: getStructure(ContinuousVestingPeriod.fromPartial({})),
						AirdropVestingAccount: getStructure(AirdropVestingAccount.fromPartial({})),
						CampaignRecord: getStructure(CampaignRecord.fromPartial({})),
						ClaimRecord: getStructure(ClaimRecord.fromPartial({})),
						Campaign: getStructure(Campaign.fromPartial({})),
						InitialClaimff: getStructure(InitialClaimff.fromPartial({})),
						Mission: getStructure(Mission.fromPartial({})),
						InitialClaim: getStructure(InitialClaim.fromPartial({})),
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
				getClaimRecord: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClaimRecord[JSON.stringify(params)] ?? {}
		},
				getClaimRecordAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ClaimRecordAll[JSON.stringify(params)] ?? {}
		},
				getInitialClaim: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InitialClaim[JSON.stringify(params)] ?? {}
		},
				getInitialClaimAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InitialClaimAll[JSON.stringify(params)] ?? {}
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClaimRecord({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryClaimRecord( key.address)).data
				
					
				commit('QUERY', { query: 'ClaimRecord', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClaimRecord', payload: { options: { all }, params: {...key},query }})
				return getters['getClaimRecord']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClaimRecord API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryClaimRecordAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryClaimRecordAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryClaimRecordAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ClaimRecordAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryClaimRecordAll', payload: { options: { all }, params: {...key},query }})
				return getters['getClaimRecordAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryClaimRecordAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInitialClaim({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryInitialClaim( key.campaignId)).data
				
					
				commit('QUERY', { query: 'InitialClaim', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInitialClaim', payload: { options: { all }, params: {...key},query }})
				return getters['getInitialClaim']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInitialClaim API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInitialClaimAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryInitialClaimAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryInitialClaimAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'InitialClaimAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInitialClaimAll', payload: { options: { all }, params: {...key},query }})
				return getters['getInitialClaimAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInitialClaimAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
	}
}
