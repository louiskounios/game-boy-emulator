package cpu

import (
	"github.com/loizoskounios/game-boy-emulator/mmu"
)

// CPU is the CPU.
type CPU struct {
	c *Clock
	i *instruction
	r *Registers
	m *mmu.Memory
}

// New returns a new CPU struct.
func New() *CPU {
	c := NewClock(0)
	i := &instruction{}
	r := NewRegisters()
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

func (cpu *CPU) load8(val uint8, dst *uint8) {
	*dst = val

	cpu.c.AddT(1)
}

// LoadAIntoA loads the contents of the accumulator into the accumulator.
func (cpu *CPU) LoadAIntoA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.load8(*acc, acc)
}

// LoadRIntoA loads the contents of the provided register into the accumulator.
func (cpu *CPU) LoadRIntoA(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	acc := cpu.r.Accumulator()
	cpu.load8(*a, acc)
}

// LoadAIntoR loads the contents of the accumulator into the provided register.
func (cpu *CPU) LoadAIntoR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	a, _ := cpu.r.Auxiliary(r)
	cpu.load8(*acc, a)
}

// LoadRIntoR loads the contents of register from into register to.
func (cpu *CPU) LoadRIntoR(from, to Register) {
	cpu.r.IncrementProgramCounter(1)

	f, _ := cpu.r.Auxiliary(from)
	t, _ := cpu.r.Auxiliary(to)
	cpu.load8(*f, t)
}

// LoadNIntoA loads the contents of the memory address specified by the 8-bit
// immediate operand into the accumulator.
func (cpu *CPU) LoadNIntoA() {
	cpu.r.IncrementProgramCounter(1)

	n := cpu.memImmediateByte()
	acc := cpu.r.Accumulator()
	cpu.load8(n, acc)
}

// LoadNIntoR loads the contents of the memory address specified by the 8-bit
// immediate operand into the provided register.
func (cpu *CPU) LoadNIntoR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	n := cpu.memImmediateByte()
	a, _ := cpu.r.Auxiliary(r)
	cpu.load8(n, a)
}

// LoadRRIntoA loads the contents of the memory address specified by the
// provided paired register into the accumulator.
func (cpu *CPU) LoadRRIntoA(rr Register) {
	cpu.r.IncrementProgramCounter(1)

	pr, _ := cpu.r.Paired(rr)
	val := cpu.memByte(pr)

	acc := cpu.r.Accumulator()
	cpu.load8(val, acc)
}

// LoadHLIntoR loads the contents of the memory address specified by register
// HL into the provided register.
func (cpu *CPU) LoadHLIntoR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)

	a, _ := cpu.r.Auxiliary(r)
	cpu.load8(val, a)
}

// LoadAIntoHL loads the contents of the accumulator into the memory address
// specified by register HL.
func (cpu *CPU) LoadAIntoHL() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	hl, _ := cpu.r.Paired(RegisterHL)

	cpu.memStoreByte(hl, *acc)

	cpu.c.AddT(1)
}

// LoadRIntoHL loads the contents of the provided register into the memory
// address specified by register HL.
func (cpu *CPU) LoadRIntoHL(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	hl, _ := cpu.r.Paired(RegisterHL)

	cpu.memStoreByte(hl, *a)

	cpu.c.AddT(1)
}

// LoadNIntoHL loads the contents of the memory specified by the 8-bit immediate
// operand into the memory address specified by register HL.
func (cpu *CPU) LoadNIntoHL() {
	cpu.r.IncrementProgramCounter(1)

	n := cpu.memImmediateByte()

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.memStoreByte(hl, n)

	cpu.c.AddT(1)
}

func (cpu *CPU) offsetC() uint16 {
	c, _ := cpu.r.Auxiliary(RegisterC)
	return cpu.offset(uint16(*c))
}

