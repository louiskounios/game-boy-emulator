package cpu

type clock struct {
	t uint8
	m uint8
}

func (c clock) T() uint8 {
	return c.t
}

func (c *clock) SetT(t uint8) {
	c.t = t
	c.m = c.t * 4
}

func (c clock) M() uint8 {
	return c.m
}
