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
	af RegisterAF
	bc Register16
	de Register16
	hl Register16
	sp uint16
	pc uint16
}

// NewRegisters returns a new Registers struct.
func NewRegisters() *Registers {
	f := flags.Flags(0)
	af := RegisterAF{
		f: f,
	}

	return &Registers{af: af}
}

// Register returns a copy of register rr, and an errUnknownRegister error,
// if encountered.
func (r *Registers) Register(rr Register) (ret uint16, err error) {
	switch rr {
	case A:
		ret = uint16(r.af.a)
	case F:
		ret = uint16(r.af.f)
	case B:
		ret = uint16(r.bc.hi)
	case C:
		ret = uint16(r.bc.lo)
	case D:
		ret = uint16(r.de.hi)
	case E:
		ret = uint16(r.de.lo)
	case H:
		ret = uint16(r.hl.hi)
	case L:
		ret = uint16(r.hl.lo)
	case AF:
		ret = r.af.Word()
	case BC:
		ret = r.bc.Word()
	case DE:
		ret = r.de.Word()
	case HL:
		ret = r.hl.Word()
	case SP:
		ret = r.sp
	case PC:
		ret = r.pc
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
		r.af.a = uint8(val)
	case F:
		r.af.f = flags.Flags(val)
	case B:
		r.bc.hi = uint8(val)
	case C:
		r.bc.lo = uint8(val)
	case D:
		r.de.hi = uint8(val)
	case E:
		r.de.lo = uint8(val)
	case H:
		r.hl.hi = uint8(val)
	case L:
		r.hl.lo = uint8(val)
	case AF:
		r.af.SetWord(val)
	case BC:
		r.bc.SetWord(val)
	case DE:
		r.de.SetWord(val)
	case HL:
		r.hl.SetWord(val)
	case SP:
		r.sp = val
	case PC:
		r.pc = val
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
		r.af.Increment()
	case B:
		r.bc.hi++
	case C:
		r.bc.lo++
	case D:
		r.de.hi++
	case E:
		r.de.lo++
	case H:
		r.hl.hi++
	case L:
		r.hl.lo++
	case BC:
		r.bc.Increment()
	case DE:
		r.de.Increment()
	case HL:
		r.hl.Increment()
	case SP:
		r.sp++
	case PC:
		r.pc++
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
		r.af.Decrement()
	case B:
		r.bc.hi--
	case C:
		r.bc.lo--
	case D:
		r.de.hi--
	case E:
		r.de.lo--
	case H:
		r.hl.hi--
	case L:
		r.hl.lo--
	case BC:
		r.bc.Decrement()
	case DE:
		r.de.Decrement()
	case HL:
		r.hl.Decrement()
	case SP:
		r.sp--
	case PC:
		r.pc--
	case F, AF:
		err = errFlagsDecrement
	default:
		err = errUnknownRegister
	}

	return err
}

func (r Registers) GetFlag(flag flags.Flag) (uint8, error) {
	return r.af.f.Get(flag)
}

func (r Registers) IsFlagSet(flag flags.Flag) (bool, error) {
	return r.af.f.IsSet(flag)
}

func (r *Registers) ResetFlag(flag flags.Flag) error {
	return r.af.f.Reset(flag)
}

func (r *Registers) SetFlag(flag flags.Flag) error {
	return r.af.f.Set(flag)
}

func (r *Registers) ToggleFlag(flag flags.Flag) error {
	return r.af.f.Toggle(flag)
}

func (r Registers) String() string {
	s := "[AF=%s | BC=%s | DE=%s | HL=%s | SP=0x%04X | PC=0x%04X]"
	return fmt.Sprintf(s, r.af, r.bc, r.de, r.hl, r.sp, r.pc)
}
