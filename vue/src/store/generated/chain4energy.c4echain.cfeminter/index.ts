import { Client, registry, MissingWalletError } from 'chain4energy-c4e-chain-client-ts'

import { EventMint } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeminter/types"
import { Minter } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeminter/types"
import { NoMinting } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeminter/types"
import { LinearMinting } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeminter/types"
import { ExponentialStepMinting } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeminter/types"
import { MinterState } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeminter/types"
import { Params } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfeminter/types"


export { EventMint, Minter, NoMinting, LinearMinting, ExponentialStepMinting, MinterState, Params };

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
				Inflation: {},
				State: {},
				
				_Structure: {
						EventMint: getStructure(EventMint.fromPartial({})),
						Minter: getStructure(Minter.fromPartial({})),
						NoMinting: getStructure(NoMinting.fromPartial({})),
						LinearMinting: getStructure(LinearMinting.fromPartial({})),
						ExponentialStepMinting: getStructure(ExponentialStepMinting.fromPartial({})),
						MinterState: getStructure(MinterState.fromPartial({})),
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
				getInflation: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Inflation[JSON.stringify(params)] ?? {}
		},
				getState: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.State[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: chain4energy.c4echain.cfeminter initialized!')
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
				let value= (await client.Chain4EnergyC4EchainCfeminter.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryInflation({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeminter.query.queryInflation()).data
				
					
				commit('QUERY', { query: 'Inflation', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryInflation', payload: { options: { all }, params: {...key},query }})
				return getters['getInflation']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryInflation API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryState({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfeminter.query.queryState()).data
				
					
				commit('QUERY', { query: 'State', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryState', payload: { options: { all }, params: {...key},query }})
				return getters['getState']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryState API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgUpdateMintersParams({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeminter.tx.sendMsgUpdateMintersParams({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateMintersParams:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateMintersParams:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgBurn({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeminter.tx.sendMsgBurn({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurn:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgBurn:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateParams({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.Chain4EnergyC4EchainCfeminter.tx.sendMsgUpdateParams({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateParams:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateParams:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgUpdateMintersParams({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeminter.tx.msgUpdateMintersParams({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateMintersParams:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateMintersParams:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgBurn({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeminter.tx.msgBurn({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgBurn:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgBurn:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateParams({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfeminter.tx.msgUpdateParams({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateParams:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateParams:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}