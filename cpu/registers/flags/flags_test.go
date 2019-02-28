package flags

import (
	"fmt"
	"testing"
)

var getTests = []struct {
	val  Flags
	flag Flag
	ret  uint8
}{
	{16, C, 1},
	{16, Z, 0},
	{32, H, 1},
	{32, N, 0},
	{64, N, 1},
	{64, H, 0},
	{128, Z, 1},
	{128, C, 0},
}

func TestGet(t *testing.T) {
	for _, tt := range getTests {
		f := New()
		*f = tt.val

		t.Run(fmt.Sprintf("val=%08b flag=%s", tt.val, tt.flag), func(t *testing.T) {
			if out, _ := f.Get(tt.flag); out != tt.ret {
				t.Errorf("got %d, expected %d", out, tt.ret)
			}
		})
	}
}

var isSetTests = []struct {
	flags *Flags
	flag  Flag
	ret   bool
}{
	{newFlags(16), C, true},
	{newFlags(16), Z, false},
	{newFlags(32), H, true},
	{newFlags(32), N, false},
	{newFlags(64), N, true},
	{newFlags(64), H, false},
	{newFlags(128), Z, true},
	{newFlags(128), C, false},
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

var resetTests = []struct {
	flags *Flags
	flag  Flag
}{
	{newFlags(16), C},
	{newFlags(16), Z},
	{newFlags(32), H},
	{newFlags(32), N},
	{newFlags(64), N},
	{newFlags(64), H},
	{newFlags(128), Z},
	{newFlags(128), C},
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
	{newFlags(16), C},
	{newFlags(16), Z},
	{newFlags(32), H},
	{newFlags(32), N},
	{newFlags(64), N},
	{newFlags(64), H},
	{newFlags(128), Z},
	{newFlags(128), C},
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
	{newFlags(16), C, false},
	{newFlags(16), Z, true},
	{newFlags(32), H, false},
	{newFlags(32), N, true},
	{newFlags(64), N, false},
	{newFlags(64), H, true},
	{newFlags(128), Z, false},
	{newFlags(128), C, true},
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
	f := New()
	*f = val
	return f
}
