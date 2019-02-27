package mmu

// AddressableRegions is the number of addressable regions in memory.
const AddressableRegions = 65536

// Memory is a uint8 array of size AddressableRegions.
type Memory [AddressableRegions]uint8

// Byte returns the contents of the memory at address a.
func (m Memory) Byte(a uint16) uint8 {
	return m[a]
}

// Word combines the contents of the memory at address a+1 and a to form and
// return a word.
//
// The memory contents at address a+1 become the 8 most significant bits of the
// resulting word. The memory contents at a become the 8 least significant
// bits.
func (m Memory) Word(a uint16) uint16 {
	return uint16(m[a+1])<<8 | uint16(m[a])
}

// SetByte writes b to memory address a.
func (m *Memory) SetByte(a uint16, b uint8) {
	m[a] = b
}

// SetWord writes w to memory addresses a and a+1.
//
// The 8 most and least significant bits of w are written to memory addresses
// a+1 and a respectively.
func (m *Memory) SetWord(a uint16, w uint16) {
	m[a+1] = uint8(w >> 8)
	m[a] = uint8(w)
}
