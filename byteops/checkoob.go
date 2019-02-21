package byteops

func isOutOfBounds(n uint8) bool {
	if n > 7 {
		return true
	}

	return false
}
