package registers

import "fmt"

// RegisterPairer is the interface that wraps the functionality that must be
// provided by an implementation of a 16-bit register pair consisting of two
// uint8 sub-registers.
type RegisterPairer interface {
	Word() uint16
	SetWord(uint16)
	Hi() *uint8
	Lo() *uint8
	Decrement()
	DecrementBy(uint16)
	Increment()
	IncrementBy(uint16)
}

// RegisterPair is a type representing a register pair made up of two uint8
// sub-registers.
type RegisterPair [2]uint8

// NewRegisterPair returns a pointer to a new register pair.
func NewRegisterPair() *RegisterPair {
	return &RegisterPair{}
}

// Hi returns a pointer to the most significant sub-register.
func (rp *RegisterPair) Hi() *uint8 {
	return &rp[0]
}

// Lo returns pointer to the least significant sub-register.
func (rp *RegisterPair) Lo() *uint8 {
	return &rp[1]
}

// Word returns the resulting 16-bit value of the register pair when its
// sub-registers are combined.
func (rp *RegisterPair) Word() uint16 {
	return uint16(rp[0])<<8 | uint16(rp[1])
}

// SetWord updates the two sub-registers in rp so that the combined value of the
// two sub-registers is equal to the received value.
func (rp *RegisterPair) SetWord(val uint16) {
	rp[0] = uint8(val >> 8)
	rp[1] = uint8(val)
}

// IncrementBy increments the paired register by the provided amount.
func (rp *RegisterPair) IncrementBy(amount uint16) {
	w := rp.Word()
	w += amount
	rp.SetWord(w)
}

// Increment increments the paired register by 1.
func (rp *RegisterPair) Increment() {
	rp.IncrementBy(1)
}

// DecrementBy decrements the paired register by the provided amount.
func (rp *RegisterPair) DecrementBy(amount uint16) {
	w := rp.Word()
	w -= amount
	rp.SetWord(w)
}

// Decrement decrements the paired register by 1.
func (rp *RegisterPair) Decrement() {
	rp.DecrementBy(1)
}

func (rp *RegisterPair) String() string {
	return fmt.Sprintf("0x%04X", rp.Word())
}
