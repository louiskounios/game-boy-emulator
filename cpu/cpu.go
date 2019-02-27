package cpu

import (
	"github.com/loizoskounios/game-boy-emulator/cpu/registers"
	"github.com/loizoskounios/game-boy-emulator/mmu"
)

// CPU is the CPU.
type CPU struct {
	c *clock
	i *instruction
	r *registers.Registers
	m *mmu.Memory
}

// NewCPU returns a new CPU struct.
func NewCPU() *CPU {
	c := &clock{}
	i := &instruction{}
	r := registers.NewRegisters()
	m := &mmu.Memory{}

	return &CPU{
		c: c,
		i: i,
		r: r,
		m: m,
	}
}

func (cpu *CPU) nop() {

}

/*
8-bit loads
*/

// PutRIntoR puts the value stored in register from into register to.
func (cpu *CPU) PutRIntoR(from registers.Register, to registers.Register) {
	val, _ := cpu.r.Register(from)
	cpu.r.SetRegister(to, val)
}

// PutRIntoHLAddress puts the value stored in register from into the memory
// location referenced by the HL register.
func (cpu *CPU) PutRIntoHLAddress(from registers.Register) {
	hl, _ := cpu.r.Register(registers.HL)
	val, _ := cpu.r.Register(from)
	cpu.m.SetByte(hl, uint8(val))
}

// PutNIntoR puts the value stored in the memory location referenced by the
// program counter into register to.
func (cpu *CPU) PutNIntoR(to registers.Register) {
	n := uint16(cpu.m.Byte(cpu.r.PC))
	cpu.r.SetRegister(to, n)
}

// PutNIntoHLAddress puts the value stored in the memory location referenced by
// the program counter into the memory location referenced by the HL register.
func (cpu *CPU) PutNIntoHLAddress() {
	hl, _ := cpu.r.Register(registers.HL)
	n := cpu.m.Byte(cpu.r.PC)
	cpu.m.SetByte(hl, n)
}

// PutHLDereferenceIntoR puts the value stored in the memory location referenced
// by register HL into register r.
func (cpu *CPU) PutHLDereferenceIntoR(to registers.Register) {
	hl, _ := cpu.r.Register(registers.HL)
	n := uint16(cpu.m.Byte(hl))
	cpu.r.SetRegister(to, n)
}
