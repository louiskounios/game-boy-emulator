package cpu

import (
	"fmt"

	"github.com/loizoskounios/game-boy-emulator/mmu"
)

// CPU is the CPU.
type CPU struct {
	c   *Clock
	i   *instruction
	r   *Registers
	mmu *mmu.MemoryManagementUnit
}

// New returns a new CPU struct.
func New() *CPU {
	c := NewClock(0)
	i := &instruction{}
	r := NewRegisters()
	mmu := mmu.New()

	return &CPU{
		c:   c,
		i:   i,
		r:   r,
		mmu: mmu,
	}
}

// Rst resets the CPU to the start of execution.
func (cpu *CPU) Rst() {
	cpu = New()
}

// Dispatch loop.
func (cpu *CPU) Dispatch() {
	opcode := uint8(0)
	currInstructionSet := &instructions
	pc := cpu.r.ProgramCounter()

	for i := 0; i < 256; i++ {
		if opcode == 0xCB {
			currInstructionSet = &instructionsCB
		} else {
			currInstructionSet = &instructions
		}

		opcode = cpu.memByte(*pc)
		cpu.i = currInstructionSet[opcode]
		fmt.Println(cpu.i)
		cpu.i.execute(cpu)
	}
}

// Nop does nothing.
func (cpu *CPU) Nop() {
	cpu.r.IncrementProgramCounter(1)
	cpu.c.AddM(1)
}

// CB switches to the CB instruction set.
func (cpu *CPU) CB() {
	cpu.r.IncrementProgramCounter(1)
	cpu.c.AddM(1)
}

/**
 * 8-bit loads
 */

func (cpu *CPU) load8(val uint8, dst *uint8) {
	*dst = val

	cpu.c.AddM(1)
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

	cpu.c.AddM(1)
}

// LoadRIntoHL loads the contents of the provided register into the memory
// address specified by register HL.
func (cpu *CPU) LoadRIntoHL(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	hl, _ := cpu.r.Paired(RegisterHL)

	cpu.memStoreByte(hl, *a)

	cpu.c.AddM(1)
}

// LoadNIntoHL loads the contents of the memory specified by the 8-bit immediate
// operand into the memory address specified by register HL.
func (cpu *CPU) LoadNIntoHL() {
	cpu.r.IncrementProgramCounter(1)

	n := cpu.memImmediateByte()

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.memStoreByte(hl, n)

	cpu.c.AddM(1)
}

func (cpu *CPU) offsetC() uint16 {
	c, _ := cpu.r.Auxiliary(RegisterC)
	return offset(uint16(*c))
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

	cpu.c.AddM(1)
}

func (cpu *CPU) offsetImmediate() uint16 {
	return offset(uint16(cpu.memImmediateByte()))
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

	cpu.c.AddM(1)
}

// LoadNNIntoA loads the contents of the memory address specified by the 16-bit
// immediate operand into register A.
func (cpu *CPU) LoadNNIntoA() {
	cpu.r.IncrementProgramCounter(1)

	address := cpu.memImmediateWord()
	acc := cpu.r.Accumulator()
	cpu.memStoreByte(address, *acc)

	cpu.c.AddM(1)
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

	cpu.c.AddM(1)
}

// LoadAIntoDE loads the contents of the accumulator into the memory address
// specified by register DE.
func (cpu *CPU) LoadAIntoDE() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	de, _ := cpu.r.Paired(RegisterDE)

	cpu.memStoreByte(de, *acc)

	cpu.c.AddM(1)
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

func (cpu *CPU) load16(val uint16, rr Register) {
	cpu.r.SetPaired(rr, val)

	cpu.c.AddM(1)
}

// LoadNNIntoRR loads the 16-bit immediate operand into the provided paired
// register.
func (cpu *CPU) LoadNNIntoRR(rr Register) {
	cpu.r.IncrementProgramCounter(1)

	val := cpu.memImmediateWord()
	cpu.load16(val, rr)
}

// LoadHLIntoSP loads the contents of register HL into the stack pointer
// register.
func (cpu *CPU) LoadHLIntoSP() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	sp := cpu.r.StackPointer()
	*sp = hl

	cpu.c.AddM(2)
}

// Pushes the provided word onto the stack, then decrements the stack pointer
// by 2.
// The 8 most significant bits of the word are stored in Memory[SP-1].
// The 8 least significant bits of the word are stored in Memory[SP-2].
func (cpu *CPU) pushWordOntoStack(word uint16) {
	sp := cpu.r.StackPointer()
	*sp -= 2
	cpu.memStoreWord(*sp, word)
}

