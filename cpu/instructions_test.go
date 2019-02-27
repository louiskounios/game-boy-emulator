package cpu

import (
	"fmt"
	"testing"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers"
)

var instructionTests = []struct {
	opcode   uint16
	pc       uint16
	val      uint8
	register registers.Register
}{
	{0x06, 0xF0F0, 200, registers.B},
	{0x0E, 0xF0F0, 200, registers.C},
	{0x16, 0xF0F0, 200, registers.D},
	{0x1E, 0xF0F0, 200, registers.E},
	{0x26, 0xF0F0, 200, registers.H},
	{0x2E, 0xF0F0, 200, registers.L},
	{0x3E, 0xF0F0, 200, registers.A},
}

func TestInstructions(t *testing.T) {
	cpu := NewCPU()

	for _, tt := range instructionTests {
		t.Run(fmt.Sprintf("opcode=0x%02X pc=%d val=%d", tt.opcode, tt.pc, tt.val), func(t *testing.T) {
			cpu.r.PC = tt.pc
			cpu.m.SetByte(tt.pc, tt.val)
			instructions[tt.opcode].execute(cpu)
			if out, err := cpu.r.Register(tt.register); uint8(out) != tt.val || err != nil {
				t.Errorf("expected %d, got %d", tt.val, out)
			}
		})
	}
}
