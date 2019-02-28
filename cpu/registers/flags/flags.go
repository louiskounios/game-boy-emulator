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
// makes it easy to pass this same enum as an argument to the bitshift
// operators.
const (
	C Flag = iota + 4
	H
	N
	Z
)

// Bitmask used for shifting operations when manipulating flags.
const bitmask uint8 = 1

// Flags is an 8-bit register.
type Flags uint8

// New returns a new 8-bit flags register.
func New() *Flags {
	return new(Flags)
}

// Get returns the value flag f is holding. An errUnknownFlag error is returned,
// if encountered.
func (flags Flags) Get(f Flag) (ret uint8, err error) {
	switch f {
	case C, H, N, Z:
		ret = uint8(flags>>f) & bitmask
	default:
		err = errUnknownFlag
	}

	return ret, err
}

// IsSet returns true if flag f is set and false otherwise. An errUnknownFlag
// error is returned, if encountered.
func (flags Flags) IsSet(f Flag) (bool, error) {
	ret, err := flags.Get(f)
	return ret == 1, err
}

// Reset resets flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Reset(f Flag) (err error) {
	switch f {
	case C, H, N, Z:
		*flags = *flags &^ Flags(bitmask<<f)
	default:
		err = errUnknownFlag
	}

	return err
}

// Set sets flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Set(f Flag) (err error) {
	switch f {
	case C, H, N, Z:
		*flags = *flags | Flags(bitmask<<f)
	default:
		err = errUnknownFlag
	}

	return err
}

// Toggle toggles flag f. flags is modified in place. An errUnknownFlag error is
// returned, if encountered.
func (flags *Flags) Toggle(f Flag) (err error) {
	switch f {
	case C, H, N, Z:
		*flags = *flags ^ Flags(bitmask<<f)
	default:
		err = errUnknownFlag
	}

	return err
}

func (flags Flags) String() string {
	z, _ := flags.Get(Z)
	n, _ := flags.Get(N)
	h, _ := flags.Get(H)
	c, _ := flags.Get(C)
	return fmt.Sprintf("%s:%d %s:%d %s:%d %s:%d", Z, z, N, n, H, h, C, c)
}
