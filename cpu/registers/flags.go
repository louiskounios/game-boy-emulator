package registers

import (
	"errors"

	"github.com/loizoskounios/game-boy-emulator/byteops"
)

var errUnknownFlag = errors.New("unknown flag")

type flag uint8

// Enumerates individual flags in the flags register.
const (
	CY flag = iota
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
	case CY:
		return mutator(&(flags.val), 4)
	case H:
		return mutator(&(flags.val), 5)
	case N:
		return mutator(&(flags.val), 6)
	case ZF:
		return mutator(&(flags.val), 7)
	default:
		return errUnknownFlag
	}
}
