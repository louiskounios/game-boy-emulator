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
			af: RegisterAF{
				a: 1,
				f: 1,
			},
			bc: Register16{
				hi: 255,
				lo: 255,
			},
			de: Register16{
				hi: 120,
				lo: 208,
			},
			hl: Register16{
				hi: 217,
				lo: 3,
			},
			sp: 10930,
			pc: 48031,
		}

		t.Run(fmt.Sprintf("r=%d", tt.r), func(t *testing.T) {
			ret, err := r.Register(tt.r)
			if ret != tt.ret {
				t.Errorf("got %v, expected %v", ret, tt.ret)
			}
			if err != tt.err {
				t.Errorf("got %#v, expected %#v", err, tt.err)
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

			ret, err := r.Register(tt.r)
			if ret != tt.val {
				t.Errorf("got %v, expected %v", ret, tt.val)
			}
			if err != tt.err {
				t.Errorf("got %#v, expected %#v", err, tt.err)
			}
		})
	}
}

var incrementTests = []struct {
	r   Register
	val uint16
	out uint16
	err error
}{
	{A, 255, 0, nil},
	{F, 255, 255, errFlagsIncrement},
	{B, 255, 0, nil},
	{C, 255, 0, nil},
	{D, 255, 0, nil},
	{E, 255, 0, nil},
	{H, 255, 0, nil},
	{L, 255, 0, nil},
	{AF, 255, 255, errFlagsIncrement},
	{BC, 255, 256, nil},
	{DE, 255, 256, nil},
	{SP, 255, 256, nil},
	{PC, 255, 256, nil},
	{Register(30), 255, 0, errUnknownRegister},
}

func TestIncrement(t *testing.T) {
	for _, tt := range incrementTests {
		t.Run(fmt.Sprintf("r=%d val=%d", tt.r, tt.val), func(t *testing.T) {
			r := Registers{}
			r.SetRegister(tt.r, tt.val)

			if err := r.Increment(tt.r); err != tt.err {
				t.Errorf("got %#v, expected %#v", err, tt.err)
			}

			if out, _ := r.Register(tt.r); out != tt.out {
				t.Errorf("got %v, expected %v", out, tt.out)
			}
		})
	}
}

var decrementTests = []struct {
	r   Register
	val uint16
	out uint16
	err error
}{
	{A, 0, 255, nil},
	{F, 0, 0, errFlagsDecrement},
	{B, 0, 255, nil},
	{C, 0, 255, nil},
	{D, 0, 255, nil},
	{E, 0, 255, nil},
	{H, 0, 255, nil},
	{L, 0, 255, nil},
	{AF, 256, 256, errFlagsDecrement},
	{BC, 256, 255, nil},
	{DE, 256, 255, nil},
	{SP, 256, 255, nil},
	{PC, 256, 255, nil},
	{Register(30), 256, 0, errUnknownRegister},
}

func TestDecrement(t *testing.T) {
	for _, tt := range decrementTests {
		t.Run(fmt.Sprintf("r=%d val=%d", tt.r, tt.val), func(t *testing.T) {
			r := Registers{}
			r.SetRegister(tt.r, tt.val)

			if err := r.Decrement(tt.r); err != tt.err {
				t.Errorf("got %#v, expected %#v", err, tt.err)
			}

			if out, _ := r.Register(tt.r); out != tt.out {
				t.Errorf("got %v, expected %v", out, tt.out)
			}
		})
	}
}
