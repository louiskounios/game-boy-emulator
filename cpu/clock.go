package cpu

import "fmt"

// Clock keeps track of the amount of machine cycles (m) and clock periods (t)
// passed. t is equal to m*4.
type Clock struct {
	m uint64
	t uint64
}

// NewClock returns a pointer to a new clock with m set to the provided value.
func NewClock(m uint64) *Clock {
	return &Clock{m: m}
}

// Reset sets the machine cycles and clock periods to 0.
func (c *Clock) Reset() {
	c.SetM(0)
}

// M returns the current amount of machine cycles.
func (c *Clock) M() uint64 {
	return c.m
}

// AddM adds the provided amount of machine cycles to the clock.
func (c *Clock) AddM(m uint64) {
	c.SetM(c.m + m)
}

// SetM sets the machine cycles of the clock to the provided value. Clock
// periods are set to m*4.
func (c *Clock) SetM(m uint64) {
	c.m = m
	c.t = m * 4
}

// T returns the current amount of clock periods.
func (c *Clock) T() uint64 {
	return c.t
}

func (c *Clock) String() string {
	return fmt.Sprintf("m=%d t=%d", c.m, c.t)
}