// PushAFOntoStack pushes the paired register AF onto the stack.
func (cpu *CPU) PushAFOntoStack() {
	cpu.r.IncrementProgramCounter(1)

	cpu.pushWordOntoStack(cpu.r.AF())

	cpu.c.AddM(2)
}

// PushRROntoStack pushes the contents of the provided paired register onto the stack.
func (cpu *CPU) PushRROntoStack(rr Register) {
	cpu.r.IncrementProgramCounter(1)

	word, _ := cpu.r.Paired(rr)
	cpu.pushWordOntoStack(word)

	cpu.c.AddM(2)
}

// Pops a word from the stack, then increments the stack pointer by 2.
// The 8 most significant bits of the word come from Memory[SP+1].
// The 8 least significant bits of the word come from Memory[SP].
func (cpu *CPU) popStack() uint16 {
	sp := cpu.r.StackPointer()
	val := cpu.memWord(*sp)
	*sp += 2

	return val
}

// PopStackIntoAF pops a word from the stack and loads it into paired
// register AF.
func (cpu *CPU) PopStackIntoAF() {
	cpu.r.IncrementProgramCounter(1)

	cpu.r.SetAF(cpu.popStack())

	cpu.c.AddM(1)
}

// PopStackIntoRR pops a word from the stack and loads it into the provided
// paired register.
func (cpu *CPU) PopStackIntoRR(to Register) {
	cpu.r.IncrementProgramCounter(1)

	cpu.r.SetPaired(to, cpu.popStack())

	cpu.c.AddM(1)
}

// LoadOffsetSPIntoHL loads the result of the addition of the stack pointer and
// the 8-bit operand, with the operand being treated as a signed integer in
// the range [-128, 127], into register HL. Flags are updated accordingly.
func (cpu *CPU) LoadOffsetSPIntoHL() {
	cpu.r.IncrementProgramCounter(1)

	sp := cpu.r.StackPointer()
	b := cpu.memImmediateByte()
	offsetSP := cpu.add16S8(*sp, b)
	cpu.r.SetPaired(RegisterHL, offsetSP)
}

// LoadSPIntoNN loads the stack pointer register into the memory address
// specified by the immediate 16-bit operand.
func (cpu *CPU) LoadSPIntoNN() {
	cpu.r.IncrementProgramCounter(1)

	address := cpu.memImmediateWord()
	sp := cpu.r.StackPointer()
	cpu.memStoreWord(address, *sp)

	cpu.c.AddM(1)
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

	cpu.c.AddM(1)

	return result
}

func (cpu *CPU) add8Helper(y uint8, useCarry bool, f func(Flag) error) {
	acc := cpu.r.Accumulator()
	*acc = cpu.add8(*acc, y, useCarry)
	f(FlagN)
}

// AddA adds the accumulator to itself, storing the result in the accumulator.
// Flags are updated accordingly.
func (cpu *CPU) AddA() {
	cpu.r.IncrementProgramCounter(1)

	cpu.add8Helper(*cpu.r.Accumulator(), false, cpu.r.ResetFlag)
}

// AddR adds the provided register to the accumulator, storing the result in the
// accumulator. Flags are updated accordingly.
func (cpu *CPU) AddR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.add8Helper(*a, false, cpu.r.ResetFlag)
}

// AddN adds the immediate byte to the accumulator, storing the result in the
// accumulator. Flags are updated accordingly.
func (cpu *CPU) AddN() {
	cpu.r.IncrementProgramCounter(1)

	cpu.add8Helper(cpu.memImmediateByte(), false, cpu.r.ResetFlag)
}

// AddHL adds the value stored in the memory location referenced by register HL
// to the accumulator, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) AddHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.add8Helper(cpu.memByte(hl), false, cpu.r.ResetFlag)
}

// AdcA adds the accumulator and the contents of the carry flag to the
// accumulator itself, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) AdcA() {
	cpu.r.IncrementProgramCounter(1)

	cpu.add8Helper(*cpu.r.Accumulator(), true, cpu.r.ResetFlag)
}

// AdcR adds the provided register and the contents of the carry flag to the
// accumulator, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) AdcR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.add8Helper(*a, true, cpu.r.ResetFlag)
}

// AdcN adds the immediate byte and the contents of the carry flag to the
// accumulator, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) AdcN() {
	cpu.r.IncrementProgramCounter(1)

	cpu.add8Helper(cpu.memImmediateByte(), true, cpu.r.ResetFlag)
}

