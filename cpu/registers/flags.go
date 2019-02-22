package registers

import (
	"errors"

	"github.com/loizoskounios/game-boy-emulator/byteops"
)

var errUnknownFlag = errors.New("unknown flag")

type flag uint8

// Enumerates individual flags in the flags register.
//
// Bits 0-3 are unused. Starting the enumeration of the used flags from 4
// makes it easy to pass this same enum as an argument to the bitshift
// operators.
const (
	CY flag = iota + 4
	H
	N
	ZF
)

type flags struct {
	val byte
}

// NewFlags returns a new 8-bit flags register.
func NewFlags() *flags {
	return &flags{}
}

// Flags returns a copy of the value of the 8-bit flags register.
func (flags *flags) Flags() byte {
	return flags.val
}

func (flags *flags) UpdateFlag(flag flag, mutator byteops.Mutator) error {
	switch flag {
	case CY, H, N, ZF:
		return mutator(&(flags.val), uint8(flag))
	default:
		return errUnknownFlag
	}
}
