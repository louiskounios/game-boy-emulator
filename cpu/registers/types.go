package registers

import (
	"fmt"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers/flags"
)

// CompoundRegisterHandler is the interface that wraps the functionality that
// must be provided by an implementation of a 16-bit compound register.
type CompoundRegisterHandler interface {
	Equals(interface{}) bool
	Hi() uint8
	SetHi(uint8)
	Lo() uint8
	SetLo(uint8)
	Word() uint16
	SetWord(uint16)
}

// Register16 is a type representing 16-bit compound registers which are
// made up of two uint8 fields.
type Register16 struct {
	hi uint8 // Most significant.
	lo uint8 // Least significant.
}

// Equals compares receiver r and parameter i for equality by comparing their
// fields. It returns false if i is not of type Register16.
func (r *Register16) Equals(i interface{}) bool {
	switch t := i.(type) {
	case Register16:
		return (r.hi == t.hi) && (r.lo == t.lo)
	case *Register16:
		return (r.hi == t.hi) && (r.lo == t.lo)
	default:
		return false
	}
}

// Hi returns a copy of the 8 most significant bits in the 16-bit compound
// register.
func (r *Register16) Hi() uint8 {
	return r.hi
}

// SetHi sets r.hi to be equal to val.
func (r *Register16) SetHi(val uint8) {
	r.hi = val
}

// Lo returns a copy of the 8 least significant bits in the 16-bit compound
// register.
func (r *Register16) Lo() uint8 {
	return r.lo
}

// SetLo sets r.lo to be equal to val.
func (r *Register16) SetLo(val uint8) {
	r.lo = val
}

// Word combines the two fields in r and returns a uint16 variable.
//
// The 8 bits from r.hi and r.lo are used as the most and least significant bits
// respectively.
func (r *Register16) Word() uint16 {
	return uint16(r.hi)<<8 | uint16(r.lo)
}

// SetWord updates the two fields in r so that the uint16 variable that is the
// result of combining them is equal to val.
//
// r.hi gets the 8 most significant bits of val.
// R.lo gets the 8 least significant bits of val.
func (r *Register16) SetWord(val uint16) {
	// Shift the 8 most significant bits to the right, then cast to 8-bit uint.
	r.hi = uint8(val >> 8)
	// Cast to 8-bit uint. The 8 most significant bits will be truncated.
	r.lo = uint8(val)
}

// IncrementBy increments the 16-bit combined version of the register by `by`.
func (r *Register16) IncrementBy(by uint16) {
	w := r.Word()
	w += by
	r.SetWord(w)
}

// Increment increments the 16-bit combined version of the register by 1.
func (r *Register16) Increment() {
	r.IncrementBy(1)
}

// DecrementBy decrements the 16-bit combined version of the register by `by`.
func (r *Register16) DecrementBy(by uint16) {
	w := r.Word()
	w -= by
	r.SetWord(w)
}

// Decrement decrements the 16-bit combined version of the register by 1.
func (r *Register16) Decrement() {
	r.DecrementBy(1)
}

func (r Register16) String() string {
	return fmt.Sprintf("0x%04X", r.Word())
}

// RegisterAF is a type representing the AF compound register which is made up
// of a uint8 variable and an 8-bit flags register.
type RegisterAF struct {
	a uint8       // Most significant.
	f flags.Flags // Least significant.
}

// Equals compares receiver r and parameter i for equality by comparing their
// fields. It returns false if i is not of type RegisterAF.
func (r RegisterAF) Equals(i interface{}) bool {
	switch t := i.(type) {
	case RegisterAF:
		return (r.a == t.a) && (r.f == t.f)
	case *RegisterAF:
		return (r.a == t.a) && (r.f == t.f)
	default:
		return false
	}
}

// Hi returns a copy of the 8 most significant bits in the 16-bit compound
// register.
func (r *RegisterAF) Hi() uint8 {
	return r.a
}

// SetHi sets r.a to be equal to val.
func (r *RegisterAF) SetHi(val uint8) {
	r.a = val
}

// Lo returns a copy of the 8 least significant bits in the 16-bit compound
// register.
func (r *RegisterAF) Lo() uint8 {
	return uint8(r.f)
}

// SetLo sets r.f to be equal to val.
func (r *RegisterAF) SetLo(val uint8) {
	r.f = flags.Flags(val)
}

// Word combines the two fields in r and returns a uint16 variable.
//
// The 8 bits from r.a and r.f are used as the most and least significant bits
// respectively.
func (r RegisterAF) Word() uint16 {
	return uint16(r.a)<<8 | uint16(r.f)
}

// SetWord updates the two fields in r so that the uint16 variable that is the
// result of combining them is equal to val.
//
// r.a gets the 8 most significant bits of val.
// R.f gets the 8 least significant bits of val.
func (r *RegisterAF) SetWord(val uint16) {
	// Shift the 8 most significant bits to the right, then cast to 8-bit uint.
	r.a = uint8(val >> 8)
	// Cast to flags.Flags. The 8 most significant bits will be truncated.
	r.f = flags.Flags(val)
}

// IncrementBy increments the 8-bit register A by `by`.
func (r *RegisterAF) IncrementBy(by uint16) {
	r.a += uint8(by)
}

// Increment increments the 8-bit register A by 1.
func (r *RegisterAF) Increment() {
	r.IncrementBy(1)
}

// DecrementBy decrements the 8-bit register A by `by`.
func (r *RegisterAF) DecrementBy(by uint16) {
	r.a -= uint8(by)
}

// Decrement decrements the 8-bit register A by 1.
func (r *RegisterAF) Decrement() {
	r.DecrementBy(1)
}

func (r RegisterAF) String() string {
	return fmt.Sprintf("0x%04X", r.Word())
}
