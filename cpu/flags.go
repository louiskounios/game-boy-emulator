package cpu

import (
	"errors"
	"fmt"
)

var errUnknownFlag = errors.New("unknown flag")

// FlagBearer is the interface that wraps the functionality that must be
// provided by an implementation of an 8-bit flags register.
type FlagBearer interface {
	IsSet(Flag) (bool, error)
	Put(Flag, bool) error
	Reset(Flag) error
	Set(Flag) error
	Toggle(Flag) error
	Value() uint8
	SetValue(uint8)
}

// Flags is an 8-bit register.
type Flags uint8

// Bitmask used for shifting operations when manipulating flags.
const bitmask Flags = 1

// NewFlags returns a new 8-bit flags register.
func NewFlags() *Flags {
	return new(Flags)
}

// Returns the value flag f is holding. An errUnknownFlag error is returned,
// if encountered.
func (flags *Flags) get(f Flag) (ret uint8, err error) {
	switch f {
	case FlagC, FlagH, FlagN, FlagZ:
		ret = uint8((*flags >> f) & bitmask)
	default:
		err = errUnknownFlag
	}

	return ret, err
}

// IsSet returns true if flag f is set and false otherwise. An errUnknownFlag
// error is returned, if encountered.
func (flags *Flags) IsSet(f Flag) (bool, error) {
	ret, err := flags.get(f)
	return ret == 1, err
}

// Put sets flag f if set is true, and resets it otherwise. An error is
// returned, if encountered.
func (flags *Flags) Put(f Flag, set bool) error {
	if set {
		return flags.Set(f)
	}

	return flags.Reset(f)
}

// Reset resets flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Reset(f Flag) (err error) {
	switch f {
	case FlagC, FlagH, FlagN, FlagZ:
		*flags = *flags &^ Flags(bitmask<<f)
	default:
		err = errUnknownFlag
	}

	return err
}

// Set sets flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Set(f Flag) (err error) {
	switch Flag(f) {
	case FlagC, FlagH, FlagN, FlagZ:
		*flags = *flags | (bitmask << f)
	default:
		err = errUnknownFlag
	}

	return err
}

// Toggle toggles flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Toggle(f Flag) (err error) {
	switch Flag(f) {
	case FlagC, FlagH, FlagN, FlagZ:
		*flags = *flags ^ (bitmask << f)
	default:
		err = errUnknownFlag
	}

	return err
}

// Value returns a uint8 representation of the flags register.
func (flags *Flags) Value() uint8 {
	return uint8(*flags)
}

// SetValue sets the flag register to be equal to the provided value.
func (flags *Flags) SetValue(val uint8) {
	*flags = Flags(val)
}

func (flags *Flags) String() string {
	c, _ := flags.get(FlagC)
	h, _ := flags.get(FlagH)
	n, _ := flags.get(FlagN)
	z, _ := flags.get(FlagZ)
	return fmt.Sprintf("%s:%d %s:%d %s:%d %s:%d", FlagZ, z, FlagN, n, FlagH, h, FlagC, c)
}
