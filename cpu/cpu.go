package cpu

import (
	"github.com/loizoskounios/game-boy-emulator/cpu/registers"
	"github.com/loizoskounios/game-boy-emulator/cpu/registers/flags"
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
	r := registers.New()
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
	f, _ := cpu.r.Auxiliary(from)
	t, _ := cpu.r.Auxiliary(to)
	*t = *f
}

// PutAIntoNNAddress calculates a 16-bit memory address by combining the two
// 8-bit values that are stored in memory locations referenced by the program
// counter and [PC+1].
// It then saves the contents of register A into that address in memory.
func (cpu *CPU) PutAIntoNNAddress() {
	address := cpu.immediateWord()
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
	n := cpu.immediateByte()
	t, _ := cpu.r.Auxiliary(to)
	*t = n
}

// PutNNDereferenceIntoA calculates a 16-bit memory address by combining the two
// 8-bit values that are stored in memory locations referenced by the program
// counter and [PC+1].
// It then saves the contents of the memory at that address into register A.
func (cpu *CPU) PutNNDereferenceIntoA() {
	address := cpu.immediateWord()
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
	n := cpu.immediateByte()
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.m.SetByte(hl, n)
}

/**
 * 16-bit loads
 */

// PutHLIntoSP puts the value stored in register HL into register SP.
func (cpu *CPU) PutHLIntoSP() {
	hl, _ := cpu.r.Paired(registers.HL)
	sp := cpu.r.StackPointer()
	*sp = hl
}

// Pushes the provided word onto the stack, then decrements the stack pointer
// by 2.
// The 8 most significant bits of the word are stored in Memory[SP-1].
// The 8 least significant bits of the word are stored in Memory[SP-2].
func (cpu *CPU) pushWordOntoStack(word uint16) {
	sp := cpu.r.StackPointer()
	*sp -= 2
	cpu.m.SetWord(*sp, word)
}

// PushAFOntoStack pushes the combined value stored in registers A and F onto
// the stack.
func (cpu *CPU) PushAFOntoStack() {
	cpu.pushWordOntoStack(cpu.r.AF())
}

// PushRROntoStack pushes the contents of paired register from onto the stack.
func (cpu *CPU) PushRROntoStack(from registers.Register) {
	word, _ := cpu.r.Paired(from)
	cpu.pushWordOntoStack(word)
}

// PutSPIntoNNAddress puts the value stored in register SP into the memory
// locations referenced by the program counter and [PC+1].
func (cpu *CPU) PutSPIntoNNAddress() {
	address := cpu.immediateWord()
	sp := cpu.r.StackPointer()
	cpu.m.SetWord(address, *sp)
}

// PutOffsetSPIntoHL puts the value resulting from the addition [SP+Memory[PC]]
// into register HL, with the value fetched from memory being treated as a
// signed integer.
// Flags are updated accordingly.
func (cpu *CPU) PutOffsetSPIntoHL() {
	sp := cpu.r.StackPointer()
	pc := cpu.r.ProgramCounter()
	offset, carry, hcarry := addSignedUnsigned(cpu.m.Byte(*pc), *sp)
	cpu.r.SetPaired(registers.HL, offset)
	cpu.r.PutFlag(uint8(flags.C), carry)
	cpu.r.PutFlag(uint8(flags.H), hcarry)
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.ResetFlag(uint8(flags.Z))
}

// Pops a word from the stack, then increments the stack pointer by 2.
// The 8 most significant bits of the word come from Memory[SP+1].
// The 8 least significant bits of the word come from Memory[SP].
func (cpu *CPU) popStack() uint16 {
	sp := cpu.r.StackPointer()
	val := cpu.m.Word(*sp)
	*sp += 2
	return val
}

// PopStackIntoAF pops a word from the stack and puts it into registers A and F.
func (cpu *CPU) PopStackIntoAF() {
	cpu.r.SetAF(cpu.popStack())
}

// PopStackIntoRR pops a word from the stack and puts it into the provided
// register pair.
func (cpu *CPU) PopStackIntoRR(to registers.Register) {
	cpu.r.SetPaired(to, cpu.popStack())
}

// PutNNIntoRR calculates a 16-bit value by combining the two 8-bit
// values that are stored in memory locations referenced by the program
// counter and [PC+1].
// It then saves that value into register to.
func (cpu *CPU) PutNNIntoRR(to registers.Register) {
	val := cpu.immediateWord()
	cpu.r.SetPaired(to, val)
}