// LoadOffsetCIntoA loads the contents of the memory address specified by the
// addition of register C and constant offset 0xFF00 into register A.
func (cpu *CPU) LoadOffsetCIntoA() {
	cpu.r.IncrementProgramCounter(1)

	address := cpu.offsetC()
	val := cpu.memByte(address)

	acc := cpu.r.Accumulator()
	cpu.load8(val, acc)
}

// LoadAIntoOffsetC loads the contents of register A into the memory address
// specified by the addition of register C and constant offset 0xFF00.
func (cpu *CPU) LoadAIntoOffsetC() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	address := cpu.offsetC()
	cpu.memStoreByte(address, *acc)

	cpu.c.AddT(1)
}

func (cpu *CPU) offsetImmediate() uint16 {
	return cpu.offset(uint16(cpu.memImmediateByte()))
}

// LoadOffsetImmediateIntoA loads the contents of the memory address specified
// by the addition of the 8-bit immediate operand and constant offset 0xFF00
// into register A.
func (cpu *CPU) LoadOffsetImmediateIntoA() {
	cpu.r.IncrementProgramCounter(1)

	address := cpu.offsetImmediate()
	val := cpu.memByte(address)

	acc := cpu.r.Accumulator()
	cpu.load8(val, acc)
}

// LoadAIntoOffsetImmediate loads the contents of register A into the memory
// address specified by the addition of the 8-bit immediate operand and constant
// offset 0xFF00.
func (cpu *CPU) LoadAIntoOffsetImmediate() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	address := cpu.offsetImmediate()
	cpu.memStoreByte(address, *acc)

	cpu.c.AddT(1)
}

// LoadNNIntoA loads the contents of the memory address specified by the 16-bit
// immediate operand into register A.
func (cpu *CPU) LoadNNIntoA() {
	cpu.r.IncrementProgramCounter(1)

	address := cpu.memImmediateWord()
	acc := cpu.r.Accumulator()
	cpu.memStoreByte(address, *acc)

	cpu.c.AddT(1)
}

// LoadAIntoNN loads the contents of register A into the memory address
// specified by the 16-bit immediate operand.
func (cpu *CPU) LoadAIntoNN() {
	cpu.r.IncrementProgramCounter(1)

	address := cpu.memImmediateWord()
	val := cpu.memByte(address)

	acc := cpu.r.Accumulator()
	cpu.load8(val, acc)
}

// LoadHLIntoAIncrementHL loads the contents of the memory address specified by
// register HL into the accumulator. Register HL is then incremented.
func (cpu *CPU) LoadHLIntoAIncrementHL() {
	cpu.LoadRRIntoA(RegisterHL)
	cpu.incrementRegister(RegisterHL)
}

// LoadHLIntoADecrementHL loads the contents of the memory address specified by
// register HL into the accumulator. Register HL is then decremented.
func (cpu *CPU) LoadHLIntoADecrementHL() {
	cpu.LoadRRIntoA(RegisterHL)
	cpu.decrementRegister(RegisterHL)
}

// LoadAIntoBC loads the contents of the accumulator into the memory address
// specified by register BC.
func (cpu *CPU) LoadAIntoBC() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	bc, _ := cpu.r.Paired(RegisterBC)

	cpu.memStoreByte(bc, *acc)

	cpu.c.AddT(1)
}

// LoadAIntoDE loads the contents of the accumulator into the memory address
// specified by register DE.
func (cpu *CPU) LoadAIntoDE() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	de, _ := cpu.r.Paired(RegisterDE)

	cpu.memStoreByte(de, *acc)

	cpu.c.AddT(1)
}

// LoadAIntoHLIncrementHL loads the contents of the accumulator into the memory
// address specified by register HL. Register HL is then incremented.
func (cpu *CPU) LoadAIntoHLIncrementHL() {
	cpu.LoadAIntoHL()
	cpu.incrementRegister(RegisterHL)
}

// LoadAIntoHLDecrementHL loads the contents of the accumulator into the memory
// address specified by register HL. Register HL is then decremented.
func (cpu *CPU) LoadAIntoHLDecrementHL() {
	cpu.LoadAIntoHL()
	cpu.decrementRegister(RegisterHL)
}

