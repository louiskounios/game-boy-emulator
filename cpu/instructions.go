package cpu

import "fmt"

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
	&instruction{0x00, 0, "mnemonic", func(cpu *CPU) { return }},
	&instruction{0x01, 0, "mnemonic", func(cpu *CPU) { return }},
	&instruction{0x02, 0, "mnemonic", func(cpu *CPU) { return }},
}
