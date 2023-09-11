package cosmossdk

import (
	"bytes"
	"cosmossdk.io/errors"
	"encoding/hex"
	"fmt"
	appparams "github.com/chain4energy/c4e-chain/v2/app/params"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateTestPubKeys returns a total of numPubKeys public keys in ascending order.
func CreateTestPubKeys(numPubKeys int) []cryptotypes.PubKey {
	var publicKeys []cryptotypes.PubKey
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := 0; i < numPubKeys; i++ {
		privkeySeed := make([]byte, 15)
		r.Read(privkeySeed)
		publicKeys = append(publicKeys, secp256k1.GenPrivKeyFromSecret(privkeySeed).PubKey())
	}

	return publicKeys
}

// NewPubKeyFromHex returns a PubKey from a hex string.
func NewPubKeyFromHex(pk string) (res cryptotypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}
	if len(pkBytes) != ed25519.PubKeySize {
		panic(errors.Wrap(sdkerrors.ErrInvalidPubKey, "invalid pubkey size"))
	}
	return &ed25519.PubKey{Key: pkBytes}
}

func CreateAccounts(accNum int, valAccNum int) (acountsAddresses []sdk.AccAddress, validatorsAddresses []sdk.ValAddress) {
	if accNum > 0 {
		acountsAddresses = CreateIncrementalAccounts(accNum, 0)
	}
	if valAccNum > 0 {
		validatorsAddresses = ConvertAddrsToValAddrs(CreateIncrementalAccounts(valAccNum, accNum))
	}
	return acountsAddresses, validatorsAddresses
}

func CreateIncrementalAccounts(accNum int, genInitNumber int) []sdk.AccAddress {
	var addresses []sdk.AccAddress
	// var buffer bytes.Buffer

	// start at 100 so we can make up to 999 test addresses with valid test addresses
	for i := 100; i < (accNum + 100); i++ {
		hex, bech := CreateRandomAccAddressHexAndBechNoBalance(int64(i + genInitNumber))
		addr, _ := TestAddr(hex, bech)
		addresses = append(addresses, addr)
	}

	return addresses
}

func CreateRandomAccAddressHexAndBechNoBalance(i int64) (hex string, bech string) {
	var buffer bytes.Buffer
	numString := strconv.Itoa(int(i))
	buffer.WriteString("A58856F0FD53BF058B4909A21AEC019107BA6") // base address string
	buffer.WriteString(numString)                               // adding on final two digits to make addresses unique
	res, _ := sdk.AccAddressFromHexUnsafe(buffer.String())
	return buffer.String(), res.String()
}

func CreateRandomAccAddress() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	_, bech := CreateRandomAccAddressHexAndBechNoBalance(r.Int63())
	return bech
}

func ConvertAddrsToValAddrs(addrs []sdk.AccAddress) []sdk.ValAddress {
	valAddrs := make([]sdk.ValAddress, len(addrs))

	for i, addr := range addrs {
		valAddrs[i] = sdk.ValAddress(addr)
	}

	return valAddrs
}

func TestAddr(addr string, bech string) (sdk.AccAddress, error) {
	res, err := sdk.AccAddressFromHexUnsafe(addr)
	if err != nil {
		return nil, err
	}
	bechexpected := res.String()
	if bech != bechexpected {
		return nil, fmt.Errorf("bech encoding doesn't match reference")
	}

	bechres, err := AccAddressFromBech32(bech)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(bechres, res) {
		return nil, err
	}

	return res, nil
}

// AccAddressFromBech32 creates an AccAddress from a Bech32 string.
func AccAddressFromBech32(address string) (addr sdk.AccAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return sdk.AccAddress{}, fmt.Errorf("empty address string is not allowed")
	}

	bech32PrefixAccAddr := appparams.Bech32PrefixAccAddr

	bz, err := sdk.GetFromBech32(address, bech32PrefixAccAddr)
	if err != nil {
		return nil, err
	}

	err = sdk.VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return bz, nil
}
