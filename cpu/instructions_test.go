package cpu

import (
	"fmt"
	"testing"
)

func TestClockCycles(t *testing.T) {
	for _, instruction := range instructions {
		t.Run(fmt.Sprintf("opcode=0x%02X mnemonic=%s t=%d", instruction.opcode, instruction.mnemonic, instruction.clockCycles), func(t *testing.T) {
			cpu := New()
			instruction.execute(cpu)

			if cpu.c.T() != uint64(instruction.clockCycles) {
				t.Errorf("got %d, expected %d", cpu.c.T(), instruction.clockCycles)
			}
		})
	}

	for _, instruction := range instructionsCB {
		t.Run(fmt.Sprintf("opcode=0x%02X mnemonic=%s t=%d", instruction.opcode, instruction.mnemonic, instruction.clockCycles), func(t *testing.T) {
			cpu := New()
			instruction.execute(cpu)

			if cpu.c.T() != uint64(instruction.clockCycles) {
				t.Errorf("got %d, expected %d", cpu.c.T(), instruction.clockCycles)
			}
		})
	}
}
