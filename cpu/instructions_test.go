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
	cpu := NewCPU()

	for _, tt := range instructionTests {
		t.Run(fmt.Sprintf("in=%d", tt.in), func(t *testing.T) {
			instructions[tt.in].execute(cpu)
		})
	}
}
