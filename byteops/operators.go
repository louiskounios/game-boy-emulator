package byteops

import "errors"

var errOutOfBounds = errors.New("requested bit is out of bounds")

// Mutator is the function prototype for the operators that modify the provided
// byte.
type Mutator func(*byte, uint8) error

// Clear modifies b in place by setting the value of the nth bit to 0.
func Clear(b *byte, n uint8) error {
	if isOutOfBounds(n) {
		return errOutOfBounds
	}

	*b &^= (byte(1) << n)
	return nil
}

// Get returns a copy of the value of the nth bit.
func Get(b byte, n uint8) (uint8, error) {
	if isOutOfBounds(n) {
		return 0, errOutOfBounds
	}

	return (b >> n) & byte(1), nil
}

// Set modifies b in place by setting the value of the nth bit to 1.
func Set(b *byte, n uint8) error {
	if isOutOfBounds(n) {
		return errOutOfBounds
	}

	*b |= (byte(1) << n)
	return nil
}

// Toggle modifies b in place by toggling the value of the nth bit.
func Toggle(b *byte, n uint8) error {
	if isOutOfBounds(n) {
		return errOutOfBounds
	}

	*b ^= (byte(1) << n)
	return nil
}

func isOutOfBounds(n uint8) bool {
	if n > 7 {
		return true
	}

	return false
}
