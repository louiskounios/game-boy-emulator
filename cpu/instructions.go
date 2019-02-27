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
	 * 8-bit Operations
	 */

	// Register (A, B, C, D, E, H, L) -> Register (B)
	0x40: &instruction{0x40, 1, "LD B,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.B) }},
	0x41: &instruction{0x41, 1, "LD B,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.B) }},
	0x42: &instruction{0x42, 1, "LD B,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.B) }},
	0x43: &instruction{0x43, 1, "LD B,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.B) }},
	0x44: &instruction{0x44, 1, "LD B,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.B) }},
	0x45: &instruction{0x45, 1, "LD B,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.B) }},
	0x47: &instruction{0x47, 1, "LD B,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.B) }},

	// Register (A, B, C, D, E, H, L) -> Register (C)
	0x48: &instruction{0x48, 1, "LD C,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.C) }},
	0x49: &instruction{0x49, 1, "LD C,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.C) }},
	0x4A: &instruction{0x4A, 1, "LD C,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.C) }},
	0x4B: &instruction{0x4B, 1, "LD C,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.C) }},
	0x4C: &instruction{0x4C, 1, "LD C,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.C) }},
	0x4D: &instruction{0x4D, 1, "LD C,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.C) }},
	0x4F: &instruction{0x4F, 1, "LD C,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.C) }},

	// Register (A, B, C, D, E, H, L) -> Register (D)
	0x50: &instruction{0x50, 1, "LD D,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.D) }},
	0x51: &instruction{0x51, 1, "LD D,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.D) }},
	0x52: &instruction{0x52, 1, "LD D,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.D) }},
	0x53: &instruction{0x53, 1, "LD D,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.D) }},
	0x54: &instruction{0x54, 1, "LD D,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.D) }},
	0x55: &instruction{0x55, 1, "LD D,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.D) }},
	0x57: &instruction{0x57, 1, "LD D,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.D) }},

	// Register (A, B, C, D, E, H, L) -> Register (E)
	0x58: &instruction{0x58, 1, "LD E,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.E) }},
	0x59: &instruction{0x59, 1, "LD E,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.E) }},
	0x5A: &instruction{0x5A, 1, "LD E,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.E) }},
	0x5B: &instruction{0x5B, 1, "LD E,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.E) }},
	0x5C: &instruction{0x5C, 1, "LD E,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.E) }},
	0x5D: &instruction{0x5D, 1, "LD E,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.E) }},
	0x5F: &instruction{0x5F, 1, "LD E,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.E) }},

	// Register (A, B, C, D, E, H, L) -> Register (H)
	0x60: &instruction{0x60, 1, "LD H,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.H) }},
	0x61: &instruction{0x61, 1, "LD H,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.H) }},
	0x62: &instruction{0x62, 1, "LD H,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.H) }},
	0x63: &instruction{0x63, 1, "LD H,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.H) }},
	0x64: &instruction{0x64, 1, "LD H,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.H) }},
	0x65: &instruction{0x65, 1, "LD H,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.H) }},
	0x67: &instruction{0x67, 1, "LD H,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.H) }},

	// Register (A, B, C, D, E, H, L) -> Register (L)
	0x68: &instruction{0x68, 1, "LD L,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.L) }},
	0x69: &instruction{0x69, 1, "LD L,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.L) }},
	0x6A: &instruction{0x6A, 1, "LD L,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.L) }},
	0x6B: &instruction{0x6B, 1, "LD L,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.L) }},
	0x6C: &instruction{0x6C, 1, "LD L,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.L) }},
	0x6D: &instruction{0x6D, 1, "LD L,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.L) }},
	0x6F: &instruction{0x6F, 1, "LD L,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.L) }},

	// Register (A, B, C, D, E, H, L) -> Register (A)
	0x78: &instruction{0x78, 1, "LD A,B", func(cpu *CPU) { cpu.PutRIntoR(registers.B, registers.A) }},
	0x79: &instruction{0x79, 1, "LD A,C", func(cpu *CPU) { cpu.PutRIntoR(registers.C, registers.A) }},
	0x7A: &instruction{0x7A, 1, "LD A,D", func(cpu *CPU) { cpu.PutRIntoR(registers.D, registers.A) }},
	0x7B: &instruction{0x7B, 1, "LD A,E", func(cpu *CPU) { cpu.PutRIntoR(registers.E, registers.A) }},
	0x7C: &instruction{0x7C, 1, "LD A,H", func(cpu *CPU) { cpu.PutRIntoR(registers.H, registers.A) }},
	0x7D: &instruction{0x7D, 1, "LD A,L", func(cpu *CPU) { cpu.PutRIntoR(registers.L, registers.A) }},
	0x7F: &instruction{0x7F, 1, "LD A,A", func(cpu *CPU) { cpu.PutRIntoR(registers.A, registers.A) }},

	// Register (A, B, C, D, E, H, L) -> Memory (address=Register HL)
	0x70: &instruction{0x70, 2, "LD (HL),B", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.B) }},
	0x71: &instruction{0x71, 2, "LD (HL),C", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.C) }},
	0x72: &instruction{0x72, 2, "LD (HL),D", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.D) }},
	0x73: &instruction{0x73, 2, "LD (HL),E", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.E) }},
	0x74: &instruction{0x74, 2, "LD (HL),H", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.H) }},
	0x75: &instruction{0x75, 2, "LD (HL),L", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.L) }},
	0x77: &instruction{0x77, 2, "LD (HL),A", func(cpu *CPU) { cpu.PutRIntoHLAddress(registers.A) }},

	// Register (A) -> Memory (address=Register BC)
	0x02: &instruction{0x02, 2, "LD (BC),A", func(cpu *CPU) { cpu.PutAIntoBCAddress() }},

	// Register (A) -> Memory (address=Register DE)
	0x12: &instruction{0x12, 2, "LD (DE),A", func(cpu *CPU) { cpu.PutAIntoDEAddress() }},

	// Memory (address=PC) -> Register (B, C, D, E, H, L, A)
	0x06: &instruction{0x06, 2, "LD B,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.B) }},
	0x0E: &instruction{0x0E, 2, "LD C,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.C) }},
	0x16: &instruction{0x16, 2, "LD D,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.D) }},
	0x1E: &instruction{0x1E, 2, "LD E,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.E) }},
	0x26: &instruction{0x26, 2, "LD H,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.H) }},
	0x2E: &instruction{0x2E, 2, "LD L,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.L) }},
	0x3E: &instruction{0x3E, 2, "LD A,d8", func(cpu *CPU) { cpu.PutNIntoR(registers.A) }},

	// Memory (address=PC) -> Memory (address=Register HL)
	0x36: &instruction{0x36, 3, "LD (HL),d8", func(cpu *CPU) { cpu.PutNIntoHLAddress() }},

	// Memory (address=Register BC) -> Register (A)
	0x0A: &instruction{0x0A, 2, "LD A,(BC)", func(cpu *CPU) { cpu.PutBCDereferenceIntoA() }},

	// Memory (address=Register DE) -> Register (A)
	0x1A: &instruction{0x1A, 2, "LD A,(DE)", func(cpu *CPU) { cpu.PutDEDereferenceIntoA() }},

	// Memory (address=Register HL) -> Register (B, C, D, E, H, L, A)
	0x46: &instruction{0x46, 2, "LD B,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.B) }},
	0x4E: &instruction{0x4E, 2, "LD C,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.C) }},
	0x56: &instruction{0x56, 2, "LD D,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.D) }},
	0x5E: &instruction{0x5E, 2, "LD E,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.E) }},
	0x66: &instruction{0x66, 2, "LD H,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.H) }},
	0x6E: &instruction{0x6E, 2, "LD L,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.L) }},
	0x7E: &instruction{0x7E, 2, "LD A,(HL)", func(cpu *CPU) { cpu.PutHLDereferenceIntoR(registers.A) }},
}
