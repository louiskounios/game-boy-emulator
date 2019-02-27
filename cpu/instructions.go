package cpu

import (
	"fmt"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers"
)

type (
	instruction struct {
		opcode      uint8
		clockCycles uint8
		mnemonic    string
		execute     func(cpu *CPU)
	}

	instructionSet [256]*instruction
)

func (i instruction) String() string {
	return fmt.Sprintf("0x%02X %s", i.opcode, i.mnemonic)
}

var instructions = instructionSet{
	0x00: &instruction{0x00, 0, "mnemonic", func(cpu *CPU) { return }},
	0x01: &instruction{0x01, 0, "mnemonic", func(cpu *CPU) { return }},
	0x02: &instruction{0x02, 0, "mnemonic", func(cpu *CPU) { return }},
	0x06: &instruction{0x06, 2, "LD B,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.B) }},
	0x0E: &instruction{0x0E, 2, "LD C,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.C) }},
	0x16: &instruction{0x16, 2, "LD D,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.D) }},
	0x1E: &instruction{0x1E, 2, "LD E,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.E) }},
	0x26: &instruction{0x26, 2, "LD H,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.H) }},
	0x2E: &instruction{0x2E, 2, "LD L,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.L) }},
}
