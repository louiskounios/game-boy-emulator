package byteops

import (
	"fmt"
	"testing"
)

var getTests = []struct {
	b   byte
	n   uint8
	out uint8
}{
	{1, 0, 1},
	{3, 1, 1},
	{128, 7, 1},
	{2, 0, 0},
}

func TestGet(t *testing.T) {
	for _, tt := range getTests {
		t.Run(fmt.Sprintf("b=%08b n=%d", tt.b, tt.n), func(t *testing.T) {
			bit, _ := Get(tt.b, tt.n)
			if bit != tt.out {
				t.Errorf("got %d, expected %d", bit, tt.out)
			}
		})
	}
}
