package helpers

import (
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
	v2 "github.com/chain4energy/c4e-chain/x/cfeminter/migrations/v2"
	cfemintertypes "github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"time"
)

func GetModuleAccountAddress(moduleName string) string {
	return authtypes.NewModuleAddress(moduleName).String()
}

var oneYearDuration = time.Hour * 365

var MainnetSubdistributors = []cfedistributortypes.SubDistributor{
	{
		Name: "tx_fee_distributor",
		Sources: []*cfedistributortypes.Account{
			{
				Id:   "fee_collector",
				Type: cfedistributortypes.ModuleAccount,
			},
		},
		Destinations: cfedistributortypes.Destinations{
			PrimaryShare: cfedistributortypes.Account{
				Id:   "c4e_distributor",
				Type: cfedistributortypes.Main,
			},
			BurnShare: sdk.ZeroDec(),
		},
	},
	{
		Name: "inflation_and_fee_distributor",
		Sources: []*cfedistributortypes.Account{
			{
				Id:   "c4e_distributor",
				Type: cfedistributortypes.Main,
			},
		},
		Destinations: cfedistributortypes.Destinations{
			PrimaryShare: cfedistributortypes.Account{
				Id:   cfedistributortypes.ValidatorsRewardsCollector,
				Type: cfedistributortypes.ModuleAccount,
			},
			BurnShare: sdk.ZeroDec(),
			Shares: []*cfedistributortypes.DestinationShare{
				{
					Name:  "usage_incentives",
					Share: sdk.MustNewDecFromStr("0.3"),
					Destination: cfedistributortypes.Account{
						Id:   "usage_incentives_collector",
						Type: cfedistributortypes.InternalAccount,
					},
				},
			},
		},
	},
	{
		Name: "usage_incentives_distributor",
		Sources: []*cfedistributortypes.Account{
			{
				Id:   "usage_incentives_collector",
				Type: cfedistributortypes.InternalAccount,
			},
		},
		Destinations: cfedistributortypes.Destinations{
			PrimaryShare: cfedistributortypes.Account{
				Id:   "green_energy_booster_collector",
				Type: cfedistributortypes.ModuleAccount,
			},
			BurnShare: sdk.ZeroDec(),
			Shares: []*cfedistributortypes.DestinationShare{
				{
					Name:  "governance_booster",
					Share: sdk.MustNewDecFromStr("0.33"),
					Destination: cfedistributortypes.Account{
						Id:   "governance_booster_collector",
						Type: cfedistributortypes.ModuleAccount,
					},
				},
			},
		},
	},
}

var minterConfigLongEndTime = time.Now().Add(16 * oneYearDuration).UTC()
var MainnetMinterConfigLong = cfemintertypes.MinterConfig{
	StartTime: time.Now().UTC(),
	Minters: []*cfemintertypes.LegacyMinter{
		{
			SequenceId: 1,
			Type:       v2.ExponentialStepMintingType,
			ExponentialStepMinting: &cfemintertypes.ExponentialStepMinting{
				Amount:           sdk.NewInt(32000000000000),
				AmountMultiplier: sdk.MustNewDecFromStr("0.5"),
				StepDuration:     oneYearDuration * 4,
			},
			EndTime: &minterConfigLongEndTime,
		},
	},
}
var timeNow = time.Now()
var minterConfigShortEndTime = timeNow.Add(time.Second * 48).UTC()
var MainnetMinterConfigShort = cfemintertypes.MinterConfig{
	StartTime: time.Now().UTC(),
	Minters: []*cfemintertypes.LegacyMinter{
		{
			SequenceId: 1,
			Type:       v2.ExponentialStepMintingType,
			ExponentialStepMinting: &cfemintertypes.ExponentialStepMinting{
				Amount:           sdk.NewInt(160000000000000),
				AmountMultiplier: sdk.MustNewDecFromStr("0.5"),
				StepDuration:     time.Second * 12,
			},
			EndTime: &minterConfigShortEndTime,
		},
		{
			SequenceId: 2,
			Type:       v2.NoMintingType,
		},
	},
}
