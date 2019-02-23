package registers

type (
	// Register8 is a type representing 8-bit registers.
	Register8 uint8

	// Register16 is a type representing 16-bit registers.
	Register16 uint16

	// Register16C is a type representing 16-bit compound registers which are
	// made up from two 8-bit registers.
	Register16C struct {
		r1 Register8
		r2 Register8
	}
)
