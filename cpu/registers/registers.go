package registers

import (
	"errors"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers/flags"
)

var errUnknownRegister = errors.New("unknown register")

// Register is the type for our individual registers enumeration.
type Register uint8

// Enumerates individual registers.
const (
	A Register = iota
	F
	B
	C
	D
	E
	H
	L
	AF
	BC
	DE
	HL
	SP
	PC
)

// Registers consists of six 16-bit registers.
type Registers struct {
	AF RegisterAF
	BC Register16
	DE Register16
	HL Register16
	SP uint16
	PC uint16
}

// NewRegisters returns a new Registers struct.
func NewRegisters() *Registers {
	f := flags.Flags(0)
	af := RegisterAF{
		f: f,
	}

	return &Registers{AF: af}
}

// Register returns a copy of register rr, and an errUnknownRegister error,
// if encountered.
func (r *Registers) Register(rr Register) (ret uint16, err error) {
	switch rr {
	case A:
		ret = uint16(r.AF.a)
	case F:
		ret = uint16(r.AF.f)
	case B:
		ret = uint16(r.BC.hi)
	case C:
		ret = uint16(r.BC.lo)
	case D:
		ret = uint16(r.DE.hi)
	case E:
		ret = uint16(r.DE.lo)
	case H:
		ret = uint16(r.HL.hi)
	case L:
		ret = uint16(r.HL.lo)
	case AF:
		ret = r.AF.Word()
	case BC:
		ret = r.BC.Word()
	case DE:
		ret = r.DE.Word()
	case HL:
		ret = r.HL.Word()
	case SP:
		ret = r.SP
	case PC:
		ret = r.PC
	default:
		err = errUnknownRegister
	}

	return ret, err
}
