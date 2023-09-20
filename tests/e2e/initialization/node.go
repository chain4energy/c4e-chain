package initialization

import (
	"cosmossdk.io/math"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"os"
	"path"
	"path/filepath"
	"strings"

	tmconfig "github.com/cometbft/cometbft/config"
	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	tmtypes "github.com/cometbft/cometbft/types"
	sdkcrypto "github.com/cosmos/cosmos-sdk/crypto"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/go-bip39"
	"github.com/spf13/viper"

	c4eapp "github.com/chain4energy/c4e-chain/v2/app"
	"github.com/chain4energy/c4e-chain/v2/tests/e2e/util"
)

const configTomlFile = "config.toml"

type internalNode struct {
	chain        *internalChain
	moniker      string
	mnemonic     string
	keyInfo      keyring.Record
	privateKey   cryptotypes.PrivKey
	consensusKey privval.FilePVKey
	nodeKey      p2p.NodeKey
	peerId       string
	isValidator  bool
}

func newNode(chain *internalChain, nodeConfig *NodeConfig, appState map[string]json.RawMessage) (*internalNode, error) {
	node := &internalNode{
		chain:       chain,
		moniker:     fmt.Sprintf("%s-node-%s", chain.chainMeta.Id, nodeConfig.Name),
		isValidator: nodeConfig.IsValidator,
	}
	// generate genesis files
	if err := node.init(appState); err != nil {
		return nil, err
	}
	// create keys
	if err := node.createKey(ValidatorWalletName); err != nil {
		return nil, err
	}
	if err := node.createNodeKey(); err != nil {
		return nil, err
	}
	if err := node.createConsensusKey(); err != nil {
		return nil, err
	}
	node.createAppConfig(nodeConfig)
	return node, nil
}

func (n *internalNode) configDir() string {
	return fmt.Sprintf("%s/%s", n.chain.chainMeta.configDir(), n.moniker)
}

func (n *internalNode) buildCreateValidatorMsg(amount sdk.Coin) (sdk.Msg, error) {
	description := stakingtypes.NewDescription(n.moniker, "", "", "", "")
	commissionRates := stakingtypes.CommissionRates{
		Rate:          sdk.MustNewDecFromStr("0.1"),
		MaxRate:       sdk.MustNewDecFromStr("0.2"),
		MaxChangeRate: sdk.MustNewDecFromStr("0.01"),
	}

	// get the initial validator min self delegation
	minSelfDelegation, _ := math.NewIntFromString("1")

	valPubKey, err := cryptocodec.FromTmPubKeyInterface(n.consensusKey.PubKey)
	if err != nil {
		return nil, err
	}
	address, _ := n.keyInfo.GetAddress()
	return stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(address),
		valPubKey,
		amount,
		description,
		commissionRates,
		minSelfDelegation,
	)
}

func (n *internalNode) createConfig() error {
	p := path.Join(n.configDir(), "config")
	return os.MkdirAll(p, 0o755)
}

func (n *internalNode) createAppConfig(nodeConfig *NodeConfig) {
	// set application configuration
	appCfgPath := filepath.Join(n.configDir(), "config", "app.toml")

	appConfig := srvconfig.DefaultConfig()
	appConfig.BaseConfig.Pruning = nodeConfig.Pruning
	appConfig.BaseConfig.PruningKeepRecent = nodeConfig.PruningKeepRecent
	appConfig.BaseConfig.PruningInterval = nodeConfig.PruningInterval
	appConfig.API.Enable = true
	appConfig.API.Address = "tcp://0.0.0.0:1317"
	appConfig.MinGasPrices = fmt.Sprintf("%s%s", MinGasPrice, C4eDenom)
	appConfig.StateSync.SnapshotInterval = nodeConfig.SnapshotInterval
	appConfig.StateSync.SnapshotKeepRecent = nodeConfig.SnapshotKeepRecent

	srvconfig.WriteConfigFile(appCfgPath, appConfig)
}

