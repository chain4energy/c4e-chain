package helpers

import (
	cfedistributortypes "github.com/chain4energy/c4e-chain/x/cfedistributor/types"
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
		Destinations: cfedistributortypes.Destinations{
			PrimaryShare: cfedistributortypes.Account{
				Id:   "c4e_distributor",
				Type: cfedistributortypes.MAIN,
			},
			BurnShare: sdk.ZeroDec(),
		},
		Sources: []*cfedistributortypes.Account{
			{
				Id:   "fee_collector",
				Type: cfedistributortypes.MODULE_ACCOUNT,
			},
		},
	},
	{
		Name: "inflation_and_fee_distributor",
		Destinations: cfedistributortypes.Destinations{
			PrimaryShare: cfedistributortypes.Account{
				Id:   cfedistributortypes.ValidatorsRewardsCollector,
				Type: cfedistributortypes.MODULE_ACCOUNT,
			},
			BurnShare: sdk.ZeroDec(),
			Shares: []*cfedistributortypes.DestinationShare{
				{
					Name:  "development_fund",
					Share: sdk.MustNewDecFromStr("0.05"),
					Destination: cfedistributortypes.Account{
						Id:   "c4e10ep2sxpf2kj6j26w7f4uuafedkuf9sf9xqq3sl",
						Type: cfedistributortypes.BASE_ACCOUNT,
					},
				},
				{
					Name:  "usage_incentives",
					Share: sdk.MustNewDecFromStr("0.35"),
					Destination: cfedistributortypes.Account{
						Id:   "usage_incentives_collector",
						Type: cfedistributortypes.INTERNAL_ACCOUNT,
					},
				},
			},
		},
		Sources: []*cfedistributortypes.Account{
			{
				Id:   "c4e_distributor",
				Type: cfedistributortypes.MAIN,
			},
		},
	},
	{
		Name: "usage_incentives_distributor",
		Destinations: cfedistributortypes.Destinations{
			PrimaryShare: cfedistributortypes.Account{
				Id:   "c4e1q5vgy0r3w9q4cclucr2kl8nwmfe2mgr6g0jlph",
				Type: cfedistributortypes.BASE_ACCOUNT,
			},
			BurnShare: sdk.ZeroDec(),
			Shares: []*cfedistributortypes.DestinationShare{
				{
					Name:  "green_energy_booster",
					Share: sdk.MustNewDecFromStr("0.34"),
					Destination: cfedistributortypes.Account{
						Id:   "green_energy_booster_collector",
						Type: cfedistributortypes.MODULE_ACCOUNT,
					},
				},
				{
					Name:  "governance_booster",
					Share: sdk.MustNewDecFromStr("0.33"),
					Destination: cfedistributortypes.Account{
						Id:   "governance_booster_collector",
						Type: cfedistributortypes.MODULE_ACCOUNT,
					},
				},
			},
		},
		Sources: []*cfedistributortypes.Account{
			{
				Id:   "usage_incentives_collector",
				Type: cfedistributortypes.INTERNAL_ACCOUNT,
			},
		},
	},
}

var minterConfigLongEndTime = time.Now().Add(16 * oneYearDuration).UTC()
var MainnetMinterConfigLong = cfemintertypes.MinterConfig{
	StartTime: time.Now().UTC(),
	Minters: []*cfemintertypes.Minter{
		{
			SequenceId: 1,
			Type:       cfemintertypes.EXPONENTIAL_STEP_MINTING,
			ExponentialStepMinting: &cfemintertypes.ExponentialStepMinting{
				Amount:           sdk.NewInt(160000000000000),
				AmountMultiplier: sdk.MustNewDecFromStr("0.5"),
				StepDuration:     oneYearDuration * 4,
			},
			EndTime: &minterConfigLongEndTime,
		},
		{
			SequenceId: 2,
			Type:       cfemintertypes.NO_MINTING,
		},
	},
}
var timeNow = time.Now()
var minterConfigShortEndTime = timeNow.Add(time.Minute*4 + time.Second).UTC()
var MainnetMinterConfigShort = cfemintertypes.MinterConfig{
	StartTime: time.Now().UTC(),
	Minters: []*cfemintertypes.Minter{
		{
			SequenceId: 1,
			Type:       cfemintertypes.EXPONENTIAL_STEP_MINTING,
			ExponentialStepMinting: &cfemintertypes.ExponentialStepMinting{
				//	160000000000000 + 160000000000000/2 + 160000000000000/4 + 160000000000000/8 = 300000000000000

				Amount:           sdk.NewInt(160000000000000),
				AmountMultiplier: sdk.MustNewDecFromStr("0.5"),
				StepDuration:     time.Second * 60,
			},
			EndTime: &minterConfigShortEndTime,
		},
		{
			SequenceId: 2,
			Type:       cfemintertypes.NO_MINTING,
		},
	},
}
