package cpu

import "fmt"

// Clock keeps track of the amount of machine cycles (t) and clock cycles (m)
// passed. m is equal to 4*t.
type Clock struct {
	t uint64
	m uint64
}

// NewClock returns a pointer to a new clock with t set to the provided value.
func NewClock(t uint64) *Clock {
	return &Clock{t: t}
}

// T returns the current amount of machine cycles.
func (c *Clock) T() uint64 {
	return c.t
}

// AddT adds the provided amount of machine cycles to the clock.
func (c *Clock) AddT(t uint64) {
	c.SetT(c.t + t)
}

// SetT sets the machine cycles of the clock to the provided value. Clock cycles
// are set to t*4.
func (c *Clock) SetT(t uint64) {
	c.t = t
	c.m = c.t * 4
}

// M returns the current amount of clock cycles.
func (c *Clock) M() uint64 {
	return c.m
}

func (c *Clock) String() string {
	return fmt.Sprintf("t=%d m=%d", c.t, c.m)
}
