// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import Chain4EnergyC4EchainCfedistributor from './chain4energy.c4echain.cfedistributor'
import Chain4EnergyC4EchainCfeminter from './chain4energy.c4echain.cfeminter'
import Chain4EnergyC4EchainCfesignature from './chain4energy.c4echain.cfesignature'


export default { 
  Chain4EnergyC4EchainCfedistributor: load(Chain4EnergyC4EchainCfedistributor, 'chain4energy.c4echain.cfedistributor'),
  Chain4EnergyC4EchainCfeminter: load(Chain4EnergyC4EchainCfeminter, 'chain4energy.c4echain.cfeminter'),
  Chain4EnergyC4EchainCfesignature: load(Chain4EnergyC4EchainCfesignature, 'chain4energy.c4echain.cfesignature'),
  
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