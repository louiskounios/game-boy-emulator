package cpu

// Register is the type for our individual registers enumeration.
type Register uint8

// Enumerates individual registers.
const (
	RegisterA Register = iota
	RegisterF
	RegisterB
	RegisterC
	RegisterD
	RegisterE
	RegisterH
	RegisterL
	RegisterBC
	RegisterDE
	RegisterHL
	RegisterSP
	RegisterPC
)

func (register Register) String() string {
	switch register {
	case 0:
		return "A"
	case 1:
		return "F"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	case 5:
		return "E"
	case 6:
		return "H"
	case 7:
		return "L"
	case 8:
		return "BC"
	case 9:
		return "DE"
	case 10:
		return "HL"
	case 11:
		return "SP"
	case 12:
		return "PC"
	default:
		return "?"
	}
}

// Flag is the type for our individual flags enumeration.
type Flag uint8

func (flag Flag) String() string {
	switch flag {
	case 4:
		return "C"
	case 5:
		return "H"
	case 6:
		return "N"
	case 7:
		return "Z"
	default:
		return "?"
	}
}

// Enumerates individual flags in the flags register.
//
// Bits 0-3 are unused. Starting the enumeration of the used flags from 4
// makes it easy to use the integer representation of this enum in bitwise
// operations.
const (
	FlagC Flag = iota + 4
	FlagH
	FlagN
	FlagZ
)
