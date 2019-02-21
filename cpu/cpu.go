package cpu

func setBit(b *byte, pos uint) {
	*b |= (byte(1) << pos)
}

func clearBit(b *byte, pos uint) {
	*b &^= (byte(1) << pos)
}
