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
// signed integer. Flags are updated accordingly.
func (cpu *CPU) PutOffsetSPIntoHL() {
	sp := cpu.r.StackPointer()
	pc := cpu.r.ProgramCounter()
	offset := cpu.add16S8(*sp, cpu.m.Byte(*pc))
	cpu.r.SetPaired(registers.HL, offset)
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

// AddA adds the accumulator to itself, storing the result in the accumulator,
// and updates the flags.
func (cpu *CPU) AddA() {
	cpu.add8Helper(*cpu.r.Accumulator(), false, cpu.r.ResetFlag)
}

// AddR adds the provided register to the accumulator, storing the result in the
// accumulator, and updates the flags.
func (cpu *CPU) AddR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.add8Helper(*a, false, cpu.r.ResetFlag)
}

// AddN adds the immediate byte to the accumulator, storing the result in the
// accumulator, and updates the flags.
func (cpu *CPU) AddN() {
	cpu.add8Helper(cpu.immediateByte(), false, cpu.r.ResetFlag)
}

// AddHLDereference adds the value stored in the memory location referenced by
// register HL to the accumulator, storing the result in the accumulator, and
// updates the flags.
func (cpu *CPU) AddHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.add8Helper(cpu.m.Byte(hl), false, cpu.r.ResetFlag)
}

// AdcA adds the accumulator and the contents of the carry flag to the
// accumulator itself, storing the result in the accumulator, and updates the
// flags.
func (cpu *CPU) AdcA() {
	cpu.add8Helper(*cpu.r.Accumulator(), true, cpu.r.ResetFlag)
}

// AdcR adds the provided register and the contents of the carry flag to the
// accumulator, storing the result in the accumulator, and updates the flags.
func (cpu *CPU) AdcR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.add8Helper(*a, true, cpu.r.ResetFlag)
}

// AdcN adds the immediate byte and the contents of the carry flag to the
// accumulator, storing the result in the accumulator, and updates the flags.
func (cpu *CPU) AdcN() {
	cpu.add8Helper(cpu.immediateByte(), true, cpu.r.ResetFlag)
}

// AdcHLDereference adds the value stored in the memory location referenced by
// register HL and the contents of the carry flag to the accumulator, storing
// the result in the accumulator, and updates the flags.
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

// SubA subtracts the accumulator from itself, storing the result in the
// accumulator, and updates the flags.
func (cpu *CPU) SubA() {
	cpu.sub8Helper(*cpu.r.Accumulator(), false, cpu.r.SetFlag)
}

// SubR subtracts the provided register from the accumulator, storing the result
// in the accumulator, and updates the flags.
func (cpu *CPU) SubR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.sub8Helper(*a, false, cpu.r.SetFlag)
}

// SubN subtracts the immediate byte from the accumulator, storing the result in
// the accumulator, and updates the flags.
func (cpu *CPU) SubN() {
	cpu.sub8Helper(cpu.immediateByte(), false, cpu.r.SetFlag)
}

// SubHLDereference subtracts the value stored in the memory location referenced
// by register HL from the accumulator, storing the result in the accumulator,
// and updates the flags.
func (cpu *CPU) SubHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.sub8Helper(cpu.m.Byte(hl), false, cpu.r.SetFlag)
}

// SbcA subtracts the accumulator and the contents of the carry flag from
// itself, storing the result in the accumulator, and updates the flags.
func (cpu *CPU) SbcA() {
	cpu.sub8Helper(*cpu.r.Accumulator(), true, cpu.r.SetFlag)
}

// SbcR subtracts the provided register and the contents of the carry flag from
// the accumulator, storing the result in the accumulator, and updates the
// flags.
func (cpu *CPU) SbcR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.sub8Helper(*a, true, cpu.r.SetFlag)
}

// SbcN subtracts the immediate byte and the contents of the carry flag from the
// accumulator, storing the result in the accumulator, and updates the flags.
func (cpu *CPU) SbcN() {
	cpu.sub8Helper(cpu.immediateByte(), true, cpu.r.SetFlag)
}

// SbcHLDereference subtracts the value stored in the memory location referenced
// by register HL and the contents of the carry flag from the accumulator,
// storing the result in the accumulator, and updates the flags.
func (cpu *CPU) SbcHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.sub8Helper(cpu.m.Byte(hl), true, cpu.r.SetFlag)
}