// AdcHL adds the value stored in the memory location referenced by register HL
// and the contents of the carry flag to the accumulator, storing the result in
// the accumulator. Flags are updated accordingly.
func (cpu *CPU) AdcHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.add8Helper(cpu.memByte(hl), true, cpu.r.ResetFlag)
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
// accumulator. Flags are updated accordingly.
func (cpu *CPU) SubA() {
	cpu.r.IncrementProgramCounter(1)

	cpu.sub8Helper(*cpu.r.Accumulator(), false, cpu.r.SetFlag)
}

// SubR subtracts the provided register from the accumulator, storing the result
// in the accumulator. Flags are updated accordingly.
func (cpu *CPU) SubR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.sub8Helper(*a, false, cpu.r.SetFlag)
}

// SubN subtracts the immediate byte from the accumulator, storing the result in
// the accumulator. Flags are updated accordingly.
func (cpu *CPU) SubN() {
	cpu.r.IncrementProgramCounter(1)

	cpu.sub8Helper(cpu.memImmediateByte(), false, cpu.r.SetFlag)
}

// SubHL subtracts the value stored in the memory location referenced by
// register HL from the accumulator, storing the result in the accumulator.
// Flags are updated accordingly.
func (cpu *CPU) SubHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.sub8Helper(cpu.memByte(hl), false, cpu.r.SetFlag)
}

// SbcA subtracts the accumulator and the contents of the carry flag from
// itself, storing the result in the accumulator. Flags are updated accordingly.
func (cpu *CPU) SbcA() {
	cpu.r.IncrementProgramCounter(1)

	cpu.sub8Helper(*cpu.r.Accumulator(), true, cpu.r.SetFlag)
}

// SbcR subtracts the provided register and the contents of the carry flag from
// the accumulator, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) SbcR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.sub8Helper(*a, true, cpu.r.SetFlag)
}

// SbcN subtracts the immediate byte and the contents of the carry flag from the
// accumulator, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) SbcN() {
	cpu.r.IncrementProgramCounter(1)

	cpu.sub8Helper(cpu.memImmediateByte(), true, cpu.r.SetFlag)
}

// SbcHL subtracts the value stored in the memory location referenced by
// register HL and the contents of the carry flag from the accumulator,
// storing the result in the accumulator. Flags are updated accordingly.
func (cpu *CPU) SbcHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.sub8Helper(cpu.memByte(hl), true, cpu.r.SetFlag)
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

	cpu.c.AddM(1)

	return result
}

// AndA performs bitwise AND between the contents of the accumulator and itself,
// storing the result in the accumulator. Flags are updated accordingly.
func (cpu *CPU) AndA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.bitwise8Helper(*acc, cpu.and8)
}

// AndR performs bitwise AND between the contents of the accumulator and the
// provided register, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) AndR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.bitwise8Helper(*a, cpu.and8)
}

// AndN performs bitwise AND between the contents of the accumulator and the
// immediate byte, storing the result in the accumulator. Flags are updated
// accordingly
func (cpu *CPU) AndN() {
	cpu.r.IncrementProgramCounter(1)

	cpu.bitwise8Helper(cpu.memImmediateByte(), cpu.and8)
}

// AndHL performs bitwise AND between the contents of the accumulator and the
// value stored in the memory location referenced by register HL, storing the
// result in the accumulator. Flags are updated accordingly.
func (cpu *CPU) AndHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.bitwise8Helper(cpu.memByte(hl), cpu.and8)
}

func (cpu *CPU) xor8(x, y uint8) (result uint8) {
	result = x ^ y

	cpu.r.ResetFlag(FlagC)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, result == 0)

	cpu.c.AddM(1)

	return result
}

// XorA performs bitwise XOR between the contents of the accumulator and itself,
// storing the result in the accumulator. Flags are updated accordingly.
func (cpu *CPU) XorA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.bitwise8Helper(*acc, cpu.xor8)
}

// XorR performs bitwise XOR between the contents of the accumulator and the
// provided register, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) XorR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.bitwise8Helper(*a, cpu.xor8)
}

// XorN performs bitwise XOR between the contents of the accumulator and the
// immediate byte, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) XorN() {
	cpu.r.IncrementProgramCounter(1)

	cpu.bitwise8Helper(cpu.memImmediateByte(), cpu.xor8)
}

// XorHL performs bitwise XOR between the contents of the accumulator and the
// value stored in the memory location referenced by register HL, storing the
// result in the accumulator. Flags are updated accordingly.
func (cpu *CPU) XorHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.bitwise8Helper(cpu.memByte(hl), cpu.xor8)
}

func (cpu *CPU) or8(x, y uint8) (result uint8) {
	result = x | y

	cpu.r.ResetFlag(FlagC)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, result == 0)

	cpu.c.AddM(1)

	return result
}

