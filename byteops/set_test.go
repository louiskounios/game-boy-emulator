package byteops

import (
	"fmt"
	"testing"
)

var setTests = []struct {
	b   byte
	n   uint8
	out uint8
}{
	{0, 0, 1},
	{0, 1, 2},
	{0, 2, 4},
	{128, 7, 128},
	{127, 7, 255},
}

func TestSet(t *testing.T) {
	for _, tt := range setTests {
		t.Run(fmt.Sprintf("b=%08b n=%d", tt.b, tt.n), func(t *testing.T) {
			Set(&tt.b, tt.n)
			if tt.b != tt.out {
				t.Errorf("got %d, expected %d", tt.b, tt.out)
			}
		})
	}
}
