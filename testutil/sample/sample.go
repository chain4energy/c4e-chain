package sample

import (
	"cosmossdk.io/math"
	testenv "github.com/chain4energy/c4e-chain/testutil/env"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// AccAddress returns a sample account address
func Coins() sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000)))
}
