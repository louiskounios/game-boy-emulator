package byteops

// Toggle modifies b in place by toggling the value of the nth bit.
func Toggle(b *byte, n uint8) error {
	if isOutOfBounds(n) {
		return errOutOfBounds
	}

	*b ^= (byte(1) << n)
	return nil
}
