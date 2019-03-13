package cpu

import (
	"fmt"
	"testing"
)

var getTests = []struct {
	val  Flags
	flag Flag
	ret  uint8
}{
	{16, FlagC, 1},
	{16, FlagZ, 0},
	{32, FlagH, 1},
	{32, FlagN, 0},
	{64, FlagN, 1},
	{64, FlagH, 0},
	{128, FlagZ, 1},
	{128, FlagC, 0},
}

var isSetTests = []struct {
	flags *Flags
	flag  Flag
	ret   bool
}{
	{newFlags(16), FlagC, true},
	{newFlags(16), FlagZ, false},
	{newFlags(32), FlagH, true},
	{newFlags(32), FlagN, false},
	{newFlags(64), FlagN, true},
	{newFlags(64), FlagH, false},
	{newFlags(128), FlagZ, true},
	{newFlags(128), FlagC, false},
}

func TestIsSet(t *testing.T) {
	for _, tt := range isSetTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			if out, _ := tt.flags.IsSet(tt.flag); out != tt.ret {
				t.Errorf("got %t, expected %t", out, tt.ret)
			}
		})
	}
}

var putTests = []struct {
	flags *Flags
	flag  Flag
	set   bool
}{
	{newFlags(0), FlagC, true},
	{newFlags(0), FlagC, false},
	{newFlags(16), FlagC, true},
	{newFlags(16), FlagC, false},
	{newFlags(0), FlagH, true},
	{newFlags(0), FlagH, false},
	{newFlags(32), FlagH, true},
	{newFlags(32), FlagH, false},
	{newFlags(0), FlagN, true},
	{newFlags(0), FlagN, false},
	{newFlags(64), FlagN, true},
	{newFlags(64), FlagN, false},
	{newFlags(0), FlagZ, true},
	{newFlags(0), FlagZ, false},
	{newFlags(128), FlagZ, true},
	{newFlags(128), FlagZ, false},
}

func TestPut(t *testing.T) {
	for _, tt := range putTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			tt.flags.Put(tt.flag, tt.set)

			if out, _ := tt.flags.IsSet(tt.flag); out != tt.set {
				t.Errorf("got %t, expected %t", out, false)
			}
		})
	}
}

var resetTests = []struct {
	flags *Flags
	flag  Flag
}{
	{newFlags(16), FlagC},
	{newFlags(16), FlagZ},
	{newFlags(32), FlagH},
	{newFlags(32), FlagN},
	{newFlags(64), FlagN},
	{newFlags(64), FlagH},
	{newFlags(128), FlagZ},
	{newFlags(128), FlagC},
}

func TestReset(t *testing.T) {
	for _, tt := range resetTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			tt.flags.Reset(tt.flag)

			if out, _ := tt.flags.IsSet(tt.flag); out {
				t.Errorf("got %t, expected %t", out, false)
			}
		})
	}
}

var setTests = []struct {
	flags *Flags
	flag  Flag
}{
	{newFlags(16), FlagC},
	{newFlags(16), FlagZ},
	{newFlags(32), FlagH},
	{newFlags(32), FlagN},
	{newFlags(64), FlagN},
	{newFlags(64), FlagH},
	{newFlags(128), FlagZ},
	{newFlags(128), FlagC},
}

func TestSet(t *testing.T) {
	for _, tt := range setTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			tt.flags.Set(tt.flag)

			if out, _ := tt.flags.IsSet(tt.flag); !out {
				t.Errorf("got %t, expected %t", out, false)
			}
		})
	}
}

var toggleTests = []struct {
	flags *Flags
	flag  Flag
	isSet bool
}{
	{newFlags(16), FlagC, false},
	{newFlags(16), FlagZ, true},
	{newFlags(32), FlagH, false},
	{newFlags(32), FlagN, true},
	{newFlags(64), FlagN, false},
	{newFlags(64), FlagH, true},
	{newFlags(128), FlagZ, false},
	{newFlags(128), FlagC, true},
}

func TestToggle(t *testing.T) {
	for _, tt := range toggleTests {
		t.Run(fmt.Sprintf("flags=[%s] flag=%s", tt.flags, tt.flag), func(t *testing.T) {
			tt.flags.Toggle(tt.flag)

			if out, _ := tt.flags.IsSet(tt.flag); out != tt.isSet {
				t.Errorf("got %t, expected %t", out, tt.isSet)
			}
		})
	}
}

func newFlags(val Flags) *Flags {
	f := NewFlags()
	*f = val
	return f
}
