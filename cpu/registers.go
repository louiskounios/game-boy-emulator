package cpu

import (
	"errors"
	"fmt"
)

var (
	errUnknownAuxiliaryRegister = errors.New("unknown auxiliary register")
	errUnknownPairedRegister    = errors.New("unknown paired register")
)

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

// NewRegisters returns new registers.
func NewRegisters() *Registers {
	return &Registers{
		new(uint8),
		NewFlags(),
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
	case RegisterA:
		ret = r.Accumulator()
	case RegisterB:
		ret = r.bc.Hi()
	case RegisterC:
		ret = r.bc.Lo()
	case RegisterD:
		ret = r.de.Hi()
	case RegisterE:
		ret = r.de.Lo()
	case RegisterH:
		ret = r.hl.Hi()
	case RegisterL:
		ret = r.hl.Lo()
	default:
		err = errUnknownAuxiliaryRegister
	}

	return ret, err
}

func (r *Registers) getRegisterPairer(rr Register) (RegisterPairer, error) {
	var rp RegisterPairer

	switch rr {
	case RegisterBC:
		rp = r.bc
	case RegisterDE:
		rp = r.de
	case RegisterHL:
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
func (r *Registers) IsFlagSet(flag Flag) (bool, error) {
	return r.f.IsSet(flag)
}

// PutFlag sets flag if set is true, and resets it otherwise. An error is
// returned, if encountered.
func (r *Registers) PutFlag(flag Flag, set bool) error {
	return r.f.Put(flag, set)
}

// ResetFlag resets flag, and returns an error, if encountered.
func (r *Registers) ResetFlag(flag Flag) error {
	return r.f.Reset(flag)
}

// SetFlag sets flag, and returns an error, if encountered.
func (r *Registers) SetFlag(flag Flag) error {
	return r.f.Set(flag)
}

// ToggleFlag toggles flag, and returns an error, if encountered.
func (r *Registers) ToggleFlag(flag Flag) error {
	return r.f.Toggle(flag)
}

func (r *Registers) String() string {
	s := "[acc=0x%02X | flags=%s | BC=%s | DE=%s | HL=%s | SP=0x%04X | PC=0x%04X]"
	return fmt.Sprintf(s, *r.a, r.f, r.bc, r.de, r.hl, *r.sp, *r.pc)
}
