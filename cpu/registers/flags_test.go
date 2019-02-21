package registers

import (
	"fmt"
	"testing"
)

func TestNewFlags(t *testing.T) {
	f := NewFlags()

	if f == nil {
		t.Error("got nil, expected not nil")
	}
}

func TestNewFlagsVal(t *testing.T) {
	f := NewFlags()

	expected := byte(0)
	actual := f.Flags()

	if actual != expected {
		t.Errorf("got %d, expected %d", expected, actual)
	}
}

var updateFlagTests = []struct {
	in  string
	out byte
}{
	{"cy", 16},
	{"h", 32},
	{"n", 64},
	{"zf", 128},
}

func TestUpdateFlag(t *testing.T) {
	f := NewFlags()

	for _, tt := range updateFlagTests {
		t.Run(fmt.Sprintf("Updating flag %s", tt.in), func(t *testing.T) {
			f.UpdateFlag(tt.in, 1)
			if f.Flags() != tt.out {
				t.Errorf("got %d, expected %d", f.Flags(), tt.out)
			}
			f.UpdateFlag(tt.in, 0)
			if f.Flags() != 0 {
				t.Errorf("got %d, expected %d", f.Flags(), 0)
			}
		})
	}
}

var updateBitOneTests = []struct {
	in  uint
	out byte
}{
	{0, 1},
	{1, 3},
	{2, 7},
	{3, 15},
	{4, 31},
	{5, 63},
	{6, 127},
	{7, 255},
}

func TestUpdateBitOne(t *testing.T) {
	f := NewFlags()

	for _, tt := range updateBitOneTests {
		t.Run(fmt.Sprintf("Setting bits 0-%d", tt.in), func(t *testing.T) {
			f.updateBit(tt.in, 1)
			if f.Flags() != tt.out {
				t.Errorf("got %d, expected %d", f.Flags(), tt.out)
			}
		})
	}
}

var updateBitZeroTests = []struct {
	in  uint
	out byte
}{
	{0, 254},
	{1, 252},
	{2, 248},
	{3, 240},
	{4, 224},
	{5, 192},
	{6, 128},
	{7, 0},
}

func TestUpdateBitZero(t *testing.T) {
	f := NewFlags()
	f.val = 255

	for _, tt := range updateBitZeroTests {
		t.Run(fmt.Sprintf("Unsetting bits 0-%d", tt.in), func(t *testing.T) {
			f.updateBit(tt.in, 0)
			if f.Flags() != tt.out {
				t.Errorf("got %d, expected %d", f.Flags(), tt.out)
			}
		})
	}
}
