package mmu

// AddressableRegions is the number of addressable regions in memory.
const AddressableRegions = 65536

// memory is a uint8 array of size AddressableRegions.
type memory [AddressableRegions]uint8

// newMemory returns a pointer to a new memory region.
func newMemory() *memory {
	return &memory{}
}

// Load returns the memory contents at the provided address.
func (m *memory) Load(addr uint16) uint8 {
	return m[addr]
}

// Store stores the provided byte to the provided address.
func (m *memory) Store(addr uint16, b uint8) {
	m[addr] = b
}
