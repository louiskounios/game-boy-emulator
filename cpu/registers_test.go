package cpu

import (
	"fmt"
	"testing"
)

func TestAccumulator(t *testing.T) {
	r := NewRegisters()
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
	r := NewRegisters()
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
	r := NewRegisters()
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
		register Register
		value    uint8
		expected uint8
	}{
		{RegisterA, 255, 255},
		{RegisterB, 255, 255},
		{RegisterC, 255, 255},
		{RegisterD, 255, 255},
		{RegisterE, 255, 255},
		{RegisterH, 255, 255},
		{RegisterL, 255, 255},
	}

	r := NewRegisters()
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
		register Register
		expected uint16
	}{
		{RegisterBC, 0},
		{RegisterDE, 0},
		{RegisterHL, 0},
	}

	r := NewRegisters()
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
		register Register
		value    uint16
		expected uint16
	}{
		{RegisterBC, 65535, 65535},
		{RegisterDE, 65535, 65535},
		{RegisterHL, 65535, 65535},
	}

	r := NewRegisters()
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("register=%s value=%d", tc.register, tc.value), func(t *testing.T) {
			r.SetPaired(tc.register, tc.value)
			if output, _ := r.Paired(tc.register); output != tc.expected {
				t.Errorf("got %d, expected %d", output, tc.expected)
			}
		})
	}
}
