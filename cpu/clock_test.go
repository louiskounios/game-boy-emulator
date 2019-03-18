package cpu

import (
	"fmt"
	"testing"
)

func TestSetM(t *testing.T) {
	var testCases = []struct {
		inM  uint64
		outM uint64
		outT uint64
	}{
		{1, 1, 4},
		{2, 2, 8},
		{3, 3, 12},
		{4, 4, 16},
	}

	for _, tc := range testCases {
		clock := NewClock(0)

		t.Run(fmt.Sprintf("t=%d", tc.inM), func(t *testing.T) {
			clock.SetM(tc.inM)

			if clock.M() != tc.outM {
				t.Errorf("got %d, expected %d", clock.T(), tc.outT)
			}

			if clock.T() != tc.outT {
				t.Errorf("got %d, expected %d", clock.M(), tc.outM)
			}
		})
	}
}

func TestAddT(t *testing.T) {
	var testCases = []struct {
		baseClock *Clock
		addM1     uint64
		addM2     uint64
		outM      uint64
		outT      uint64
	}{
		{NewClock(0), 2, 3, 5, 20},
		{NewClock(1000), 1, 2, 1003, 4012},
		{NewClock(2000), 0, 0, 2000, 8000},
		{NewClock(3000), 23, 0, 3023, 12092},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("baseClock=%s addM1=%d addM2=%d", tc.baseClock, tc.addM1, tc.addM2), func(t *testing.T) {
			tc.baseClock.AddM(tc.addM1)
			tc.baseClock.AddM(tc.addM2)

			if tc.baseClock.M() != tc.outM {
				t.Errorf("got %d, expected %d", tc.baseClock.M(), tc.outM)
			}

			if tc.baseClock.T() != tc.outT {
				t.Errorf("got %d, expected %d", tc.baseClock.T(), tc.outT)
			}
		})
	}
}
