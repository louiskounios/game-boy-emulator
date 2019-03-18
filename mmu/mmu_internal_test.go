package mmu

import "testing"

func TestNew(t *testing.T) {
	mmu := New()

	for i := bios.start; i <= bios.end; i++ {
		if val := mmu.m.Load(i); val != BIOS[i] {
			t.Errorf("got %d, expected %d", val, BIOS[i])
		}
	}

	for i := bios.end + 1; i >= bios.end && i <= interruptsEnable.end; i++ {
		if val := mmu.m.Load(i); val != 0 {
			t.Errorf("got %d, expected %d", val, 0)
		}
	}
}
