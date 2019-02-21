package byteops

import (
	"fmt"
	"testing"
)

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
