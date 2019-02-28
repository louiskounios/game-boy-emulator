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

/**
 * 8-bit loads
 */

// PutRIntoR puts the value stored in register from into register to.
func (cpu *CPU) PutRIntoR(from, to registers.Register) {
	val, _ := cpu.r.Register(from)
	cpu.r.SetRegister(to, val)
}

// PutAIntoNNAddress calculates a 16-bit memory address by combining the two
// 8-bit values that are stored in memory locations referenced by the program
// counter and [PC+1].
// It then saves the contents of register A into that address in memory.
func (cpu *CPU) PutAIntoNNAddress() {
	address := cpu.wordFromProgramCounter()
	cpu.putRegisterIntoMemory(registers.A, address)
}

// PutAIntoBCAddress puts the value stored in register from into the memory
// location referenced by the BC register.
func (cpu *CPU) PutAIntoBCAddress() {
	cpu.putRegisterIntoAddressInRegister(registers.BC, registers.A)
}

// PutAIntoDEAddress puts the value stored in register from into the memory
// location referenced by the DE register.
func (cpu *CPU) PutAIntoDEAddress() {
	cpu.putRegisterIntoAddressInRegister(registers.DE, registers.A)
}

// PutRIntoHLAddress puts the value stored in register from into the memory
// location referenced by the HL register.
func (cpu *CPU) PutRIntoHLAddress(from registers.Register) {
	cpu.putRegisterIntoAddressInRegister(registers.HL, from)
}

// PutAIntoHLAddressThenIncrementHL puts the value stored in register A into
// the memory location referenced by the HL register, then increments register
// HL.
func (cpu *CPU) PutAIntoHLAddressThenIncrementHL() {
	cpu.PutRIntoHLAddress(registers.A)
	cpu.incrementRegister(registers.HL)
}

// PutAIntoHLAddressThenDecrementHL puts the value stored in register A into
// the memory location referenced by the HL register, then increments register
// HL.
func (cpu *CPU) PutAIntoHLAddressThenDecrementHL() {
	cpu.PutRIntoHLAddress(registers.A)
	cpu.decrementRegister(registers.HL)
}

// PutAIntoOffsetCAddress puts the value stored in Register A into the offset
// memory location resulting from the addition [C+0xFF00].
func (cpu *CPU) PutAIntoOffsetCAddress() {
	address := cpu.offsetAddressFromC()
	cpu.putRegisterIntoMemory(registers.A, address)
}

// PutAIntoOffsetImmediateAddress puts the value stored in Register A into the
// offset memory location resulting from the addition [Memory[PC]+0xFF00].
func (cpu *CPU) PutAIntoOffsetImmediateAddress() {
	address := cpu.offsetAddressFromImmediate()
	cpu.putRegisterIntoMemory(registers.A, address)
}

// PutNIntoR puts the value stored in the memory location referenced by the
// program counter into register to.
func (cpu *CPU) PutNIntoR(to registers.Register) {
	val := uint16(cpu.m.Byte(cpu.r.PC))
	cpu.r.SetRegister(to, val)
}

// PutNNDereferenceIntoA calculates a 16-bit memory address by combining the two
// 8-bit values that are stored in memory locations referenced by the program
// counter and [PC+1].
// It then saves the contents of the memory at that address into register A.
func (cpu *CPU) PutNNDereferenceIntoA() {
	address := cpu.wordFromProgramCounter()
	cpu.putMemoryIntoRegister(address, registers.A)
}

// PutOffsetCDereferenceIntoA puts the value stored in the offset memory
// location resulting from the addition [C+0xFF00] into register A.
func (cpu *CPU) PutOffsetCDereferenceIntoA() {
	address := cpu.offsetAddressFromC()
	cpu.putMemoryIntoRegister(address, registers.A)
}

// PutOffsetImmediateDereferenceIntoA puts the value stored in the offset
// memory location resulting from the addition [Memory[PC]+0xFF00] into
// register A.
func (cpu *CPU) PutOffsetImmediateDereferenceIntoA() {
	address := cpu.offsetAddressFromImmediate()
	cpu.putMemoryIntoRegister(address, registers.A)
}

// PutBCDereferenceIntoA puts the value stored in the memory location referenced
// by register BC into register A.
func (cpu *CPU) PutBCDereferenceIntoA() {
	cpu.putRegisterDereferenceIntoRegister(registers.BC, registers.A)
}

