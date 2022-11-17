package cfeairdrop

import (
	"testing"
	cfeairdropmodulekeeper "github.com/chain4energy/c4e-chain/x/cfeairdrop/keeper"
)

type C4eAirdropKeeperUtils struct {
	t                           *testing.T
	helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper
}

func NewC4eAirdropKeeperUtils(t *testing.T, helpeCfeairdropmodulekeeper *cfeairdropmodulekeeper.Keeper) C4eAirdropKeeperUtils {
	return C4eAirdropKeeperUtils{t: t, helpeCfeairdropmodulekeeper: helpeCfeairdropmodulekeeper}
}