// OrA performs bitwise OR between the contents of the accumulator and itself,
// storing the result in the accumulator. Flags are updated accordingly.
func (cpu *CPU) OrA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.bitwise8Helper(*acc, cpu.or8)
}

// OrR performs bitwise OR between the contents of the accumulator and the
// provided register, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) OrR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.bitwise8Helper(*a, cpu.or8)
}

// OrN performs bitwise OR between the contents of the accumulator and the
// immediate byte, storing the result in the accumulator. Flags are updated
// accordingly.
func (cpu *CPU) OrN() {
	cpu.r.IncrementProgramCounter(1)

	cpu.bitwise8Helper(cpu.memImmediateByte(), cpu.or8)
}

// OrHL performs bitwise OR between the contents of the accumulator and the
// value stored in the memory location referenced by register HL, storing the
// result in the accumulator. Flags are updated accordingly.
func (cpu *CPU) OrHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.bitwise8Helper(cpu.memByte(hl), cpu.or8)
}

// See cpu.sub8Helper function.
func (cpu *CPU) compare8Helper(y uint8) {
	acc := cpu.r.Accumulator()
	cpu.add8(*acc, ^y+1, false)
	cpu.r.SetFlag(FlagN)
}

// CompareA subtracts the accumulator from itself, discarding the result. Flags
// are updated accordingly.
func (cpu *CPU) CompareA() {
	cpu.r.IncrementProgramCounter(1)

	cpu.compare8Helper(*cpu.r.Accumulator())
}

// CompareR subtracts the provided register from the accumulator, discarding the
// result. Flags are updated accordingly.
func (cpu *CPU) CompareR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.compare8Helper(*a)
}

// CompareN subtracts the immediate byte from the accumulator, discarding the
// result. Flags are updated accordingly.
func (cpu *CPU) CompareN() {
	cpu.r.IncrementProgramCounter(1)

	cpu.compare8Helper(cpu.memImmediateByte())
}

// CompareHL subtracts the value stored in the memory location referenced by
// register HL from the accumulator, discarding the result. Flags are updated
// accordingly.
func (cpu *CPU) CompareHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.compare8Helper(cpu.memByte(hl))
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

	cpu.c.AddM(1)
}

// IncrementA increments the accumulator register by 1. Flags are updated
// accordingly.
func (cpu *CPU) IncrementA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.increment8Helper(acc)
}

// IncrementR increments the provided register by 1. Flags are updated
// accordingly.
func (cpu *CPU) IncrementR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.increment8Helper(a)
}

// IncrementHL increments the memory contents referenced by register HL by 1.
// Flags are updated accordingly.
func (cpu *CPU) IncrementHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.increment8Helper(&val)
	cpu.memStoreByte(hl, val)
}

func (cpu *CPU) decrement8Helper(x *uint8) {
	by := uint8(1)
	by = ^by + 1
	*x = cpu.increment8(*x, by)
	cpu.r.SetFlag(FlagN)

	cpu.c.AddM(1)
}

// DecrementA decrements the accumulator register by 1. Flags are updated
// accordingly.
func (cpu *CPU) DecrementA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.decrement8Helper(acc)
}

// DecrementR decrements the provided register by 1. Flags are updated
// accordingly.
func (cpu *CPU) DecrementR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.decrement8Helper(a)
}

// DecrementHL decrements the memory contents referenced by register HL by 1.
// Flags are updated accordingly.
func (cpu *CPU) DecrementHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.decrement8Helper(&val)
	cpu.memStoreByte(hl, val)
}

// DecimalAdjustA adjusts the contents of the accumulator following a binary
// addition / subtraction. It retroactively turns the previous operation into
// a BCD addition / subtraction. This is achieved by adding / subtracting 6
// to / from the accumulator's upper and/or lower nybble. The N, C and H flags
// are utilised to determine whether the correction must be added or subtracted
// (depending on if the previous operation was an addition or subtraction) and
// how / if each nybble should be affected (depending on if a carry or
// half-carry occurred in the previous operation).
//
// Adapted from: https://forums.nesdev.com/viewtopic.php?f=20&t=15944#p196282
func (cpu *CPU) DecimalAdjustA() {
	cpu.r.IncrementProgramCounter(1)

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

	cpu.c.AddM(1)
}

// ComplementA sets the accumulator to the one's complement of itself. Flags are
// updated accordingly.
func (cpu *CPU) ComplementA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	*acc = ^*acc

	cpu.r.SetFlag(FlagH)
	cpu.r.SetFlag(FlagN)

	cpu.c.AddM(1)
}

