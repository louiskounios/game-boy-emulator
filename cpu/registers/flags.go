package registers

type flags struct {
	val byte
}

// NewFlags returns a new 8-bit flags register.
func NewFlags() *flags {
	return &flags{}
}

// Value returns a copy of the value of the 8-bit flags register.
func (f *flags) Flags() byte {
	return f.val
}

// How many bits the bitmask must be shifted by to mask the bit corresponding
// to the flag.
var flagShift = map[string]uint{
	"cy": 4,
	"h":  5,
	"n":  6,
	"zf": 7,
}

// UpdateFlag sets the value of a flag to the given value.
//
// Accepted values for flag are "cy", "h", "n", "zf".
// Accepted values for val are 0 or 1.
func (f *flags) UpdateFlag(flag string, val byte) {
	f.updateBit(flagShift[flag], val)
}

func (f *flags) updateBit(pos uint, val byte) {
	bitmask := byte(1) << pos
	f.val = f.val&^bitmask | (val << pos)
}