// PutDEDereferenceIntoA puts the value stored in the memory location referenced
// by register DE into register A.
func (cpu *CPU) PutDEDereferenceIntoA() {
	cpu.putRegisterDereferenceIntoRegister(registers.DE, registers.A)
}

// PutHLDereferenceIntoR puts the value stored in the memory location referenced
// by register HL into register r.
func (cpu *CPU) PutHLDereferenceIntoR(to registers.Register) {
	cpu.putRegisterDereferenceIntoRegister(registers.HL, to)
}

// PutHLDereferenceIntoAThenIncrementHL puts the value stored in the memory
// location referenced by register HL into register A, then increments
// register HL.
func (cpu *CPU) PutHLDereferenceIntoAThenIncrementHL() {
	cpu.PutHLDereferenceIntoR(registers.A)
	cpu.incrementRegister(registers.HL)
}

// PutHLDereferenceIntoAThenDecrementHL puts the value stored in the memory
// location referenced by register HL into register A, then increments
// register HL.
func (cpu *CPU) PutHLDereferenceIntoAThenDecrementHL() {
	cpu.PutHLDereferenceIntoR(registers.A)
	cpu.decrementRegister(registers.HL)
}

// PutNDereferenceIntoHLAddress puts the value stored in the memory location
// referenced by the program counter into the memory location referenced by the
// HL register.
func (cpu *CPU) PutNDereferenceIntoHLAddress() {
	hl, _ := cpu.r.Register(registers.HL)
	pc, _ := cpu.r.Register(registers.PC)
	n := cpu.m.Byte(pc)
	cpu.m.SetByte(hl, n)
}

/**
 * 16-bit loads
 */

// PutHLIntoSP puts the value stored in register HL into register SP.
func (cpu *CPU) PutHLIntoSP() {
	val, _ := cpu.r.Register(registers.HL)
	cpu.r.SetRegister(registers.SP, val)
}

// PutSPIntoNNAddress puts the value stored in register SP into the memory
// locations referenced by the program counter and [PC+1].
func (cpu *CPU) PutSPIntoNNAddress() {
	address := cpu.wordFromProgramCounter()
	val, _ := cpu.r.Register(registers.SP)
	cpu.m.SetWord(address, val)
}

// PutNNIntoRR calculates a 16-bit value by combining the two 8-bit
// values that are stored in memory locations referenced by the program
// counter and [PC+1].
// It then saves that value into register to.
func (cpu *CPU) PutNNIntoRR(to registers.Register) {
	val := cpu.wordFromProgramCounter()
	cpu.r.SetRegister(to, val)
}

/**
 * Common operations
 */

func (cpu *CPU) putRegisterIntoAddressInRegister(ar, vr registers.Register) {
	address, _ := cpu.r.Register(ar)
	cpu.putRegisterIntoMemory(vr, address)
}

func (cpu *CPU) putRegisterDereferenceIntoRegister(fr, tr registers.Register) {
	address, _ := cpu.r.Register(fr)
	val := uint16(cpu.m.Byte(address))
	cpu.r.SetRegister(tr, val)
}

func (cpu *CPU) putRegisterIntoMemory(r registers.Register, address uint16) {
	val, _ := cpu.r.Register(r)
	cpu.m.SetByte(address, uint8(val))
}

func (cpu *CPU) putMemoryIntoRegister(address uint16, r registers.Register) {
	val := uint16(cpu.m.Byte(address))
	cpu.r.SetRegister(r, val)
}

func (cpu *CPU) incrementRegister(r registers.Register) {
	cpu.r.Increment(r)
}

func (cpu *CPU) decrementRegister(r registers.Register) {
	cpu.r.Decrement(r)
}

func (cpu *CPU) wordFromProgramCounter() (word uint16) {
	pc, _ := cpu.r.Register(registers.PC)
	word = cpu.m.Word(pc)
	return word
}

func (cpu *CPU) offsetAddressFromImmediate() uint16 {
	pc, _ := cpu.r.Register(registers.PC)
	address := uint16(cpu.m.Byte(pc))
	return cpu.offsetAddress(address)
}

func (cpu *CPU) offsetAddressFromC() uint16 {
	c, _ := cpu.r.Register(registers.C)
	return cpu.offsetAddress(c)
}

func (cpu *CPU) offsetAddress(address uint16) uint16 {
	return address + 0xFF00
}
