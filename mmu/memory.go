package mmu

// AddressableRegions is the number of addressable regions in memory.
const AddressableRegions = 65536

// Memory is a uint8 array of size AddressableRegions.
type Memory [AddressableRegions]uint8

// Load returns the memory contents at the provided address.
func (m Memory) Load(addr uint16) uint8 {
	return m[addr]
}

// Store stores the provided byte to the provided address.
func (m *Memory) Store(addr uint16, b uint8) {
	m[addr] = b
}
