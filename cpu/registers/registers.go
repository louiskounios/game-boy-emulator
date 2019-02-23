package registers

// The AF register is a 16-bit compound register that consists of a standard
// 8-bit register and the 8-bit flags register.
type afRegister struct {
	R1 Register8
	R2 *Flags
}

// Registers consists of six 16-bit registers.
type Registers struct {
	AF afRegister
	BC Register16C
	DE Register16C
	HL Register16C
	SP Register16
	PC Register16
}

// NewRegisters returns a new Registers struct.
func NewRegisters() *Registers {
	f := NewFlags()
	af := afRegister{
		R2: f,
	}

	return &Registers{AF: af}
}
