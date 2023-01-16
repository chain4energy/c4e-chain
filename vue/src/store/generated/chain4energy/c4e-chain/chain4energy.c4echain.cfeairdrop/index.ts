import { txClient, queryClient, MissingWalletError , registry} from './module'

import { ContinuousVestingPeriod } from "./module/types/cfeairdrop/account"
import { AirdropVestingAccount } from "./module/types/cfeairdrop/account"
import { UserAirdropEntries } from "./module/types/cfeairdrop/airdrop"
import { AirdropEntry } from "./module/types/cfeairdrop/airdrop"
import { AirdropEntries } from "./module/types/cfeairdrop/airdrop"
import { Campaign } from "./module/types/cfeairdrop/airdrop"
import { InitialClaim } from "./module/types/cfeairdrop/airdrop"
import { Mission } from "./module/types/cfeairdrop/airdrop"
import { Params } from "./module/types/cfeairdrop/params"


export { ContinuousVestingPeriod, AirdropVestingAccount, UserAirdropEntries, AirdropEntry, AirdropEntries, Campaign, InitialClaim, Mission, Params };

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
				UserAirdropEntries: {},
				UsersAirdropEntries: {},
				InitialClaim: {},
				InitialClaims: {},
				Mission: {},
				MissionAll: {},
				Campaigns: {},
				Campaign: {},
				
				_Structure: {
						ContinuousVestingPeriod: getStructure(ContinuousVestingPeriod.fromPartial({})),
						AirdropVestingAccount: getStructure(AirdropVestingAccount.fromPartial({})),
						UserAirdropEntries: getStructure(UserAirdropEntries.fromPartial({})),
						AirdropEntry: getStructure(AirdropEntry.fromPartial({})),
						AirdropEntries: getStructure(AirdropEntries.fromPartial({})),
						Campaign: getStructure(Campaign.fromPartial({})),
						InitialClaim: getStructure(InitialClaim.fromPartial({})),
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
				getUserAirdropEntries: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UserAirdropEntries[JSON.stringify(params)] ?? {}
		},
				getUsersAirdropEntries: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UsersAirdropEntries[JSON.stringify(params)] ?? {}
		},
				getInitialClaim: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InitialClaim[JSON.stringify(params)] ?? {}
		},
				getInitialClaims: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.InitialClaims[JSON.stringify(params)] ?? {}
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
		
		
		
		
		 		
		
		
		async QueryUserAirdropEntries({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUserAirdropEntries( key.address)).data
				
					
				commit('QUERY', { query: 'UserAirdropEntries', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUserAirdropEntries', payload: { options: { all }, params: {...key},query }})
				return getters['getUserAirdropEntries']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUserAirdropEntries API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUsersAirdropEntries({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUsersAirdropEntries(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryUsersAirdropEntries({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'UsersAirdropEntries', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUsersAirdropEntries', payload: { options: { all }, params: {...key},query }})
				return getters['getUsersAirdropEntries']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUsersAirdropEntries API Node Unavailable. Could not perform query: ' + e.message)
				
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
		
		
		
		
		 		
		
		
		async QueryInitialClaims({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryInitialClaims(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryInitialClaims({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'InitialClaims', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInitialClaims', payload: { options: { all }, params: {...key},query }})
				return getters['getInitialClaims']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInitialClaims API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMission({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryMission( key.campaignId,  key.missionId)).data
				
					
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryMissionAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryMissionAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCampaigns(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryCampaigns({...query, 'pagination.key':(<any> value).pagination.next_key})).data
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
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryCampaign( key.campaignId)).data
				
					
				commit('QUERY', { query: 'Campaign', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCampaign', payload: { options: { all }, params: {...key},query }})
				return getters['getCampaign']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCampaign API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgClaim({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgClaim(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaim:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgClaim:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddAirdropEntries({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddAirdropEntries(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAirdropEntries:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddAirdropEntries:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateAirdropCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateAirdropCampaign(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateAirdropCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateAirdropCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteAirdropEntry({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteAirdropEntry(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteAirdropEntry:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteAirdropEntry:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddMissionToAidropCampaign({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddMissionToAidropCampaign(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddMissionToAidropCampaign:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddMissionToAidropCampaign:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgClaim({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgClaim(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgClaim:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgClaim:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddAirdropEntries({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddAirdropEntries(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddAirdropEntries:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddAirdropEntries:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateAirdropCampaign({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateAirdropCampaign(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateAirdropCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateAirdropCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteAirdropEntry({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteAirdropEntry(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteAirdropEntry:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteAirdropEntry:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddMissionToAidropCampaign({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgAddMissionToAidropCampaign(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddMissionToAidropCampaign:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddMissionToAidropCampaign:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