/**
 * 8-bit arithmetic / logical operations
 */

// Adapted from: https://stackoverflow.com/a/8037485/1283818
func (cpu *CPU) add8(x, y uint8, useCarry bool) (result uint8) {
	var (
		carryOut     bool
		halfCarryOut bool
	)

	if s, _ := cpu.r.IsFlagSet(uint8(flags.C)); s && useCarry {
		carryOut = (x >= 0xFF-y)
		result = x + y + 1
	} else {
		carryOut = (x > 0xFF-y)
		result = x + y
	}

	carryIns := result ^ x ^ y
	halfCarryOut = (carryIns>>4)&1 == 1

	cpu.r.PutFlag(uint8(flags.C), carryOut)
	cpu.r.PutFlag(uint8(flags.H), halfCarryOut)
	cpu.r.PutFlag(uint8(flags.Z), result == 0)

	return result
}

func (cpu *CPU) add8Helper(y uint8, useCarry bool, f func(uint8) error) {
	acc := cpu.r.Accumulator()
	*acc = cpu.add8(*acc, y, useCarry)
	f(uint8(flags.N))
}

// AddA adds the accumulator to itself and updates the flags.
func (cpu *CPU) AddA() {
	cpu.add8Helper(*cpu.r.Accumulator(), false, cpu.r.ResetFlag)
}

// AddR adds the provided register to the accumulator and updates the flags.
func (cpu *CPU) AddR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.add8Helper(*a, false, cpu.r.ResetFlag)
}

// AddN adds the immediate byte to the accumulator and updates the flags.
func (cpu *CPU) AddN() {
	cpu.add8Helper(cpu.immediateByte(), false, cpu.r.ResetFlag)
}

// AddHLDereference adds the value stored in the memory location referenced by
// register HL to the accumulator and updates the flags.
func (cpu *CPU) AddHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.add8Helper(cpu.m.Byte(hl), false, cpu.r.ResetFlag)
}

// AdcA adds the accumulator and the contents of the carry flag to the
// accumulator itself and updates the flags.
func (cpu *CPU) AdcA() {
	cpu.add8Helper(*cpu.r.Accumulator(), true, cpu.r.ResetFlag)
}

// AdcR adds the provided register and the contents of the carry flag to the
// accumulator and updates the flags.
func (cpu *CPU) AdcR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.add8Helper(*a, true, cpu.r.ResetFlag)
}

// AdcN adds the immediate byte and the contents of the carry flag to the
// accumulator and updates the flags.
func (cpu *CPU) AdcN() {
	cpu.add8Helper(cpu.immediateByte(), true, cpu.r.ResetFlag)
}

// AdcHLDereference adds the value stored in the memory location referenced by
// register HL and the contents of the carry flag to the accumulator and updates
// the flags.
func (cpu *CPU) AdcHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.add8Helper(cpu.m.Byte(hl), true, cpu.r.ResetFlag)
}

// Adapted from: https://stackoverflow.com/a/8037485/1283818
func (cpu *CPU) sub8Helper(y uint8, useCarry bool, f func(uint8) error) {
	// a - b - c = a + ^b + 1 - c = a + ^b + !c
	// a - b = a + ^b + 1
	y = ^y
	if useCarry {
		cpu.r.ToggleFlag(uint8(flags.C))
		defer cpu.r.ToggleFlag(uint8(flags.C))
	} else {
		y++
	}

	cpu.add8Helper(y, useCarry, f)
}

// SubA subtracts the accumulator from itself and updates the flags.
func (cpu *CPU) SubA() {
	cpu.sub8Helper(*cpu.r.Accumulator(), false, cpu.r.SetFlag)
}

// SubR subtracts the provided register from the accumulator and updates the
// flags.
func (cpu *CPU) SubR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.sub8Helper(*a, false, cpu.r.SetFlag)
}

// SubN subtracts the immediate byte from the accumulator and updates the flags.
func (cpu *CPU) SubN() {
	cpu.sub8Helper(cpu.immediateByte(), false, cpu.r.SetFlag)
}

// SubHLDereference subtracts the value stored in the memory location referenced
// by register HL from the accumulator and updates the flags.
func (cpu *CPU) SubHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.sub8Helper(cpu.m.Byte(hl), false, cpu.r.SetFlag)
}

