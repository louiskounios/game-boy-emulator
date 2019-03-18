package cpu_test

import (
	"fmt"
	"testing"

	"github.com/loizoskounios/game-boy-emulator/cpu"
)

func TestClockSetM(t *testing.T) {
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
		clock := cpu.NewClock(0)

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

func TestClockAddM(t *testing.T) {
	var testCases = []struct {
		baseClock *cpu.Clock
		addM1     uint64
		addM2     uint64
		outM      uint64
		outT      uint64
	}{
		{cpu.NewClock(0), 2, 3, 5, 20},
		{cpu.NewClock(1000), 1, 2, 1003, 4012},
		{cpu.NewClock(2000), 0, 0, 2000, 8000},
		{cpu.NewClock(3000), 23, 0, 3023, 12092},
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

func TestClockReset(t *testing.T) {
	var testCases = []struct {
		baseClock *cpu.Clock
		outM      uint64
		outT      uint64
	}{
		{cpu.NewClock(1000), 0, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("baseClock=%s", tc.baseClock), func(t *testing.T) {
			tc.baseClock.Reset()

			if tc.baseClock.M() != tc.outM {
				t.Errorf("got %d, expected %d", tc.baseClock.M(), tc.outM)
			}

			if tc.baseClock.T() != tc.outT {
				t.Errorf("got %d, expected %d", tc.baseClock.T(), tc.outT)
			}
		})
	}
}
