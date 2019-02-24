package registers

import (
	"errors"

	"github.com/loizoskounios/game-boy-emulator/byteops"
)

var errUnknownFlag = errors.New("unknown flag")

// Enumerates individual flags in the flags register.
//
// Bits 0-3 are unused. Starting the enumeration of the used flags from 4
// makes it easy to pass this same enum as an argument to the bitshift
// operators.
const (
	C uint8 = iota + 4
	H
	N
	Z
)

// Flags is an 8-bit register.
type Flags uint8

// NewFlags returns a new 8-bit flags register.
func NewFlags() *Flags {
	return new(Flags)
}

// UpdateFlag updates the nth flag of flags using the mutator function provided.
func (flags *Flags) UpdateFlag(n uint8, mutator byteops.Mutator) error {
	switch n {
	case C, H, N, Z:
		return nil
	default:
		return errUnknownFlag
	}
}
