package flags

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/loizoskounios/game-boy-emulator/byteops"
// )

// func TestNewFlags(t *testing.T) {
// 	f := NewFlags()

// 	if f == nil {
// 		t.Error("got nil, expected not nil")
// 	}
// }

// var updateFlagTests = []struct {
// 	f   flag
// 	m   byteops.Mutator
// 	out byte
// }{
// 	{CY, byteops.Clear, 0},
// 	{CY, byteops.Set, 16},
// 	{CY, byteops.Toggle, 16},
// 	{H, byteops.Clear, 0},
// 	{H, byteops.Set, 32},
// 	{H, byteops.Toggle, 32},
// 	{N, byteops.Clear, 0},
// 	{N, byteops.Set, 64},
// 	{N, byteops.Toggle, 64},
// 	{ZF, byteops.Clear, 0},
// 	{ZF, byteops.Set, 128},
// 	{ZF, byteops.Toggle, 128},
// }

// func TestUpdateFlag(t *testing.T) {
// 	for _, tt := range updateFlagTests {
// 		flags := NewFlags()

// 		t.Run(fmt.Sprintf("b=%v m=%v", tt.f, tt.m), func(t *testing.T) {
// 			flags.UpdateFlag(tt.f, tt.m)
// 			if flags.Flags() != tt.out {
// 				t.Errorf("got %d, expected %d", flags.Flags(), tt.out)
// 			}
// 		})
// 	}
// }
