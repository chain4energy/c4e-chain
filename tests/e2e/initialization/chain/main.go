package main_package

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chain4energy/c4e-chain/tests/e2e/initialization"
	"os"
	"time"
)

func main() {
	var (
		valConfig    []*initialization.NodeConfig
		dataDir      string
		chainId      string
		config       string
		votingPeriod time.Duration
		forkHeight   int
	)

	flag.StringVar(&dataDir, "data-dir", "", "chain data directory")
	flag.StringVar(&chainId, "chain-id", "", "chain ID")
	flag.StringVar(&config, "config", "", "serialized config")
	flag.DurationVar(&votingPeriod, "voting-period", 30000000000, "voting period")
	flag.IntVar(&forkHeight, "fork-height", 0, "fork height")

	flag.Parse()

	err := json.Unmarshal([]byte(config), &valConfig)
	if err != nil {
		panic(err)
	}

	if len(dataDir) == 0 {
		panic("data-dir is required")
	}

	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		panic(err)
	}

	createdChain, err := initialization.InitChain(chainId, dataDir, valConfig, votingPeriod, forkHeight)
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(createdChain)
	fileName := fmt.Sprintf("%v/%v-encode", dataDir, chainId)
	if err = os.WriteFile(fileName, b, 0o777); err != nil {
		panic(err)
	}
}
