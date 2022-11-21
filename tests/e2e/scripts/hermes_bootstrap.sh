#!/bin/bash

set -ex

# initialize Hermes relayer configuration
mkdir -p /root/.hermes/
touch /root/.hermes/config.toml

# setup Hermes relayer configuration
tee /root/.hermes/config.toml <<EOF
[global]
log_level = 'info'
[mode]
[mode.clients]
enabled = true
refresh = true
misbehaviour = false
[mode.connections]
enabled = false
[mode.channels]
enabled = false
[mode.packets]
enabled = true
clear_interval = 100
clear_on_start = true
tx_confirmation = true
auto_register_counterparty_payee = false
[rest]
enabled = true
host = '0.0.0.0'
port = 3031
[telemetry]
enabled = true
host = '127.0.0.1'
port = 3001
[[chains]]
id = '$OSMO_A_E2E_CHAIN_ID'
rpc_addr = 'http://$OSMO_A_E2E_VAL_HOST:26657'
grpc_addr = 'http://$OSMO_A_E2E_VAL_HOST:9090'
websocket_addr = 'ws://$OSMO_A_E2E_VAL_HOST:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'c4e'
key_name = 'val01-osmosis-a'
store_prefix = 'ibc'
default_gas = 100000
max_gas = 400000
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 180000
clock_drift = '5s'
max_block_time = '30s'
memo_prefix = ''
sequential_batch_tx = false
[chains.trust_threshold]
numerator = '1'
denominator = '3'

[chains.gas_price]
price = 0.1
denom = 'uc4e'

[chains.packet_filter]
 policy = 'allow'
 list = [
   ['transfer', 'channel-0'],
]

[chains.address_type]
derivation = 'cosmos'
[[chains]]
id = '$OSMO_B_E2E_CHAIN_ID'
rpc_addr = 'http://$OSMO_B_E2E_VAL_HOST:26657'
grpc_addr = 'http://$OSMO_B_E2E_VAL_HOST:9090'
websocket_addr = 'ws://$OSMO_B_E2E_VAL_HOST:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'c4e'
key_name = 'val01-osmosis-b'
store_prefix = 'ibc'
max_gas = 400000
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 180000
clock_drift = '5s'
max_block_time = '30s'
memo_prefix = ''
sequential_batch_tx = false

[chains.trust_threshold]
numerator = '1'
denominator = '3'

[chains.gas_price]
price = 0.1
denom = 'uosmo'

[chains.packet_filter]
 policy = 'allow'
 list = [
   ['transfer', 'channel-1490'],
 ]

[chains.address_type]
derivation = 'cosmos'
EOF

# import keys
hermes keys restore ${OSMO_B_E2E_CHAIN_ID} -n "val01-osmosis-b" -m "${OSMO_B_E2E_VAL_MNEMONIC}"
hermes keys restore ${OSMO_A_E2E_CHAIN_ID} -n "val01-osmosis-a" -m "${OSMO_A_E2E_VAL_MNEMONIC}"

# start Hermes relayer
hermes start
