#!/bin/bash
PARALLEL_SIMULATIONS_NUMBER=1
SIMULATION_RESULT_DIRECTORY=simulationResults
mkdir -p $SIMULATION_RESULT_DIRECTORY

run_simulation(){
  while true
  do
    START_TIME=$(date '+%Y-%m-%d_%H:%M:%S')
    echo "Running simulation $i.. Current date: $START_TIME"
    SEED=$(shuf -i 0-6500 -n 1)
    SIM_NUM_BLOCKS=$(shuf -i 200-20000 -n 1)
    SIM_BLOCK_SIZE=$(shuf -i 200-400 -n 1)
    SIMAPP=./app
    simulationResultFile=$SIMULATION_RESULT_DIRECTORY/.simulation-result.$i.$START_TIME

    echo "Running application benchmark for numBlocks=$SIM_NUM_BLOCKS, blockSize=$SIM_BLOCK_SIZE, seed=$SEED current time=$START_TIME. This may take awhile!" > "$simulationResultFile"

    go test -run=^$ -bench ^BenchmarkSimulation $SIMAPP -Seed="$SEED" -NumBlocks="$SIM_NUM_BLOCKS" -BlockSize="$SIM_BLOCK_SIZE" \
    -Commit=true -Verbose=true -Enabled=true -PrintAllInvariants -timeout 24h >> "$simulationResultFile"

    if grep -E "FAIL   github.com/chain4energy/c4e-chain/app|panic|exit status" "$simulationResultFile"
    then
        simulationResultFileError=$simulationResultFile."ERROR"
        mv "$simulationResultFile" "$simulationResultFileError"
        echo "Error while running simulation"
        curl -F file=@"$simulationResultFileError" -F "initial_comment=$START_TIME" -F channels=C049DAHR884 -H "Authorization: Bearer $SLACK_BOT_BEARER_TOKEN" https://slack.com/api/files.upload
    fi
    done
}

for ((i=1;i<=PARALLEL_SIMULATIONS_NUMBER;i++))
do
	run_simulation i &
done

while true
  do
    git pull
    sleep 100 # 1 hour
done