func (cpu *CPU) bitwise8Helper(y uint8, f func(uint8, uint8) uint8) {
	acc := cpu.r.Accumulator()
	*acc = f(*acc, y)
}

func (cpu *CPU) and8(x, y uint8) (result uint8) {
	result = x & y

	cpu.r.ResetFlag(uint8(flags.C))
	cpu.r.SetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.PutFlag(uint8(flags.Z), result == 0)

	return result
}

// AndA performs bitwise AND between the contents of the accumulator and itself,
// storing the result in the accumulator. It also updates the flags.
func (cpu *CPU) AndA() {
	acc := cpu.r.Accumulator()
	cpu.bitwise8Helper(*acc, cpu.and8)
}

// AndR performs bitwise AND between the contents of the accumulator and the
// provided register, storing the result in the accumulator. It also updates the
// flags.
func (cpu *CPU) AndR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.bitwise8Helper(*a, cpu.and8)
}

// AndN performs bitwise AND between the contents of the accumulator and the
// immediate byte, storing the result in the accumulator. It also updates the
// flags.
func (cpu *CPU) AndN() {
	cpu.bitwise8Helper(cpu.immediateByte(), cpu.and8)
}

// AndHLDereference performs bitwise AND between the contents of the accumulator
// and the value stored in the memory location referenced by register HL,
// storing the result in the accumulator. It also updates the flags.
func (cpu *CPU) AndHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.bitwise8Helper(cpu.m.Byte(hl), cpu.and8)
}

func (cpu *CPU) xor8(x, y uint8) (result uint8) {
	result = x ^ y

	cpu.r.ResetFlag(uint8(flags.C))
	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.PutFlag(uint8(flags.Z), result == 0)

	return result
}

// XorA performs bitwise XOR between the contents of the accumulator and itself,
// storing the result in the accumulator. It also updates the flags.
func (cpu *CPU) XorA() {
	acc := cpu.r.Accumulator()
	cpu.bitwise8Helper(*acc, cpu.xor8)
}

// XorR performs bitwise XOR between the contents of the accumulator and the
// provided register, storing the result in the accumulator. It also updates the
// flags.
func (cpu *CPU) XorR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.bitwise8Helper(*a, cpu.xor8)
}

// XorN performs bitwise XOR between the contents of the accumulator and the
// immediate byte, storing the result in the accumulator. It also updates the
// flags.
func (cpu *CPU) XorN() {
	cpu.bitwise8Helper(cpu.immediateByte(), cpu.xor8)
}

// XorHLDereference performs bitwise XOR between the contents of the accumulator
// and the value stored in the memory location referenced by register HL,
// storing the result in the accumulator. It also updates the flags.
func (cpu *CPU) XorHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.bitwise8Helper(cpu.m.Byte(hl), cpu.xor8)
}

func (cpu *CPU) or8(x, y uint8) (result uint8) {
	result = x | y

	cpu.r.ResetFlag(uint8(flags.C))
	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.PutFlag(uint8(flags.Z), result == 0)

	return result
}

// OrA performs bitwise OR between the contents of the accumulator and itself,
// storing the result in the accumulator. It also updates the flags.
func (cpu *CPU) OrA() {
	acc := cpu.r.Accumulator()
	cpu.bitwise8Helper(*acc, cpu.or8)
}

// OrR performs bitwise OR between the contents of the accumulator and the
// provided register, storing the result in the accumulator. It also updates the
// flags.
func (cpu *CPU) OrR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.bitwise8Helper(*a, cpu.or8)
}

// OrN performs bitwise OR between the contents of the accumulator and the
// immediate byte, storing the result in the accumulator. It also updates the
// flags.
func (cpu *CPU) OrN() {
	cpu.bitwise8Helper(cpu.immediateByte(), cpu.or8)
}

// OrHLDereference performs bitwise OR between the contents of the accumulator
// and the value stored in the memory location referenced by register HL,
// storing the result in the accumulator. It also updates the flags.
func (cpu *CPU) OrHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.bitwise8Helper(cpu.m.Byte(hl), cpu.or8)
}

// See cpu.sub8Helper function.
func (cpu *CPU) compare8Helper(y uint8) {
	acc := cpu.r.Accumulator()
	cpu.add8(*acc, ^y+1, false)
	cpu.r.SetFlag(uint8(flags.N))
}

// CompareA subtracts the accumulator from itself, discarding the result, and
// updates the flags.
func (cpu *CPU) CompareA() {
	cpu.compare8Helper(*cpu.r.Accumulator())
}

