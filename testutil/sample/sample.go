package sample

import (
	"cosmossdk.io/math"
	"fmt"
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

// Coins returns a sample sdk.Coins
func Coins() sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(testenv.DefaultTestDenom, math.NewInt(10000)))
}

func PrepareDifferentDenomCoins(n int, amount math.Int) sdk.Coins {
	var coins sdk.Coins
	for i := 0; i < n; i++ {
		coins = coins.Add(sdk.NewCoin(fmt.Sprintf("uc4e%d", i), amount))
	}
	return coins
}
