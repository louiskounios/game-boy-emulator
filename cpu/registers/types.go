package registers

type (
	// Register8 is a type representing 8-bit registers.
	Register8 uint8

	// Register16 is a type representing 16-bit registers.
	Register16 uint16

	// Register16C is a type representing 16-bit compound registers which are
	// made up from two 8-bit registers.
	Register16C struct {
		R1 Register8 // Most significant.
		R2 Register8 // Least significant.
	}
)

// Equals compares the receiver r and the parameter rr for equality.
func (r Register16C) Equals(rr *Register16C) bool {
	return (r.R1 == rr.R1) && (r.R2 == rr.R2)
}

// Word returns a 16-bit register that is the result of combining the two 8-bit
// registers that exist in a 16-bit compound register.
//
// The 8 bits from r.R1 and r.R2 become the most and least significant bits of
// the resulting 16-bit register.
func (r Register16C) Word() Register16 {
	return Register16(r.R1)<<8 ^ Register16(r.R2)
}

// SetWord takes a 16-bit register rr and performs the necessary bit
// manipulations and modifies r in place to generate the equivalent
// 16-bit compound register.
//
// The first 8-bit register in r gets the 8 most significant bits of rr.
// The second 8-bit register in r gets the 8 least significant bits of rr.
func (r *Register16C) SetWord(rr Register16) {
	// Shift the 8 most significant bits to the right, then cast to 8-bit uint.
	r.R1 = Register8(rr >> 8)
	// Cast to 8-bit uint. The 8 most significant bits will be truncated.
	r.R2 = Register8(rr)
}
