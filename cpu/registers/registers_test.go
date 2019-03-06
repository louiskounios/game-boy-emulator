package registers_test

import (
	"fmt"
	"testing"

	"github.com/loizoskounios/game-boy-emulator/cpu/registers"
)

func TestAccumulator(t *testing.T) {
	r := registers.New()
	acc := r.Accumulator()

	if acc == nil {
		t.Error("got nil, expected pointer")
	}

	if *acc != 0 {
		t.Errorf("got %d, expected 0", *acc)
	}

	*acc = 255
	if *acc != 255 {
		t.Errorf("got %d, expected 255", *acc)
	}
}

func TestStackPointer(t *testing.T) {
	r := registers.New()
	sp := r.StackPointer()

	if sp == nil {
		t.Error("got nil, expected pointer")
	}

	if *sp != 0 {
		t.Errorf("got %d, expected 0", *sp)
	}

	*sp = 255
	if *sp != 255 {
		t.Errorf("got %d, expected 255", *sp)
	}
}

func TestProgramCounter(t *testing.T) {
	r := registers.New()
	pc := r.ProgramCounter()

	if pc == nil {
		t.Error("got nil, expected pointer")
	}

	if *pc != 0 {
		t.Errorf("got %d, expected 0", *pc)
	}

	*pc = 255
	if *pc != 255 {
		t.Errorf("got %d, expected 255", *pc)
	}
}

func TestAuxiliary(t *testing.T) {
	testCases := []struct {
		register registers.Register
		value    uint8
		expected uint8
	}{
		{registers.A, 255, 255},
		{registers.B, 255, 255},
		{registers.C, 255, 255},
		{registers.D, 255, 255},
		{registers.E, 255, 255},
		{registers.H, 255, 255},
		{registers.L, 255, 255},
	}

	r := registers.New()
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("register=%s value=%d", tc.register, tc.value), func(t *testing.T) {
			aux, _ := r.Auxiliary(tc.register)

			if aux == nil {
				t.Error("got nil, expected pointer")
			}

			if *aux != 0 {
				t.Errorf("got %d, expected 0", *aux)
			}

			*aux = tc.value
			if *aux != tc.expected {
				t.Errorf("got %d, expected %d", *aux, tc.expected)
			}
		})
	}
}

func TestPaired(t *testing.T) {
	testCases := []struct {
		register registers.Register
		expected uint16
	}{
		{registers.BC, 0},
		{registers.DE, 0},
		{registers.HL, 0},
	}

	r := registers.New()
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("register=%s", tc.register), func(t *testing.T) {
			if output, _ := r.Paired(tc.register); output != tc.expected {
				t.Errorf("got %d, expected %d", output, tc.expected)
			}
		})
	}
}

func TestSetPaired(t *testing.T) {
	testCases := []struct {
		register registers.Register
		value    uint16
		expected uint16
	}{
		{registers.BC, 65535, 65535},
		{registers.DE, 65535, 65535},
		{registers.HL, 65535, 65535},
	}

	r := registers.New()
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("register=%s value=%d", tc.register, tc.value), func(t *testing.T) {
			r.SetPaired(tc.register, tc.value)
			if output, _ := r.Paired(tc.register); output != tc.expected {
				t.Errorf("got %d, expected %d", output, tc.expected)
			}
		})
	}
}
