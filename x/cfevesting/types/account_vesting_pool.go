package types

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/v2/types/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (av AccountVestingPools) Validate() error {
	vs := av.VestingPools
	_, err := sdk.AccAddressFromBech32(av.Owner)
	if err != nil {
		return fmt.Errorf("account vesting pools address: %s: %s", av.Owner, err.Error())
	}
	for _, v := range vs {
		if err = v.Validate(av.Owner); err != nil {
			return err
		}
		err = av.checkDuplications(vs, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (av AccountVestingPools) checkDuplications(vs []*VestingPool, v *VestingPool) error {
	numOfNames := 0
	for _, vCheck := range vs {
		if v.Name == vCheck.Name {
			numOfNames++
		}
		if numOfNames > 1 {
			return fmt.Errorf("vesting pool with name: %s defined more than once for account: %s", v.Name, av.Owner)
		}
	}

	return nil
}

func (av AccountVestingPools) ValidateAgainstVestingTypes(vestingTypes []GenesisVestingType) error {
	vs := av.VestingPools
	for _, v := range vs {
		found := false
		for _, vtCheck := range vestingTypes {
			if v.VestingType == vtCheck.Name {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("vesting pool with name: %s defined for account: %s - vesting type not found: %s", v.Name, av.Owner, v.VestingType)
		}
	}
	return nil
}

func (pool *VestingPool) GetLockedNotReserved() math.Int {
	return pool.InitiallyLocked.Sub(pool.Sent).Sub(pool.Withdrawn).Sub(pool.GetAllReserved())
}

func (pool *VestingPool) GetCurrentlyLocked() math.Int {
	return pool.InitiallyLocked.Sub(pool.Sent).Sub(pool.Withdrawn)
}

func (pool *VestingPool) GetReserved(reservationId uint64) math.Int {
	return pool.GetReservation(reservationId).Amount
}

func (pool *VestingPool) GetAllReserved() math.Int {
	amountSum := math.ZeroInt()
	for _, reservation := range pool.Reservations {
		amountSum = amountSum.Add(reservation.Amount)
	}
	return amountSum
}

func (pool *VestingPool) Validate(accountAdd string) error {
	if len(pool.Name) == 0 {
		return fmt.Errorf("vesting pool defined for account: %s has no name", accountAdd)
	}
	if pool.InitiallyLocked.IsNegative() {
		return fmt.Errorf("vesting pool with name: %s defined for account: %s has InitiallyLocked value negative %s", pool.Name, accountAdd, pool.InitiallyLocked)
	}
	if pool.Withdrawn.IsNegative() {
		return fmt.Errorf("vesting pool with name: %s defined for account: %s has Withdrawn value negative %s", pool.Name, accountAdd, pool.Withdrawn)
	}
	if pool.Sent.IsNegative() {
		return fmt.Errorf("vesting pool with name: %s defined for account: %s has Sent value negative %s", pool.Name, accountAdd, pool.Sent)
	}
	if pool.GetCurrentlyLocked().IsNegative() {
		return fmt.Errorf("vesting pool with name: %s defined for account: %s has InitiallyLocked (%s) < Withdrawn (%s) + Sent (%s)",
			pool.Name, accountAdd, pool.InitiallyLocked, pool.Withdrawn, pool.Sent)
	}
	return nil
}

type AccountVestingPoolsList []AccountVestingPools

func (avpl AccountVestingPoolsList) GetGenesisAmount() math.Int {
	result := math.ZeroInt()
	for _, avp := range avpl {
		for _, vp := range avp.VestingPools {
			if vp.GenesisPool {
				result = result.Add(vp.GetCurrentlyLocked())
			}
		}
	}
	return result
}

func (pool *VestingPool) GetReservation(reservationId uint64) *VestingPoolReservation {
	for _, reservation := range pool.Reservations {
		if reservation.Id == reservationId {
			return reservation
		}
	}
	return nil
}

func (pool *VestingPool) AddReservation(reservationId uint64, amount math.Int) {
	for _, reservation := range pool.Reservations {
		if reservation.Id == reservationId {
			reservation.Amount = reservation.Amount.Add(amount)
			return
		}
	}
	pool.Reservations = append(pool.Reservations, &VestingPoolReservation{
		Id:     reservationId,
		Amount: amount,
	})
}

func (pool *VestingPool) SubstractFromReservation(reservationId uint64, amount math.Int) error {
	for i, reservation := range pool.Reservations {
		if reservation.Id == reservationId {
			if amount.GT(reservation.Amount) {
				return errors.Wrapf(c4eerrors.ErrAmount, "cannot substract from reservation, amount too big (%s > %s)", amount, reservation.Amount)
			}

			if amount.Equal(reservation.Amount) {
				pool.Reservations = append(pool.Reservations[:i], pool.Reservations[i+1:]...)
			} else {
				reservation.Amount = reservation.Amount.Sub(amount)
			}
			return nil
		}
	}
	return errors.Wrapf(sdkerrors.ErrNotFound, "reservation with id %d not found", reservationId)
}

func (pool *VestingPool) DecrementReservedAndSent(reservationId uint64, amount math.Int) error {
	if err := pool.SubstractFromReservation(reservationId, amount); err != nil {
		return err
	}
	pool.Sent = pool.Sent.Add(amount)

	return nil
}

func (pool *VestingPool) IncrementSent(amount math.Int) error {
	available := pool.GetLockedNotReserved()

	if available.LT(amount) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"send to new vesting account - vesting available: %s is smaller than requested amount: %s", available, amount)
	}
	pool.Sent = pool.Sent.Add(amount)
	return nil
}

func (vestingType *VestingType) ValidateVestingFree(free sdk.Dec) error {
	if free.GT(vestingType.Free) {
		return errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the free decimal must be equal to or lower than the vesting type free (%s < %s)", vestingType.Free.String(), free.String()))
	}

	return nil
}

func (vestingType *VestingType) ValidateVestingPeriods(lockupPeriod time.Duration, vestingPeriod time.Duration) error {
	if vestingType.LockupPeriod > lockupPeriod {
		return errors.Wrapf(c4eerrors.ErrParam,
			"the duration of lockup period must be equal to or greater than the vesting type lockup period (%s > %s)", vestingType.LockupPeriod.String(), lockupPeriod.String())
	}

	if vestingType.VestingPeriod > vestingPeriod {
		return errors.Wrapf(c4eerrors.ErrParam,
			"the duration of vesting period must be equal to or greater than the vesting type vesting period (%s > %s)", vestingType.VestingPeriod.String(), vestingPeriod.String())
	}

	return nil
}