// CompareR subtracts the provided register from the accumulator, discarding the
// result, and updates the flags.
func (cpu *CPU) CompareR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.compare8Helper(*a)
}

// CompareN subtracts the immediate byte from the accumulator, discarding the
// result, and updates the flags.
func (cpu *CPU) CompareN() {
	cpu.compare8Helper(cpu.immediateByte())
}

// CompareHLDereference subtracts the value stored in the memory location referenced
// by register HL from the accumulator, discarding the result, and updates the
// flags.
func (cpu *CPU) CompareHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.compare8Helper(cpu.m.Byte(hl))
}

func (cpu *CPU) increment8(x, by uint8) (result uint8) {
	result = x + by
	carryIns := result ^ x ^ by
	halfCarryOut := (carryIns>>4)&1 == 1

	cpu.r.PutFlag(uint8(flags.H), halfCarryOut)
	cpu.r.PutFlag(uint8(flags.Z), result == 0)

	return result
}

func (cpu *CPU) increment8Helper(x *uint8) {
	*x = cpu.increment8(*x, 1)
	cpu.r.ResetFlag(uint8(flags.N))
}

// IncrementA increments the accumulator register by 1, and updates the flags.
func (cpu *CPU) IncrementA() {
	acc := cpu.r.Accumulator()
	cpu.increment8Helper(acc)
}

// IncrementR increments the provided register by 1, and updates the flags.
func (cpu *CPU) IncrementR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.increment8Helper(a)
}

// IncrementHLDereference increments the memory contents referenced by register
// HL by 1, and updates the flags.
func (cpu *CPU) IncrementHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	val := cpu.m.Byte(hl)
	cpu.increment8Helper(&val)
	cpu.m.SetByte(hl, val)
}

func (cpu *CPU) decrement8Helper(x *uint8) {
	by := uint8(1)
	by = ^by + 1
	*x = cpu.increment8(*x, by)
	cpu.r.SetFlag(uint8(flags.N))
}

// DecrementA decrements the accumulator register by 1, and updates the flags.
func (cpu *CPU) DecrementA() {
	acc := cpu.r.Accumulator()
	cpu.decrement8Helper(acc)
}

// DecrementR decrements the provided register by 1, and updates the flags.
func (cpu *CPU) DecrementR(r registers.Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.decrement8Helper(a)
}

// DecrementHLDereference decrements the memory contents referenced by register
// HL by 1, and updates the flags.
func (cpu *CPU) DecrementHLDereference() {
	hl, _ := cpu.r.Paired(registers.HL)
	val := cpu.m.Byte(hl)
	cpu.decrement8Helper(&val)
	cpu.m.SetByte(hl, val)
}

// DecimalAdjustA adjusts the contents of the accumulator following a binary
// addition / subtraction. It retroactively turns the previous operation into
// a BCD addition / subtraction. This is achieved by subtracting 6 from the
// accumulator's upper and/or lower nybble. The N, C and H flags are utilised
// to determine whether the correction must be added or subtracted (depending
// on if the previous operation was an addition or subtraction) and how / if
// each nybble should be affected (depending on if a carry or half-carry
// occurred in the previous operation).
//
// Adapted from: https://forums.nesdev.com/viewtopic.php?f=20&t=15944#p196282
func (cpu *CPU) DecimalAdjustA() {
	acc := cpu.r.Accumulator()

	if n, _ := cpu.r.IsFlagSet(uint8(flags.N)); !n {
		if c, _ := cpu.r.IsFlagSet(uint8(flags.C)); c || *acc > 0x99 {
			*acc += 0x60
			cpu.r.SetFlag(uint8(flags.C))
		}

		if h, _ := cpu.r.IsFlagSet(uint8(flags.H)); h || (*acc&0x0F) > 0x09 {
			*acc += 0x06
		}
	} else {
		if c, _ := cpu.r.IsFlagSet(uint8(flags.C)); c {
			*acc -= 0x60
		}

		if h, _ := cpu.r.IsFlagSet(uint8(flags.H)); h {
			*acc -= 0x06
		}
	}

	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.PutFlag(uint8(flags.Z), *acc == 0)
}

// ComplementA sets the accumulator to the one's complement of itself, and
// updates the flags.
func (cpu *CPU) ComplementA() {
	acc := cpu.r.Accumulator()
	*acc = ^*acc

	cpu.r.SetFlag(uint8(flags.H))
	cpu.r.SetFlag(uint8(flags.N))
}

