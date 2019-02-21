package byteops

import (
	"fmt"
	"testing"
)

var clearTests = []struct {
	b   byte
	n   uint8
	out uint8
}{
	{1, 0, 0},
	{2, 1, 0},
	{4, 2, 0},
	{127, 6, 63},
	{127, 0, 126},
}

func TestClear(t *testing.T) {
	for _, tt := range clearTests {
		t.Run(fmt.Sprintf("b=%08b n=%d", tt.b, tt.n), func(t *testing.T) {
			Clear(&tt.b, tt.n)
			if tt.b != tt.out {
				t.Errorf("got %d, expected %d", tt.b, tt.out)
			}
		})
	}
}

var getTests = []struct {
	b   byte
	n   uint8
	out uint8
}{
	{1, 0, 1},
	{3, 1, 1},
	{128, 7, 1},
	{2, 0, 0},
}

func TestGet(t *testing.T) {
	for _, tt := range getTests {
		t.Run(fmt.Sprintf("b=%08b n=%d", tt.b, tt.n), func(t *testing.T) {
			bit, _ := Get(tt.b, tt.n)
			if bit != tt.out {
				t.Errorf("got %d, expected %d", bit, tt.out)
			}
		})
	}
}

var setTests = []struct {
	b   byte
	n   uint8
	out uint8
}{
	{0, 0, 1},
	{0, 1, 2},
	{0, 2, 4},
	{128, 7, 128},
	{127, 7, 255},
}

func TestSet(t *testing.T) {
	for _, tt := range setTests {
		t.Run(fmt.Sprintf("b=%08b n=%d", tt.b, tt.n), func(t *testing.T) {
			Set(&tt.b, tt.n)
			if tt.b != tt.out {
				t.Errorf("got %d, expected %d", tt.b, tt.out)
			}
		})
	}
}

var toggleTests = []struct {
	b   byte
	n   uint8
	out uint8
}{
	{0, 0, 1},
	{1, 0, 0},
	{1, 1, 3},
	{11, 2, 15},
	{128, 7, 0},
}

func TestToggle(t *testing.T) {
	for _, tt := range toggleTests {
		t.Run(fmt.Sprintf("b=%08b n=%d", tt.b, tt.n), func(t *testing.T) {
			Toggle(&tt.b, tt.n)
			if tt.b != tt.out {
				t.Errorf("got %d, expected %d", tt.b, tt.out)
			}
		})
	}
}

var oobTests = []struct {
	in  uint8
	out bool
}{
	{0, false},
	{5, false},
	{7, false},
	{8, true},
}

func TestOutOfBounds(t *testing.T) {
	for _, tt := range oobTests {
		t.Run(fmt.Sprintf("n=%d", tt.in), func(t *testing.T) {
			oob := isOutOfBounds(tt.in)
			if oob != tt.out {
				t.Errorf("got %v, expected %v", oob, tt.out)
			}
		})
	}
}
