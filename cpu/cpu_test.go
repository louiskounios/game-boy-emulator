package cpu

import (
	"fmt"
	"testing"
)

func TestAddSignedUnsigned(t *testing.T) {
	var tests = []struct {
		unsigned uint16
		asSigned uint8
		result   uint16
		carry    bool
	}{
		{50, 10, 60, false},
		{50, 246, 40, false},
		{0, 0, 0, false},
		{65535, 0, 65535, false},
		{65535, 1, 0, true},
		{65535, 127, 126, true},
		{0, 255, 65535, true},
		{0, 128, 65408, true},
		{128, 128, 0, false},
		{127, 128, 65535, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("unsigned=%d asSigned=%d", test.unsigned, test.asSigned), func(t *testing.T) {
			r, c, hc := addSignedUnsigned(test.asSigned, test.unsigned)
			if r != test.result || c != test.carry || true {
				t.Errorf("got %d %t %t, expected %d %t", r, c, hc, test.result, test.carry)
			}
		})
	}
}
