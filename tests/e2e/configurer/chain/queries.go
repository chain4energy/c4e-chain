package chain

import (
	"context"
	"cosmossdk.io/math"
	"encoding/json"
	"fmt"
	cfeclaimmoduletypes "github.com/chain4energy/c4e-chain/x/cfeclaim/types"
	cfedistributormoduletypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	cfeevtypes "github.com/chain4energy/c4e-chain/x/cfeev/types"
	cfemintermoduletypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"io"
	"net/http"
	"time"

	cfevestingmoduletypes "github.com/chain4energy/c4e-chain/x/cfevesting/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/stretchr/testify/require"
	tmabcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/chain4energy/c4e-chain/tests/e2e/util"
)

const outputJsonFlag = "--output=json"

func (n *NodeConfig) QueryGRPCGateway(path string, parameters ...string) ([]byte, error) {
	if len(parameters)%2 != 0 {
		return nil, fmt.Errorf("invalid number of parameters, must follow the format of key + value")
	}

	// add the URL for the given validator ID, and pre-pend to to path.
	hostPort, err := n.containerManager.GetHostPort(n.Name, "1317/tcp")
	require.NoError(n.t, err)
	endpoint := fmt.Sprintf("http://%s", hostPort)
	fullQueryPath := fmt.Sprintf("%s/%s", endpoint, path)

	var resp *http.Response
	require.Eventually(n.t, func() bool {
		req, err := http.NewRequest("GET", fullQueryPath, nil)
		if err != nil {
			return false
		}

		if len(parameters) > 0 {
			q := req.URL.Query()
			for i := 0; i < len(parameters); i += 2 {
				q.Add(parameters[i], parameters[i+1])
			}
			req.URL.RawQuery = q.Encode()
		}

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			n.t.Logf("error while executing HTTP request: %s", err.Error())
			return false
		}

		return resp.StatusCode != http.StatusServiceUnavailable
	}, time.Minute, time.Millisecond*10, "failed to execute HTTP request")

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bz))
	}
	return bz, nil
}

// QueryBalances returns balances at the address.
func (n *NodeConfig) QueryBalances(address string) (sdk.Coins, error) {
	path := fmt.Sprintf("cosmos/bank/v1beta1/balances/%s", address)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var balancesResp banktypes.QueryAllBalancesResponse
	if err := util.Cdc.UnmarshalJSON(bz, &balancesResp); err != nil {
		return sdk.Coins{}, err
	}
	return balancesResp.GetBalances(), nil
}

func (n *NodeConfig) QuerySupplyOf(denom string) (math.Int, error) {
	path := fmt.Sprintf("cosmos/bank/v1beta1/supply/by_denom?denom=%s", denom)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var supplyResp banktypes.QuerySupplyOfResponse
	if err := util.Cdc.UnmarshalJSON(bz, &supplyResp); err != nil {
		return math.NewInt(0), err
	}
	return supplyResp.Amount.Amount, nil
}

func (n *NodeConfig) QueryPropTally(proposalNumber int) (math.Int, math.Int, math.Int, math.Int, error) {
	path := fmt.Sprintf("cosmos/gov/v1beta1/proposals/%d/tally", proposalNumber)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var balancesResp govv1.QueryTallyResultResponse
	if err := util.Cdc.UnmarshalJSON(bz, &balancesResp); err != nil {
		return math.ZeroInt(), math.ZeroInt(), math.ZeroInt(), math.ZeroInt(), err
	}
	noTotal, _ := math.NewIntFromString(balancesResp.Tally.NoCount)
	yesTotal, _ := math.NewIntFromString(balancesResp.Tally.YesCount)
	noWithVetoTotal, _ := math.NewIntFromString(balancesResp.Tally.NoWithVetoCount)
	abstainTotal, _ := math.NewIntFromString(balancesResp.Tally.AbstainCount)

	return noTotal, yesTotal, noWithVetoTotal, abstainTotal, nil
}

func (n *NodeConfig) QueryPropStatus(proposalNumber int) (string, error) {
	path := fmt.Sprintf("cosmos/gov/v1/proposals/%d", proposalNumber)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var propResp govv1.QueryProposalResponse
	if err := util.Cdc.UnmarshalJSON(bz, &propResp); err != nil {
		return "", err
	}
	proposalStatus := propResp.Proposal.Status

	return proposalStatus.String(), nil
}

