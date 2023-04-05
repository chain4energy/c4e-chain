// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Chain4EnergyC4EchainCfeairdrop from './chain4energy.c4echain.cfeairdrop'
import Chain4EnergyC4EchainCfedistributor from './chain4energy.c4echain.cfedistributor'
import Chain4EnergyC4EchainCfeminter from './chain4energy.c4echain.cfeminter'
import Chain4EnergyC4EchainCfesignature from './chain4energy.c4echain.cfesignature'
import Chain4EnergyC4EchainCfevesting from './chain4energy.c4echain.cfevesting'


export default { 
  Chain4EnergyC4EchainCfeairdrop: load(Chain4EnergyC4EchainCfeairdrop, 'chain4energy.c4echain.cfeairdrop'),
  Chain4EnergyC4EchainCfedistributor: load(Chain4EnergyC4EchainCfedistributor, 'chain4energy.c4echain.cfedistributor'),
  Chain4EnergyC4EchainCfeminter: load(Chain4EnergyC4EchainCfeminter, 'chain4energy.c4echain.cfeminter'),
  Chain4EnergyC4EchainCfesignature: load(Chain4EnergyC4EchainCfesignature, 'chain4energy.c4echain.cfesignature'),
  Chain4EnergyC4EchainCfevesting: load(Chain4EnergyC4EchainCfevesting, 'chain4energy.c4echain.cfevesting'),
  
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