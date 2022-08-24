// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Chain4EnergyC4EChainChain4EnergyC4EchainCfeenergybank from './chain4energy/c4e-chain/chain4energy.c4echain.cfeenergybank'
import Chain4EnergyC4EChainChain4EnergyC4EchainCfeminter from './chain4energy/c4e-chain/chain4energy.c4echain.cfeminter'
import Chain4EnergyC4EChainChain4EnergyC4EchainCferoutingdistributor from './chain4energy/c4e-chain/chain4energy.c4echain.cferoutingdistributor'
import Chain4EnergyC4EChainChain4EnergyC4EchainCfesignature from './chain4energy/c4e-chain/chain4energy.c4echain.cfesignature'
import Chain4EnergyC4EChainChain4EnergyC4EchainCfevesting from './chain4energy/c4e-chain/chain4energy.c4echain.cfevesting'
import Chain4EnergyC4EChainChain4EnergyC4EchainEnergybank from './chain4energy/c4e-chain/chain4energy.c4echain.energybank'


export default { 
  Chain4EnergyC4EChainChain4EnergyC4EchainCfeenergybank: load(Chain4EnergyC4EChainChain4EnergyC4EchainCfeenergybank, 'chain4energy.c4echain.cfeenergybank'),
  Chain4EnergyC4EChainChain4EnergyC4EchainCfeminter: load(Chain4EnergyC4EChainChain4EnergyC4EchainCfeminter, 'chain4energy.c4echain.cfeminter'),
  Chain4EnergyC4EChainChain4EnergyC4EchainCferoutingdistributor: load(Chain4EnergyC4EChainChain4EnergyC4EchainCferoutingdistributor, 'chain4energy.c4echain.cferoutingdistributor'),
  Chain4EnergyC4EChainChain4EnergyC4EchainCfesignature: load(Chain4EnergyC4EChainChain4EnergyC4EchainCfesignature, 'chain4energy.c4echain.cfesignature'),
  Chain4EnergyC4EChainChain4EnergyC4EchainCfevesting: load(Chain4EnergyC4EChainChain4EnergyC4EchainCfevesting, 'chain4energy.c4echain.cfevesting'),
  Chain4EnergyC4EChainChain4EnergyC4EchainEnergybank: load(Chain4EnergyC4EChainChain4EnergyC4EchainEnergybank, 'chain4energy.c4echain.energybank'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
