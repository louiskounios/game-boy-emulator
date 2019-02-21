package byteops

import (
	"fmt"
	"testing"
)

var clearTests = []struct {
	b   byte
	n   uint8
	out uint8
}{
	{1, 0, 0},
	{2, 1, 0},
	{4, 2, 0},
	{127, 6, 63},
	{127, 0, 126},
}

func TestClear(t *testing.T) {
	for _, tt := range clearTests {
		t.Run(fmt.Sprintf("b=%08b n=%d", tt.b, tt.n), func(t *testing.T) {
			Clear(&tt.b, tt.n)
			if tt.b != tt.out {
				t.Errorf("got %d, expected %d", tt.b, tt.out)
			}
		})
	}
}
