package cpu

import (
	"fmt"
	"testing"
)

var instructionTests = []struct {
	in  uint8
	out bool
}{
	{0x00, false},
	{0x01, true},
}

func TestInstructions(t *testing.T) {
	for _, tt := range instructionTests {
		t.Run(fmt.Sprintf("in=%d", tt.in), func(t *testing.T) {
			if out := instructions[tt.in].execute(); out != tt.out {
				t.Errorf("got %t, expected %t", out, tt.out)
			}
		})
	}
}