// ComplementCarryFlag toggles the carry flag. Flags are updated accordingly.
func (cpu *CPU) ComplementCarryFlag() {
	cpu.r.IncrementProgramCounter(1)

	cpu.r.ToggleFlag(FlagC)

	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)

	cpu.c.AddM(1)
}

// SetCarryFlag sets the carry flag. Flags are updated accordingly.
func (cpu *CPU) SetCarryFlag() {
	cpu.r.IncrementProgramCounter(1)

	cpu.r.SetFlag(FlagC)

	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)

	cpu.c.AddM(1)
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

	cpu.c.AddM(2)

	return result
}

func (cpu *CPU) add16Helper(y uint16) {
	hl, _ := cpu.r.Paired(RegisterHL)
	cpu.r.SetPaired(RegisterHL, cpu.add16(hl, y))
}

// AddRR adds the provided register to register HL, storing the result in
// register HL. Flags are updated accordingly.
func (cpu *CPU) AddRR(rr Register) {
	cpu.r.IncrementProgramCounter(1)

	pr, _ := cpu.r.Paired(rr)
	cpu.add16Helper(pr)
}

// AddSP adds the stack pointer register to register HL, storing the result in
// register HL. Flags are updated accordingly.
func (cpu *CPU) AddSP() {
	cpu.r.IncrementProgramCounter(1)

	sp := cpu.r.StackPointer()
	cpu.add16Helper(*sp)
}

// IncrementRR increments the provided register by 1.
func (cpu *CPU) IncrementRR(rr Register) {
	cpu.r.IncrementProgramCounter(1)

	cpu.r.IncrementPaired(rr)

	cpu.c.AddM(2)
}

// IncrementSP increments the stack pointer register by 1.
func (cpu *CPU) IncrementSP() {
	cpu.r.IncrementProgramCounter(1)

	sp := cpu.r.StackPointer()
	*sp++

	cpu.c.AddM(2)
}

// DecrementRR decrements the provided register by 1.
func (cpu *CPU) DecrementRR(rr Register) {
	cpu.r.IncrementProgramCounter(1)

	cpu.r.DecrementPaired(rr)

	cpu.c.AddM(2)
}

