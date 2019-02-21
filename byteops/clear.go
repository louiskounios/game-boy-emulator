package byteops

// Clear modifies b in place by setting the value of the nth bit to 0.
func Clear(b *byte, n uint8) error {
	if isOutOfBounds(n) {
		return errOutOfBounds
	}

	*b &^= (byte(1) << n)
	return nil
}
