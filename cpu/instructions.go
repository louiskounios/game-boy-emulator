package cpu

import (
	"fmt"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers"
)

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
	/**
	 * 8-bit loads
	 */

	// Register (A, B, C, D, E, H, L) -> Register (B)
	0x47: &instruction{0x47, 1, "LD B,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.B) }},
	0x40: &instruction{0x40, 1, "LD B,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.B) }},
	0x41: &instruction{0x41, 1, "LD B,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.B) }},
	0x42: &instruction{0x42, 1, "LD B,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.B) }},
	0x43: &instruction{0x43, 1, "LD B,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.B) }},
	0x44: &instruction{0x44, 1, "LD B,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.B) }},
	0x45: &instruction{0x45, 1, "LD B,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.B) }},

	// Register (A, B, C, D, E, H, L) -> Register (C)
	0x4F: &instruction{0x4F, 1, "LD C,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.C) }},
	0x48: &instruction{0x48, 1, "LD C,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.C) }},
	0x49: &instruction{0x49, 1, "LD C,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.C) }},
	0x4A: &instruction{0x4A, 1, "LD C,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.C) }},
	0x4B: &instruction{0x4B, 1, "LD C,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.C) }},
	0x4C: &instruction{0x4C, 1, "LD C,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.C) }},
	0x4D: &instruction{0x4D, 1, "LD C,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.C) }},

	// Register (A, B, C, D, E, H, L) -> Register (D)
	0x57: &instruction{0x57, 1, "LD D,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.D) }},
	0x50: &instruction{0x50, 1, "LD D,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.D) }},
	0x51: &instruction{0x51, 1, "LD D,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.D) }},
	0x52: &instruction{0x52, 1, "LD D,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.D) }},
	0x53: &instruction{0x53, 1, "LD D,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.D) }},
	0x54: &instruction{0x54, 1, "LD D,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.D) }},
	0x55: &instruction{0x55, 1, "LD D,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.D) }},

	// Register (A, B, C, D, E, H, L) -> Register (E)
	0x5F: &instruction{0x5F, 1, "LD E,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.E) }},
	0x58: &instruction{0x58, 1, "LD E,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.E) }},
	0x59: &instruction{0x59, 1, "LD E,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.E) }},
	0x5A: &instruction{0x5A, 1, "LD E,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.E) }},
	0x5B: &instruction{0x5B, 1, "LD E,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.E) }},
	0x5C: &instruction{0x5C, 1, "LD E,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.E) }},
	0x5D: &instruction{0x5D, 1, "LD E,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.E) }},

	// Register (A, B, C, D, E, H, L) -> Register (H)
	0x67: &instruction{0x67, 1, "LD H,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.H) }},
	0x60: &instruction{0x60, 1, "LD H,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.H) }},
	0x61: &instruction{0x61, 1, "LD H,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.H) }},
	0x62: &instruction{0x62, 1, "LD H,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.H) }},
	0x63: &instruction{0x63, 1, "LD H,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.H) }},
	0x64: &instruction{0x64, 1, "LD H,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.H) }},
	0x65: &instruction{0x65, 1, "LD H,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.H) }},

	// Register (A, B, C, D, E, H, L) -> Register (L)
	0x6F: &instruction{0x6F, 1, "LD L,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.L) }},
	0x68: &instruction{0x68, 1, "LD L,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.L) }},
	0x69: &instruction{0x69, 1, "LD L,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.L) }},
	0x6A: &instruction{0x6A, 1, "LD L,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.L) }},
	0x6B: &instruction{0x6B, 1, "LD L,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.L) }},
	0x6C: &instruction{0x6C, 1, "LD L,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.L) }},
	0x6D: &instruction{0x6D, 1, "LD L,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.L) }},

	// Register (A, B, C, D, E, H, L) -> Register (A)
	0x7F: &instruction{0x7F, 1, "LD A,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.A) }},
	0x78: &instruction{0x78, 1, "LD A,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.A) }},
	0x79: &instruction{0x79, 1, "LD A,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.A) }},
	0x7A: &instruction{0x7A, 1, "LD A,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.A) }},
	0x7B: &instruction{0x7B, 1, "LD A,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.A) }},
	0x7C: &instruction{0x7C, 1, "LD A,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.A) }},
	0x7D: &instruction{0x7D, 1, "LD A,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.A) }},

	// Register (A) -> Memory[Memory[PC and PC+1]]
	0xEA: &instruction{0xEA, 4, "LD (a16),A", func(cpu *CPU) { cpu.PutAIntoNNAddress() }},

	// Register (A) -> Memory[BC]
	0x02: &instruction{0x02, 2, "LD (BC),A", func(cpu *CPU) { cpu.PutAIntoBCAddress() }},

	// Register (A) -> Memory[DE]
	0x12: &instruction{0x12, 2, "LD (DE),A", func(cpu *CPU) { cpu.PutAIntoDEAddress() }},

	// Register (A, B, C, D, E, H, L) -> Memory[HL]
	0x77: &instruction{0x77, 2, "LD (HL),A", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.A) }},
	0x70: &instruction{0x70, 2, "LD (HL),B", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.B) }},
	0x71: &instruction{0x71, 2, "LD (HL),C", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.C) }},
	0x72: &instruction{0x72, 2, "LD (HL),D", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.D) }},
	0x73: &instruction{0x73, 2, "LD (HL),E", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.E) }},
	0x74: &instruction{0x74, 2, "LD (HL),H", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.H) }},
	0x75: &instruction{0x75, 2, "LD (HL),L", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.L) }},

	// Register (A) -> Memory[HL++]
	0x22: &instruction{0x22, 2, "LD (HL+),A", func(cpu *CPU) { cpu.PutAIntoHLAddressThenIncrementHL() }},

	// Register (A) -> Memory[HL--]
	0x32: &instruction{0x32, 2, "LD (HL-),A", func(cpu *CPU) { cpu.PutAIntoHLAddressThenDecrementHL() }},

	// Register (A) -> Memory[C+0xFF0]
	0xE2: &instruction{0xE2, 2, "LD (C),A", func(cpu *CPU) { cpu.PutAIntoOffsetCAddress() }},

	// Register (A) -> Memory[Memory[PC]+0xFF00]
	0xE0: &instruction{0xE0, 3, "LD (a8),A", func(cpu *CPU) { cpu.PutAIntoOffsetImmediateAddress() }},

	// Memory[PC] -> Register (B, C, D, E, H, L, A)
	0x3E: &instruction{0x3E, 2, "LD A,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.A) }},
	0x06: &instruction{0x06, 2, "LD B,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.B) }},
	0x0E: &instruction{0x0E, 2, "LD C,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.C) }},
	0x16: &instruction{0x16, 2, "LD D,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.D) }},
	0x1E: &instruction{0x1E, 2, "LD E,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.E) }},
	0x26: &instruction{0x26, 2, "LD H,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.H) }},
	0x2E: &instruction{0x2E, 2, "LD L,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.L) }},

	// Memory[Memory[PC and PC+1]] -> Register (A)
	0xFA: &instruction{0xFA, 4, "LD A,(a16)", func(cpu *CPU) { cpu.PutNNDereferenceIntoA() }},

	// Memory[C+0xFF00] -> Register A
	0xF2: &instruction{0xF2, 2, "LD A,(C)", func(cpu *CPU) { cpu.PutOffsetCDereferenceIntoA() }},

	// Memory[Memory[PC]+0xFF00] -> Register (A)
	0xF0: &instruction{0xF0, 3, "LDH A,(a8)", func(cpu *CPU) { cpu.PutOffsetImmediateDereferenceIntoA() }},

	// Memory[BC] -> Register (A)
	0x0A: &instruction{0x0A, 2, "LD A,(BC)", func(cpu *CPU) { cpu.PutBCDereferenceIntoA() }},

	// Memory[DE] -> Register (A)
	0x1A: &instruction{0x1A, 2, "LD A,(DE)", func(cpu *CPU) { cpu.PutDEDereferenceIntoA() }},

	// Memory[HL] -> Register (B, C, D, E, H, L, A)
	0x7E: &instruction{0x7E, 2, "LD A,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.A) }},
	0x46: &instruction{0x46, 2, "LD B,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.B) }},
	0x4E: &instruction{0x4E, 2, "LD C,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.C) }},
	0x56: &instruction{0x56, 2, "LD D,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.D) }},
	0x5E: &instruction{0x5E, 2, "LD E,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.E) }},
	0x66: &instruction{0x66, 2, "LD H,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.H) }},
	0x6E: &instruction{0x6E, 2, "LD L,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.L) }},

	// Memory[HL++] -> Register (A)
	0x2A: &instruction{0x2A, 2, "LD A,(HL+)", func(cpu *CPU) { cpu.PutHLDereferenceIntoAThenIncrementHL() }},

	// Memory[HL--] -> Register (A)
	0x3A: &instruction{0x3A, 2, "LD A,(HL-)", func(cpu *CPU) { cpu.PutHLDereferenceIntoAThenDecrementHL() }},

	// Memory[PC] -> Memory[HL]
	0x36: &instruction{0x36, 3, "LD (HL),d8", func(cpu *CPU) { cpu.PutNDereferenceIntoHLAddress() }},

	/**
	 * 16-bit loads
	 */

	// Register (HL) -> Register (SP)
	0xF9: &instruction{0xF9, 2, "LD SP,HL", func(cpu *CPU) { cpu.PutHLIntoSP() }},

	// Register (SP) -> Memory[PC and PC+1]
	0x08: &instruction{0x08, 5, "LD (a16),SP", func(cpu *CPU) { cpu.PutSPIntoNNAddress() }},

	// Register (AF, BC, DE, HL) -> Memory[--SP and --SP]
	0xF5: &instruction{0xF5, 4, "PUSH AF", func(cpu *CPU) { cpu.PushAFOntoStack() }},
	0xC5: &instruction{0xC5, 4, "PUSH BC", func(cpu *CPU) { cpu.PushRROntoStack(registers.BC) }},
	0xD5: &instruction{0xD5, 4, "PUSH DE", func(cpu *CPU) { cpu.PushRROntoStack(registers.DE) }},
	0xE5: &instruction{0xE5, 4, "PUSH HL", func(cpu *CPU) { cpu.PushRROntoStack(registers.HL) }},

	// Memory[SP++ and SP++] -> Register (AF, BC, DE, HL)
	0xF1: &instruction{0xF1, 3, "POP AF", func(cpu *CPU) { cpu.PopStackIntoAF() }},
	0xC1: &instruction{0xC1, 3, "POP BC", func(cpu *CPU) { cpu.PopStackIntoRR(registers.BC) }},
	0xD1: &instruction{0xD1, 3, "POP DE", func(cpu *CPU) { cpu.PopStackIntoRR(registers.DE) }},
	0xE1: &instruction{0xE1, 3, "POP HL", func(cpu *CPU) { cpu.PopStackIntoRR(registers.HL) }},

	// Memory[PC and PC+1] -> Register (BC, DE, HL, SP)
	0x01: &instruction{0x01, 3, "LD BC,d16", func(cpu *CPU) { cpu.PutNNIntoRR(registers.BC) }},
	0x11: &instruction{0x11, 3, "LD DE,d16", func(cpu *CPU) { cpu.PutNNIntoRR(registers.DE) }},
	0x21: &instruction{0x21, 3, "LD HL,d16", func(cpu *CPU) { cpu.PutNNIntoRR(registers.HL) }},
	0x31: &instruction{0x31, 3, "LD SP,d16", func(cpu *CPU) { cpu.PutNNIntoRR(registers.SP) }},

	// Memory[PC] + Register (SP) -> Register (HL)
	0xF8: &instruction{0xF8, 3, "LD HL,SP+r8", func(cpu *CPU) { cpu.PutOffsetSPIntoHL() }},

	/**
	 * 8-bit arithmetic / logical operations
	 */

	// Register (A) <- Register (A) + Register (A, B, C, D, E, H, L)
	0x87: &instruction{0x87, 1, "ADD A,A", func(cpu *CPU) { cpu.AddA() }},
	0x80: &instruction{0x80, 1, "ADD A,B", func(cpu *CPU) { cpu.AddR(registers.B) }},
	0x81: &instruction{0x81, 1, "ADD A,C", func(cpu *CPU) { cpu.AddR(registers.C) }},
	0x82: &instruction{0x82, 1, "ADD A,D", func(cpu *CPU) { cpu.AddR(registers.D) }},
	0x83: &instruction{0x83, 1, "ADD A,E", func(cpu *CPU) { cpu.AddR(registers.E) }},
	0x84: &instruction{0x84, 1, "ADD A,H", func(cpu *CPU) { cpu.AddR(registers.H) }},
	0x85: &instruction{0x85, 1, "ADD A,L", func(cpu *CPU) { cpu.AddR(registers.L) }},

	// Register (A) <- Register (A) + Memory[HL]
	0x86: &instruction{0x86, 2, "ADD A,(HL)", func(cpu *CPU) { cpu.AddHLDereference() }},

	// Register (A) <- Register (A) + Memory[PC]
	0xC6: &instruction{0xC6, 2, "ADD A,d8", func(cpu *CPU) { cpu.AddN() }},

	// Register (A) <- Register (A) + Register (A, B, C, D, E, H, L) + Flag (C)
	0x8F: &instruction{0x8F, 1, "ADC A,A", func(cpu *CPU) { cpu.AdcA() }},
	0x88: &instruction{0x88, 1, "ADC A,B", func(cpu *CPU) { cpu.AdcR(registers.B) }},
	0x89: &instruction{0x89, 1, "ADC A,C", func(cpu *CPU) { cpu.AdcR(registers.C) }},
	0x8A: &instruction{0x8A, 1, "ADC A,D", func(cpu *CPU) { cpu.AdcR(registers.D) }},
	0x8B: &instruction{0x8B, 1, "ADC A,E", func(cpu *CPU) { cpu.AdcR(registers.E) }},
	0x8C: &instruction{0x8C, 1, "ADC A,H", func(cpu *CPU) { cpu.AdcR(registers.H) }},
	0x8D: &instruction{0x8D, 1, "ADC A,L", func(cpu *CPU) { cpu.AdcR(registers.L) }},

	// Register (A) <- Register (A) + Memory[HL] + Flag (C)
	0x8E: &instruction{0x8E, 2, "ADC A,(HL)", func(cpu *CPU) { cpu.AdcHLDereference() }},

	// Register (A) <- Register (A) + Memory[PC] + Flag (C)
	0xCE: &instruction{0xCE, 2, "ADC A,d8", func(cpu *CPU) { cpu.AdcN() }},

	// Register (A) <- Register (A) - Register (A, B, C, D, E, H, L)
	0x97: &instruction{0x97, 1, "SUB A", func(cpu *CPU) { cpu.SubA() }},
	0x90: &instruction{0x90, 1, "SUB B", func(cpu *CPU) { cpu.SubR(registers.B) }},
	0x91: &instruction{0x91, 1, "SUB C", func(cpu *CPU) { cpu.SubR(registers.C) }},
	0x92: &instruction{0x92, 1, "SUB D", func(cpu *CPU) { cpu.SubR(registers.D) }},
	0x93: &instruction{0x93, 1, "SUB E", func(cpu *CPU) { cpu.SubR(registers.E) }},
	0x94: &instruction{0x94, 1, "SUB H", func(cpu *CPU) { cpu.SubR(registers.H) }},
	0x95: &instruction{0x95, 1, "SUB L", func(cpu *CPU) { cpu.SubR(registers.L) }},

	// Register (A) <- Register (A) - Memory[HL]
	0x96: &instruction{0x96, 2, "SUB (HL)", func(cpu *CPU) { cpu.SubHLDereference() }},

	// Register (A) <- Register (A) - Memory[PC]
	0xD6: &instruction{0xD6, 2, "SUB d8", func(cpu *CPU) { cpu.SubN() }},

	// Register (A) <- Register (A) - Register (A, B, C, D, E, H, L) - Flag (C)
	0x9F: &instruction{0x9F, 1, "SBC A,A", func(cpu *CPU) { cpu.SbcA() }},
	0x98: &instruction{0x98, 1, "SBC A,B", func(cpu *CPU) { cpu.SbcR(registers.B) }},
	0x99: &instruction{0x99, 1, "SBC A,C", func(cpu *CPU) { cpu.SbcR(registers.C) }},
	0x9A: &instruction{0x9A, 1, "SBC A,D", func(cpu *CPU) { cpu.SbcR(registers.D) }},
	0x9B: &instruction{0x9B, 1, "SBC A,E", func(cpu *CPU) { cpu.SbcR(registers.E) }},
	0x9C: &instruction{0x9C, 1, "SBC A,H", func(cpu *CPU) { cpu.SbcR(registers.H) }},
	0x9D: &instruction{0x9D, 1, "SBC A,L", func(cpu *CPU) { cpu.SbcR(registers.L) }},

	// Register (A) <- Register (A) - Memory[HL] - Flag (C)
	0x9E: &instruction{0x9E, 2, "SBC A,(HL)", func(cpu *CPU) { cpu.SbcHLDereference() }},

	// Register (A) <- Register (A) - Memory[PC] - Flag (C)
	0xDE: &instruction{0xDE, 2, "SBC A,d8", func(cpu *CPU) { cpu.SbcN() }},
}
