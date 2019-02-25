package mmu

import (
	"fmt"
	"testing"
)

var byteTests = []struct {
	address uint16
	out     uint8
}{
	{0, 0},
	{500, 128},
	{65535, 255},
}

func TestByte(t *testing.T) {
	m := Memory{
		0:     0,
		500:   128,
		65535: 255,
	}

	for _, tt := range byteTests {
		t.Run(fmt.Sprintf("address=%d", tt.address), func(t *testing.T) {
			if out := m.Byte(tt.address); out != tt.out {
				t.Errorf("got %d, expected %d", out, tt.out)
			}
		})
	}
}

var wordTests = []struct {
	address uint16
	out     uint16
}{
	{0, 257},
	{500, 32896},
	{65534, 65535},
}

func TestWord(t *testing.T) {
	m := Memory{
		0:     1,
		1:     1,
		500:   128,
		501:   128,
		65534: 255,
		65535: 255,
	}

	for _, tt := range wordTests {
		t.Run(fmt.Sprintf("address=%d", tt.address), func(t *testing.T) {
			if out := m.Word(tt.address); out != tt.out {
				t.Errorf("got %d, expected %d", out, tt.out)
			}
		})
	}
}

var setByteTests = []struct {
	address uint16
	b       uint8
	out     uint8
}{
	{0, 255, 255},
	{30000, 128, 128},
	{65535, 193, 193},
}

func TestSetByte(t *testing.T) {
	m := Memory{}

	for _, tt := range setByteTests {
		t.Run(fmt.Sprintf("address=%d b=%d", tt.address, tt.b), func(t *testing.T) {
			m.SetByte(tt.address, tt.b)
			if out := m.Byte(tt.address); out != tt.out {
				t.Errorf("got %d, expected %d", out, tt.out)
			}
		})
	}
}

var setWordTests = []struct {
	address uint16
	w       uint16
	out     uint16
}{
	{0, 1024, 1024},
	{30001, 3091, 3091},
	{65534, 65535, 65535},
}

func TestSetWord(t *testing.T) {
	m := Memory{}

	for _, tt := range setWordTests {
		t.Run(fmt.Sprintf("address=%d w=%d", tt.address, tt.w), func(t *testing.T) {
			m.SetWord(tt.address, tt.w)
			if out := m.Word(tt.address); out != tt.out {
				t.Errorf("got %d, expected %d", out, tt.out)
			}
		})
	}
}
