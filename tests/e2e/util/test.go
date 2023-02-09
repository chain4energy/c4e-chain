package util

import (
	"errors"
	"time"
)

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
// copied and modified this function from testify assertion package
func Eventually(condition func() bool, waitFor time.Duration, tick time.Duration) error {
	ch := make(chan bool, 1)

	timer := time.NewTimer(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for tick := ticker.C; ; {
		select {
		case <-timer.C:
			return errors.New("condition never satisfied")
		case <-tick:
			tick = nil
			go func() { ch <- condition() }()
		case v := <-ch:
			if v {
				return nil
			}
			tick = ticker.C
		}
	}
}
