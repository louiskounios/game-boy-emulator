package cpu

import (
	"fmt"
	"testing"
)

func TestMachineCycles(t *testing.T) {
	for _, instruction := range instructions {
		t.Run(fmt.Sprintf("opcode=0x%02X mnemonic=%s m=%d", instruction.opcode, instruction.mnemonic, instruction.machineCycles), func(t *testing.T) {
			cpu := New()
			instruction.execute(cpu)

			if cpu.c.M() != uint64(instruction.machineCycles) {
				t.Errorf("got %d, expected %d", cpu.c.M(), instruction.machineCycles)
			}
		})
	}

	for _, instruction := range instructionsCB {
		t.Run(fmt.Sprintf("opcode=0x%02X mnemonic=%s m=%d", instruction.opcode, instruction.mnemonic, instruction.machineCycles), func(t *testing.T) {
			cpu := New()
			instruction.execute(cpu)

			if cpu.c.M() != uint64(instruction.machineCycles) {
				t.Errorf("got %d, expected %d", cpu.c.T(), instruction.machineCycles)
			}
		})
	}
}
