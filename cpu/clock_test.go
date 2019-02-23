package cpu

import (
	"fmt"
	"testing"
)

var clockTests = []struct {
	inT  uint8
	outT uint8
	outM uint8
}{
	{1, 1, 4},
	{2, 2, 8},
	{3, 3, 12},
	{4, 4, 16},
}

func TestClock(t *testing.T) {
	for _, tt := range clockTests {
		clock := clock{}

		t.Run(fmt.Sprintf("t=%d", tt.inT), func(t *testing.T) {
			clock.SetT(tt.inT)

			if clock.T() != tt.outT {
				t.Errorf("got %d, expected %d", clock.T(), tt.outT)
			}

			if clock.M() != tt.outM {
				t.Errorf("got %d, expected %d", clock.M(), tt.outM)
			}
		})
	}
}
