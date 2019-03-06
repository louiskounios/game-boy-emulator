package flags_test

import (
	"fmt"
	"testing"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers/flags"
)

var getTests = []struct {
	val  flags.Flags
	flag flags.Flag
	ret  uint8
}{
	{16, flags.C, 1},
	{16, flags.Z, 0},
	{32, flags.H, 1},
	{32, flags.N, 0},
	{64, flags.N, 1},
	{64, flags.H, 0},
	{128, flags.Z, 1},
	{128, flags.C, 0},
}

var isSetTests = []struct {
	flags *flags.Flags
	flag  flags.Flag
	ret   bool
}{
	{newFlags(16), flags.C, true},
	{newFlags(16), flags.Z, false},
	{newFlags(32), flags.H, true},
	{newFlags(32), flags.N, false},
	{newFlags(64), flags.N, true},
	{newFlags(64), flags.H, false},
	{newFlags(128), flags.Z, true},
	{newFlags(128), flags.C, false},
}

func TestIsSet(t *testing.T) {
	for _, tt := range isSetTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			if out, _ := tt.flags.IsSet(uint8(tt.flag)); out != tt.ret {
				t.Errorf("got %t, expected %t", out, tt.ret)
			}
		})
	}
}

var putTests = []struct {
	flags *flags.Flags
	flag  flags.Flag
	set   bool
}{
	{newFlags(0), flags.C, true},
	{newFlags(0), flags.C, false},
	{newFlags(16), flags.C, true},
	{newFlags(16), flags.C, false},
	{newFlags(0), flags.H, true},
	{newFlags(0), flags.H, false},
	{newFlags(32), flags.H, true},
	{newFlags(32), flags.H, false},
	{newFlags(0), flags.N, true},
	{newFlags(0), flags.N, false},
	{newFlags(64), flags.N, true},
	{newFlags(64), flags.N, false},
	{newFlags(0), flags.Z, true},
	{newFlags(0), flags.Z, false},
	{newFlags(128), flags.Z, true},
	{newFlags(128), flags.Z, false},
}

func TestPut(t *testing.T) {
	for _, tt := range putTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			tt.flags.Put(uint8(tt.flag), tt.set)

			if out, _ := tt.flags.IsSet(uint8(tt.flag)); out != tt.set {
				t.Errorf("got %t, expected %t", out, false)
			}
		})
	}
}

var resetTests = []struct {
	flags *flags.Flags
	flag  flags.Flag
}{
	{newFlags(16), flags.C},
	{newFlags(16), flags.Z},
	{newFlags(32), flags.H},
	{newFlags(32), flags.N},
	{newFlags(64), flags.N},
	{newFlags(64), flags.H},
	{newFlags(128), flags.Z},
	{newFlags(128), flags.C},
}

func TestReset(t *testing.T) {
	for _, tt := range resetTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			tt.flags.Reset(uint8(tt.flag))

			if out, _ := tt.flags.IsSet(uint8(tt.flag)); out {
				t.Errorf("got %t, expected %t", out, false)
			}
		})
	}
}

var setTests = []struct {
	flags *flags.Flags
	flag  flags.Flag
}{
	{newFlags(16), flags.C},
	{newFlags(16), flags.Z},
	{newFlags(32), flags.H},
	{newFlags(32), flags.N},
	{newFlags(64), flags.N},
	{newFlags(64), flags.H},
	{newFlags(128), flags.Z},
	{newFlags(128), flags.C},
}

func TestSet(t *testing.T) {
	for _, tt := range setTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			tt.flags.Set(uint8(tt.flag))

			if out, _ := tt.flags.IsSet(uint8(tt.flag)); !out {
				t.Errorf("got %t, expected %t", out, false)
			}
		})
	}
}

var toggleTests = []struct {
	flags *flags.Flags
	flag  flags.Flag
	isSet bool
}{
	{newFlags(16), flags.C, false},
	{newFlags(16), flags.Z, true},
	{newFlags(32), flags.H, false},
	{newFlags(32), flags.N, true},
	{newFlags(64), flags.N, false},
	{newFlags(64), flags.H, true},
	{newFlags(128), flags.Z, false},
	{newFlags(128), flags.C, true},
}

func TestToggle(t *testing.T) {
	for _, tt := range toggleTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			tt.flags.Toggle(uint8(tt.flag))

			if out, _ := tt.flags.IsSet(uint8(tt.flag)); out != tt.isSet {
				t.Errorf("got %t, expected %t", out, tt.isSet)
			}
		})
	}
}

func newFlags(val flags.Flags) *flags.Flags {
	f := flags.New()
	*f = val
	return f
}
