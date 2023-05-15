package types

import (
	"cosmossdk.io/errors"
	"fmt"
	c4eerrors "github.com/chain4energy/c4e-chain/types/errors"
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

func (m *VestingPool) GetCurrentlyLockedWithoutReservations() math.Int {
	return m.InitiallyLocked.Sub(m.Sent).Sub(m.Withdrawn).Sub(m.GetCurrentlyLockedInReservations())
}

func (m *VestingPool) GetCurrentlyLocked() math.Int {
	return m.InitiallyLocked.Sub(m.Sent).Sub(m.Withdrawn)
}

func (m *VestingPool) GetCurrentlyLockedInReservation(reservationId uint64) math.Int {
	return m.GetReservation(reservationId).Amount
}

func (pool *VestingPool) GetCurrentlyLockedInReservations() math.Int {
	amountSum := math.ZeroInt()
	for _, reservation := range pool.Reservations {
		amountSum = amountSum.Add(reservation.Amount)
	}
	return amountSum
}

func (m *VestingPool) Validate(accountAdd string) error {
	if len(m.Name) == 0 {
		return fmt.Errorf("vesting pool defined for account: %s has no name", accountAdd)
	}
	if m.InitiallyLocked.IsNegative() {
		return fmt.Errorf("vesting pool with name: %s defined for account: %s has InitiallyLocked value negative %s", m.Name, accountAdd, m.InitiallyLocked)
	}
	if m.Withdrawn.IsNegative() {
		return fmt.Errorf("vesting pool with name: %s defined for account: %s has Withdrawn value negative %s", m.Name, accountAdd, m.Withdrawn)
	}
	if m.Sent.IsNegative() {
		return fmt.Errorf("vesting pool with name: %s defined for account: %s has Sent value negative %s", m.Name, accountAdd, m.Sent)
	}
	if m.GetCurrentlyLocked().IsNegative() {
		return fmt.Errorf("vesting pool with name: %s defined for account: %s has InitiallyLocked (%s) < Withdrawn (%s) + Sent (%s)",
			m.Name, accountAdd, m.InitiallyLocked, m.Withdrawn, m.Sent)
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
	return errors.Wrapf(c4eerrors.ErrNotExists, "reservation with id %d not found", reservationId)
}

func (vestingPool *VestingPool) SendFromReservedTokens(reservationId uint64, amount math.Int) error { // TODO mylaca nazwa metody
	available := vestingPool.GetCurrentlyLockedInReservations() // TODO po co to?

	if available.LT(amount) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds, "send to new vesting account - vesting available: %s is smaller than requested amount: %s", available, amount)
	}
	if err := vestingPool.SubstractFromReservation(reservationId, amount); err != nil {
		return err
	}
	return nil
}

func (vestingPool *VestingPool) SendFromLockedTokens(amount math.Int) error { // TODO mylaca nazwa metody
	available := vestingPool.GetCurrentlyLockedWithoutReservations()

	if available.LT(amount) {
		return errors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"send to new vesting account - vesting available: %s is smaller than requested amount: %s", available, amount)
	}
	vestingPool.Sent = vestingPool.Sent.Add(amount)
	return nil
}

func (vestingType *VestingType) ValidateVestingFree(free sdk.Dec) error {
	if free.GT(vestingType.Free) {
		return errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the free decimal must be equal to or greater than the vesting type free (%s > %s)", vestingType.Free.String(), free.String()))
	}

	return nil
}

func (vestingType *VestingType) ValidateVestingPeriods(lockupPeriod time.Duration, vestingPeriod time.Duration) error {
	if vestingType.LockupPeriod > lockupPeriod {
		return errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the duration of lockup period must be equal to or greater than the vesting type lockup period (%s > %s)", vestingType.LockupPeriod.String(), lockupPeriod.String()))
	}

	if vestingType.VestingPeriod > vestingPeriod {
		return errors.Wrapf(c4eerrors.ErrParam,
			fmt.Sprintf("the duration of vesting period must be equal to or greater than the vesting type vesting period (%s > %s)", vestingType.VestingPeriod.String(), vestingPeriod.String()))
	}

	return nil
}
