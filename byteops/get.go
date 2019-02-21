package byteops

// Get returns a copy of the value of the nth bit.
func Get(b byte, n uint8) (uint8, error) {
	if isOutOfBounds(n) {
		return 0, errOutOfBounds
	}

	return (b >> n) & byte(1), nil
}
