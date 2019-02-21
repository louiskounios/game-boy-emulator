package byteops

import (
	"fmt"
	"testing"
)

var toggleTests = []struct {
	b   byte
	n   uint8
	out uint8
}{
	{0, 0, 1},
	{1, 0, 0},
	{1, 1, 3},
	{11, 2, 15},
	{128, 7, 0},
}

func TestToggle(t *testing.T) {
	for _, tt := range toggleTests {
		t.Run(fmt.Sprintf("b=%08b n=%d", tt.b, tt.n), func(t *testing.T) {
			Toggle(&tt.b, tt.n)
			if tt.b != tt.out {
				t.Errorf("got %d, expected %d", tt.b, tt.out)
			}
		})
	}
}
