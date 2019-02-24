package registers

import (
	"fmt"
	"testing"
)

var equalsTests = []struct {
	r1  Register16C
	r2  Register16C
	out bool
}{
	{Register16C{0, 0}, Register16C{0, 0}, true},
	{Register16C{0, 1}, Register16C{0, 0}, false},
	{Register16C{128, 0}, Register16C{0, 128}, false},
	{Register16C{255, 0}, Register16C{255, 0}, true},
	{Register16C{255, 255}, Register16C{255, 255}, true},
}

func TestEquals(t *testing.T) {
	for _, tt := range equalsTests {
		t.Run(fmt.Sprintf("r1=%v r2=%v", tt.r1, tt.r2), func(t *testing.T) {
			if out := tt.r1.Equals(&tt.r1) && tt.r1.Equals(&tt.r1) && tt.r1.Equals(&tt.r2) && tt.r2.Equals(&tt.r1); out != tt.out {
				t.Errorf("got %t, expected %t", out, tt.out)
			}
		})
	}
}

var getWordTests = []struct {
	in  Register16C
	out Register16
}{
	{Register16C{1, 1}, Register16(257)},
	{Register16C{128, 1}, Register16(32769)},
	{Register16C{128, 0}, Register16(32768)},
	{Register16C{128, 128}, Register16(32896)},
	{Register16C{255, 255}, Register16(65535)},
}

func TestWordGetter(t *testing.T) {
	for _, tt := range getWordTests {
		t.Run(fmt.Sprintf("in=%v", tt.in), func(t *testing.T) {
			if r := tt.in.Word(); r != tt.out {
				t.Errorf("got %d, expected %d", r, tt.out)
			}
		})
	}
}

var setWordTests = []struct {
	in  Register16
	out Register16C
}{
	{Register16(257), Register16C{1, 1}},
	{Register16(32769), Register16C{128, 1}},
	{Register16(32768), Register16C{128, 0}},
	{Register16(32896), Register16C{128, 128}},
	{Register16(65535), Register16C{255, 255}},
}

func TestWordSetter(t *testing.T) {
	for _, tt := range setWordTests {
		r := Register16C{}

		t.Run(fmt.Sprintf("in=%d", tt.in), func(t *testing.T) {
			if r.SetWord(tt.in); !r.Equals(&tt.out) {
				t.Errorf("got %v, expected %v", r, tt.out)
			}
		})
	}
}