/**
 * 16-bit loads
 */

// PutHLIntoSP puts the value stored in register HL into register SP.
func (cpu *CPU) PutHLIntoSP() {
	hl, _ := cpu.r.Paired(RegisterHL)
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
func (cpu *CPU) PushRROntoStack(from Register) {
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
	cpu.r.SetPaired(RegisterHL, offset)
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
func (cpu *CPU) PopStackIntoRR(to Register) {
	cpu.r.SetPaired(to, cpu.popStack())
}

// PutNNIntoRR calculates a 16-bit value by combining the two 8-bit
// values that are stored in memory locations referenced by the program
// counter and [PC+1].
// It then saves that value into register to.
func (cpu *CPU) PutNNIntoRR(to Register) {
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

	if s, _ := cpu.r.IsFlagSet(FlagC); s && useCarry {
		carryOut = (x >= 0xFF-y)
		result = x + y + 1
	} else {
		carryOut = (x > 0xFF-y)
		result = x + y
	}

	carryIns := result ^ x ^ y
	halfCarryOut = (carryIns>>4)&1 == 1

	cpu.r.PutFlag(FlagC, carryOut)
	cpu.r.PutFlag(FlagH, halfCarryOut)
	cpu.r.PutFlag(FlagZ, result == 0)

	return result
}

func (cpu *CPU) add8Helper(y uint8, useCarry bool, f func(Flag) error) {
	acc := cpu.r.Accumulator()
	*acc = cpu.add8(*acc, y, useCarry)
	f(FlagN)
}

// AddA adds the accumulator to itself, storing the result in the accumulator,
// and updates the flags.
func (cpu *CPU) AddA() {
	cpu.add8Helper(*cpu.r.Accumulator(), false, cpu.r.ResetFlag)
}

// AddR adds the provided register to the accumulator, storing the result in the
// accumulator, and updates the flags.
func (cpu *CPU) AddR(r Register) {
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
	hl, _ := cpu.r.Paired(RegisterHL)
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
func (cpu *CPU) AdcR(r Register) {
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
	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.add8Helper(cpu.m.Byte(hl), true, cpu.r.ResetFlag)
}

// Adapted from: https://stackoverflow.com/a/8037485/1283818
func (cpu *CPU) sub8Helper(y uint8, useCarry bool, f func(Flag) error) {
	// a - b - c = a + ^b + 1 - c = a + ^b + !c
	// a - b = a + ^b + 1
	y = ^y
	if useCarry {
		cpu.r.ToggleFlag(FlagC)
		defer cpu.r.ToggleFlag(FlagC)
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
func (cpu *CPU) SubR(r Register) {
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
	hl, _ := cpu.r.Paired(RegisterHL)
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
func (cpu *CPU) SbcR(r Register) {
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
	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.sub8Helper(cpu.m.Byte(hl), true, cpu.r.SetFlag)
}

func (cpu *CPU) bitwise8Helper(y uint8, f func(uint8, uint8) uint8) {
	acc := cpu.r.Accumulator()
	*acc = f(*acc, y)
}

func (cpu *CPU) and8(x, y uint8) (result uint8) {
	result = x & y

	cpu.r.ResetFlag(FlagC)
	cpu.r.SetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, result == 0)

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
func (cpu *CPU) AndR(r Register) {
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
	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.bitwise8Helper(cpu.m.Byte(hl), cpu.and8)
}

func (cpu *CPU) xor8(x, y uint8) (result uint8) {
	result = x ^ y

	cpu.r.ResetFlag(FlagC)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, result == 0)

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
func (cpu *CPU) XorR(r Register) {
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
	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.bitwise8Helper(cpu.m.Byte(hl), cpu.xor8)
}

func (cpu *CPU) or8(x, y uint8) (result uint8) {
	result = x | y

	cpu.r.ResetFlag(FlagC)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, result == 0)

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
func (cpu *CPU) OrR(r Register) {
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
	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.bitwise8Helper(cpu.m.Byte(hl), cpu.or8)
}

// See cpu.sub8Helper function.
func (cpu *CPU) compare8Helper(y uint8) {
	acc := cpu.r.Accumulator()
	cpu.add8(*acc, ^y+1, false)
	cpu.r.SetFlag(FlagN)
}

// CompareA subtracts the accumulator from itself, discarding the result, and
// updates the flags.
func (cpu *CPU) CompareA() {
	cpu.compare8Helper(*cpu.r.Accumulator())
}

// CompareR subtracts the provided register from the accumulator, discarding the
// result, and updates the flags.
func (cpu *CPU) CompareR(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.compare8Helper(*a)
}

// CompareN subtracts the immediate byte from the accumulator, discarding the
// result, and updates the flags.
func (cpu *CPU) CompareN() {
	cpu.compare8Helper(cpu.immediateByte())
}

// CompareHLDereference subtracts the value stored in the memory location
// referenced by register HL from the accumulator, discarding the result, and
// updates the flags.
func (cpu *CPU) CompareHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.compare8Helper(cpu.m.Byte(hl))
}

func (cpu *CPU) increment8(x, by uint8) (result uint8) {
	result = x + by
	carryIns := result ^ x ^ by
	halfCarryOut := (carryIns>>4)&1 == 1

	cpu.r.PutFlag(FlagH, halfCarryOut)
	cpu.r.PutFlag(FlagZ, result == 0)

	return result
}

func (cpu *CPU) increment8Helper(x *uint8) {
	*x = cpu.increment8(*x, 1)
	cpu.r.ResetFlag(FlagN)
}

// IncrementA increments the accumulator register by 1, and updates the flags.
func (cpu *CPU) IncrementA() {
	acc := cpu.r.Accumulator()
	cpu.increment8Helper(acc)
}

// IncrementR increments the provided register by 1, and updates the flags.
func (cpu *CPU) IncrementR(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.increment8Helper(a)
}

// IncrementHLDereference increments the memory contents referenced by register
// HL by 1, and updates the flags.
func (cpu *CPU) IncrementHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.increment8Helper(&val)
	cpu.m.SetByte(hl, val)
}

func (cpu *CPU) decrement8Helper(x *uint8) {
	by := uint8(1)
	by = ^by + 1
	*x = cpu.increment8(*x, by)
	cpu.r.SetFlag(FlagN)
}

// DecrementA decrements the accumulator register by 1, and updates the flags.
func (cpu *CPU) DecrementA() {
	acc := cpu.r.Accumulator()
	cpu.decrement8Helper(acc)
}

// DecrementR decrements the provided register by 1, and updates the flags.
func (cpu *CPU) DecrementR(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.decrement8Helper(a)
}

// DecrementHLDereference decrements the memory contents referenced by register
// HL by 1, and updates the flags.
func (cpu *CPU) DecrementHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
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

	if n, _ := cpu.r.IsFlagSet(FlagN); !n {
		if c, _ := cpu.r.IsFlagSet(FlagC); c || *acc > 0x99 {
			*acc += 0x60
			cpu.r.SetFlag(FlagC)
		}

		if h, _ := cpu.r.IsFlagSet(FlagH); h || (*acc&0x0F) > 0x09 {
			*acc += 0x06
		}
	} else {
		if c, _ := cpu.r.IsFlagSet(FlagC); c {
			*acc -= 0x60
		}

		if h, _ := cpu.r.IsFlagSet(FlagH); h {
			*acc -= 0x06
		}
	}

	cpu.r.ResetFlag(FlagH)
	cpu.r.PutFlag(FlagZ, *acc == 0)
}

// ComplementA sets the accumulator to the one's complement of itself, and
// updates the flags.
func (cpu *CPU) ComplementA() {
	acc := cpu.r.Accumulator()
	*acc = ^*acc

	cpu.r.SetFlag(FlagH)
	cpu.r.SetFlag(FlagN)
}

// ComplementCarryFlag toggles the carry flag, and updates the flags.
func (cpu *CPU) ComplementCarryFlag() {
	cpu.r.ToggleFlag(FlagC)

	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
}

// SetCarryFlag sets the carry flag, and updates the flags.
func (cpu *CPU) SetCarryFlag() {
	cpu.r.SetFlag(FlagC)

	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
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

	cpu.r.PutFlag(FlagC, carryOut)
	cpu.r.PutFlag(FlagH, halfCarryOut)
	cpu.r.ResetFlag(FlagN)

	return result
}

func (cpu *CPU) add16Helper(y uint16) {
	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.r.SetPaired(RegisterHL, cpu.add16(hl, y))
}

// AddRR adds the provided register to register HL, storing the result in
// register HL, and updates the flags.
func (cpu *CPU) AddRR(rr Register) {
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
func (cpu *CPU) IncrementRR(rr Register) {
	cpu.r.IncrementPaired(rr)
}

// IncrementSP increments the stack pointer register by 1.
func (cpu *CPU) IncrementSP() {
	sp := cpu.r.StackPointer()
	*sp = *sp + 1
}

// DecrementRR decrements the provided register by 1.
func (cpu *CPU) DecrementRR(rr Register) {
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
	cpu.r.ResetFlag(FlagZ)

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

func (cpu *CPU) shouldJump(flag Flag, isSet bool) bool {
	fs, _ := cpu.r.IsFlagSet(flag)
	return fs == isSet
}

// JumpHL sets the program counter to be equal to the contents of register HL.
func (cpu *CPU) JumpHL() {
	pc := cpu.r.ProgramCounter()
	hl, _ := cpu.r.Paired(RegisterHL)
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
func (cpu *CPU) JumpOffsetConditionally(flag Flag, isSet bool) {
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
func (cpu *CPU) JumpNNConditionally(flag Flag, isSet bool) {
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
func (cpu *CPU) CallNNConditionally(flag Flag, isSet bool) {
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
func (cpu *CPU) ReturnConditionally(flag Flag, isSet bool) {
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

func rotate8(x uint8, right bool) (result uint8, rotatedBitset bool) {
	if right {
		rotatedBitset = x&0x01 == 0x01
		result = x>>1 | x<<7

	} else {
		rotatedBitset = x&0x80 == 0x80
		result = x<<1 | x>>7
	}

	return result, rotatedBitset
}

func (cpu *CPU) rotate8SwapHelper(x *uint8, right bool) {
	var rotatedBitset bool
	*x, rotatedBitset = rotate8(*x, right)

	carry, _ := cpu.r.IsFlagSet(FlagC)
	if carry {
		if right {
			*x |= 0x80
		} else {
			*x |= 0x01
		}
	}

	cpu.r.PutFlag(FlagC, rotatedBitset)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, *x == 0)
}

func (cpu *CPU) rotate8BothHelper(x *uint8, right bool) {
	var rotatedBitset bool
	*x, rotatedBitset = rotate8(*x, right)

	cpu.r.PutFlag(FlagC, rotatedBitset)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, *x == 0)
}

// RLCA rotates the contents of the accumulator to the left. The MSB becomes the
// LSB and the carry flag. All other flags are reset.
func (cpu *CPU) RLCA() {
	acc := cpu.r.Accumulator()
	cpu.rotate8BothHelper(acc, false)
	cpu.r.ResetFlag(FlagZ)
}

// RLA rotates the contents of the accumulator to the left. The MSB becomes the
// carry flag and the carry flag becomes the LSB. All other flags are reset.
func (cpu *CPU) RLA() {
	acc := cpu.r.Accumulator()
	cpu.rotate8SwapHelper(acc, false)
	cpu.r.ResetFlag(FlagZ)
}

// RRCA rotates the contents of the accumulator to the right. The LSB becomes
// the MSB and the carry flag. All other flags are reset.
func (cpu *CPU) RRCA() {
	acc := cpu.r.Accumulator()
	cpu.rotate8BothHelper(acc, true)
	cpu.r.ResetFlag(FlagZ)
}

// RRA rotates the contents of the accumulator to the right. The LSB becomes the
// carry flag and the carry flag becomes the MSB. All other flags are reset.
func (cpu *CPU) RRA() {
	acc := cpu.r.Accumulator()
	cpu.rotate8SwapHelper(acc, true)
	cpu.r.ResetFlag(FlagZ)
}

// RLCACB rotates the contents of the accumulator to the left. The MSB becomes
// the LSB and the carry flag. Flags are updated accordingly.
func (cpu *CPU) RLCACB() {
	acc := cpu.r.Accumulator()
	cpu.rotate8BothHelper(acc, false)
}

// RLC rotates the contents of the provided register to the left. The MSB
// becomes the LSB and the carry flag. Flags are updated accordingly.
func (cpu *CPU) RLC(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.rotate8BothHelper(a, false)
}

// RLCHLDereference rotates the contents of the memory location referenced by
// register HL to the left. The MSB becomes the the LSB and the carry flag.
// Flags are updated accordingly.
func (cpu *CPU) RLCHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.rotate8BothHelper(&val, false)
	cpu.m.SetByte(hl, val)
}

// RRCACB rotates the contents of the accumulator to the right. The LSB becomes
// the MSB and the carry flag. Flags are updated accordingly.
func (cpu *CPU) RRCACB() {
	acc := cpu.r.Accumulator()
	cpu.rotate8BothHelper(acc, true)
}

// RRC rotates the contents of the provided register to the right. The LSB
// becomes the MSB and the carry flag. Flags are updated accordingly.
func (cpu *CPU) RRC(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.rotate8BothHelper(a, true)
}

// RRCHLDereference rotates the contents of the memory location referenced by
// register HL to the right. The LSB becomes the the MSB and the carry flag.
// Flags are updated accordingly.
func (cpu *CPU) RRCHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.rotate8BothHelper(&val, true)
	cpu.m.SetByte(hl, val)
}

// RLACB rotates the contents of the accumulator to the left. The MSB becomes
// the carry flag and the carry flag becomes the LSB. Flags are updated
// accordingly.
func (cpu *CPU) RLACB() {
	acc := cpu.r.Accumulator()
	cpu.rotate8SwapHelper(acc, false)
}

// RL rotates the contents of the provided register to the left. The MSB becomes
// the carry flag and the carry flag becomes the LSB. Flags are updated
// accordingly.
func (cpu *CPU) RL(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.rotate8SwapHelper(a, false)
}

// RLHLDereference rotates the contents of the memory location referenced by
// register HL to the left. The MSB becomes the carry flag and the carry flag
// becomes the LSB. Flags are updated accordingly.
func (cpu *CPU) RLHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.rotate8SwapHelper(&val, false)
	cpu.m.SetByte(hl, val)
}

// RRACB rotates the contents of the accumulator to the right. The LSB becomes
// the carry flag and the carry flag becomes the MSB. Flags are updated
// accordingly.
func (cpu *CPU) RRACB() {
	acc := cpu.r.Accumulator()
	cpu.rotate8SwapHelper(acc, true)
}

// RR rotates the contents of the provided register to the right. The LSB
// becomes the carry flag and the carry flag becomes the MSB. Flags are updated
// accordingly.
func (cpu *CPU) RR(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.rotate8SwapHelper(a, true)
}

// RRHLDereference rotates the contents of the memory location referenced by
// register HL to the right. The LSB becomes the carry flag and the carry flag
// becomes the MSB. Flags are updated accordingly.
func (cpu *CPU) RRHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.rotate8SwapHelper(&val, true)
	cpu.m.SetByte(hl, val)
}

func shift8(x uint8, right bool) (result uint8, shiftedBitSet bool) {
	if right {
		shiftedBitSet = x&0x01 == 0x01
		result = x >> 1
	} else {
		shiftedBitSet = x&0x80 == 0x80
		result = x << 1
	}

	return result, shiftedBitSet
}

func (cpu *CPU) shift8Helper(x *uint8, right bool) {
	var shiftedBitSet bool
	*x, shiftedBitSet = shift8(*x, right)

	cpu.r.PutFlag(FlagC, shiftedBitSet)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, *x == 0)
}

// SLAA shifts the contents of the accumulator to the left. The MSB becomes the
// carry flag and the LSB is reset. Flags are updated accordingly.
func (cpu *CPU) SLAA() {
	acc := cpu.r.Accumulator()
	cpu.shift8Helper(acc, false)
}

// SLA shifts the contents of the provided register to the left. The MSB becomes
// the carry flag and the LSB is reset. Flags are updated accordingly.
func (cpu *CPU) SLA(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.shift8Helper(a, false)
}

// SLAHLDereference shifts the contents of the memory location referenced by
// register HL to the left. The MSB becomes the carry flag and the LSB is reset.
// Flags are updated accordingly.
func (cpu *CPU) SLAHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.shift8Helper(&val, false)
	cpu.m.SetByte(hl, val)
}

// SRAA shifts the contents of the accumulator to the right. The LSB becomes the
// carry flag and the MSB retains its value. Flags are updated accordingly.
func (cpu *CPU) SRAA() {
	acc := cpu.r.Accumulator()
	bit7mask := *acc & 0x80
	cpu.shift8Helper(acc, true)
	*acc = *acc | bit7mask
}

// SRA shifts the contents of the provided register to the right. The LSB
// becomes the carry flag and the MSB retains its value. Flags are updated
// accordingly.
func (cpu *CPU) SRA(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	bit7mask := *a & 0x80
	cpu.shift8Helper(a, true)
	*a = *a | bit7mask
}

// SRAHLDereference shifts the contents of the memory location referenced by
// register HL to the right. The LSB becomes the carry flag and the MSB retains
// its value. Flags are updated accordingly.
func (cpu *CPU) SRAHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	bit7mask := val & 0x80
	cpu.shift8Helper(&val, true)
	val |= bit7mask
	cpu.m.SetByte(hl, val)
}

// SRLA shifts the contents of the accumulator to the right. The LSB becomes the
// carry flag and the MSB is reset. Flags are updated accordingly.
func (cpu *CPU) SRLA() {
	acc := cpu.r.Accumulator()
	cpu.shift8Helper(acc, true)
}

// SRL shifts the contents of the provided register to the right. The LSB
// becomes the carry flag and the MSB is reset. Flags are updated accordingly.
func (cpu *CPU) SRL(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.shift8Helper(a, true)
}

// SRLHLDereference shifts the contents of the memory location referenced by
// register HL to the right. The LSB becomes the carry flag and the MSB is
// reset. Flags are updated accordingly.
func (cpu *CPU) SRLHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.shift8Helper(&val, true)
	cpu.m.SetByte(hl, val)
}

