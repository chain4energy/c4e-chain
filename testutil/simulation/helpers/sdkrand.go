package helpers

import (
	"cosmossdk.io/math"
	"errors"
	"math/big"
	"math/rand"
	"time"
	"unsafe"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// shamelessly copied from cosmos sdk github
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang#31832326
func RandStringOfLengthCustomSeed(r *rand.Rand, n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, r.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func RandStringOfLength(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return RandStringOfLengthCustomSeed(r, n)
}

func RandPositiveInt(r *rand.Rand, max math.Int) (math.Int, error) {
	if !max.GTE(sdk.OneInt()) {
		return math.Int{}, errors.New("max too small")
	}

	max = max.Sub(sdk.OneInt())

	return math.NewIntFromBigInt(new(big.Int).Rand(r, max.BigInt())).Add(sdk.OneInt()), nil
}

func RandomAmount(r *rand.Rand, max math.Int) math.Int {
	randInt := big.NewInt(0)

	switch r.Intn(10) {
	case 0:
		// randInt = big.NewInt(0)
	case 1:
		randInt = max.BigInt()
	default: // NOTE: there are 10 total cases.
		randInt = big.NewInt(0).Rand(r, max.BigInt()) // up to max - 1
	}

	return math.NewIntFromBigInt(randInt)
}

func RandomDecAmount(r *rand.Rand, max sdk.Dec) sdk.Dec {
	randInt := big.NewInt(0)
	switch r.Intn(10) {
	case 1:
		randInt = max.BigInt()
	default:
		randInt = big.NewInt(0).Rand(r, max.BigInt())
	}

	return sdk.NewDecFromBigIntWithPrec(randInt, sdk.Precision)
}

func RandTimestamp(r *rand.Rand) time.Time {
	unixTime := r.Int63n(253373529600)
	return time.Unix(unixTime, 0)
}

func RandIntBetween(r *rand.Rand, min, max int) int {
	return r.Intn(max-min) + min
}

func RandDurationBetween(r *rand.Rand, min, max int) time.Duration {
	return time.Duration(r.Intn(max-min)+min) * time.Second
}

func RandIntBetweenWith0(r *rand.Rand, min, max int64) int64 {
	diff := int(max - min)
	if diff <= 0 {
		return min
	}
	return int64(r.Intn(diff)) + min
}

func RandomInt(r *rand.Rand, max int) int64 {
	return int64(r.Intn(max))
}

func RandomBool(r *rand.Rand) bool {
	return r.Intn(2) == 0
}

func RandIntWith0(r *rand.Rand, max int) int {
	if max == 0 {
		return 0
	}
	return r.Intn(max)
}

func RandSubsetCoins(r *rand.Rand, coins sdk.Coins) sdk.Coins {
	if len(coins) == 0 {
		return sdk.Coins{}
	}
	denomIdx := r.Intn(len(coins))
	coin := coins[denomIdx]
	amt, err := RandPositiveInt(r, coin.Amount)
	if err != nil {
		return sdk.Coins{}
	}

	subset := sdk.Coins{sdk.NewCoin(coin.Denom, amt)}

	for i, c := range coins {
		if i == denomIdx {
			continue
		}
		if r.Intn(2) == 0 && len(coins) != 1 {
			continue
		}

		amt, err := RandPositiveInt(r, c.Amount)
		if err != nil {
			continue
		}

		subset = append(subset, sdk.NewCoin(c.Denom, amt))
	}

	return subset.Sort()
}

type multiSource []rand.Source

func (ms multiSource) Int63() (r int64) {
	for _, source := range ms {
		r ^= source.Int63()
	}

	return r
}

func (ms multiSource) Seed(seed int64) {
	panic("multiSource Seed should not be called")
}