// QueryHashFromBlock gets block hash at a specific height. Otherwise, error.
func (n *NodeConfig) QueryHashFromBlock(height int64) (string, error) {
	block, err := n.rpcClient.Block(context.Background(), &height)
	if err != nil {
		return "", err
	}
	return block.BlockID.Hash.String(), nil
}

// QueryCurrentHeight returns the current block height of the node or error.
func (n *NodeConfig) QueryCurrentHeight() (int64, error) {
	status, err := n.rpcClient.Status(context.Background())
	if err != nil {
		return 0, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}

// QueryLatestBlockTime returns the latest block time.
func (n *NodeConfig) QueryLatestBlockTime() time.Time {
	status, err := n.rpcClient.Status(context.Background())
	require.NoError(n.t, err)
	return status.SyncInfo.LatestBlockTime
}

// QueryListSnapshots gets all snapshots currently created for a node.
func (n *NodeConfig) QueryListSnapshots() ([]*tmabcitypes.Snapshot, error) {
	abciResponse, err := n.rpcClient.ABCIQuery(context.Background(), "/app/snapshots", nil)
	if err != nil {
		return nil, err
	}

	var listSnapshots tmabcitypes.ResponseListSnapshots
	if err := json.Unmarshal(abciResponse.Response.Value, &listSnapshots); err != nil {
		return nil, err
	}

	return listSnapshots.Snapshots, nil
}

func (n *NodeConfig) QueryVestingPoolsInfo(address string) []*cfevestingmoduletypes.VestingPoolInfo {
	path := "/c4e/vesting/v1beta1/vesting_pools/" + address

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfevestingmoduletypes.QueryVestingPoolsResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.VestingPools
}

func (n *NodeConfig) QueryVestingPoolsNotFound(address string) {
	path := "/c4e/vesting/v1beta1/vesting_pools/" + address

	_, err := n.QueryGRPCGateway(path)
	require.Error(n.t, err)
	require.EqualError(n.t, err, "unexpected status code: 404, body: {\n  \"code\": 5,\n  \"message\": \"vesting pools not found\",\n  \"details\": [\n  ]\n}")

}

func (n *NodeConfig) QueryVestingTypes() []cfevestingmoduletypes.GenesisVestingType {
	path := "/c4e/vesting/v1beta1/vesting_type"

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfevestingmoduletypes.QueryVestingTypeResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.VestingTypes
}
func (n *NodeConfig) QueryFailedProposal(proposalNumber int) {
	path := fmt.Sprintf("cosmos/gov/v1beta1/proposals/%d", proposalNumber)
	_, err := n.QueryGRPCGateway(path)
	fmt.Println(err)
	require.Error(n.t, err)
}

func (n *NodeConfig) QueryAccount(address string) authtypes.AccountI {
	path := "/cosmos/auth/v1beta1/accounts/" + address

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response authtypes.QueryAccountResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	acc := response.Account
	if acc == nil {
		return nil
	}
	return acc.GetCachedValue().(authtypes.AccountI)
}

func (n *NodeConfig) QueryAccountNotFound(address string) {
	path := "/cosmos/auth/v1beta1/accounts/" + address

	_, err := n.QueryGRPCGateway(path)
	require.Error(n.t, err)
	require.EqualError(n.t, err, "unexpected status code: 404, body: {\n  \"code\": 5,\n  \"message\": \"account "+address+" not found\",\n  \"details\": [\n  ]\n}")

}

func (n *NodeConfig) QueryCampaign(campaignId string) cfeclaimmoduletypes.Campaign {
	path := "/c4e/claim/v1beta1/campaign/" + campaignId

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeclaimmoduletypes.QueryCampaignResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.Campaign
}

func (n *NodeConfig) QueryCampaigns() []cfeclaimmoduletypes.Campaign {
	path := "/c4e/claim/v1beta1/campaigns"

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeclaimmoduletypes.QueryCampaignsResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.Campaigns
}

func (n *NodeConfig) QueryLastCampaignsId() uint64 {
	path := "/c4e/claim/v1beta1/campaigns"

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeclaimmoduletypes.QueryCampaignsResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.Campaigns[len(response.Campaigns)-1].Id
}

func (n *NodeConfig) QueryCampaignMission(campaignId, missionId string) cfeclaimmoduletypes.Mission {
	path := "/c4e/claim/v1beta1/mission/" + campaignId + "/" + missionId

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeclaimmoduletypes.QueryMissionResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.Mission
}

func (n *NodeConfig) QueryUserEntry(address string) cfeclaimmoduletypes.UserEntry {
	path := "/c4e/claim/v1beta1/user_entry/" + address

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeclaimmoduletypes.QueryUserEntryResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.UserEntry
}

func (n *NodeConfig) QueryUserEntries() []cfeclaimmoduletypes.UserEntry {
	path := "/c4e/claim/v1beta1/users_entries"

	bz, err := n.QueryGRPCGateway(path, "pagination.limit", "1000000000000")
	require.NoError(n.t, err)

	var response cfeclaimmoduletypes.QueryUsersEntriesResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.UsersEntries
}

func (n *NodeConfig) QueryCfevestingParams(moduleParams *cfevestingmoduletypes.QueryParamsResponse) {
	cmd := []string{"c4ed", "query", "cfevesting", "params", outputJsonFlag}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	err = json.Unmarshal(out.Bytes(), &moduleParams)
	require.NoError(n.t, err)
}

func (n *NodeConfig) QueryCfeminterParams(moduleParams *cfemintermoduletypes.QueryParamsResponse) {
	cmd := []string{"c4ed", "query", "cfeminter", "params", outputJsonFlag}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	err = util.Cdc.UnmarshalJSON(out.Bytes(), moduleParams)
	require.NoError(n.t, err)
}

func (n *NodeConfig) QueryCfedistributorParams(moduleParams *cfedistributormoduletypes.QueryParamsResponse) {
	cmd := []string{"c4ed", "query", "cfedistributor", "params", outputJsonFlag}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	err = json.Unmarshal(out.Bytes(), &moduleParams)
	require.NoError(n.t, err)
}

func (n *NodeConfig) QueryCfeEvParams(moduleParams *cfeevtypes.QueryParamsResponse) {
	cmd := []string{"c4ed", "query", "cfeev", "params", outputJsonFlag}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	err = json.Unmarshal(out.Bytes(), &moduleParams)
	require.NoError(n.t, err)
}

func (n *NodeConfig) QueryFeegrant(granter, grantee string, feegrantResponse *feegrant.QueryAllowanceRequest) {
	cmd := []string{"c4ed", "query", "feegrant", "grant", granter, grantee, outputJsonFlag}

	out, _, err := n.containerManager.ExecCmd(n.t, n.Name, cmd, "")
	require.NoError(n.t, err)
	err = json.Unmarshal(out.Bytes(), &feegrantResponse)
	require.NoError(n.t, err)
}

func (n *NodeConfig) QueryEnergyTransferOffers(owner string) []cfeevtypes.EnergyTransferOffer {
	path := "/c4e/ev/v1beta1/energy_transfer_offers/" + owner

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeevtypes.QueryEnergyTransferOffersResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.EnergyTransferOffers
}

func (n *NodeConfig) QueryEnergyTransfers(owner string) []cfeevtypes.EnergyTransfer {
	path := "/c4e/ev/v1beta1/energy_transfers/" + owner

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeevtypes.QueryEnergyTransfersResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.EnergyTransfers
}

func (n *NodeConfig) QueryEnergyTransferOffer(id string) cfeevtypes.EnergyTransferOffer {
	path := "/c4e/ev/v1beta1/energy_transfer_offer/" + id

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeevtypes.QueryEnergyTransferOfferResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.EnergyTransferOffer
}

func (n *NodeConfig) QueryEnergyTransfer(id string) cfeevtypes.EnergyTransfer {
	path := "/c4e/ev/v1beta1/energy_transfer/" + id

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response cfeevtypes.QueryEnergyTransferResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.EnergyTransfer
}

func (n *NodeConfig) QueryEnergyTransferOfferNotFound(id string) {
	path := "/c4e/ev/v1beta1/energy_transfer_offer/" + id

	_, err := n.QueryGRPCGateway(path)
	require.Error(n.t, err)
}

func (n *NodeConfig) QueryEnergyTransferNotFound(id string) {
	path := "/c4e/ev/v1beta1/energy_transfer/" + id
	_, err := n.QueryGRPCGateway(path)
	require.Error(n.t, err)
}

func (n *NodeConfig) QueryPropStatusTimed(proposalNumber int, desiredStatus string, totalTime chan time.Duration) {
	start := time.Now()
	require.Eventually(
		n.t,
		func() bool {
			status, err := n.QueryPropStatus(proposalNumber)
			if err != nil {
				return false
			}

			return status == desiredStatus
		},
		1*time.Minute,
		10*time.Millisecond,
		"C4e node failed to retrieve prop tally",
	)
	elapsed := time.Since(start)
	totalTime <- elapsed
}