func (n *internalNode) createNodeKey() error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(n.configDir())
	config.Moniker = n.moniker

	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return err
	}

	n.nodeKey = *nodeKey
	return nil
}

func (n *internalNode) createConsensusKey() error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(n.configDir())
	config.Moniker = n.moniker

	pvKeyFile := config.PrivValidatorKeyFile()
	if err := tmos.EnsureDir(filepath.Dir(pvKeyFile), 0o777); err != nil {
		return err
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err := tmos.EnsureDir(filepath.Dir(pvStateFile), 0o777); err != nil {
		return err
	}

	filePV := privval.LoadOrGenFilePV(pvKeyFile, pvStateFile)
	n.consensusKey = filePV.Key

	return nil
}

func (n *internalNode) createKeyFromMnemonic(name, mnemonic string) error {
	encoding := c4eapp.MakeEncodingConfig()
	legacyCdc := codec.NewProtoCodec(encoding.InterfaceRegistry)
	kb, err := keyring.New(keyringAppName, keyring.BackendTest, n.configDir(), nil, legacyCdc)

	if err != nil {
		return err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return err
	}

	info, err := kb.NewAccount(name, mnemonic, "", sdk.FullFundraiserPath, algo)
	if err != nil {
		return err
	}

	privKeyArmor, err := kb.ExportPrivKeyArmor(name, keyringPassphrase)
	if err != nil {
		return err
	}

	privKey, _, err := sdkcrypto.UnarmorDecryptPrivKey(privKeyArmor, keyringPassphrase)
	if err != nil {
		return err
	}

	n.keyInfo = *info
	n.mnemonic = mnemonic
	n.privateKey = privKey

	return nil
}

func (n *internalNode) createKey(name string) error {
	mnemonic, err := n.createMnemonic()
	if err != nil {
		return err
	}

	return n.createKeyFromMnemonic(name, mnemonic)
}

func (n *internalNode) export() *Node {
	address, _ := n.keyInfo.GetAddress()
	pubKey, _ := n.keyInfo.GetPubKey()
	return &Node{
		Name:          n.moniker,
		ConfigDir:     n.configDir(),
		Mnemonic:      n.mnemonic,
		PublicAddress: address.String(),
		PublicKey:     pubKey.Address().String(),
		PeerId:        n.peerId,
		IsValidator:   n.isValidator,
	}
}

func (n *internalNode) getNodeKey() *p2p.NodeKey {
	return &n.nodeKey
}

func (n *internalNode) getGenesisDoc() (*tmtypes.GenesisDoc, error) {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config
	config.SetRoot(n.configDir())

	genFile := config.GenesisFile()
	doc := &tmtypes.GenesisDoc{}

	if _, err := os.Stat(genFile); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else {
		var err error

		doc, err = tmtypes.GenesisDocFromFile(genFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read genesis doc from file: %w", err)
		}
	}

	return doc, nil
}

func (n *internalNode) init(appState map[string]json.RawMessage) error {
	if err := n.createConfig(); err != nil {
		return err
	}

	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(n.configDir())
	config.Moniker = n.moniker

	genDoc, err := n.getGenesisDoc()
	if err != nil {
		return err
	}

	genesisToMarshal := c4eapp.ModuleBasics.DefaultGenesis(util.Cdc)

	if len(appState) > 0 {
		genesisToMarshal = appState
	}

	appStateBytes, err := json.MarshalIndent(genesisToMarshal, "", " ")
	if err != nil {
		return fmt.Errorf("failed to JSON encode app genesis state: %w", err)
	}

	genDoc.ChainID = n.chain.chainMeta.Id
	genDoc.Validators = nil
	genDoc.AppState = appStateBytes

	if err = genutil.ExportGenesisFile(genDoc, config.GenesisFile()); err != nil {
		return fmt.Errorf("failed to export app genesis state: %w", err)
	}

	tmconfig.WriteConfigFile(filepath.Join(config.RootDir, "config", configTomlFile), config)
	return nil
}