// DecrementSP decrements the stack pointer register by 1.
func (cpu *CPU) DecrementSP() {
	cpu.r.IncrementProgramCounter(1)

	sp := cpu.r.StackPointer()
	*sp--

	cpu.c.AddM(2)
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

// AddOffsetImmediateToSP adds the 8-bit immediate operand to the stack pointer
// register, with the operand being treated as a signed integer in the range
// [-128, 127]. Flags are updated accordingly.
func (cpu *CPU) AddOffsetImmediateToSP() {
	cpu.r.IncrementProgramCounter(1)

	sp := cpu.r.StackPointer()
	offsetSP := cpu.add16S8(*sp, cpu.memImmediateByte())
	*sp = offsetSP

	cpu.c.AddM(1)
}

/**
 * Jumps / calls
 */

func (cpu *CPU) shouldJump(flag Flag, isSet bool) bool {
	fs, _ := cpu.r.IsFlagSet(flag)
	return fs == isSet
}

// JumpHL loads the contents of paired register HL into the program counter.
func (cpu *CPU) JumpHL() {
	pc := cpu.r.ProgramCounter()
	hl, _ := cpu.r.Paired(RegisterHL)
	*pc = hl

	cpu.c.AddM(1)
}

// JumpOffset loads the result of the addition of the program counter and the
// 8-bit immediate operand, with the operand being treated as a signed integer
// in the range [-128, 127], into the program counter.
func (cpu *CPU) JumpOffset() {
	ib := cpu.memImmediateByte()
	pc := cpu.r.ProgramCounter()
	*pc = cpu.add16S8(*pc, ib)
}

// JumpOffsetConditionally loads the result of the addition of the program
// counter and the 8-bit immediate operand, with the operand being treated as a
// signed integer in the range [-128, 127], into the program counter. The
// condition is that the status of the provided flag must match the provided
// status.
func (cpu *CPU) JumpOffsetConditionally(flag Flag, isSet bool) {
	if cpu.shouldJump(flag, isSet) {
		cpu.JumpOffset()
	} else {
		cpu.r.IncrementProgramCounter(1)
		cpu.c.AddM(2)
	}
}

// JumpNN loads the 16-bit immediate operand into the program counter.
func (cpu *CPU) JumpNN() {
	pc := cpu.r.ProgramCounter()
	*pc = cpu.memImmediateWord()

	cpu.c.AddM(2)
}

// JumpNNConditionally loads the 16-bit immediate operand into the program
// counter if the status of the provided flag matches the provided status.
func (cpu *CPU) JumpNNConditionally(flag Flag, isSet bool) {
	if cpu.shouldJump(flag, isSet) {
		cpu.JumpNN()
	} else {
		cpu.r.IncrementProgramCounter(1)
		cpu.c.AddM(3)
	}
}

// CallNN pushes the program counter onto the stack, then loads the 16-bit
// immediate operand into the program counter.
func (cpu *CPU) CallNN() {
	pc := cpu.r.ProgramCounter()
	cpu.pushWordOntoStack(*pc)
	cpu.JumpNN()
}

// CallNNConditionally pushes the program counter onto the stack, then loads the
// 16-bit immediate operand into the program counter. The condition is that the
// status of the provided flag must match the provided status.
func (cpu *CPU) CallNNConditionally(flag Flag, isSet bool) {
	if cpu.shouldJump(flag, isSet) {
		cpu.CallNN()
	} else {
		cpu.r.IncrementProgramCounter(1)
		cpu.c.AddM(3)
	}
}

// Return loads a word popped from the stack into the program counter.
func (cpu *CPU) Return() {
	pc := cpu.r.ProgramCounter()
	*pc = cpu.popStack()

	cpu.c.AddM(2)
}

// ReturnConditionally loads a word popped from the stack into the program
// counter if the status of the provided flag matches the provided status.
func (cpu *CPU) ReturnConditionally(flag Flag, isSet bool) {
	if cpu.shouldJump(flag, isSet) {
		cpu.Return()
		cpu.c.AddM(1)
	} else {
		cpu.r.IncrementProgramCounter(1)
		cpu.c.AddM(2)
	}
}

// ReturnPostInterrupt is a placeholder.
// TODO: Revisit when emulating interrupts.
func (cpu *CPU) ReturnPostInterrupt() {
	cpu.Return()
}

// Restart pushes the program counter onto the stack, then loads the provided
// value into the program counter.
func (cpu *CPU) Restart(t uint8) {
	pc := cpu.r.ProgramCounter()
	cpu.pushWordOntoStack(*pc)
	*pc = uint16(t)

	cpu.c.AddM(2)
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

	cpu.c.AddM(1)
}

func (cpu *CPU) rotate8BothHelper(x *uint8, right bool) {
	var rotatedBitset bool
	*x, rotatedBitset = rotate8(*x, right)

	cpu.r.PutFlag(FlagC, rotatedBitset)
	cpu.r.ResetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, *x == 0)

	cpu.c.AddM(1)
}

// RLCA rotates the contents of the accumulator to the left. The MSB becomes the
// LSB and the carry flag. All other flags are reset.
func (cpu *CPU) RLCA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.rotate8BothHelper(acc, false)
	cpu.r.ResetFlag(FlagZ)
}

// RLA rotates the contents of the accumulator to the left. The MSB becomes the
// carry flag and the carry flag becomes the LSB. All other flags are reset.
func (cpu *CPU) RLA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.rotate8SwapHelper(acc, false)
	cpu.r.ResetFlag(FlagZ)
}

// RRCA rotates the contents of the accumulator to the right. The LSB becomes
// the MSB and the carry flag. All other flags are reset.
func (cpu *CPU) RRCA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.rotate8BothHelper(acc, true)
	cpu.r.ResetFlag(FlagZ)
}

// RRA rotates the contents of the accumulator to the right. The LSB becomes the
// carry flag and the carry flag becomes the MSB. All other flags are reset.
func (cpu *CPU) RRA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.rotate8SwapHelper(acc, true)
	cpu.r.ResetFlag(FlagZ)
}

// RLCACB rotates the contents of the accumulator to the left. The MSB becomes
// the LSB and the carry flag. Flags are updated accordingly.
func (cpu *CPU) RLCACB() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.rotate8BothHelper(acc, false)
}

// RLC rotates the contents of the provided register to the left. The MSB
// becomes the LSB and the carry flag. Flags are updated accordingly.
func (cpu *CPU) RLC(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.rotate8BothHelper(a, false)
}

// RLCHL rotates the contents of the memory location referenced by register HL
// to the left. The MSB becomes the the LSB and the carry flag. Flags are
// updated accordingly.
func (cpu *CPU) RLCHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.rotate8BothHelper(&val, false)
	cpu.memStoreByte(hl, val)
}

// RRCACB rotates the contents of the accumulator to the right. The LSB becomes
// the MSB and the carry flag. Flags are updated accordingly.
func (cpu *CPU) RRCACB() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.rotate8BothHelper(acc, true)
}

