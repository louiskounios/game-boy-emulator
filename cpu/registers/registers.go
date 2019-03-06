package registers

import (
	"errors"
	"fmt"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers/flags"
)

var (
	errUnknownAuxiliaryRegister = errors.New("unknown auxiliary register")
	errUnknownPairedRegister    = errors.New("unknown paired register")
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
	BC
	DE
	HL
	SP
	PC
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

// FlagBearer is the interface that wraps the functionality that must be
// provided by an implementation of an 8-bit flags register.
type FlagBearer interface {
	IsSet(uint8) (bool, error)
	Put(uint8, bool) error
	Reset(uint8) error
	Set(uint8) error
	Toggle(uint8) error
	Value() uint8
	SetValue(uint8)
}

// Registers consists of the accumulator and flags registers, three paired
// registers consisting of two 8-bit sub-registers, a stack pointer,
// and a program counter.
type Registers struct {
	a  *uint8
	f  FlagBearer
	bc RegisterPairer
	de RegisterPairer
	hl RegisterPairer
	sp *uint16
	pc *uint16
}

// New returns new registers.
func New() *Registers {
	return &Registers{
		new(uint8),
		flags.New(),
		NewRegisterPair(),
		NewRegisterPair(),
		NewRegisterPair(),
		new(uint16),
		new(uint16),
	}
}

// AF returns the combined 16-bit value of the accumulator and flags registers.
func (r *Registers) AF() uint16 {
	return uint16(*r.a)<<8 | uint16(r.f.Value())
}

// SetAF sets the 16-bit value when combining the accumulator and flags
// registers to be equal to the provided value.
func (r *Registers) SetAF(val uint16) {
	*r.a = uint8(val >> 8)
	r.f.SetValue(uint8(val))
}

// Accumulator returns a pointer to the accumulator register.
func (r *Registers) Accumulator() *uint8 {
	return r.a
}

// Auxiliary returns a pointer to the requested auxiliary register, and an
// error if encountered.
func (r *Registers) Auxiliary(rr Register) (ret *uint8, err error) {
	switch rr {
	case A:
		ret = r.Accumulator()
	case B:
		ret = r.bc.Hi()
	case C:
		ret = r.bc.Lo()
	case D:
		ret = r.de.Hi()
	case E:
		ret = r.de.Lo()
	case H:
		ret = r.hl.Hi()
	case L:
		ret = r.hl.Lo()
	default:
		err = errUnknownAuxiliaryRegister
	}

	return ret, err
}

func (r *Registers) getRegisterPairer(rr Register) (RegisterPairer, error) {
	var rp RegisterPairer

	switch rr {
	case BC:
		rp = r.bc
	case DE:
		rp = r.de
	case HL:
		rp = r.hl
	default:
		return nil, errUnknownPairedRegister
	}

	return rp, nil
}

// Paired returns the 16-bit value of a paired register.
func (r *Registers) Paired(rr Register) (uint16, error) {
	rp, err := r.getRegisterPairer(rr)
	if err != nil {
		return 0, err
	}
	return rp.Word(), nil
}

// DecrementPairedBy decrements paired register rr by the provided amount.
func (r *Registers) DecrementPairedBy(rr Register, amount uint16) error {
	rp, err := r.getRegisterPairer(rr)
	if err != nil {
		return err
	}
	rp.DecrementBy(amount)
	return nil
}

// DecrementPaired decrements paired register rr by 1.
func (r *Registers) DecrementPaired(rr Register) error {
	rp, err := r.getRegisterPairer(rr)
	if err != nil {
		return err
	}
	rp.Decrement()
	return nil
}

// IncrementPairedBy increments paired register rr by the provided amount.
func (r *Registers) IncrementPairedBy(rr Register, amount uint16) error {
	rp, err := r.getRegisterPairer(rr)
	if err != nil {
		return err
	}
	rp.IncrementBy(amount)
	return nil
}

// IncrementPaired increments paired register rr by 1.
func (r *Registers) IncrementPaired(rr Register) error {
	rp, err := r.getRegisterPairer(rr)
	if err != nil {
		return err
	}
	rp.Increment()
	return nil
}

// SetPaired sets the 16-bit value of a paired register to be equal to the
// provided value.
func (r *Registers) SetPaired(rr Register, val uint16) error {
	rp, err := r.getRegisterPairer(rr)
	if err != nil {
		return err
	}
	rp.SetWord(val)
	return nil
}

// StackPointer returns a pointer to the stack pointer.
func (r *Registers) StackPointer() *uint16 {
	return r.sp
}

// ProgramCounter returns a pointer to the program counter.
func (r *Registers) ProgramCounter() *uint16 {
	return r.pc
}

// IsFlagSet returns whether flag is set or not and an error, if encountered.
func (r *Registers) IsFlagSet(flag uint8) (bool, error) {
	return r.f.IsSet(flag)
}

// PutFlag sets flag if set is true, and resets it otherwise. An error is
// returned, if encountered.
func (r *Registers) PutFlag(flag uint8, set bool) error {
	return r.f.Put(flag, set)
}

// ResetFlag resets flag, and returns an error, if encountered.
func (r *Registers) ResetFlag(flag uint8) error {
	return r.f.Reset(flag)
}

// SetFlag sets flag, and returns an error, if encountered.
func (r *Registers) SetFlag(flag uint8) error {
	return r.f.Set(flag)
}

// ToggleFlag toggles flag, and returns an error, if encountered.
func (r *Registers) ToggleFlag(flag uint8) error {
	return r.f.Toggle(flag)
}

func (r *Registers) String() string {
	s := "[acc=0x%02X | flags=%s | BC=%s | DE=%s | HL=%s | SP=0x%04X | PC=0x%04X]"
	return fmt.Sprintf(s, *r.a, r.f, r.bc, r.de, r.hl, *r.sp, *r.pc)
}
