package v120

import (
	"time"

	cfeupgradetypes "github.com/chain4energy/c4e-chain/v2/app/upgrades"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

const (
	Account1 = "c4e1dsm96gwcv35m4rqd93pzcsztpkrqe0ev7getj8"
	Account2 = "c4e10wjj2qmn4zjg2sdxq9mfyj5v4yukwyhzdtf2zp"
	Account3 = "c4e1zrd0783g8qa5659apw5tpuqmz2ct6j20t4ymx3"
	Account4 = "c4e1y8lndj6jz5z93g4xd05nmwyc3wtn39dfgfx7r7"
)

func ModifyVestingAccountsState(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers) error {
	if err := upgradeVestingAccounnt(ctx, appKeepers, Account1); err != nil {
		return err
	}
	if err := upgradeVestingAccounnt(ctx, appKeepers, Account2); err != nil {
		return err
	}
	if err := upgradeVestingAccounnt(ctx, appKeepers, Account3); err != nil {
		return err
	}
	return upgradeVestingAccounnt(ctx, appKeepers, Account4)
}

func upgradeVestingAccounnt(ctx sdk.Context, appKeepers cfeupgradetypes.AppKeepers, address string) error {
	accAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}
	account := appKeepers.GetAccountKeeper().GetAccount(ctx, accAddr)
	if account == nil {
		ctx.Logger().Info("account does not exist", "address", address)
		return nil
	}
	vestingAccount, ok := account.(*vestingtypes.ContinuousVestingAccount)
	if !ok {
		ctx.Logger().Info("account is not ContinuousVestingAccount", "address", address)
		return nil
	}
	startTime := time.Unix(vestingAccount.StartTime, 0)
	endTime := time.Unix(vestingAccount.EndTime, 0)
	vestingAccount.StartTime = startTime.AddDate(1, 0, 0).Unix()
	vestingAccount.EndTime = endTime.AddDate(1, 0, 0).Unix()
	appKeepers.GetAccountKeeper().SetAccount(ctx, vestingAccount)
	return nil
}