// SbcA subtracts the accumulator from itself and updates the flags.
func (cpu *CPU) SbcA() {
	cpu.sub8Helper(*cpu.r.Accumulator(), true, cpu.r.SetFlag)
}

// SbcR subtracts the provided register from the accumulator and updates the
// flags.
func (cpu *CPU) SbcR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.sub8Helper(*a, true, cpu.r.SetFlag)
}

// SbcN subtracts the immediate byte from the accumulator and updates the flags.
func (cpu *CPU) SbcN() {
	cpu.sub8Helper(cpu.immediateByte(), true, cpu.r.SetFlag)
}

// SbcHLDereference subtracts the value stored in the memory location referenced
// by register HL from the accumulator and updates the flags.
func (cpu *CPU) SbcHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.sub8Helper(cpu.m.Byte(hl), true, cpu.r.SetFlag)
}

func (cpu *CPU) and8(x, y uint8) (result uint8) {
	result = x & y

	cpu.r.ResetFlag(uint8(flags.C))
	cpu.r.SetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.PutFlag(uint8(flags.Z), result == 0)

	return result
}

func (cpu *CPU) or8(x, y uint8) (result uint8) {
	result = x | y

	cpu.r.ResetFlag(uint8(flags.C))
	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.PutFlag(uint8(flags.Z), result == 0)

	return result
}

func (cpu *CPU) xor8(x, y uint8) (result uint8) {
	result = x ^ y

	cpu.r.ResetFlag(uint8(flags.C))
	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.PutFlag(uint8(flags.Z), result == 0)

	return result
}

/**
 * Common operations
 */

func (cpu *CPU) putRegisterIntoAddressInRegister(ar, vr registers.Register) {
	address, _ := cpu.r.Paired(ar)
	cpu.putRegisterIntoMemory(vr, address)
}

func (cpu *CPU) putRegisterDereferenceIntoRegister(fr, tr registers.Register) {
	address, _ := cpu.r.Paired(fr)
	val := cpu.m.Byte(address)
	t, _ := cpu.r.Auxiliary(tr)
	*t = val
}

func (cpu *CPU) putRegisterIntoMemory(r registers.Register, address uint16) {
	v, _ := cpu.r.Auxiliary(r)
	cpu.m.SetByte(address, *v)
}

func (cpu *CPU) putMemoryIntoRegister(address uint16, r registers.Register) {
	val := cpu.m.Byte(address)
	t, _ := cpu.r.Auxiliary(r)
	*t = val
}

func (cpu *CPU) incrementRegister(r registers.Register) {
	cpu.r.IncrementPaired(r)
}

func (cpu *CPU) decrementRegister(r registers.Register) {
	cpu.r.DecrementPaired(r)
}

func (cpu *CPU) immediateByte() uint8 {
	pc := cpu.r.ProgramCounter()
	return cpu.m.Byte(*pc)
}

func (cpu *CPU) immediateWord() uint16 {
	pc := cpu.r.ProgramCounter()
	return cpu.m.Word(*pc)
}

func (cpu *CPU) offsetAddressFromImmediate() uint16 {
	return cpu.offsetAddress(uint16(cpu.immediateByte()))
}

func (cpu *CPU) offsetAddressFromC() uint16 {
	c, _ := cpu.r.Auxiliary(registers.C)
	return cpu.offsetAddress(uint16(*c))
}

func (cpu *CPU) offsetAddress(address uint16) uint16 {
	return address + 0xFF00
}

// Adds uint8 s to uint16 u, with s being treated as a signed variable.
// Returns three values.
// result is the result of the addition; carry is true if the result overflowed
// underflowed, false otherwise; hcarry is true if the result cannot fit within
// 11 bits.
func addSignedUnsigned(s uint8, u uint16) (result uint16, carry, hcarry bool) {
	// Stores the addition of the two numbers masked on their 11 least significant
	// bits.
	var partialResult uint16

	if s > 127 {
		ss := uint16(-s)
		result = u - ss
		if result > u {
			carry = true
		}
		partialResult = (u & 0x7FF) - (ss & 0x7FF)
	} else {
		ss := uint16(s)
		result = u + ss
		if result < u {
			carry = true
		}
		partialResult = (u & 0x7FF) + (ss & 0x7FF)
	}

	// We have a half carry if the partial result cannot fit within 11 bits.
	if partialResult > 0x7FF {
		hcarry = true
	}

	return result, carry, hcarry
}