func swap8(x uint8) uint8 {
	return x<<4 | x>>4
}

func (cpu *CPU) swap8Helper(x *uint8) {
	*x = swap8(*x)

	cpu.r.ResetFlag(FlagC)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, *x == 0)
}

// SwapA swaps the accumulator nibbles. Flags are updated accordingly.
func (cpu *CPU) SwapA() {
	acc := cpu.r.Accumulator()
	cpu.swap8Helper(acc)
}

// Swap swaps the provided register's nibbles. Flags are updated accordingly.
func (cpu *CPU) Swap(r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.swap8Helper(a)
}

// SwapHLDereference swaps the nibbles of the contents of the memory location
// referenced by register HL. Flags are updated accordingly.
func (cpu *CPU) SwapHLDereference() {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.swap8Helper(&val)
	cpu.m.SetByte(hl, val)
}

func (cpu *CPU) bit8(x, b uint8) {
	isSet := (x>>b)&0x01 == 1

	cpu.r.SetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, !isSet)
}

// BitA sets the Z flag to the complement of the contents of the provided bit in
// the accumulator. The H and N flags are set and reset respectively.
func (cpu *CPU) BitA(b uint8) {
	acc := cpu.r.Accumulator()
	cpu.bit8(*acc, b)
}

// Bit sets the Z flag to the complement of the contents of the provided bit in
// the provided register. The H and N flags are set and reset respectively.
func (cpu *CPU) Bit(b uint8, r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.bit8(*a, b)
}