// ComplementCarryFlag toggles the carry flag, and updates the flags.
func (cpu *CPU) ComplementCarryFlag() {
	cpu.r.ToggleFlag(uint8(flags.C))

	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
}

// SetCarryFlag sets the carry flag, and updates the flags.
func (cpu *CPU) SetCarryFlag() {
	cpu.r.SetFlag(uint8(flags.C))

	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
}

/**
 * 16-bit arithmetic / logical operations
 */

// Adapted from: https://stackoverflow.com/a/8037485/1283818
func (cpu *CPU) add16(x, y uint16) (result uint16) {
	var (
		carryOut     bool
		halfCarryOut bool
	)

	carryOut = (x > 0xFFFF-y)
	result = x + y

	carryIns := result ^ x ^ y
	halfCarryOut = (carryIns>>12)&1 == 1

	cpu.r.PutFlag(uint8(flags.C), carryOut)
	cpu.r.PutFlag(uint8(flags.H), halfCarryOut)
	cpu.r.ResetFlag(uint8(flags.N))

	return result
}

func (cpu *CPU) add16Helper(y uint16) {
	hl, _ := cpu.r.Paired(registers.HL)
	cpu.r.SetPaired(registers.HL, cpu.add16(hl, y))
}

// AddRR adds the provided register to register HL, storing the result in
// register HL, and updates the flags.
func (cpu *CPU) AddRR(rr registers.Register) {
	pr, _ := cpu.r.Paired(rr)
	cpu.add16Helper(pr)
}

// AddSP adds the stack pointer register to register HL, storing the result in
// register HL, and updates the flags.
func (cpu *CPU) AddSP() {
	sp := cpu.r.StackPointer()
	cpu.add16Helper(*sp)
}

// IncrementRR increments the provided register by 1.
func (cpu *CPU) IncrementRR(rr registers.Register) {
	cpu.r.IncrementPaired(rr)
}

// IncrementSP increments the stack pointer register by 1.
func (cpu *CPU) IncrementSP() {
	sp := cpu.r.StackPointer()
	*sp = *sp + 1
}

// DecrementRR decrements the provided register by 1.
func (cpu *CPU) DecrementRR(rr registers.Register) {
	cpu.r.DecrementPaired(rr)
}

// DecrementSP decrements the stack pointer register by 1.
func (cpu *CPU) DecrementSP() {
	sp := cpu.r.StackPointer()
	*sp = *sp - 1
}

// Adds uint8 to uint16 with uint8 being treated as a signed number in the range
// [-128, 127].
func (cpu *CPU) add16S8(x uint16, y uint8) (result uint16) {
	var yy uint16
	if y > 127 {
		// Adding this to x is equal to subtracting ^y+1 from x.
		yy = ^uint16(^y+1) + 1
	} else {
		yy = uint16(y)
	}

	result = cpu.add16(x, yy)
	cpu.r.ResetFlag(uint8(flags.Z))

	return result
}

// AddOffsetImmediateToSP adds the immediate byte to the stack pointer register,
// with the immediate byte being treated as a signed integer in the range
// [-128, 127]. Flags are updated accordingly.
func (cpu *CPU) AddOffsetImmediateToSP() {
	sp := cpu.r.StackPointer()
	pc := cpu.r.ProgramCounter()
	offset := cpu.add16S8(*sp, cpu.m.Byte(*pc))
	*sp = offset
}

/**
 * Jumps / calls
 */

func (cpu *CPU) shouldJump(flag flags.Flag, isSet bool) bool {
	fs, _ := cpu.r.IsFlagSet(uint8(flag))
	return fs == isSet
}

// JumpHL sets the program counter to be equal to the contents of register HL.
func (cpu *CPU) JumpHL() {
	pc := cpu.r.ProgramCounter()
	hl, _ := cpu.r.Paired(registers.HL)
	*pc = hl
}

// JumpOffset sets the program counter to the result of adding the program
// counter to the immediate byte, with the immediate byte being treated as a
// signed integer in the range [-128, 127].
func (cpu *CPU) JumpOffset() {
	pc := cpu.r.ProgramCounter()
	ib := cpu.m.Byte(*pc)
	*pc = cpu.add16S8(*pc, ib)
}