// RRC rotates the contents of the provided register to the right. The LSB
// becomes the MSB and the carry flag. Flags are updated accordingly.
func (cpu *CPU) RRC(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.rotate8BothHelper(a, true)
}

// RRCHL rotates the contents of the memory location referenced by register HL
// to the right. The LSB becomes the the MSB and the carry flag. Flags are
// updated accordingly.
func (cpu *CPU) RRCHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.rotate8BothHelper(&val, true)
	cpu.memStoreByte(hl, val)
}

// RLACB rotates the contents of the accumulator to the left. The MSB becomes
// the carry flag and the carry flag becomes the LSB. Flags are updated
// accordingly.
func (cpu *CPU) RLACB() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.rotate8SwapHelper(acc, false)
}

// RL rotates the contents of the provided register to the left. The MSB becomes
// the carry flag and the carry flag becomes the LSB. Flags are updated
// accordingly.
func (cpu *CPU) RL(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.rotate8SwapHelper(a, false)
}

// RLHL rotates the contents of the memory location referenced by register HL
// to the left. The MSB becomes the carry flag and the carry flag becomes the
// LSB. Flags are updated accordingly.
func (cpu *CPU) RLHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.rotate8SwapHelper(&val, false)
	cpu.memStoreByte(hl, val)
}

// RRACB rotates the contents of the accumulator to the right. The LSB becomes
// the carry flag and the carry flag becomes the MSB. Flags are updated
// accordingly.
func (cpu *CPU) RRACB() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.rotate8SwapHelper(acc, true)
}

// RR rotates the contents of the provided register to the right. The LSB
// becomes the carry flag and the carry flag becomes the MSB. Flags are updated
// accordingly.
func (cpu *CPU) RR(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.rotate8SwapHelper(a, true)
}

// RRHL rotates the contents of the memory location referenced by register HL
// to the right. The LSB becomes the carry flag and the carry flag becomes the
// MSB. Flags are updated accordingly.
func (cpu *CPU) RRHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.rotate8SwapHelper(&val, true)
	cpu.memStoreByte(hl, val)
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

	cpu.c.AddM(1)
}

// SLAA shifts the contents of the accumulator to the left. The MSB becomes the
// carry flag and the LSB is reset. Flags are updated accordingly.
func (cpu *CPU) SLAA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.shift8Helper(acc, false)
}

// SLA shifts the contents of the provided register to the left. The MSB becomes
// the carry flag and the LSB is reset. Flags are updated accordingly.
func (cpu *CPU) SLA(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.shift8Helper(a, false)
}

// SLAHL shifts the contents of the memory location referenced by register HL
// to the left. The MSB becomes the carry flag and the LSB is reset. Flags are
// updated accordingly.
func (cpu *CPU) SLAHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.shift8Helper(&val, false)
	cpu.memStoreByte(hl, val)
}

// SRAA shifts the contents of the accumulator to the right. The LSB becomes the
// carry flag and the MSB retains its value. Flags are updated accordingly.
func (cpu *CPU) SRAA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	bit7mask := *acc & 0x80
	cpu.shift8Helper(acc, true)
	*acc = *acc | bit7mask
}

// SRA shifts the contents of the provided register to the right. The LSB
// becomes the carry flag and the MSB retains its value. Flags are updated
// accordingly.
func (cpu *CPU) SRA(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	bit7mask := *a & 0x80
	cpu.shift8Helper(a, true)
	*a = *a | bit7mask
}

// SRAHL shifts the contents of the memory location referenced by register HL
// to the right. The LSB becomes the carry flag and the MSB retains its value.
// Flags are updated accordingly.
func (cpu *CPU) SRAHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	bit7mask := val & 0x80
	cpu.shift8Helper(&val, true)
	val |= bit7mask
	cpu.memStoreByte(hl, val)
}

// SRLA shifts the contents of the accumulator to the right. The LSB becomes the
// carry flag and the MSB is reset. Flags are updated accordingly.
func (cpu *CPU) SRLA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.shift8Helper(acc, true)
}

// SRL shifts the contents of the provided register to the right. The LSB
// becomes the carry flag and the MSB is reset. Flags are updated accordingly.
func (cpu *CPU) SRL(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.shift8Helper(a, true)
}

// SRLHL shifts the contents of the memory location referenced by register HL
// to the right. The LSB becomes the carry flag and the MSB is reset. Flags are
// updated accordingly.
func (cpu *CPU) SRLHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.shift8Helper(&val, true)
	cpu.memStoreByte(hl, val)
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

	cpu.c.AddM(1)
}

// SwapA swaps the accumulator nibbles. Flags are updated accordingly.
func (cpu *CPU) SwapA() {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.swap8Helper(acc)
}

