#!/bin/bash
PARALLEL_SIMULATIONS_NUMER=1

run_simulation(){
  while true
  do
    echo "Running simulation $i"
    start_time=$(date '+%Y-%m-%d_%H:%M:%S')
    SEED=$(shuf -i 0-6500 -n 1)
    SIM_NUM_BLOCKS=$(shuf -i 200-1000 -n 1)
    SIM_BLOCK_SIZE=$(shuf -i 200-400 -n 1)
    SIMAPP=./app
    resultFile=.simulation-result.$i.$start_time
    echo "Running application benchmark for numBlocks=$SIM_NUM_BLOCKS, blockSize=$SIM_BLOCK_SIZE, seed=$SEED current time=$start_time. This may take awhile!" > $resultFile
    go test -run=^$ -bench ^BenchmarkSimulation $SIMAPP -Seed=$SEED -NumBlocks=$SIM_NUM_BLOCKS -BlockSize=$SIM_BLOCK_SIZE \
    -Commit=true -Verbose=true -Enabled=true >> $resultFile
    if grep -E "FAIL   github.com/chain4energy/c4e-chain/app|panic|exit status" $resultFile
    then
        echo "Error"
        #Error curl
    fi
    curl -F file=@.simulation-result.start_time -F "initial_comment=start_time" -F channels=C049DAHR884 -H "Authorization: Bearer xoxb-89841793824-4319406644228-odhbQhFoRpzxJxDbfooS23u8" https://slack.com/api/files.upload
  done
}

for ((i=1;i<=$PARALLEL_SIMULATIONS_NUMER;i++))
do
	run_simulation $i &
done

while true
  do
    git pull
    sleep 1000
done