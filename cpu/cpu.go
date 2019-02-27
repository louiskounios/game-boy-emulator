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
// counter and [Program Counter+1].
// It then saves the contents of register A into that address in memory.
func (cpu *CPU) PutAIntoNNAddress() {
	address := cpu.addressFromProgramCounter()
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

// PutAIntoCPlusFF00Address puts the value stored in Register A into the memory
// location resulting from the addition [Register C+0xFF00].
func (cpu *CPU) PutAIntoCPlusFF00Address() {
	cpu.putRegisterIntoMaskedAddress(registers.C, registers.A)
}

// PutAIntoNPlusFF00Address puts the value stored in register A into the memory
// location resulting from the addition [Program Counter+0xFF00].
func (cpu *CPU) PutAIntoNPlusFF00Address() {
	cpu.putRegisterIntoMaskedAddress(registers.PC, registers.A)
}

// PutNIntoR puts the value stored in the memory location referenced by the
// program counter into register to.
func (cpu *CPU) PutNIntoR(to registers.Register) {
	val := uint16(cpu.m.Byte(cpu.r.PC))
	cpu.r.SetRegister(to, val)
}

// PutNNIntoA calculates a 16-bit memory address by combining the two 8-bit
// values that are stored in memory locations referenced by the program
// counter and [Program Counter+1].
// It then saves the contents of the memory at that address into register A.
func (cpu *CPU) PutNNIntoA() {
	address := cpu.addressFromProgramCounter()
	cpu.putMemoryIntoRegister(address, registers.A)
}

// PutCPlusFF00IntoA puts the value stored in the memory location resulting from
// the addition [Program Counter+0xFF00] into register A.
func (cpu *CPU) PutCPlusFF00IntoA() {
	cpu.putMaskedAddressValueIntoRegister(registers.C, registers.A)
}

// PutNPlusFF00IntoA puts the value stored in the memory location resulting from
// the addition [Program Counter+0xFF00] into register A.
func (cpu *CPU) PutNPlusFF00IntoA() {
	cpu.putMaskedAddressValueIntoRegister(registers.PC, registers.A)
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

// PutNIntoHLAddress puts the value stored in the memory location referenced by
// the program counter into the memory location referenced by the HL register.
func (cpu *CPU) PutNIntoHLAddress() {
	hl, _ := cpu.r.Register(registers.HL)
	pc, _ := cpu.r.Register(registers.PC)
	n := cpu.m.Byte(pc)
	cpu.m.SetByte(hl, n)
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

func (cpu *CPU) putRegisterIntoMaskedAddress(ar, vr registers.Register) {
	address := cpu.maskedAddress(ar)
	cpu.putRegisterIntoMemory(vr, address)
}

func (cpu *CPU) putMaskedAddressValueIntoRegister(ar, tr registers.Register) {
	address := cpu.maskedAddress(ar)
	cpu.putMemoryIntoRegister(address, tr)
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

func (cpu *CPU) addressFromProgramCounter() (address uint16) {
	pc, _ := cpu.r.Register(registers.PC)
	address = cpu.m.Word(pc)
	return address
}

func (cpu *CPU) maskedAddress(r registers.Register) (address uint16) {
	address, _ = cpu.r.Register(r)
	address += 0xFF00
	return address
}
