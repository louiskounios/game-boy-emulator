package mmu

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	m := Memory{
		0:     0,
		500:   128,
		65535: 255,
	}

	var testCases = []struct {
		address uint16
		out     uint8
	}{
		{0, 0},
		{500, 128},
		{65535, 255},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("address=%d", tc.address), func(t *testing.T) {
			if out := m.Load(tc.address); out != tc.out {
				t.Errorf("got %d, expected %d", out, tc.out)
			}
		})
	}
}

func TestStore(t *testing.T) {
	m := Memory{}

	var testCases = []struct {
		address uint16
		b       uint8
		out     uint8
	}{
		{0, 255, 255},
		{30000, 128, 128},
		{65535, 193, 193},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("address=%d b=%d", tc.address, tc.b), func(t *testing.T) {
			m.Store(tc.address, tc.b)
			if out := m.Load(tc.address); out != tc.out {
				t.Errorf("got %d, expected %d", out, tc.out)
			}
		})
	}
}
