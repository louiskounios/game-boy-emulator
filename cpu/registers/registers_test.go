package registers

import (
	"fmt"
	"testing"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers/flags"
)

func TestAccumulator(t *testing.T) {
	r := New()
	acc := r.Accumulator()
	*acc = 255
	if r.af.a != 255 {
		t.Errorf("got %d, expected %d", r.af.a, 255)
	}
}

func TestFlags(t *testing.T) {
	r := New()
	f := r.Flags()
	*f = flags.Flags(255)
	if *r.af.f != flags.Flags(255) {
		t.Errorf("got %d, expected %d", *r.af.f, 255)
	}
}

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
	{HL, 55555, nil},
	{SP, 10930, nil},
	{PC, 48031, nil},
	{Register(30), 0, errUnknownRegister},
}

func TestRegister(t *testing.T) {
	for _, tt := range registersTests {
		r := Registers{
			af: RegisterAF{
				a: 1,
				f: newFlags(1),
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

		t.Run(fmt.Sprintf("r=%s", tt.r), func(t *testing.T) {
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
	{HL, 55555, nil},
	{SP, 10930, nil},
	{PC, 48031, nil},
	{Register(30), 0, errUnknownRegister},
}

func TestSetRegister(t *testing.T) {
	for _, tt := range setRegistersTests {
		r := New()

		t.Run(fmt.Sprintf("r=%s val=%d", tt.r, tt.val), func(t *testing.T) {
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

var incrementByTests = []struct {
	r   Register
	val uint16
	by  uint8
	out uint16
	err error
}{
	{A, 255, 1, 0, nil},
	{F, 254, 2, 254, errFlagsIncrement},
	{B, 253, 3, 0, nil},
	{C, 252, 4, 0, nil},
	{D, 251, 5, 0, nil},
	{E, 250, 6, 0, nil},
	{H, 249, 7, 0, nil},
	{L, 248, 8, 0, nil},
	{AF, 247, 9, 247, errFlagsIncrement},
	{BC, 246, 10, 256, nil},
	{DE, 245, 11, 256, nil},
	{HL, 244, 12, 256, nil},
	{SP, 243, 13, 256, nil},
	{PC, 242, 14, 256, nil},
	{Register(30), 241, 15, 0, errUnknownRegister},
}

func TestIncrementBy(t *testing.T) {
	for _, tt := range incrementByTests {
		t.Run(fmt.Sprintf("r=%s val=%d by=%d", tt.r, tt.val, tt.by), func(t *testing.T) {
			r := New()
			r.SetRegister(tt.r, tt.val)

			if err := r.IncrementBy(tt.r, tt.by); err != tt.err {
				t.Errorf("got %#v, expected %#v", err, tt.err)
			}

			if out, _ := r.Register(tt.r); out != tt.out {
				t.Errorf("got %v, expected %v", out, tt.out)
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
	{HL, 255, 256, nil},
	{SP, 255, 256, nil},
	{PC, 255, 256, nil},
	{Register(30), 255, 0, errUnknownRegister},
}

func TestIncrement(t *testing.T) {
	for _, tt := range incrementTests {
		t.Run(fmt.Sprintf("r=%s val=%d", tt.r, tt.val), func(t *testing.T) {
			r := New()
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

var decrementByTests = []struct {
	r   Register
	val uint16
	by  uint8
	out uint16
	err error
}{
	{A, 0, 1, 255, nil},
	{F, 1, 2, 1, errFlagsDecrement},
	{B, 2, 3, 255, nil},
	{C, 3, 4, 255, nil},
	{D, 4, 5, 255, nil},
	{E, 5, 6, 255, nil},
	{H, 6, 7, 255, nil},
	{L, 7, 8, 255, nil},
	{AF, 8, 9, 8, errFlagsDecrement},
	{BC, 9, 10, 65535, nil},
	{DE, 10, 11, 65535, nil},
	{HL, 11, 12, 65535, nil},
	{SP, 12, 13, 65535, nil},
	{PC, 13, 14, 65535, nil},
	{Register(30), 14, 15, 0, errUnknownRegister},
}

func TestDecrementBy(t *testing.T) {
	for _, tt := range decrementByTests {
		t.Run(fmt.Sprintf("r=%s val=%d by=%d", tt.r, tt.val, tt.by), func(t *testing.T) {
			r := New()
			r.SetRegister(tt.r, tt.val)

			if err := r.DecrementBy(tt.r, tt.by); err != tt.err {
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
	{HL, 256, 255, nil},
	{SP, 256, 255, nil},
	{PC, 256, 255, nil},
	{Register(30), 256, 0, errUnknownRegister},
}

func TestDecrement(t *testing.T) {
	for _, tt := range decrementTests {
		t.Run(fmt.Sprintf("r=%s val=%d", tt.r, tt.val), func(t *testing.T) {
			r := New()
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

var getComponentsTests = []struct {
	r  Register
	hi uint8
	lo uint8
}{
	{AF, 1, 1},
	{BC, 255, 255},
	{DE, 120, 208},
	{HL, 217, 3},
}

func TestGetComponents(t *testing.T) {
	r := Registers{
		af: RegisterAF{
			a: 1,
			f: newFlags(1),
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
	}

	for _, tt := range getComponentsTests {
		t.Run(fmt.Sprintf("r=%s", tt.r), func(t *testing.T) {
			if hi, lo, _ := r.GetComponents(tt.r); hi != tt.hi || lo != tt.lo {
				t.Errorf("got %d %d, expected %d %d", hi, lo, tt.hi, tt.lo)
			}
		})
	}
}
