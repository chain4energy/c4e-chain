import { Client, registry, MissingWalletError } from 'chain4energy-c4e-chain-client-ts'

import { AccountVestingPools } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { VestingPool } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { NewVestingAccount } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { NewVestingPool } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { NewVestingAccountFromVestingPool } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { WithdrawAvailable } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { VestingSplit } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { GenesisVestingType } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { Params } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { VestingPoolInfo } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { VestingAccountTrace } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { ContinuousVestingPeriod } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { RepeatedContinuousVestingAccount } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { BaseVestingAccount } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { AuthBaseAccount } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { VestingTypes } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"
import { VestingType } from "chain4energy-c4e-chain-client-ts/chain4energy.c4echain.cfevesting/types"


export { AccountVestingPools, VestingPool, NewVestingAccount, NewVestingPool, NewVestingAccountFromVestingPool, WithdrawAvailable, VestingSplit, GenesisVestingType, Params, VestingPoolInfo, VestingAccountTrace, ContinuousVestingPeriod, RepeatedContinuousVestingAccount, BaseVestingAccount, AuthBaseAccount, VestingTypes, VestingType };

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
				VestingType: {},
				VestingPools: {},
				VestingsSummary: {},
				GenesisVestingsSummary: {},
				
				_Structure: {
						AccountVestingPools: getStructure(AccountVestingPools.fromPartial({})),
						VestingPool: getStructure(VestingPool.fromPartial({})),
						NewVestingAccount: getStructure(NewVestingAccount.fromPartial({})),
						NewVestingPool: getStructure(NewVestingPool.fromPartial({})),
						NewVestingAccountFromVestingPool: getStructure(NewVestingAccountFromVestingPool.fromPartial({})),
						WithdrawAvailable: getStructure(WithdrawAvailable.fromPartial({})),
						VestingSplit: getStructure(VestingSplit.fromPartial({})),
						GenesisVestingType: getStructure(GenesisVestingType.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						VestingPoolInfo: getStructure(VestingPoolInfo.fromPartial({})),
						VestingAccountTrace: getStructure(VestingAccountTrace.fromPartial({})),
						ContinuousVestingPeriod: getStructure(ContinuousVestingPeriod.fromPartial({})),
						RepeatedContinuousVestingAccount: getStructure(RepeatedContinuousVestingAccount.fromPartial({})),
						BaseVestingAccount: getStructure(BaseVestingAccount.fromPartial({})),
						AuthBaseAccount: getStructure(AuthBaseAccount.fromPartial({})),
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
				getVestingPools: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VestingPools[JSON.stringify(params)] ?? {}
		},
				getVestingsSummary: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VestingsSummary[JSON.stringify(params)] ?? {}
		},
				getGenesisVestingsSummary: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GenesisVestingsSummary[JSON.stringify(params)] ?? {}
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
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfevesting.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVestingType({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfevesting.query.queryVestingType()).data
				
					
				commit('QUERY', { query: 'VestingType', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVestingType', payload: { options: { all }, params: {...key},query }})
				return getters['getVestingType']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVestingType API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVestingPools({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfevesting.query.queryVestingPools( key.owner)).data
				
					
				commit('QUERY', { query: 'VestingPools', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVestingPools', payload: { options: { all }, params: {...key},query }})
				return getters['getVestingPools']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVestingPools API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryVestingsSummary({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfevesting.query.queryVestingsSummary()).data
				
					
				commit('QUERY', { query: 'VestingsSummary', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVestingsSummary', payload: { options: { all }, params: {...key},query }})
				return getters['getVestingsSummary']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVestingsSummary API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryGenesisVestingsSummary({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.Chain4EnergyC4EchainCfevesting.query.queryGenesisVestingsSummary()).data
				
					
				commit('QUERY', { query: 'GenesisVestingsSummary', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryGenesisVestingsSummary', payload: { options: { all }, params: {...key},query }})
				return getters['getGenesisVestingsSummary']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryGenesisVestingsSummary API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgMoveAvailableVesting({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfevesting.tx.sendMsgMoveAvailableVesting({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMoveAvailableVesting:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgMoveAvailableVesting:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateVestingPool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfevesting.tx.sendMsgCreateVestingPool({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVestingPool:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateVestingPool:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSplitVesting({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfevesting.tx.sendMsgSplitVesting({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSplitVesting:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSplitVesting:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawAllAvailable({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfevesting.tx.sendMsgWithdrawAllAvailable({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawAllAvailable:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdrawAllAvailable:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateDenomParam({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfevesting.tx.sendMsgUpdateDenomParam({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateDenomParam:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateDenomParam:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgMoveAvailableVestingByDenoms({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfevesting.tx.sendMsgMoveAvailableVestingByDenoms({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMoveAvailableVestingByDenoms:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgMoveAvailableVestingByDenoms:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateVestingAccount({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfevesting.tx.sendMsgCreateVestingAccount({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVestingAccount:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateVestingAccount:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSendToVestingAccount({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.Chain4EnergyC4EchainCfevesting.tx.sendMsgSendToVestingAccount({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendToVestingAccount:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSendToVestingAccount:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgMoveAvailableVesting({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfevesting.tx.msgMoveAvailableVesting({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMoveAvailableVesting:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgMoveAvailableVesting:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateVestingPool({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfevesting.tx.msgCreateVestingPool({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVestingPool:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateVestingPool:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSplitVesting({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfevesting.tx.msgSplitVesting({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSplitVesting:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSplitVesting:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgWithdrawAllAvailable({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfevesting.tx.msgWithdrawAllAvailable({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawAllAvailable:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdrawAllAvailable:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateDenomParam({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfevesting.tx.msgUpdateDenomParam({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateDenomParam:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateDenomParam:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgMoveAvailableVestingByDenoms({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfevesting.tx.msgMoveAvailableVestingByDenoms({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgMoveAvailableVestingByDenoms:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgMoveAvailableVestingByDenoms:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateVestingAccount({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfevesting.tx.msgCreateVestingAccount({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVestingAccount:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateVestingAccount:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSendToVestingAccount({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.Chain4EnergyC4EchainCfevesting.tx.msgSendToVestingAccount({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendToVestingAccount:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSendToVestingAccount:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
