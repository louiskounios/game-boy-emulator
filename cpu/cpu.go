package cpu

import "github.com/loizoskounios/game-boy-emulator/cpu/registers"

// CPU is the CPU.
type CPU struct {
	clock     clock
	registers registers.Registers
}

// NewCPU returns a new CPU struct.
func NewCPU() *CPU {
	return &CPU{}
}
