package config

import (
	"cosmossdk.io/math"
	appparams "github.com/chain4energy/c4e-chain/v2/app/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

const (
	// if not skipping upgrade, how many blocks we allow for fork to run pre upgrade state creation
	ForkHeightPreUpgradeOffset int64 = 60
	// estimated number of blocks it takes to submit for a proposal
	PropSubmitBlocks float32 = 5
	// estimated number of blocks it takes to deposit for a proposal
	PropDepositBlocks float32 = 5
	// number of blocks it takes to vote for a single validator to vote for a proposal
	PropVoteBlocks float32 = 1.2
	// number of blocks used as a calculation buffer
	PropBufferBlocks float32 = 6
	// number of blocks used as a calculation buffer after an upgrade
	UpgradeBufferBlocks int64 = 1
	// max retries for json unmarshalling
	MaxRetries = 60
)

var (
	MinDepositValue   = govv1.DefaultMinDepositTokens
	InitialMinDeposit = MinDepositValue.Int64() / 4
	BaseBalance       = sdk.NewCoin(appparams.MicroC4eUnit, math.NewInt(10000000000))
)
