import { txClient, queryClient, MissingWalletError , registry} from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'

import { AccountVestingsList } from "./module/types/cfevesting/account_vesting"
import { AccountVestings } from "./module/types/cfevesting/account_vesting"
import { Vesting } from "./module/types/cfevesting/account_vesting"
import { Params } from "./module/types/cfevesting/params"
import { VestingInfo } from "./module/types/cfevesting/query"
import { VestingTypes } from "./module/types/cfevesting/vesting_types"
import { VestingType } from "./module/types/cfevesting/vesting_types"


export { AccountVestingsList, AccountVestings, Vesting, Params, VestingInfo, VestingTypes, VestingType };

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
				VestingType: {},
				Vesting: {},
				
				_Structure: {
						AccountVestingsList: getStructure(AccountVestingsList.fromPartial({})),
						AccountVestings: getStructure(AccountVestings.fromPartial({})),
						Vesting: getStructure(Vesting.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						VestingInfo: getStructure(VestingInfo.fromPartial({})),
						VestingTypes: getStructure(VestingTypes.fromPartial({})),
						VestingType: getStructure(VestingType.fromPartial({})),
						
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
				getVestingType: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VestingType[JSON.stringify(params)] ?? {}
		},
				getVesting: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Vesting[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: chain4energy.c4echain.cfevesting initialized!')
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
					throw new SpVuexError('Subscriptions: ' + e.message)
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
				throw new SpVuexError('QueryClient:QueryParams', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVestingType({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVestingType()).data
				
					
				commit('QUERY', { query: 'VestingType', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVestingType', payload: { options: { all }, params: {...key},query }})
				return getters['getVestingType']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryVestingType', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVesting({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVesting( key.address)).data
				
					
				commit('QUERY', { query: 'Vesting', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVesting', payload: { options: { all }, params: {...key},query }})
				return getters['getVesting']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryVesting', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgDelegate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDelegate(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgDelegate:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDelegate:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawDelegatorReward({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgWithdrawDelegatorReward(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgWithdrawDelegatorReward:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgWithdrawDelegatorReward:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgVest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgVest(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgVest:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgVest:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawAllAvailable({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgWithdrawAllAvailable(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgWithdrawAllAvailable:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgWithdrawAllAvailable:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBeginRedelegate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBeginRedelegate(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgBeginRedelegate:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgBeginRedelegate:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUndelegate({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUndelegate(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUndelegate:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUndelegate:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSendVesting({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendVesting(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSendVesting:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSendVesting:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgDelegate({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDelegate(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgDelegate:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDelegate:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgWithdrawDelegatorReward({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgWithdrawDelegatorReward(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgWithdrawDelegatorReward:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgWithdrawDelegatorReward:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgVest({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgVest(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgVest:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgVest:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgWithdrawAllAvailable({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgWithdrawAllAvailable(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgWithdrawAllAvailable:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgWithdrawAllAvailable:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgBeginRedelegate({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgBeginRedelegate(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgBeginRedelegate:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgBeginRedelegate:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgUndelegate({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUndelegate(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUndelegate:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUndelegate:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgSendVesting({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendVesting(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgSendVesting:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgSendVesting:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		
	}
}