// BitHLDereference sets the Z flag to the complement of the contents of the
// provided bit in the contents of the memory location referenced by register
// HL. The H and N flags are set and reset respectively.
func (cpu *CPU) BitHLDereference(b uint8) {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.bit8(val, b)
}

func (cpu *CPU) reset8(x *uint8, b uint8) {
	*x &^= (1 << b)
}

// ResetA resets the provided bit in the accumulator.
func (cpu *CPU) ResetA(b uint8) {
	acc := cpu.r.Accumulator()
	cpu.reset8(acc, b)
}

// Reset resets the provided bit in the provided register.
func (cpu *CPU) Reset(b uint8, r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.reset8(a, b)
}

// ResetHLDereference resets the provided bit in the contents of the memory
// location referenced by register HL.
func (cpu *CPU) ResetHLDereference(b uint8) {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.reset8(&val, b)
	cpu.m.SetByte(hl, val)
}

func (cpu *CPU) set8(x *uint8, b uint8) {
	*x |= (1 << b)
}

// SetA sets the provided bit in the accumulator.
func (cpu *CPU) SetA(b uint8) {
	acc := cpu.r.Accumulator()
	cpu.set8(acc, b)
}

// Set sets the provided bit in the provided register.
func (cpu *CPU) Set(b uint8, r Register) {
	a, _ := cpu.r.Auxiliary(r)
	cpu.set8(a, b)
}