func (n *internalNode) createMnemonic() (string, error) {
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

func (n *internalNode) initNodeConfigs(persistentPeers []string) error {
	tmCfgPath := filepath.Join(n.configDir(), "config", configTomlFile)

	vpr := viper.New()
	vpr.SetConfigFile(tmCfgPath)
	if err := vpr.ReadInConfig(); err != nil {
		return err
	}

	valConfig := tmconfig.DefaultConfig()
	if err := vpr.Unmarshal(valConfig); err != nil {
		return err
	}

	valConfig.P2P.ListenAddress = "tcp://0.0.0.0:26656"
	valConfig.P2P.AddrBookStrict = false
	valConfig.P2P.ExternalAddress = fmt.Sprintf("%s:%d", n.moniker, 26656)
	valConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"
	valConfig.StateSync.Enable = false
	valConfig.LogLevel = "info"
	valConfig.P2P.PersistentPeers = strings.Join(persistentPeers, ",")
	valConfig.Storage.DiscardABCIResponses = true

	tmconfig.WriteConfigFile(tmCfgPath, valConfig)
	return nil
}

func (n *internalNode) initStateSyncConfig(trustHeight int64, trustHash string, stateSyncRPCServers []string) error {
	tmCfgPath := filepath.Join(n.configDir(), "config", configTomlFile)

	vpr := viper.New()
	vpr.SetConfigFile(tmCfgPath)
	if err := vpr.ReadInConfig(); err != nil {
		return err
	}

	valConfig := tmconfig.DefaultConfig()
	if err := vpr.Unmarshal(valConfig); err != nil {
		return err
	}

	valConfig.StateSync = tmconfig.DefaultStateSyncConfig()
	valConfig.StateSync.Enable = true
	valConfig.StateSync.TrustHeight = trustHeight
	valConfig.StateSync.TrustHash = trustHash
	valConfig.StateSync.RPCServers = stateSyncRPCServers

	tmconfig.WriteConfigFile(tmCfgPath, valConfig)
	return nil
}

// signMsg returns a signed tx of the provided messages,
// signed by the validator, using 0 fees, a high gas limit, and a common memo.
func (n *internalNode) signMsg(msgs ...sdk.Msg) (*sdktx.Tx, error) {
	txBuilder := util.EncodingConfig.TxConfig.NewTxBuilder()

	if err := txBuilder.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	txBuilder.SetMemo(fmt.Sprintf("%s@%s:26656", n.nodeKey.ID(), n.moniker))
	txBuilder.SetFeeAmount(sdk.NewCoins())
	txBuilder.SetGasLimit(uint64(200000 * len(msgs)))

	signerData := authsigning.SignerData{
		ChainID:       n.chain.chainMeta.Id,
		AccountNumber: 0,
		Sequence:      0,
	}
	pubKey, _ := n.keyInfo.GetPubKey()
	sig := txsigning.SignatureV2{
		PubKey: pubKey,
		Data: &txsigning.SingleSignatureData{
			SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
			Signature: nil,
		},
		Sequence: 0,
	}

	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	bytesToSign, err := util.EncodingConfig.TxConfig.SignModeHandler().GetSignBytes(
		txsigning.SignMode_SIGN_MODE_DIRECT,
		signerData,
		txBuilder.GetTx(),
	)
	if err != nil {
		return nil, err
	}

	sigBytes, err := n.privateKey.Sign(bytesToSign)
	if err != nil {
		return nil, err
	}
	pubKey, _ = n.keyInfo.GetPubKey()
	sig = txsigning.SignatureV2{
		PubKey: pubKey,
		Data: &txsigning.SingleSignatureData{
			SignMode:  txsigning.SignMode_SIGN_MODE_DIRECT,
			Signature: sigBytes,
		},
		Sequence: 0,
	}
	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	signedTx := txBuilder.GetTx()
	bz, err := util.EncodingConfig.TxConfig.TxEncoder()(signedTx)
	if err != nil {
		return nil, err
	}

	return decodeTx(bz)
}
