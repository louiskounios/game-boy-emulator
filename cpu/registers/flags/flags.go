package flags

import (
	"errors"
	"fmt"
)

var errUnknownFlag = errors.New("unknown flag")

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
// makes it easy to pass this same enum when shifting.
const (
	C Flag = iota + 4
	H
	N
	Z
)

// Flags is an 8-bit register.
type Flags uint8

// Bitmask used for shifting operations when manipulating flags.
const bitmask Flags = 1

// New returns a new 8-bit flags register.
func New() *Flags {
	return new(Flags)
}

// Returns the value flag f is holding. An errUnknownFlag error is returned,
// if encountered.
func (flags *Flags) get(f Flag) (ret uint8, err error) {
	switch f {
	case C, H, N, Z:
		ret = uint8((*flags >> f) & bitmask)
	default:
		err = errUnknownFlag
	}

	return ret, err
}

// IsSet returns true if flag f is set and false otherwise. An errUnknownFlag
// error is returned, if encountered.
func (flags *Flags) IsSet(f uint8) (bool, error) {
	ret, err := flags.get(Flag(f))
	return ret == 1, err
}

// Put sets flag f if set is true, and resets it otherwise. An error is
// returned, if encountered.
func (flags *Flags) Put(f uint8, set bool) error {
	if set {
		return flags.Set(f)
	}

	return flags.Reset(f)
}

// Reset resets flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Reset(f uint8) (err error) {
	switch Flag(f) {
	case C, H, N, Z:
		*flags = *flags &^ Flags(bitmask<<f)
	default:
		err = errUnknownFlag
	}

	return err
}

// Set sets flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Set(f uint8) (err error) {
	switch Flag(f) {
	case C, H, N, Z:
		*flags = *flags | Flags(bitmask<<f)
	default:
		err = errUnknownFlag
	}

	return err
}

// Toggle toggles flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Toggle(f uint8) (err error) {
	switch Flag(f) {
	case C, H, N, Z:
		*flags = *flags ^ Flags(bitmask<<f)
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
	z, _ := flags.get(Z)
	n, _ := flags.get(N)
	h, _ := flags.get(H)
	c, _ := flags.get(C)
	return fmt.Sprintf("%s:%d %s:%d %s:%d %s:%d", Z, z, N, n, H, h, C, c)
}