// JumpOffsetConditionally sets the program counter to the result of adding the
// program counter to the immediate byte, with the immediate byte being treated
// as a signed integer in the range [-128, 127]. The condition is that the
// status of the provided flag must match the provided status.
func (cpu *CPU) JumpOffsetConditionally(flag flags.Flag, isSet bool) {
	if cpu.shouldJump(flag, isSet) {
		cpu.JumpOffset()
	}
}

// JumpNN sets the program counter to the immediate word. Memory[PC+1] is
// treated as the most significant byte and Memory[PC] as the least significant
// byte.
func (cpu *CPU) JumpNN() {
	pc := cpu.r.ProgramCounter()
	*pc = cpu.immediateWord()
}

// JumpNNConditionally sets the program counter to the immediate word if the
// status of the provided flag matches the provided status.
func (cpu *CPU) JumpNNConditionally(flag flags.Flag, isSet bool) {
	if cpu.shouldJump(flag, isSet) {
		cpu.JumpNN()
	}
}

// CallNN pushes the program counter onto the stack, then sets the program
// counter to the immediate word.
func (cpu *CPU) CallNN() {
	pc := cpu.r.ProgramCounter()
	cpu.pushWordOntoStack(*pc)
	cpu.JumpNN()
}

// CallNNConditionally pushes the program counter onto the stack, then sets the
// program counter to the immediate word. The condition is that the status of
// the provided flag must match the provided status.
func (cpu *CPU) CallNNConditionally(flag flags.Flag, isSet bool) {
	if cpu.shouldJump(flag, isSet) {
		cpu.CallNN()
	}
}

// Return pops a word from the stack and puts it into the program counter.
func (cpu *CPU) Return() {
	pc := cpu.r.ProgramCounter()
	*pc = cpu.popStack()
}

// ReturnConditionally pops a word from the stack and puts it into the program
// counter if the status of the provided flag matches the provided status.
func (cpu *CPU) ReturnConditionally(flag flags.Flag, isSet bool) {
	if cpu.shouldJump(flag, isSet) {
		cpu.Return()
	}
}

// ReturnPostInterrupt is a placeholder.
// TODO: Revisit when emulating interrupts.
func (cpu *CPU) ReturnPostInterrupt() {
	cpu.Return()
}

// Restart pushes the program counter onto the stack, then sets the program
// counter to the provided value.
func (cpu *CPU) Restart(t uint8) {
	pc := cpu.r.ProgramCounter()
	cpu.pushWordOntoStack(*pc)
	*pc = uint16(t)
}

/**
 * 8-bit rotation / shifts and bit instructions
 */

// RLCA rotates the contents of the accumulator to the left. The MSB becomes the
// the LSB and the carry flag. Other flags are reset.
func (cpu *CPU) RLCA() {
	acc := cpu.r.Accumulator()

	*acc = (*acc << 1) | (*acc >> 7)

	cpu.r.PutFlag(uint8(flags.C), *acc&0x01 == 0x01)
	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.ResetFlag(uint8(flags.Z))
}

// RLA rotates the contents of the accumulator to the left. The MSB becomes the
// carry flag and the carry flag becomes the LSB.
func (cpu *CPU) RLA() {
	acc := cpu.r.Accumulator()

	msb := *acc >> 7
	carry, _ := cpu.r.IsFlagSet(uint8(flags.C))

	*acc = *acc << 1
	if carry {
		*acc |= 0x01
	}

	cpu.r.PutFlag(uint8(flags.C), msb == 1)
	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.ResetFlag(uint8(flags.Z))
}

// RRCA rotates the contents of the accumulator to the right. The LSB becomes
// the MSB and the carry flag. Other flags are reset.
func (cpu *CPU) RRCA() {
	acc := cpu.r.Accumulator()

	*acc = (*acc >> 1) | (*acc << 7)

	cpu.r.PutFlag(uint8(flags.C), *acc&0x80 == 0x80)
	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.ResetFlag(uint8(flags.Z))
}

// RRA rotates the contents of the accumulator to the right. The LSB becomes the
// carry flag and the carry flag becomes the MSB.
func (cpu *CPU) RRA() {
	acc := cpu.r.Accumulator()

	lsb := *acc << 7
	carry, _ := cpu.r.IsFlagSet(uint8(flags.C))

	*acc = *acc >> 1
	if carry {
		*acc |= 0x80
	}

	cpu.r.PutFlag(uint8(flags.C), lsb == 1)
	cpu.r.ResetFlag(uint8(flags.H))
	cpu.r.ResetFlag(uint8(flags.N))
	cpu.r.ResetFlag(uint8(flags.Z))
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
