package byteops

// Set modifies b in place by setting the value of the nth bit to 1.
func Set(b *byte, n uint8) error {
	if isOutOfBounds(n) {
		return errOutOfBounds
	}

	*b |= (byte(1) << n)
	return nil
}