// Swap swaps the provided register's nibbles. Flags are updated accordingly.
func (cpu *CPU) Swap(r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.swap8Helper(a)
}

// SwapHL swaps the nibbles of the contents of the memory location referenced by
// register HL. Flags are updated accordingly.
func (cpu *CPU) SwapHL() {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.swap8Helper(&val)
	cpu.memStoreByte(hl, val)
}

func (cpu *CPU) bit8(x, b uint8) {
	isSet := (x>>b)&0x01 == 1

	cpu.r.SetFlag(FlagH)
	cpu.r.ResetFlag(FlagN)
	cpu.r.PutFlag(FlagZ, !isSet)

	cpu.c.AddM(1)
}

// BitA sets the Z flag to the complement of the contents of the provided bit in
// the accumulator. The H and N flags are set and reset respectively.
func (cpu *CPU) BitA(b uint8) {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.bit8(*acc, b)
}

// Bit sets the Z flag to the complement of the contents of the provided bit in
// the provided register. The H and N flags are set and reset respectively.
func (cpu *CPU) Bit(b uint8, r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.bit8(*a, b)
}

// BitHL sets the Z flag to the complement of the contents of the provided bit
// in the contents of the memory location referenced by register HL. The H and
// N flags are set and reset respectively.
func (cpu *CPU) BitHL(b uint8) {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.bit8(val, b)
}

func (cpu *CPU) reset8(x *uint8, b uint8) {
	*x &^= (1 << b)

	cpu.c.AddM(1)
}

// ResetA resets the provided bit in the accumulator.
func (cpu *CPU) ResetA(b uint8) {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.reset8(acc, b)
}

// Reset resets the provided bit in the provided register.
func (cpu *CPU) Reset(b uint8, r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.reset8(a, b)
}

// ResetHL resets the provided bit in the contents of the memory location
// referenced by register HL.
func (cpu *CPU) ResetHL(b uint8) {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.reset8(&val, b)
	cpu.memStoreByte(hl, val)
}

func (cpu *CPU) set8(x *uint8, b uint8) {
	*x |= (1 << b)

	cpu.c.AddM(1)
}

// SetA sets the provided bit in the accumulator.
func (cpu *CPU) SetA(b uint8) {
	cpu.r.IncrementProgramCounter(1)

	acc := cpu.r.Accumulator()
	cpu.set8(acc, b)
}

// Set sets the provided bit in the provided register.
func (cpu *CPU) Set(b uint8, r Register) {
	cpu.r.IncrementProgramCounter(1)

	a, _ := cpu.r.Auxiliary(r)
	cpu.set8(a, b)
}

// SetHL sets the provided bit in the contents of the memory location referenced
// by register HL.
func (cpu *CPU) SetHL(b uint8) {
	cpu.r.IncrementProgramCounter(1)

	hl, _ := cpu.r.Paired(RegisterHL)
	val := cpu.memByte(hl)
	cpu.set8(&val, b)
	cpu.memStoreByte(hl, val)
}

/**
 * Common operations
 */

func offset(addr uint16) uint16 {
	return addr + 0xFF00
}

func (cpu *CPU) incrementRegister(r Register) {
	cpu.r.IncrementPaired(r)
}

func (cpu *CPU) decrementRegister(r Register) {
	cpu.r.DecrementPaired(r)
}

func (cpu *CPU) memByte(addr uint16) uint8 {
	val := cpu.mmu.Load(addr)

	cpu.c.AddM(1)

	return val
}

func (cpu *CPU) memImmediateByte() uint8 {
	pc := cpu.r.ProgramCounter()
	val := cpu.mmu.Load(*pc)

	*pc++
	cpu.c.AddM(1)

	return val
}

func (cpu *CPU) memStoreByte(addr uint16, b uint8) {
	cpu.mmu.Store(addr, b)

	cpu.c.AddM(1)
}

func (cpu *CPU) memWord(addr uint16) uint16 {
	lsb := cpu.memByte(addr)
	msb := cpu.memByte(addr + 1)

	return uint16(msb)<<8 | uint16(lsb)
}

func (cpu *CPU) memImmediateWord() uint16 {
	lsb := cpu.memImmediateByte()
	msb := cpu.memImmediateByte()
	return uint16(msb)<<8 | uint16(lsb)
}

func (cpu *CPU) memStoreWord(addr uint16, w uint16) {
	lsb := uint8(w)
	msb := uint8(w >> 8)
	cpu.memStoreByte(addr, lsb)
	cpu.memStoreByte(addr+1, msb)
}
