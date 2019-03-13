package cpu

import (
	"fmt"
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
	0x47: &instruction{0x47, 1, "LD B,A", func(cpu *CPU) { cpu.PutRIntoR(RegisterA, RegisterB) }},
	0x40: &instruction{0x40, 1, "LD B,B", func(cpu *CPU) { cpu.PutRIntoR(RegisterB, RegisterB) }},
	0x41: &instruction{0x41, 1, "LD B,C", func(cpu *CPU) { cpu.PutRIntoR(RegisterC, RegisterB) }},
	0x42: &instruction{0x42, 1, "LD B,D", func(cpu *CPU) { cpu.PutRIntoR(RegisterD, RegisterB) }},
	0x43: &instruction{0x43, 1, "LD B,E", func(cpu *CPU) { cpu.PutRIntoR(RegisterE, RegisterB) }},
	0x44: &instruction{0x44, 1, "LD B,H", func(cpu *CPU) { cpu.PutRIntoR(RegisterH, RegisterB) }},
	0x45: &instruction{0x45, 1, "LD B,L", func(cpu *CPU) { cpu.PutRIntoR(RegisterL, RegisterB) }},

	// Register (A, B, C, D, E, H, L) -> Register (C)
	0x4F: &instruction{0x4F, 1, "LD C,A", func(cpu *CPU) { cpu.PutRIntoR(RegisterA, RegisterC) }},
	0x48: &instruction{0x48, 1, "LD C,B", func(cpu *CPU) { cpu.PutRIntoR(RegisterB, RegisterC) }},
	0x49: &instruction{0x49, 1, "LD C,C", func(cpu *CPU) { cpu.PutRIntoR(RegisterC, RegisterC) }},
	0x4A: &instruction{0x4A, 1, "LD C,D", func(cpu *CPU) { cpu.PutRIntoR(RegisterD, RegisterC) }},
	0x4B: &instruction{0x4B, 1, "LD C,E", func(cpu *CPU) { cpu.PutRIntoR(RegisterE, RegisterC) }},
	0x4C: &instruction{0x4C, 1, "LD C,H", func(cpu *CPU) { cpu.PutRIntoR(RegisterH, RegisterC) }},
	0x4D: &instruction{0x4D, 1, "LD C,L", func(cpu *CPU) { cpu.PutRIntoR(RegisterL, RegisterC) }},

	// Register (A, B, C, D, E, H, L) -> Register (D)
	0x57: &instruction{0x57, 1, "LD D,A", func(cpu *CPU) { cpu.PutRIntoR(RegisterA, RegisterD) }},
	0x50: &instruction{0x50, 1, "LD D,B", func(cpu *CPU) { cpu.PutRIntoR(RegisterB, RegisterD) }},
	0x51: &instruction{0x51, 1, "LD D,C", func(cpu *CPU) { cpu.PutRIntoR(RegisterC, RegisterD) }},
	0x52: &instruction{0x52, 1, "LD D,D", func(cpu *CPU) { cpu.PutRIntoR(RegisterD, RegisterD) }},
	0x53: &instruction{0x53, 1, "LD D,E", func(cpu *CPU) { cpu.PutRIntoR(RegisterE, RegisterD) }},
	0x54: &instruction{0x54, 1, "LD D,H", func(cpu *CPU) { cpu.PutRIntoR(RegisterH, RegisterD) }},
	0x55: &instruction{0x55, 1, "LD D,L", func(cpu *CPU) { cpu.PutRIntoR(RegisterL, RegisterD) }},

	// Register (A, B, C, D, E, H, L) -> Register (E)
	0x5F: &instruction{0x5F, 1, "LD E,A", func(cpu *CPU) { cpu.PutRIntoR(RegisterA, RegisterE) }},
	0x58: &instruction{0x58, 1, "LD E,B", func(cpu *CPU) { cpu.PutRIntoR(RegisterB, RegisterE) }},
	0x59: &instruction{0x59, 1, "LD E,C", func(cpu *CPU) { cpu.PutRIntoR(RegisterC, RegisterE) }},
	0x5A: &instruction{0x5A, 1, "LD E,D", func(cpu *CPU) { cpu.PutRIntoR(RegisterD, RegisterE) }},
	0x5B: &instruction{0x5B, 1, "LD E,E", func(cpu *CPU) { cpu.PutRIntoR(RegisterE, RegisterE) }},
	0x5C: &instruction{0x5C, 1, "LD E,H", func(cpu *CPU) { cpu.PutRIntoR(RegisterH, RegisterE) }},
	0x5D: &instruction{0x5D, 1, "LD E,L", func(cpu *CPU) { cpu.PutRIntoR(RegisterL, RegisterE) }},

	// Register (A, B, C, D, E, H, L) -> Register (H)
	0x67: &instruction{0x67, 1, "LD H,A", func(cpu *CPU) { cpu.PutRIntoR(RegisterA, RegisterH) }},
	0x60: &instruction{0x60, 1, "LD H,B", func(cpu *CPU) { cpu.PutRIntoR(RegisterB, RegisterH) }},
	0x61: &instruction{0x61, 1, "LD H,C", func(cpu *CPU) { cpu.PutRIntoR(RegisterC, RegisterH) }},
	0x62: &instruction{0x62, 1, "LD H,D", func(cpu *CPU) { cpu.PutRIntoR(RegisterD, RegisterH) }},
	0x63: &instruction{0x63, 1, "LD H,E", func(cpu *CPU) { cpu.PutRIntoR(RegisterE, RegisterH) }},
	0x64: &instruction{0x64, 1, "LD H,H", func(cpu *CPU) { cpu.PutRIntoR(RegisterH, RegisterH) }},
	0x65: &instruction{0x65, 1, "LD H,L", func(cpu *CPU) { cpu.PutRIntoR(RegisterL, RegisterH) }},

	// Register (A, B, C, D, E, H, L) -> Register (L)
	0x6F: &instruction{0x6F, 1, "LD L,A", func(cpu *CPU) { cpu.PutRIntoR(RegisterA, RegisterL) }},
	0x68: &instruction{0x68, 1, "LD L,B", func(cpu *CPU) { cpu.PutRIntoR(RegisterB, RegisterL) }},
	0x69: &instruction{0x69, 1, "LD L,C", func(cpu *CPU) { cpu.PutRIntoR(RegisterC, RegisterL) }},
	0x6A: &instruction{0x6A, 1, "LD L,D", func(cpu *CPU) { cpu.PutRIntoR(RegisterD, RegisterL) }},
	0x6B: &instruction{0x6B, 1, "LD L,E", func(cpu *CPU) { cpu.PutRIntoR(RegisterE, RegisterL) }},
	0x6C: &instruction{0x6C, 1, "LD L,H", func(cpu *CPU) { cpu.PutRIntoR(RegisterH, RegisterL) }},
	0x6D: &instruction{0x6D, 1, "LD L,L", func(cpu *CPU) { cpu.PutRIntoR(RegisterL, RegisterL) }},

	// Register (A, B, C, D, E, H, L) -> Register (A)
	0x7F: &instruction{0x7F, 1, "LD A,A", func(cpu *CPU) { cpu.PutRIntoR(RegisterA, RegisterA) }},
	0x78: &instruction{0x78, 1, "LD A,B", func(cpu *CPU) { cpu.PutRIntoR(RegisterB, RegisterA) }},
	0x79: &instruction{0x79, 1, "LD A,C", func(cpu *CPU) { cpu.PutRIntoR(RegisterC, RegisterA) }},
	0x7A: &instruction{0x7A, 1, "LD A,D", func(cpu *CPU) { cpu.PutRIntoR(RegisterD, RegisterA) }},
	0x7B: &instruction{0x7B, 1, "LD A,E", func(cpu *CPU) { cpu.PutRIntoR(RegisterE, RegisterA) }},
	0x7C: &instruction{0x7C, 1, "LD A,H", func(cpu *CPU) { cpu.PutRIntoR(RegisterH, RegisterA) }},
	0x7D: &instruction{0x7D, 1, "LD A,L", func(cpu *CPU) { cpu.PutRIntoR(RegisterL, RegisterA) }},

	// Register (A) -> Memory[Memory[PC and PC+1]]
	0xEA: &instruction{0xEA, 4, "LD (a16),A", func(cpu *CPU) { cpu.PutAIntoNNAddress() }},

	// Register (A) -> Memory[BC]
	0x02: &instruction{0x02, 2, "LD (BC),A", func(cpu *CPU) { cpu.PutAIntoBCAddress() }},

	// Register (A) -> Memory[DE]
	0x12: &instruction{0x12, 2, "LD (DE),A", func(cpu *CPU) { cpu.PutAIntoDEAddress() }},

	// Register (A, B, C, D, E, H, L) -> Memory[HL]
	0x77: &instruction{0x77, 2, "LD (HL),A", func(cpu *CPU) { cpu.PutRIntoHLAddress(RegisterA) }},
	0x70: &instruction{0x70, 2, "LD (HL),B", func(cpu *CPU) { cpu.PutRIntoHLAddress(RegisterB) }},
	0x71: &instruction{0x71, 2, "LD (HL),C", func(cpu *CPU) { cpu.PutRIntoHLAddress(RegisterC) }},
	0x72: &instruction{0x72, 2, "LD (HL),D", func(cpu *CPU) { cpu.PutRIntoHLAddress(RegisterD) }},
	0x73: &instruction{0x73, 2, "LD (HL),E", func(cpu *CPU) { cpu.PutRIntoHLAddress(RegisterE) }},
	0x74: &instruction{0x74, 2, "LD (HL),H", func(cpu *CPU) { cpu.PutRIntoHLAddress(RegisterH) }},
	0x75: &instruction{0x75, 2, "LD (HL),L", func(cpu *CPU) { cpu.PutRIntoHLAddress(RegisterL) }},

	// Register (A) -> Memory[HL++]
	0x22: &instruction{0x22, 2, "LD (HL+),A", func(cpu *CPU) { cpu.PutAIntoHLAddressThenIncrementHL() }},

	// Register (A) -> Memory[HL--]
	0x32: &instruction{0x32, 2, "LD (HL-),A", func(cpu *CPU) { cpu.PutAIntoHLAddressThenDecrementHL() }},

	// Register (A) -> Memory[C+0xFF0]
	0xE2: &instruction{0xE2, 2, "LD (C),A", func(cpu *CPU) { cpu.PutAIntoOffsetCAddress() }},

	// Register (A) -> Memory[Memory[PC]+0xFF00]
	0xE0: &instruction{0xE0, 3, "LD (a8),A", func(cpu *CPU) { cpu.PutAIntoOffsetImmediateAddress() }},

	// Memory[PC] -> Register (B, C, D, E, H, L, A)
	0x3E: &instruction{0x3E, 2, "LD A,d8", func(cpu *CPU) { cpu.PutNIntoR(RegisterA) }},
	0x06: &instruction{0x06, 2, "LD B,d8", func(cpu *CPU) { cpu.PutNIntoR(RegisterB) }},
	0x0E: &instruction{0x0E, 2, "LD C,d8", func(cpu *CPU) { cpu.PutNIntoR(RegisterC) }},
	0x16: &instruction{0x16, 2, "LD D,d8", func(cpu *CPU) { cpu.PutNIntoR(RegisterD) }},
	0x1E: &instruction{0x1E, 2, "LD E,d8", func(cpu *CPU) { cpu.PutNIntoR(RegisterE) }},
	0x26: &instruction{0x26, 2, "LD H,d8", func(cpu *CPU) { cpu.PutNIntoR(RegisterH) }},
	0x2E: &instruction{0x2E, 2, "LD L,d8", func(cpu *CPU) { cpu.PutNIntoR(RegisterL) }},

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
	0x7E: &instruction{0x7E, 2, "LD A,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(RegisterA) }},
	0x46: &instruction{0x46, 2, "LD B,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(RegisterB) }},
	0x4E: &instruction{0x4E, 2, "LD C,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(RegisterC) }},
	0x56: &instruction{0x56, 2, "LD D,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(RegisterD) }},
	0x5E: &instruction{0x5E, 2, "LD E,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(RegisterE) }},
	0x66: &instruction{0x66, 2, "LD H,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(RegisterH) }},
	0x6E: &instruction{0x6E, 2, "LD L,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(RegisterL) }},

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
	0xC5: &instruction{0xC5, 4, "PUSH BC", func(cpu *CPU) { cpu.PushRROntoStack(RegisterBC) }},
	0xD5: &instruction{0xD5, 4, "PUSH DE", func(cpu *CPU) { cpu.PushRROntoStack(RegisterDE) }},
	0xE5: &instruction{0xE5, 4, "PUSH HL", func(cpu *CPU) { cpu.PushRROntoStack(RegisterHL) }},

	// Memory[SP++ and SP++] -> Register (AF, BC, DE, HL)
	0xF1: &instruction{0xF1, 3, "POP AF", func(cpu *CPU) { cpu.PopStackIntoAF() }},
	0xC1: &instruction{0xC1, 3, "POP BC", func(cpu *CPU) { cpu.PopStackIntoRR(RegisterBC) }},
	0xD1: &instruction{0xD1, 3, "POP DE", func(cpu *CPU) { cpu.PopStackIntoRR(RegisterDE) }},
	0xE1: &instruction{0xE1, 3, "POP HL", func(cpu *CPU) { cpu.PopStackIntoRR(RegisterHL) }},

	// Memory[PC and PC+1] -> Register (BC, DE, HL, SP)
	0x01: &instruction{0x01, 3, "LD BC,d16", func(cpu *CPU) { cpu.PutNNIntoRR(RegisterBC) }},
	0x11: &instruction{0x11, 3, "LD DE,d16", func(cpu *CPU) { cpu.PutNNIntoRR(RegisterDE) }},
	0x21: &instruction{0x21, 3, "LD HL,d16", func(cpu *CPU) { cpu.PutNNIntoRR(RegisterHL) }},
	0x31: &instruction{0x31, 3, "LD SP,d16", func(cpu *CPU) { cpu.PutNNIntoRR(RegisterSP) }},

	// Register (HL) <- Register (SP) + Memory[PC]
	0xF8: &instruction{0xF8, 3, "LD HL,SP+r8", func(cpu *CPU) { cpu.PutOffsetSPIntoHL() }},

	/**
	 * 8-bit arithmetic / logical operations
	 */

	// Register (A) <- Register (A) + Register (A, B, C, D, E, H, L)
	0x87: &instruction{0x87, 1, "ADD A,A", func(cpu *CPU) { cpu.AddA() }},
	0x80: &instruction{0x80, 1, "ADD A,B", func(cpu *CPU) { cpu.AddR(RegisterB) }},
	0x81: &instruction{0x81, 1, "ADD A,C", func(cpu *CPU) { cpu.AddR(RegisterC) }},
	0x82: &instruction{0x82, 1, "ADD A,D", func(cpu *CPU) { cpu.AddR(RegisterD) }},
	0x83: &instruction{0x83, 1, "ADD A,E", func(cpu *CPU) { cpu.AddR(RegisterE) }},
	0x84: &instruction{0x84, 1, "ADD A,H", func(cpu *CPU) { cpu.AddR(RegisterH) }},
	0x85: &instruction{0x85, 1, "ADD A,L", func(cpu *CPU) { cpu.AddR(RegisterL) }},

	// Register (A) <- Register (A) + Memory[HL]
	0x86: &instruction{0x86, 2, "ADD A,(HL)", func(cpu *CPU) { cpu.AddHLDereference() }},

	// Register (A) <- Register (A) + Memory[PC]
	0xC6: &instruction{0xC6, 2, "ADD A,d8", func(cpu *CPU) { cpu.AddN() }},

	// Register (A) <- Register (A) + Register (A, B, C, D, E, H, L) + Flag (C)
	0x8F: &instruction{0x8F, 1, "ADC A,A", func(cpu *CPU) { cpu.AdcA() }},
	0x88: &instruction{0x88, 1, "ADC A,B", func(cpu *CPU) { cpu.AdcR(RegisterB) }},
	0x89: &instruction{0x89, 1, "ADC A,C", func(cpu *CPU) { cpu.AdcR(RegisterC) }},
	0x8A: &instruction{0x8A, 1, "ADC A,D", func(cpu *CPU) { cpu.AdcR(RegisterD) }},
	0x8B: &instruction{0x8B, 1, "ADC A,E", func(cpu *CPU) { cpu.AdcR(RegisterE) }},
	0x8C: &instruction{0x8C, 1, "ADC A,H", func(cpu *CPU) { cpu.AdcR(RegisterH) }},
	0x8D: &instruction{0x8D, 1, "ADC A,L", func(cpu *CPU) { cpu.AdcR(RegisterL) }},

	// Register (A) <- Register (A) + Memory[HL] + Flag (C)
	0x8E: &instruction{0x8E, 2, "ADC A,(HL)", func(cpu *CPU) { cpu.AdcHLDereference() }},

	// Register (A) <- Register (A) + Memory[PC] + Flag (C)
	0xCE: &instruction{0xCE, 2, "ADC A,d8", func(cpu *CPU) { cpu.AdcN() }},

	// Register (A) <- Register (A) - Register (A, B, C, D, E, H, L)
	0x97: &instruction{0x97, 1, "SUB A", func(cpu *CPU) { cpu.SubA() }},
	0x90: &instruction{0x90, 1, "SUB B", func(cpu *CPU) { cpu.SubR(RegisterB) }},
	0x91: &instruction{0x91, 1, "SUB C", func(cpu *CPU) { cpu.SubR(RegisterC) }},
	0x92: &instruction{0x92, 1, "SUB D", func(cpu *CPU) { cpu.SubR(RegisterD) }},
	0x93: &instruction{0x93, 1, "SUB E", func(cpu *CPU) { cpu.SubR(RegisterE) }},
	0x94: &instruction{0x94, 1, "SUB H", func(cpu *CPU) { cpu.SubR(RegisterH) }},
	0x95: &instruction{0x95, 1, "SUB L", func(cpu *CPU) { cpu.SubR(RegisterL) }},

	// Register (A) <- Register (A) - Memory[HL]
	0x96: &instruction{0x96, 2, "SUB (HL)", func(cpu *CPU) { cpu.SubHLDereference() }},

	// Register (A) <- Register (A) - Memory[PC]
	0xD6: &instruction{0xD6, 2, "SUB d8", func(cpu *CPU) { cpu.SubN() }},

	// Register (A) <- Register (A) - Register (A, B, C, D, E, H, L) - Flag (C)
	0x9F: &instruction{0x9F, 1, "SBC A,A", func(cpu *CPU) { cpu.SbcA() }},
	0x98: &instruction{0x98, 1, "SBC A,B", func(cpu *CPU) { cpu.SbcR(RegisterB) }},
	0x99: &instruction{0x99, 1, "SBC A,C", func(cpu *CPU) { cpu.SbcR(RegisterC) }},
	0x9A: &instruction{0x9A, 1, "SBC A,D", func(cpu *CPU) { cpu.SbcR(RegisterD) }},
	0x9B: &instruction{0x9B, 1, "SBC A,E", func(cpu *CPU) { cpu.SbcR(RegisterE) }},
	0x9C: &instruction{0x9C, 1, "SBC A,H", func(cpu *CPU) { cpu.SbcR(RegisterH) }},
	0x9D: &instruction{0x9D, 1, "SBC A,L", func(cpu *CPU) { cpu.SbcR(RegisterL) }},

	// Register (A) <- Register (A) - Memory[HL] - Flag (C)
	0x9E: &instruction{0x9E, 2, "SBC A,(HL)", func(cpu *CPU) { cpu.SbcHLDereference() }},

	// Register (A) <- Register (A) - Memory[PC] - Flag (C)
	0xDE: &instruction{0xDE, 2, "SBC A,d8", func(cpu *CPU) { cpu.SbcN() }},

	// Register (A) <- Register (A) & Register (A, B, C, D, E, H, L)
	0xA7: &instruction{0xA7, 1, "AND A", func(cpu *CPU) { cpu.AndA() }},
	0xA0: &instruction{0xA0, 1, "AND B", func(cpu *CPU) { cpu.AndR(RegisterB) }},
	0xA1: &instruction{0xA1, 1, "AND C", func(cpu *CPU) { cpu.AndR(RegisterC) }},
	0xA2: &instruction{0xA2, 1, "AND D", func(cpu *CPU) { cpu.AndR(RegisterD) }},
	0xA3: &instruction{0xA3, 1, "AND E", func(cpu *CPU) { cpu.AndR(RegisterE) }},
	0xA4: &instruction{0xA4, 1, "AND H", func(cpu *CPU) { cpu.AndR(RegisterH) }},
	0xA5: &instruction{0xA5, 1, "AND L", func(cpu *CPU) { cpu.AndR(RegisterL) }},

	// Register (A) <- Register (A) & Memory[HL]
	0xA6: &instruction{0xA6, 2, "AND (HL)", func(cpu *CPU) { cpu.AndHLDereference() }},

	// Register (A) <- Register (A) & Memory[PC]
	0xE6: &instruction{0xE6, 2, "AND d8", func(cpu *CPU) { cpu.AndN() }},

	// Register (A) <- Register (A) ^ Register (A, B, C, D, E, H, L)
	0xAF: &instruction{0xAF, 1, "XOR A", func(cpu *CPU) { cpu.XorA() }},
	0xA8: &instruction{0xA8, 1, "XOR B", func(cpu *CPU) { cpu.XorR(RegisterB) }},
	0xA9: &instruction{0xA9, 1, "XOR C", func(cpu *CPU) { cpu.XorR(RegisterC) }},
	0xAA: &instruction{0xAA, 1, "XOR D", func(cpu *CPU) { cpu.XorR(RegisterD) }},
	0xAB: &instruction{0xAB, 1, "XOR E", func(cpu *CPU) { cpu.XorR(RegisterE) }},
	0xAC: &instruction{0xAC, 1, "XOR H", func(cpu *CPU) { cpu.XorR(RegisterH) }},
	0xAD: &instruction{0xAD, 1, "XOR L", func(cpu *CPU) { cpu.XorR(RegisterL) }},

	// Register (A) <- Register (A) ^ Memory[HL]
	0xAE: &instruction{0xAE, 2, "XOR (HL)", func(cpu *CPU) { cpu.XorHLDereference() }},

	// Register (A) <- Register (A) ^ Memory[PC]
	0xEE: &instruction{0xEE, 2, "XOR d8", func(cpu *CPU) { cpu.XorN() }},

	// Register (A) <- Register (A) | Register (A, B, C, D, E, H, L)
	0xB7: &instruction{0xB7, 1, "OR A", func(cpu *CPU) { cpu.OrA() }},
	0xB0: &instruction{0xB0, 1, "OR B", func(cpu *CPU) { cpu.OrR(RegisterB) }},
	0xB1: &instruction{0xB1, 1, "OR C", func(cpu *CPU) { cpu.OrR(RegisterC) }},
	0xB2: &instruction{0xB2, 1, "OR D", func(cpu *CPU) { cpu.OrR(RegisterD) }},
	0xB3: &instruction{0xB3, 1, "OR E", func(cpu *CPU) { cpu.OrR(RegisterE) }},
	0xB4: &instruction{0xB4, 1, "OR H", func(cpu *CPU) { cpu.OrR(RegisterH) }},
	0xB5: &instruction{0xB5, 1, "OR L", func(cpu *CPU) { cpu.OrR(RegisterL) }},

	// Register (A) <- Register (A) | Memory[HL]
	0xB6: &instruction{0xB6, 2, "OR (HL)", func(cpu *CPU) { cpu.OrHLDereference() }},

	// Register (A) <- Register (A) | Memory[PC]
	0xF6: &instruction{0xF6, 2, "OR d8", func(cpu *CPU) { cpu.OrN() }},

	// Register (A) - Register (A, B, C, D, E, H, L)
	0xBF: &instruction{0xBF, 1, "CP A", func(cpu *CPU) { cpu.CompareA() }},
	0xB8: &instruction{0xB8, 1, "CP B", func(cpu *CPU) { cpu.CompareR(RegisterB) }},
	0xB9: &instruction{0xB9, 1, "CP C", func(cpu *CPU) { cpu.CompareR(RegisterC) }},
	0xBA: &instruction{0xBA, 1, "CP D", func(cpu *CPU) { cpu.CompareR(RegisterD) }},
	0xBB: &instruction{0xBB, 1, "CP E", func(cpu *CPU) { cpu.CompareR(RegisterE) }},
	0xBC: &instruction{0xBC, 1, "CP H", func(cpu *CPU) { cpu.CompareR(RegisterH) }},
	0xBD: &instruction{0xBD, 1, "CP L", func(cpu *CPU) { cpu.CompareR(RegisterL) }},

	// Register (A) - Memory[HL]
	0xBE: &instruction{0xBE, 2, "CP (HL)", func(cpu *CPU) { cpu.CompareHLDereference() }},

	// Register (A) - Memory[PC]
	0xFE: &instruction{0xFE, 2, "CP d8", func(cpu *CPU) { cpu.CompareN() }},

	// Register (A, B, C, D, E, H, L) <- Register (A, B, C, D, E, H, L) + 1
	0x3C: &instruction{0x3C, 1, "INC A", func(cpu *CPU) { cpu.IncrementA() }},
	0x04: &instruction{0x04, 1, "INC B", func(cpu *CPU) { cpu.IncrementR(RegisterB) }},
	0x0C: &instruction{0x0C, 1, "INC C", func(cpu *CPU) { cpu.IncrementR(RegisterC) }},
	0x14: &instruction{0x14, 1, "INC D", func(cpu *CPU) { cpu.IncrementR(RegisterD) }},
	0x1C: &instruction{0x1C, 1, "INC E", func(cpu *CPU) { cpu.IncrementR(RegisterE) }},
	0x24: &instruction{0x24, 1, "INC H", func(cpu *CPU) { cpu.IncrementR(RegisterH) }},
	0x2C: &instruction{0x2C, 1, "INC L", func(cpu *CPU) { cpu.IncrementR(RegisterL) }},

	// Memory[HL] <- Memory[HL] + 1
	0x34: &instruction{0x34, 3, "INC (HL)", func(cpu *CPU) { cpu.IncrementHLDereference() }},

	// Register (A, B, C, D, E, H, L) <- Register (A, B, C, D, E, H, L) - 1
	0x3D: &instruction{0x3D, 1, "DEC A", func(cpu *CPU) { cpu.DecrementA() }},
	0x05: &instruction{0x05, 1, "DEC B", func(cpu *CPU) { cpu.DecrementR(RegisterB) }},
	0x0D: &instruction{0x0D, 1, "DEC C", func(cpu *CPU) { cpu.DecrementR(RegisterC) }},
	0x15: &instruction{0x15, 1, "DEC D", func(cpu *CPU) { cpu.DecrementR(RegisterD) }},
	0x1D: &instruction{0x1D, 1, "DEC E", func(cpu *CPU) { cpu.DecrementR(RegisterE) }},
	0x25: &instruction{0x25, 1, "DEC H", func(cpu *CPU) { cpu.DecrementR(RegisterH) }},
	0x2D: &instruction{0x2D, 1, "DEC L", func(cpu *CPU) { cpu.DecrementR(RegisterL) }},

	// Memory[HL] <- Memory[HL] - 1
	0x35: &instruction{0x35, 3, "DEC (HL)", func(cpu *CPU) { cpu.DecrementHLDereference() }},

	// Register (A) decimally adjusted
	0x27: &instruction{0x27, 1, "DAA", func(cpu *CPU) { cpu.DecimalAdjustA() }},

	// ^Register (A) -> Register (A)
	0x2F: &instruction{0x2F, 1, "CPL", func(cpu *CPU) { cpu.ComplementA() }},

	// 1 -> Flag (C)
	0x37: &instruction{0x37, 1, "SCF", func(cpu *CPU) { cpu.SetCarryFlag() }},

	// ^Flag (C) -> Flag (C)
	0x3F: &instruction{0x3F, 1, "CCF", func(cpu *CPU) { cpu.ComplementCarryFlag() }},

	/**
	 * 16-bit arithmetic / logical operations
	 */

	// Register (HL) <- Register (HL) + Register (BC, DE, HL, SP)
	0x09: &instruction{0x09, 2, "ADD HL,BC", func(cpu *CPU) { cpu.AddRR(RegisterBC) }},
	0x19: &instruction{0x19, 2, "ADD HL,DE", func(cpu *CPU) { cpu.AddRR(RegisterDE) }},
	0x29: &instruction{0x29, 2, "ADD HL,HL", func(cpu *CPU) { cpu.AddRR(RegisterHL) }},
	0x39: &instruction{0x39, 2, "ADD HL,SP", func(cpu *CPU) { cpu.AddSP() }},

	// Register (BC, DE, HL, SP) <- Register (BC, DE, HL, SP) + 1
	0x03: &instruction{0x03, 2, "INC BC", func(cpu *CPU) { cpu.IncrementRR(RegisterBC) }},
	0x13: &instruction{0x13, 2, "INC DE", func(cpu *CPU) { cpu.IncrementRR(RegisterDE) }},
	0x23: &instruction{0x23, 2, "INC HL", func(cpu *CPU) { cpu.IncrementRR(RegisterHL) }},
	0x33: &instruction{0x33, 2, "INC SP", func(cpu *CPU) { cpu.IncrementSP() }},

	// Register (BC, DE, HL, SP) <- Register (BC, DE, HL, SP) - 1
	0x0B: &instruction{0x0B, 2, "DEC BC", func(cpu *CPU) { cpu.DecrementRR(RegisterBC) }},
	0x1B: &instruction{0x1B, 2, "DEC DE", func(cpu *CPU) { cpu.DecrementRR(RegisterDE) }},
	0x2B: &instruction{0x2B, 2, "DEC HL", func(cpu *CPU) { cpu.DecrementRR(RegisterHL) }},
	0x3B: &instruction{0x3B, 2, "DEC SP", func(cpu *CPU) { cpu.DecrementSP() }},

	// Register (SP) <- Register (SP) + Memory[PC]
	0xE8: &instruction{0xE8, 4, "ADD SP,r8", func(cpu *CPU) { cpu.AddOffsetImmediateToSP() }},

	/**
	 * Jumps / calls
	 */

	// Register (HL) -> Register (PC)
	0xE9: &instruction{0xE9, 1, "JP HL", func(cpu *CPU) { cpu.JumpHL() }},

	// Register (PC) <- Register (PC) + Memory[PC]
	0x18: &instruction{0x18, 3, "JR r8", func(cpu *CPU) { cpu.JumpOffset() }},

	// If Flag (C, Z) == (0, 1) => Register (PC) <- Register (PC) + Memory[PC]
	0x20: &instruction{0x20, 3, "JR NZ,r8", func(cpu *CPU) { cpu.JumpOffsetConditionally(FlagZ, false) }},
	0x28: &instruction{0x28, 3, "JR Z,r8", func(cpu *CPU) { cpu.JumpOffsetConditionally(FlagZ, true) }},
	0x30: &instruction{0x30, 3, "JR NC,r8", func(cpu *CPU) { cpu.JumpOffsetConditionally(FlagC, false) }},
	0x38: &instruction{0x38, 3, "JR C,r8", func(cpu *CPU) { cpu.JumpOffsetConditionally(FlagC, true) }},

	// Memory[PC and PC+1] -> Register (PC)
	0xC3: &instruction{0xC3, 4, "JP a16", func(cpu *CPU) { cpu.JumpNN() }},

	// If Flag (C, Z) == (0, 1) => Memory[PC and PC+1] -> Register (PC)
	0xC2: &instruction{0xC2, 4, "JP NZ,a16", func(cpu *CPU) { cpu.JumpNNConditionally(FlagZ, false) }},
	0xCA: &instruction{0xCA, 4, "JP Z,a16", func(cpu *CPU) { cpu.JumpNNConditionally(FlagZ, true) }},
	0xD2: &instruction{0xD2, 4, "JP NC,a16", func(cpu *CPU) { cpu.JumpNNConditionally(FlagC, false) }},
	0xDA: &instruction{0xDA, 4, "JP C,a16", func(cpu *CPU) { cpu.JumpNNConditionally(FlagC, true) }},

	// Register (PC) -> Memory[--SP and --SP], Memory[PC and PC+1] -> Register (PC)
	0xCD: &instruction{0xCD, 6, "CALL a16", func(cpu *CPU) { cpu.CallNN() }},

	// If Flag (C, Z) == (0, 1) => Register (PC) -> Memory[--SP and --SP], Memory[PC and PC+1] -> Register (PC)
	0xC4: &instruction{0xC4, 6, "CALL NZ,a16", func(cpu *CPU) { cpu.CallNNConditionally(FlagZ, false) }},
	0xCC: &instruction{0xCC, 6, "CALL Z,a16", func(cpu *CPU) { cpu.CallNNConditionally(FlagZ, true) }},
	0xD4: &instruction{0xD4, 6, "CALL NC,a16", func(cpu *CPU) { cpu.CallNNConditionally(FlagC, false) }},
	0xDC: &instruction{0xDC, 6, "CALL C,a16", func(cpu *CPU) { cpu.CallNNConditionally(FlagC, true) }},

	// Memory[SP++ and SP++] -> Register (PC)
	0xC9: &instruction{0xC9, 4, "RET", func(cpu *CPU) { cpu.Return() }},

	// If Flag (C, Z) == (0, 1) => Memory[SP++ and SP++] -> Register (PC)
	0xC0: &instruction{0xC0, 5, "RET NZ", func(cpu *CPU) { cpu.ReturnConditionally(FlagZ, false) }},
	0xC8: &instruction{0xC8, 5, "RET Z", func(cpu *CPU) { cpu.ReturnConditionally(FlagZ, true) }},
	0xD0: &instruction{0xD0, 5, "RET NC", func(cpu *CPU) { cpu.ReturnConditionally(FlagC, false) }},
	0xD8: &instruction{0xD8, 5, "RET C", func(cpu *CPU) { cpu.ReturnConditionally(FlagC, true) }},

	// Memory[SP++ and SP++] -> Register (PC), <more to do>
	0xD9: &instruction{0xD9, 4, "RETI", func(cpu *CPU) { cpu.ReturnPostInterrupt() }},

	// Register (PC) -> Memory[--SP and --SP], (0x00, 0x08, 0x10, 0x18, 0x20, 0x28, 0x30, 0x38) -> Register (PC)
	0xC7: &instruction{0xC7, 4, "RST 00H", func(cpu *CPU) { cpu.Restart(0x00) }},
	0xCF: &instruction{0xCF, 4, "RST 08H", func(cpu *CPU) { cpu.Restart(0x08) }},
	0xD7: &instruction{0xD7, 4, "RST 10H", func(cpu *CPU) { cpu.Restart(0x10) }},
	0xDF: &instruction{0xDF, 4, "RST 18H", func(cpu *CPU) { cpu.Restart(0x18) }},
	0xE7: &instruction{0xE7, 4, "RST 20H", func(cpu *CPU) { cpu.Restart(0x20) }},
	0xEF: &instruction{0xEF, 4, "RST 28H", func(cpu *CPU) { cpu.Restart(0x28) }},
	0xF7: &instruction{0xF7, 4, "RST 30H", func(cpu *CPU) { cpu.Restart(0x30) }},
	0xFF: &instruction{0xFF, 4, "RST 38H", func(cpu *CPU) { cpu.Restart(0x38) }},

	/**
	 * 8-bit rotation / shifts and bit instructions
	 */

	// Register (A) << / >>
	0x07: &instruction{0x07, 1, "RLCA", func(cpu *CPU) { cpu.RLCA() }},
	0x0F: &instruction{0x0F, 1, "RRCA", func(cpu *CPU) { cpu.RRCA() }},
	0x17: &instruction{0x17, 1, "RLA", func(cpu *CPU) { cpu.RLA() }},
	0x1F: &instruction{0x1F, 1, "RRA", func(cpu *CPU) { cpu.RRA() }},
}

var instructionsCB = instructionSet{
	/**
	 * 8-bit rotation / shifts and bit instructions
	 */

	// Register (A, B, C, D, E, H, L) <<
	0x07: &instruction{0x07, 1, "RLC A", func(cpu *CPU) { cpu.RLCACB() }},
	0x00: &instruction{0x00, 1, "RLC B", func(cpu *CPU) { cpu.RLC(RegisterB) }},
	0x01: &instruction{0x01, 1, "RLC C", func(cpu *CPU) { cpu.RLC(RegisterC) }},
	0x02: &instruction{0x02, 1, "RLC D", func(cpu *CPU) { cpu.RLC(RegisterD) }},
	0x03: &instruction{0x03, 1, "RLC E", func(cpu *CPU) { cpu.RLC(RegisterE) }},
	0x04: &instruction{0x04, 1, "RLC H", func(cpu *CPU) { cpu.RLC(RegisterH) }},
	0x05: &instruction{0x05, 1, "RLC L", func(cpu *CPU) { cpu.RLC(RegisterL) }},

	// Memory[HL] <<
	0x06: &instruction{0x06, 3, "RLC (HL)", func(cpu *CPU) { cpu.RLCHLDereference() }},

	// Register (A, B, C, D, E, H, L) >>
	0x0F: &instruction{0x0F, 1, "RRC A", func(cpu *CPU) { cpu.RRCACB() }},
	0x08: &instruction{0x08, 1, "RRC B", func(cpu *CPU) { cpu.RRC(RegisterB) }},
	0x09: &instruction{0x09, 1, "RRC C", func(cpu *CPU) { cpu.RRC(RegisterC) }},
	0x0A: &instruction{0x0A, 1, "RRC D", func(cpu *CPU) { cpu.RRC(RegisterD) }},
	0x0B: &instruction{0x0B, 1, "RRC E", func(cpu *CPU) { cpu.RRC(RegisterE) }},
	0x0C: &instruction{0x0C, 1, "RRC H", func(cpu *CPU) { cpu.RRC(RegisterH) }},
	0x0D: &instruction{0x0D, 1, "RRC L", func(cpu *CPU) { cpu.RRC(RegisterL) }},

	// Memory[HL] >>
	0x0E: &instruction{0x0E, 3, "RRC (HL)", func(cpu *CPU) { cpu.RRCHLDereference() }},

	// Register (A, B, C, D, E, H, L) <<
	0x17: &instruction{0x17, 1, "RL A", func(cpu *CPU) { cpu.RLA() }},
	0x10: &instruction{0x10, 1, "RL B", func(cpu *CPU) { cpu.RL(RegisterB) }},
	0x11: &instruction{0x11, 1, "RL C", func(cpu *CPU) { cpu.RL(RegisterC) }},
	0x12: &instruction{0x12, 1, "RL D", func(cpu *CPU) { cpu.RL(RegisterD) }},
	0x13: &instruction{0x13, 1, "RL E", func(cpu *CPU) { cpu.RL(RegisterE) }},
	0x14: &instruction{0x14, 1, "RL H", func(cpu *CPU) { cpu.RL(RegisterH) }},
	0x15: &instruction{0x15, 1, "RL L", func(cpu *CPU) { cpu.RL(RegisterL) }},

	// Memory[HL] <<
	0x16: &instruction{0x16, 3, "RL (HL)", func(cpu *CPU) { cpu.RLHLDereference() }},

	// Register (A, B, C, D, E, H, L) >>
	0x1F: &instruction{0x1F, 1, "RR A", func(cpu *CPU) { cpu.RRA() }},
	0x18: &instruction{0x18, 1, "RR B", func(cpu *CPU) { cpu.RR(RegisterB) }},
	0x19: &instruction{0x19, 1, "RR C", func(cpu *CPU) { cpu.RR(RegisterC) }},
	0x1A: &instruction{0x1A, 1, "RR D", func(cpu *CPU) { cpu.RR(RegisterD) }},
	0x1B: &instruction{0x1B, 1, "RR E", func(cpu *CPU) { cpu.RR(RegisterE) }},
	0x1C: &instruction{0x1C, 1, "RR H", func(cpu *CPU) { cpu.RR(RegisterH) }},
	0x1D: &instruction{0x1D, 1, "RR L", func(cpu *CPU) { cpu.RR(RegisterL) }},

	// Memory[HL] >>
	0x1E: &instruction{0x1E, 3, "RR (HL)", func(cpu *CPU) { cpu.RRHLDereference() }},

	// Register (A, B, C, D, E, H, L) <<
	0x27: &instruction{0x27, 1, "SLA A", func(cpu *CPU) { cpu.SLAA() }},
	0x20: &instruction{0x20, 1, "SLA B", func(cpu *CPU) { cpu.SLA(RegisterB) }},
	0x21: &instruction{0x21, 1, "SLA C", func(cpu *CPU) { cpu.SLA(RegisterC) }},
	0x22: &instruction{0x22, 1, "SLA D", func(cpu *CPU) { cpu.SLA(RegisterD) }},
	0x23: &instruction{0x23, 1, "SLA E", func(cpu *CPU) { cpu.SLA(RegisterE) }},
	0x24: &instruction{0x24, 1, "SLA H", func(cpu *CPU) { cpu.SLA(RegisterH) }},
	0x25: &instruction{0x25, 1, "SLA L", func(cpu *CPU) { cpu.SLA(RegisterL) }},

	// Memory[HL] <<
	0x26: &instruction{0x26, 3, "SLA (HL)", func(cpu *CPU) { cpu.SLAHLDereference() }},

	// Register (A, B, C, D, E, H, L) >>
	0x2F: &instruction{0x2F, 1, "SRA A", func(cpu *CPU) { cpu.SRAA() }},
	0x28: &instruction{0x28, 1, "SRA B", func(cpu *CPU) { cpu.SRA(RegisterB) }},
	0x29: &instruction{0x29, 1, "SRA C", func(cpu *CPU) { cpu.SRA(RegisterC) }},
	0x2A: &instruction{0x2A, 1, "SRA D", func(cpu *CPU) { cpu.SRA(RegisterD) }},
	0x2B: &instruction{0x2B, 1, "SRA E", func(cpu *CPU) { cpu.SRA(RegisterE) }},
	0x2C: &instruction{0x2C, 1, "SRA H", func(cpu *CPU) { cpu.SRA(RegisterH) }},
	0x2D: &instruction{0x2D, 1, "SRA L", func(cpu *CPU) { cpu.SRA(RegisterL) }},

	// Memory[HL] >>
	0x2E: &instruction{0x2E, 3, "SRA (HL)", func(cpu *CPU) { cpu.SRAHLDereference() }},

	// Register (A, B, C, D, E, H, L) >>
	0x3F: &instruction{0x3F, 1, "SRL A", func(cpu *CPU) { cpu.SRLA() }},
	0x38: &instruction{0x38, 1, "SRL B", func(cpu *CPU) { cpu.SRL(RegisterB) }},
	0x39: &instruction{0x39, 1, "SRL C", func(cpu *CPU) { cpu.SRL(RegisterC) }},
	0x3A: &instruction{0x3A, 1, "SRL D", func(cpu *CPU) { cpu.SRL(RegisterD) }},
	0x3B: &instruction{0x3B, 1, "SRL E", func(cpu *CPU) { cpu.SRL(RegisterE) }},
	0x3C: &instruction{0x3C, 1, "SRL H", func(cpu *CPU) { cpu.SRL(RegisterH) }},
	0x3D: &instruction{0x3D, 1, "SRL L", func(cpu *CPU) { cpu.SRL(RegisterL) }},

	// Memory[HL] >>
	0x3E: &instruction{0x3E, 3, "SRL (HL)", func(cpu *CPU) { cpu.SRLHLDereference() }},

	// Register (A, B, C, D, E, H, L) <- Register (A, B, C, D, E, H, L)[0-4] | Register (A, B, C, D, E, H, L)[4-8]
	0x37: &instruction{0x37, 1, "SWAP A", func(cpu *CPU) { cpu.SwapA() }},
	0x30: &instruction{0x30, 1, "SWAP B", func(cpu *CPU) { cpu.Swap(RegisterB) }},
	0x31: &instruction{0x31, 1, "SWAP C", func(cpu *CPU) { cpu.Swap(RegisterC) }},
	0x32: &instruction{0x32, 1, "SWAP D", func(cpu *CPU) { cpu.Swap(RegisterD) }},
	0x33: &instruction{0x33, 1, "SWAP E", func(cpu *CPU) { cpu.Swap(RegisterE) }},
	0x34: &instruction{0x34, 1, "SWAP H", func(cpu *CPU) { cpu.Swap(RegisterH) }},
	0x35: &instruction{0x35, 1, "SWAP L", func(cpu *CPU) { cpu.Swap(RegisterL) }},

	// Memory[HL] <- Memory[HL][0-4] | Memory[HL][4-8]
	0x36: &instruction{0x36, 3, "SWAP (HL)", func(cpu *CPU) { cpu.SwapHLDereference() }},

	// Flag (Z) <- ^Register (A, B, C, D, E, H, L)[0]
	0x47: &instruction{0x47, 1, "BIT 0,A", func(cpu *CPU) { cpu.BitA(0) }},
	0x40: &instruction{0x40, 1, "BIT 0,B", func(cpu *CPU) { cpu.Bit(0, RegisterB) }},
	0x41: &instruction{0x41, 1, "BIT 0,C", func(cpu *CPU) { cpu.Bit(0, RegisterC) }},
	0x42: &instruction{0x42, 1, "BIT 0,D", func(cpu *CPU) { cpu.Bit(0, RegisterD) }},
	0x43: &instruction{0x43, 1, "BIT 0,E", func(cpu *CPU) { cpu.Bit(0, RegisterE) }},
	0x44: &instruction{0x44, 1, "BIT 0,H", func(cpu *CPU) { cpu.Bit(0, RegisterH) }},
	0x45: &instruction{0x45, 1, "BIT 0,L", func(cpu *CPU) { cpu.Bit(0, RegisterL) }},

	// Flag (Z) <- ^Memory[HL][0]
	0x46: &instruction{0x46, 2, "BIT 0,(HL)", func(cpu *CPU) { cpu.BitHLDereference(0) }},

	// Flag (Z) <- ^Register (A, B, C, D, E, H, L)[1]
	0x4F: &instruction{0x4F, 1, "BIT 1,A", func(cpu *CPU) { cpu.BitA(1) }},
	0x48: &instruction{0x48, 1, "BIT 1,B", func(cpu *CPU) { cpu.Bit(1, RegisterB) }},
	0x49: &instruction{0x49, 1, "BIT 1,C", func(cpu *CPU) { cpu.Bit(1, RegisterC) }},
	0x4A: &instruction{0x4A, 1, "BIT 1,D", func(cpu *CPU) { cpu.Bit(1, RegisterD) }},
	0x4B: &instruction{0x4B, 1, "BIT 1,E", func(cpu *CPU) { cpu.Bit(1, RegisterE) }},
	0x4C: &instruction{0x4C, 1, "BIT 1,H", func(cpu *CPU) { cpu.Bit(1, RegisterH) }},
	0x4D: &instruction{0x4D, 1, "BIT 1,L", func(cpu *CPU) { cpu.Bit(1, RegisterL) }},

	// Flag (Z) <- ^Memory[HL][1]
	0x4E: &instruction{0x4E, 2, "BIT 1,(HL)", func(cpu *CPU) { cpu.BitHLDereference(1) }},

	// Flag (Z) <- ^Register (A, B, C, D, E, H, L)[2]
	0x57: &instruction{0x57, 1, "BIT 2,A", func(cpu *CPU) { cpu.BitA(2) }},
	0x50: &instruction{0x50, 1, "BIT 2,B", func(cpu *CPU) { cpu.Bit(2, RegisterB) }},
	0x51: &instruction{0x51, 1, "BIT 2,C", func(cpu *CPU) { cpu.Bit(2, RegisterC) }},
	0x52: &instruction{0x52, 1, "BIT 2,D", func(cpu *CPU) { cpu.Bit(2, RegisterD) }},
	0x53: &instruction{0x53, 1, "BIT 2,E", func(cpu *CPU) { cpu.Bit(2, RegisterE) }},
	0x54: &instruction{0x54, 1, "BIT 2,H", func(cpu *CPU) { cpu.Bit(2, RegisterH) }},
	0x55: &instruction{0x55, 1, "BIT 2,L", func(cpu *CPU) { cpu.Bit(2, RegisterL) }},

	// Flag (Z) <- ^Memory[HL][2]
	0x56: &instruction{0x56, 2, "BIT 2,(HL)", func(cpu *CPU) { cpu.BitHLDereference(2) }},

	// Flag (Z) <- ^Register (A, B, C, D, E, H, L)[3]
	0x5F: &instruction{0x5F, 1, "BIT 3,A", func(cpu *CPU) { cpu.BitA(3) }},
	0x58: &instruction{0x58, 1, "BIT 3,B", func(cpu *CPU) { cpu.Bit(3, RegisterB) }},
	0x59: &instruction{0x59, 1, "BIT 3,C", func(cpu *CPU) { cpu.Bit(3, RegisterC) }},
	0x5A: &instruction{0x5A, 1, "BIT 3,D", func(cpu *CPU) { cpu.Bit(3, RegisterD) }},
	0x5B: &instruction{0x5B, 1, "BIT 3,E", func(cpu *CPU) { cpu.Bit(3, RegisterE) }},
	0x5C: &instruction{0x5C, 1, "BIT 3,H", func(cpu *CPU) { cpu.Bit(3, RegisterH) }},
	0x5D: &instruction{0x5D, 1, "BIT 3,L", func(cpu *CPU) { cpu.Bit(3, RegisterL) }},

	// Flag (Z) <- ^Memory[HL][3]
	0x5E: &instruction{0x5E, 2, "BIT 3,(HL)", func(cpu *CPU) { cpu.BitHLDereference(3) }},

	// Flag (Z) <- ^Register (A, B, C, D, E, H, L)[4]
	0x67: &instruction{0x67, 1, "BIT 4,A", func(cpu *CPU) { cpu.BitA(4) }},
	0x60: &instruction{0x60, 1, "BIT 4,B", func(cpu *CPU) { cpu.Bit(4, RegisterB) }},
	0x61: &instruction{0x61, 1, "BIT 4,C", func(cpu *CPU) { cpu.Bit(4, RegisterC) }},
	0x62: &instruction{0x62, 1, "BIT 4,D", func(cpu *CPU) { cpu.Bit(4, RegisterD) }},
	0x63: &instruction{0x63, 1, "BIT 4,E", func(cpu *CPU) { cpu.Bit(4, RegisterE) }},
	0x64: &instruction{0x64, 1, "BIT 4,H", func(cpu *CPU) { cpu.Bit(4, RegisterH) }},
	0x65: &instruction{0x65, 1, "BIT 4,L", func(cpu *CPU) { cpu.Bit(4, RegisterL) }},

	// Flag (Z) <- ^Memory[HL][4]
	0x66: &instruction{0x66, 2, "BIT 4,(HL)", func(cpu *CPU) { cpu.BitHLDereference(4) }},

	// Flag (Z) <- ^Register (A, B, C, D, E, H, L)[5]
	0x6F: &instruction{0x6F, 1, "BIT 5,A", func(cpu *CPU) { cpu.BitA(5) }},
	0x68: &instruction{0x68, 1, "BIT 5,B", func(cpu *CPU) { cpu.Bit(5, RegisterB) }},
	0x69: &instruction{0x69, 1, "BIT 5,C", func(cpu *CPU) { cpu.Bit(5, RegisterC) }},
	0x6A: &instruction{0x6A, 1, "BIT 5,D", func(cpu *CPU) { cpu.Bit(5, RegisterD) }},
	0x6B: &instruction{0x6B, 1, "BIT 5,E", func(cpu *CPU) { cpu.Bit(5, RegisterE) }},
	0x6C: &instruction{0x6C, 1, "BIT 5,H", func(cpu *CPU) { cpu.Bit(5, RegisterH) }},
	0x6D: &instruction{0x6D, 1, "BIT 5,L", func(cpu *CPU) { cpu.Bit(5, RegisterL) }},

	// Flag (Z) <- ^Memory[HL][5]
	0x6E: &instruction{0x6E, 2, "BIT 5,(HL)", func(cpu *CPU) { cpu.BitHLDereference(5) }},

	// Flag (Z) <- ^Register (A, B, C, D, E, H, L)[6]
	0x77: &instruction{0x77, 1, "BIT 6,A", func(cpu *CPU) { cpu.BitA(6) }},
	0x70: &instruction{0x70, 1, "BIT 6,B", func(cpu *CPU) { cpu.Bit(6, RegisterB) }},
	0x71: &instruction{0x71, 1, "BIT 6,C", func(cpu *CPU) { cpu.Bit(6, RegisterC) }},
	0x72: &instruction{0x72, 1, "BIT 6,D", func(cpu *CPU) { cpu.Bit(6, RegisterD) }},
	0x73: &instruction{0x73, 1, "BIT 6,E", func(cpu *CPU) { cpu.Bit(6, RegisterE) }},
	0x74: &instruction{0x74, 1, "BIT 6,H", func(cpu *CPU) { cpu.Bit(6, RegisterH) }},
	0x75: &instruction{0x75, 1, "BIT 6,L", func(cpu *CPU) { cpu.Bit(6, RegisterL) }},

	// Flag (Z) <- ^Memory[HL][6]
	0x76: &instruction{0x76, 2, "BIT 6,(HL)", func(cpu *CPU) { cpu.BitHLDereference(6) }},

	// Flag (Z) <- ^Register (A, B, C, D, E, H, L)[7]
	0x7F: &instruction{0x7F, 1, "BIT 7,A", func(cpu *CPU) { cpu.BitA(7) }},
	0x78: &instruction{0x78, 1, "BIT 7,B", func(cpu *CPU) { cpu.Bit(7, RegisterB) }},
	0x79: &instruction{0x79, 1, "BIT 7,C", func(cpu *CPU) { cpu.Bit(7, RegisterC) }},
	0x7A: &instruction{0x7A, 1, "BIT 7,D", func(cpu *CPU) { cpu.Bit(7, RegisterD) }},
	0x7B: &instruction{0x7B, 1, "BIT 7,E", func(cpu *CPU) { cpu.Bit(7, RegisterE) }},
	0x7C: &instruction{0x7C, 1, "BIT 7,H", func(cpu *CPU) { cpu.Bit(7, RegisterH) }},
	0x7D: &instruction{0x7D, 1, "BIT 7,L", func(cpu *CPU) { cpu.Bit(7, RegisterL) }},

	// Flag (Z) <- ^Memory[HL][7]
	0x7E: &instruction{0x7E, 2, "BIT 7,(HL)", func(cpu *CPU) { cpu.BitHLDereference(7) }},

	// Register (A, B, C, D, E, H, L)[0] <- 0
	0x87: &instruction{0x87, 1, "RES 0,A", func(cpu *CPU) { cpu.ResetA(0) }},
	0x80: &instruction{0x80, 1, "RES 0,B", func(cpu *CPU) { cpu.Reset(0, RegisterB) }},
	0x81: &instruction{0x81, 1, "RES 0,C", func(cpu *CPU) { cpu.Reset(0, RegisterC) }},
	0x82: &instruction{0x82, 1, "RES 0,D", func(cpu *CPU) { cpu.Reset(0, RegisterD) }},
	0x83: &instruction{0x83, 1, "RES 0,E", func(cpu *CPU) { cpu.Reset(0, RegisterE) }},
	0x84: &instruction{0x84, 1, "RES 0,H", func(cpu *CPU) { cpu.Reset(0, RegisterH) }},
	0x85: &instruction{0x85, 1, "RES 0,L", func(cpu *CPU) { cpu.Reset(0, RegisterL) }},

	// Memory[HL][0] <- 0
	0x86: &instruction{0x86, 3, "RES 0,(HL)", func(cpu *CPU) { cpu.ResetHLDereference(0) }},

	// Register (A, B, C, D, E, H, L)[1] <- 0
	0x8F: &instruction{0x8F, 1, "RES 1,A", func(cpu *CPU) { cpu.ResetA(1) }},
	0x88: &instruction{0x88, 1, "RES 1,B", func(cpu *CPU) { cpu.Reset(1, RegisterB) }},
	0x89: &instruction{0x89, 1, "RES 1,C", func(cpu *CPU) { cpu.Reset(1, RegisterC) }},
	0x8A: &instruction{0x8A, 1, "RES 1,D", func(cpu *CPU) { cpu.Reset(1, RegisterD) }},
	0x8B: &instruction{0x8B, 1, "RES 1,E", func(cpu *CPU) { cpu.Reset(1, RegisterE) }},
	0x8C: &instruction{0x8C, 1, "RES 1,H", func(cpu *CPU) { cpu.Reset(1, RegisterH) }},
	0x8D: &instruction{0x8D, 1, "RES 1,L", func(cpu *CPU) { cpu.Reset(1, RegisterL) }},

	// Memory[HL][1] <- 0
	0x8E: &instruction{0x8E, 3, "RES 1,(HL)", func(cpu *CPU) { cpu.ResetHLDereference(1) }},

	// Register (A, B, C, D, E, H, L)[2] <- 0
	0x97: &instruction{0x97, 1, "RES 2,A", func(cpu *CPU) { cpu.ResetA(2) }},
	0x90: &instruction{0x90, 1, "RES 2,B", func(cpu *CPU) { cpu.Reset(2, RegisterB) }},
	0x91: &instruction{0x91, 1, "RES 2,C", func(cpu *CPU) { cpu.Reset(2, RegisterC) }},
	0x92: &instruction{0x92, 1, "RES 2,D", func(cpu *CPU) { cpu.Reset(2, RegisterD) }},
	0x93: &instruction{0x93, 1, "RES 2,E", func(cpu *CPU) { cpu.Reset(2, RegisterE) }},
	0x94: &instruction{0x94, 1, "RES 2,H", func(cpu *CPU) { cpu.Reset(2, RegisterH) }},
	0x95: &instruction{0x95, 1, "RES 2,L", func(cpu *CPU) { cpu.Reset(2, RegisterL) }},

	// Memory[HL][2] <- 0
	0x96: &instruction{0x96, 3, "RES 2,(HL)", func(cpu *CPU) { cpu.ResetHLDereference(2) }},

	// Register (A, B, C, D, E, H, L)[3] <- 0
	0x9F: &instruction{0x9F, 1, "RES 3,A", func(cpu *CPU) { cpu.ResetA(3) }},
	0x98: &instruction{0x98, 1, "RES 3,B", func(cpu *CPU) { cpu.Reset(3, RegisterB) }},
	0x99: &instruction{0x99, 1, "RES 3,C", func(cpu *CPU) { cpu.Reset(3, RegisterC) }},
	0x9A: &instruction{0x9A, 1, "RES 3,D", func(cpu *CPU) { cpu.Reset(3, RegisterD) }},
	0x9B: &instruction{0x9B, 1, "RES 3,E", func(cpu *CPU) { cpu.Reset(3, RegisterE) }},
	0x9C: &instruction{0x9C, 1, "RES 3,H", func(cpu *CPU) { cpu.Reset(3, RegisterH) }},
	0x9D: &instruction{0x9D, 1, "RES 3,L", func(cpu *CPU) { cpu.Reset(3, RegisterL) }},

	// Memory[HL][3] <- 0
	0x9E: &instruction{0x9E, 3, "RES 3,(HL)", func(cpu *CPU) { cpu.ResetHLDereference(3) }},

	// Register (A, B, C, D, E, H, L)[4] <- 0
	0xA7: &instruction{0xA7, 1, "RES 4,A", func(cpu *CPU) { cpu.ResetA(4) }},
	0xA0: &instruction{0xA0, 1, "RES 4,B", func(cpu *CPU) { cpu.Reset(4, RegisterB) }},
	0xA1: &instruction{0xA1, 1, "RES 4,C", func(cpu *CPU) { cpu.Reset(4, RegisterC) }},
	0xA2: &instruction{0xA2, 1, "RES 4,D", func(cpu *CPU) { cpu.Reset(4, RegisterD) }},
	0xA3: &instruction{0xA3, 1, "RES 4,E", func(cpu *CPU) { cpu.Reset(4, RegisterE) }},
	0xA4: &instruction{0xA4, 1, "RES 4,H", func(cpu *CPU) { cpu.Reset(4, RegisterH) }},
	0xA5: &instruction{0xA5, 1, "RES 4,L", func(cpu *CPU) { cpu.Reset(4, RegisterL) }},

	// Memory[HL][4] <- 0
	0xA6: &instruction{0xA6, 3, "RES 4,(HL)", func(cpu *CPU) { cpu.ResetHLDereference(4) }},

	// Register (A, B, C, D, E, H, L)[5] <- 0
	0xAF: &instruction{0xAF, 1, "RES 5,A", func(cpu *CPU) { cpu.ResetA(5) }},
	0xA8: &instruction{0xA8, 1, "RES 5,B", func(cpu *CPU) { cpu.Reset(5, RegisterB) }},
	0xA9: &instruction{0xA9, 1, "RES 5,C", func(cpu *CPU) { cpu.Reset(5, RegisterC) }},
	0xAA: &instruction{0xAA, 1, "RES 5,D", func(cpu *CPU) { cpu.Reset(5, RegisterD) }},
	0xAB: &instruction{0xAB, 1, "RES 5,E", func(cpu *CPU) { cpu.Reset(5, RegisterE) }},
	0xAC: &instruction{0xAC, 1, "RES 5,H", func(cpu *CPU) { cpu.Reset(5, RegisterH) }},
	0xAD: &instruction{0xAD, 1, "RES 5,L", func(cpu *CPU) { cpu.Reset(5, RegisterL) }},

	// Memory[HL][5] <- 0
	0xAE: &instruction{0xAE, 3, "RES 5,(HL)", func(cpu *CPU) { cpu.ResetHLDereference(5) }},

	// Register (A, B, C, D, E, H, L)[6] <- 0
	0xB7: &instruction{0xB7, 1, "RES 6,A", func(cpu *CPU) { cpu.ResetA(6) }},
	0xB0: &instruction{0xB0, 1, "RES 6,B", func(cpu *CPU) { cpu.Reset(6, RegisterB) }},
	0xB1: &instruction{0xB1, 1, "RES 6,C", func(cpu *CPU) { cpu.Reset(6, RegisterC) }},
	0xB2: &instruction{0xB2, 1, "RES 6,D", func(cpu *CPU) { cpu.Reset(6, RegisterD) }},
	0xB3: &instruction{0xB3, 1, "RES 6,E", func(cpu *CPU) { cpu.Reset(6, RegisterE) }},
	0xB4: &instruction{0xB4, 1, "RES 6,H", func(cpu *CPU) { cpu.Reset(6, RegisterH) }},
	0xB5: &instruction{0xB5, 1, "RES 6,L", func(cpu *CPU) { cpu.Reset(6, RegisterL) }},

	// Memory[HL][6] <- 0
	0xB6: &instruction{0xB6, 3, "RES 6,(HL)", func(cpu *CPU) { cpu.ResetHLDereference(6) }},

	// Register (A, B, C, D, E, H, L)[7] <- 0
	0xBF: &instruction{0xBF, 1, "RES 7,A", func(cpu *CPU) { cpu.ResetA(7) }},
	0xB8: &instruction{0xB8, 1, "RES 7,B", func(cpu *CPU) { cpu.Reset(7, RegisterB) }},
	0xB9: &instruction{0xB9, 1, "RES 7,C", func(cpu *CPU) { cpu.Reset(7, RegisterC) }},
	0xBA: &instruction{0xBA, 1, "RES 7,D", func(cpu *CPU) { cpu.Reset(7, RegisterD) }},
	0xBB: &instruction{0xBB, 1, "RES 7,E", func(cpu *CPU) { cpu.Reset(7, RegisterE) }},
	0xBC: &instruction{0xBC, 1, "RES 7,H", func(cpu *CPU) { cpu.Reset(7, RegisterH) }},
	0xBD: &instruction{0xBD, 1, "RES 7,L", func(cpu *CPU) { cpu.Reset(7, RegisterL) }},

	// Memory[HL][7] <- 0
	0xBE: &instruction{0xBE, 3, "RES 7,(HL)", func(cpu *CPU) { cpu.ResetHLDereference(7) }},

	// Register (A, B, C, D, E, H, L)[0] <- 1
	0xC7: &instruction{0xC7, 1, "SET 0,A", func(cpu *CPU) { cpu.SetA(0) }},
	0xC0: &instruction{0xC0, 1, "SET 0,B", func(cpu *CPU) { cpu.Set(0, RegisterB) }},
	0xC1: &instruction{0xC1, 1, "SET 0,C", func(cpu *CPU) { cpu.Set(0, RegisterC) }},
	0xC2: &instruction{0xC2, 1, "SET 0,D", func(cpu *CPU) { cpu.Set(0, RegisterD) }},
	0xC3: &instruction{0xC3, 1, "SET 0,E", func(cpu *CPU) { cpu.Set(0, RegisterE) }},
	0xC4: &instruction{0xC4, 1, "SET 0,H", func(cpu *CPU) { cpu.Set(0, RegisterH) }},
	0xC5: &instruction{0xC5, 1, "SET 0,L", func(cpu *CPU) { cpu.Set(0, RegisterL) }},

	// Memory[HL][0] <- 1
	0xC6: &instruction{0xC6, 3, "SET 0,(HL)", func(cpu *CPU) { cpu.SetHLDereference(0) }},

	// Register (A, B, C, D, E, H, L)[1] <- 1
	0xCF: &instruction{0xCF, 1, "SET 1,A", func(cpu *CPU) { cpu.SetA(1) }},
	0xC8: &instruction{0xC8, 1, "SET 1,B", func(cpu *CPU) { cpu.Set(1, RegisterB) }},
	0xC9: &instruction{0xC9, 1, "SET 1,C", func(cpu *CPU) { cpu.Set(1, RegisterC) }},
	0xCA: &instruction{0xCA, 1, "SET 1,D", func(cpu *CPU) { cpu.Set(1, RegisterD) }},
	0xCB: &instruction{0xCB, 1, "SET 1,E", func(cpu *CPU) { cpu.Set(1, RegisterE) }},
	0xCC: &instruction{0xCC, 1, "SET 1,H", func(cpu *CPU) { cpu.Set(1, RegisterH) }},
	0xCD: &instruction{0xCD, 1, "SET 1,L", func(cpu *CPU) { cpu.Set(1, RegisterL) }},

	// Memory[HL][1] <- 1
	0xCE: &instruction{0xCE, 3, "SET 1,(HL)", func(cpu *CPU) { cpu.SetHLDereference(1) }},

	// Register (A, B, C, D, E, H, L)[2] <- 1
	0xD7: &instruction{0xD7, 1, "SET 2,A", func(cpu *CPU) { cpu.SetA(2) }},
	0xD0: &instruction{0xD0, 1, "SET 2,B", func(cpu *CPU) { cpu.Set(2, RegisterB) }},
	0xD1: &instruction{0xD1, 1, "SET 2,C", func(cpu *CPU) { cpu.Set(2, RegisterC) }},
	0xD2: &instruction{0xD2, 1, "SET 2,D", func(cpu *CPU) { cpu.Set(2, RegisterD) }},
	0xD3: &instruction{0xD3, 1, "SET 2,E", func(cpu *CPU) { cpu.Set(2, RegisterE) }},
	0xD4: &instruction{0xD4, 1, "SET 2,H", func(cpu *CPU) { cpu.Set(2, RegisterH) }},
	0xD5: &instruction{0xD5, 1, "SET 2,L", func(cpu *CPU) { cpu.Set(2, RegisterL) }},

	// Memory[HL][2] <- 1
	0xD6: &instruction{0xD6, 3, "SET 2,(HL)", func(cpu *CPU) { cpu.SetHLDereference(2) }},

	// Register (A, B, C, D, E, H, L)[3] <- 1
	0xDF: &instruction{0xDF, 1, "SET 3,A", func(cpu *CPU) { cpu.SetA(3) }},
	0xD8: &instruction{0xD8, 1, "SET 3,B", func(cpu *CPU) { cpu.Set(3, RegisterB) }},
	0xD9: &instruction{0xD9, 1, "SET 3,C", func(cpu *CPU) { cpu.Set(3, RegisterC) }},
	0xDA: &instruction{0xDA, 1, "SET 3,D", func(cpu *CPU) { cpu.Set(3, RegisterD) }},
	0xDB: &instruction{0xDB, 1, "SET 3,E", func(cpu *CPU) { cpu.Set(3, RegisterE) }},
	0xDC: &instruction{0xDC, 1, "SET 3,H", func(cpu *CPU) { cpu.Set(3, RegisterH) }},
	0xDD: &instruction{0xDD, 1, "SET 3,L", func(cpu *CPU) { cpu.Set(3, RegisterL) }},

	// Memory[HL][3] <- 1
	0xDE: &instruction{0xDE, 3, "SET 3,(HL)", func(cpu *CPU) { cpu.SetHLDereference(3) }},

	// Register (A, B, C, D, E, H, L)[4] <- 1
	0xE7: &instruction{0xE7, 1, "SET 4,A", func(cpu *CPU) { cpu.SetA(4) }},
	0xE0: &instruction{0xE0, 1, "SET 4,B", func(cpu *CPU) { cpu.Set(4, RegisterB) }},
	0xE1: &instruction{0xE1, 1, "SET 4,C", func(cpu *CPU) { cpu.Set(4, RegisterC) }},
	0xE2: &instruction{0xE2, 1, "SET 4,D", func(cpu *CPU) { cpu.Set(4, RegisterD) }},
	0xE3: &instruction{0xE3, 1, "SET 4,E", func(cpu *CPU) { cpu.Set(4, RegisterE) }},
	0xE4: &instruction{0xE4, 1, "SET 4,H", func(cpu *CPU) { cpu.Set(4, RegisterH) }},
	0xE5: &instruction{0xE5, 1, "SET 4,L", func(cpu *CPU) { cpu.Set(4, RegisterL) }},

	// Memory[HL][4] <- 1
	0xE6: &instruction{0xE6, 3, "SET 4,(HL)", func(cpu *CPU) { cpu.SetHLDereference(4) }},

	// Register (A, B, C, D, E, H, L)[5] <- 1
	0xEF: &instruction{0xEF, 1, "SET 5,A", func(cpu *CPU) { cpu.SetA(5) }},
	0xE8: &instruction{0xE8, 1, "SET 5,B", func(cpu *CPU) { cpu.Set(5, RegisterB) }},
	0xE9: &instruction{0xE9, 1, "SET 5,C", func(cpu *CPU) { cpu.Set(5, RegisterC) }},
	0xEA: &instruction{0xEA, 1, "SET 5,D", func(cpu *CPU) { cpu.Set(5, RegisterD) }},
	0xEB: &instruction{0xEB, 1, "SET 5,E", func(cpu *CPU) { cpu.Set(5, RegisterE) }},
	0xEC: &instruction{0xEC, 1, "SET 5,H", func(cpu *CPU) { cpu.Set(5, RegisterH) }},
	0xED: &instruction{0xED, 1, "SET 5,L", func(cpu *CPU) { cpu.Set(5, RegisterL) }},

	// Memory[HL][5] <- 1
	0xEE: &instruction{0xEE, 3, "SET 5,(HL)", func(cpu *CPU) { cpu.SetHLDereference(5) }},

	// Register (A, B, C, D, E, H, L)[6] <- 1
	0xF7: &instruction{0xF7, 1, "SET 6,A", func(cpu *CPU) { cpu.SetA(6) }},
	0xF0: &instruction{0xF0, 1, "SET 6,B", func(cpu *CPU) { cpu.Set(6, RegisterB) }},
	0xF1: &instruction{0xF1, 1, "SET 6,C", func(cpu *CPU) { cpu.Set(6, RegisterC) }},
	0xF2: &instruction{0xF2, 1, "SET 6,D", func(cpu *CPU) { cpu.Set(6, RegisterD) }},
	0xF3: &instruction{0xF3, 1, "SET 6,E", func(cpu *CPU) { cpu.Set(6, RegisterE) }},
	0xF4: &instruction{0xF4, 1, "SET 6,H", func(cpu *CPU) { cpu.Set(6, RegisterH) }},
	0xF5: &instruction{0xF5, 1, "SET 6,L", func(cpu *CPU) { cpu.Set(6, RegisterL) }},

	// Memory[HL][6] <- 1
	0xF6: &instruction{0xF6, 3, "SET 6,(HL)", func(cpu *CPU) { cpu.SetHLDereference(6) }},

	// Register (A, B, C, D, E, H, L)[7] <- 1
	0xFF: &instruction{0xFF, 1, "SET 7,A", func(cpu *CPU) { cpu.SetA(7) }},
	0xF8: &instruction{0xF8, 1, "SET 7,B", func(cpu *CPU) { cpu.Set(7, RegisterB) }},
	0xF9: &instruction{0xF9, 1, "SET 7,C", func(cpu *CPU) { cpu.Set(7, RegisterC) }},
	0xFA: &instruction{0xFA, 1, "SET 7,D", func(cpu *CPU) { cpu.Set(7, RegisterD) }},
	0xFB: &instruction{0xFB, 1, "SET 7,E", func(cpu *CPU) { cpu.Set(7, RegisterE) }},
	0xFC: &instruction{0xFC, 1, "SET 7,H", func(cpu *CPU) { cpu.Set(7, RegisterH) }},
	0xFD: &instruction{0xFD, 1, "SET 7,L", func(cpu *CPU) { cpu.Set(7, RegisterL) }},

	// Memory[HL][7] <- 1
	0xFE: &instruction{0xFE, 3, "SET 7,(HL)", func(cpu *CPU) { cpu.SetHLDereference(7) }},
}
