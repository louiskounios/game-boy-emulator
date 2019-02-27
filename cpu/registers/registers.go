package registers

import (
	"errors"
	"fmt"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers/flags"
)

var (
	errUnknownRegister = errors.New("unknown register")
	errFlagsIncrement  = errors.New("flags register cannot be incremented")
	errFlagsDecrement  = errors.New("flags register cannot be decremented")
)

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

// SetRegister sets register rr to value val by modifying r in place.
// It returns an errUnknownRegister error, if encountered.
func (r *Registers) SetRegister(rr Register, val uint16) (err error) {
	switch rr {
	case A:
		r.AF.a = uint8(val)
	case F:
		r.AF.f = flags.Flags(val)
	case B:
		r.BC.hi = uint8(val)
	case C:
		r.BC.lo = uint8(val)
	case D:
		r.DE.hi = uint8(val)
	case E:
		r.DE.lo = uint8(val)
	case H:
		r.HL.hi = uint8(val)
	case L:
		r.HL.lo = uint8(val)
	case AF:
		r.AF.SetWord(val)
	case BC:
		r.BC.SetWord(val)
	case DE:
		r.DE.SetWord(val)
	case HL:
		r.HL.SetWord(val)
	case SP:
		r.SP = val
	case PC:
		r.PC = val
	default:
		err = errUnknownRegister
	}

	return err
}

// Increment increments register rr by 1. r is modified in place.
// It returns an error, if encountered.
func (r *Registers) Increment(rr Register) (err error) {
	switch rr {
	case A:
		r.AF.Increment()
	case B:
		r.BC.hi++
	case C:
		r.BC.lo++
	case D:
		r.DE.hi++
	case E:
		r.DE.lo++
	case H:
		r.HL.hi++
	case L:
		r.HL.lo++
	case BC:
		r.BC.Increment()
	case DE:
		r.DE.Increment()
	case HL:
		r.HL.Increment()
	case SP:
		r.SP++
	case PC:
		r.PC++
	case F, AF:
		err = errFlagsIncrement
	default:
		err = errUnknownRegister
	}

	return err
}

// Decrement decrements register rr by by 1. r is modified in place.
// It returns an error, if encountered.
func (r *Registers) Decrement(rr Register) (err error) {
	switch rr {
	case A:
		r.AF.Decrement()
	case B:
		r.BC.hi--
	case C:
		r.BC.lo--
	case D:
		r.DE.hi--
	case E:
		r.DE.lo--
	case H:
		r.HL.hi--
	case L:
		r.HL.lo--
	case BC:
		r.BC.Decrement()
	case DE:
		r.DE.Decrement()
	case HL:
		r.HL.Decrement()
	case SP:
		r.SP--
	case PC:
		r.PC--
	case F, AF:
		err = errFlagsDecrement
	default:
		err = errUnknownRegister
	}

	return err
}

func (r Registers) String() string {
	s := "[AF=%s | BC=%s | DE=%s | HL=%s | SP=0x%04X | PC=0x%04X]"
	return fmt.Sprintf(s, r.AF, r.BC, r.DE, r.HL, r.SP, r.PC)
}