// SetHLDereference sets the provided bit in the contents of the memory location
// referenced by register HL.
func (cpu *CPU) SetHLDereference(b uint8) {
	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.m.Byte(hl)
	cpu.set8(&val, b)
	cpu.m.SetByte(hl, val)
}

/**
 * Common operations
 */

func (cpu *CPU) putRegisterIntoAddressInRegister(ar, vr Register) {
	address, _ := cpu.r.Paired(ar)
	cpu.putRegisterIntoMemory(vr, address)
}

func (cpu *CPU) putRegisterDereferenceIntoRegister(fr, tr Register) {
	address, _ := cpu.r.Paired(fr)
	val := cpu.m.Byte(address)
	t, _ := cpu.r.Auxiliary(tr)
	*t = val
}

func (cpu *CPU) putRegisterIntoMemory(r Register, address uint16) {
	v, _ := cpu.r.Auxiliary(r)
	cpu.m.SetByte(address, *v)
}

func (cpu *CPU) putMemoryIntoRegister(address uint16, r Register) {
	val := cpu.m.Byte(address)
	t, _ := cpu.r.Auxiliary(r)
	*t = val
}

func (cpu *CPU) incrementRegister(r Register) {
	cpu.r.IncrementPaired(r)
}

func (cpu *CPU) decrementRegister(r Register) {
	cpu.r.DecrementPaired(r)
}

func (cpu *CPU) offset(addr uint16) uint16 {
	return addr + 0xFF00
}

func (cpu *CPU) memByte(addr uint16) uint8 {
	val := cpu.m.Load(addr)

	cpu.c.AddT(1)

	return val
}

func (cpu *CPU) memImmediateByte() uint8 {
	pc := cpu.r.ProgramCounter()
	val := cpu.m.Load(*pc)

	*pc++
	cpu.c.AddT(1)

	return val
}

func (cpu *CPU) memStoreByte(addr uint16, b uint8) {
	cpu.m.Store(address, *acc)

	cpu.c.AddT(1)
}

func (cpu *CPU) memImmediateWord() uint16 {
	lsb := cpu.memImmediateByte()
	msb := cpu.memImmediateByte()
	return uint16(msb)<<8 | uint16(lsb)
}
