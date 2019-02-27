package registers

import (
	"fmt"
	"testing"
)

var registersTests = []struct {
	r   Register
	ret uint16
	err error
}{
	{A, 1, nil},
	{F, 1, nil},
	{B, 255, nil},
	{C, 255, nil},
	{D, 120, nil},
	{E, 208, nil},
	{H, 217, nil},
	{L, 3, nil},
	{AF, 257, nil},
	{BC, 65535, nil},
	{DE, 30928, nil},
	{SP, 10930, nil},
	{PC, 48031, nil},
	{Register(30), 0, errUnknownRegister},
}

func TestRegister(t *testing.T) {
	for _, tt := range registersTests {
		r := Registers{
			AF: RegisterAF{
				a: 1,
				f: 1,
			},
			BC: Register16{
				hi: 255,
				lo: 255,
			},
			DE: Register16{
				hi: 120,
				lo: 208,
			},
			HL: Register16{
				hi: 217,
				lo: 3,
			},
			SP: 10930,
			PC: 48031,
		}

		t.Run(fmt.Sprintf("r=%d", tt.r), func(t *testing.T) {
			if ret, err := r.Register(tt.r); ret != tt.ret || err != tt.err {
				t.Errorf("got %v %T, expected %v %T", ret, err, tt.ret, tt.err)
			}
		})
	}
}

var setRegistersTests = []struct {
	r   Register
	val uint16
	err error
}{
	{A, 1, nil},
	{F, 1, nil},
	{B, 255, nil},
	{C, 255, nil},
	{D, 120, nil},
	{E, 208, nil},
	{H, 217, nil},
	{L, 3, nil},
	{AF, 257, nil},
	{BC, 65535, nil},
	{DE, 30928, nil},
	{SP, 10930, nil},
	{PC, 48031, nil},
	{Register(30), 0, errUnknownRegister},
}

func TestSetRegister(t *testing.T) {
	for _, tt := range setRegistersTests {
		r := Registers{}

		t.Run(fmt.Sprintf("r=%d val=%d", tt.r, tt.val), func(t *testing.T) {
			r.SetRegister(tt.r, tt.val)
			if ret, err := r.Register(tt.r); ret != tt.val || err != tt.err {
				t.Errorf("got %v %T, expected %v %T", ret, err, tt.val, tt.err)
			}
		})
	}
}
