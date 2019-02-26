package registers

import (
	"fmt"
	"testing"
)

var register16EqualsTests = []struct {
	r1  CompoundRegisterHandler
	r2  CompoundRegisterHandler
	out bool
}{
	{&Register16{0, 0}, &Register16{0, 0}, true},
	{&Register16{0, 1}, &Register16{0, 0}, false},
	{&Register16{128, 0}, &Register16{0, 128}, false},
	{&Register16{255, 0}, &Register16{255, 0}, true},
	{&Register16{255, 255}, &Register16{255, 255}, true},
	{&Register16{255, 255}, &RegisterAF{255, 255}, false},
}

func TestRegister16Equals(t *testing.T) {
	for _, tt := range register16EqualsTests {
		t.Run(fmt.Sprintf("r1=%v r2=%v", tt.r1, tt.r2), func(t *testing.T) {
			if out := tt.r1.Equals(tt.r1) && tt.r1.Equals(tt.r1) && tt.r1.Equals(tt.r2) && tt.r2.Equals(tt.r1); out != tt.out {
				t.Errorf("got %t, expected %t", out, tt.out)
			}
		})
	}
}

var register16HiTests = []struct {
	r   Register16
	out uint8
}{
	{Register16{0, 0}, 0},
	{Register16{120, 0}, 120},
	{Register16{255, 0}, 255},
}

func TestRegister16Hi(t *testing.T) {
	for _, tt := range register16HiTests {
		t.Run(fmt.Sprintf("r=%v", tt.r), func(t *testing.T) {
			if out := tt.r.Hi(); out != tt.out {
				t.Errorf("got %d, expected %d", out, tt.out)
			}
		})
	}
}

var register16SetHiTests = []struct {
	val uint8
	out Register16
}{
	{0, Register16{0, 0}},
	{120, Register16{120, 0}},
	{255, Register16{255, 0}},
}

func TestRegister16SetHi(t *testing.T) {
	for _, tt := range register16SetHiTests {
		r := Register16{}

		t.Run(fmt.Sprintf("val=%d", tt.val), func(t *testing.T) {
			r.SetHi(tt.val)

			if !r.Equals(tt.out) {
				t.Errorf("got %v, expected %v", r, tt.out)
			}
		})
	}
}

var register16LoTests = []struct {
	r   Register16
	out uint8
}{
	{Register16{0, 0}, 0},
	{Register16{0, 120}, 120},
	{Register16{0, 255}, 255},
}

func TestRegister16Lo(t *testing.T) {
	for _, tt := range register16LoTests {
		t.Run(fmt.Sprintf("r=%v", tt.r), func(t *testing.T) {
			if out := tt.r.Lo(); out != tt.out {
				t.Errorf("got %d, expected %d", out, tt.out)
			}
		})
	}
}

var register16SetLoTests = []struct {
	val uint8
	out Register16
}{
	{0, Register16{0, 0}},
	{120, Register16{0, 120}},
	{255, Register16{0, 255}},
}

func TestRegister16SetLo(t *testing.T) {
	for _, tt := range register16SetLoTests {
		r := Register16{}

		t.Run(fmt.Sprintf("val=%d", tt.val), func(t *testing.T) {
			r.SetLo(tt.val)

			if !r.Equals(tt.out) {
				t.Errorf("got %v, expected %v", r, tt.out)
			}
		})
	}
}

var register16WordTests = []struct {
	in  Register16
	out uint16
}{
	{Register16{1, 1}, 257},
	{Register16{128, 1}, 32769},
	{Register16{128, 0}, 32768},
	{Register16{128, 128}, 32896},
	{Register16{255, 255}, 65535},
}

func TestRegister16Word(t *testing.T) {
	for _, tt := range register16WordTests {
		t.Run(fmt.Sprintf("in=%v", tt.in), func(t *testing.T) {
			if r := tt.in.Word(); r != tt.out {
				t.Errorf("got %d, expected %d", r, tt.out)
			}
		})
	}
}

var register16SetWordTests = []struct {
	in  uint16
	out Register16
}{
	{257, Register16{1, 1}},
	{32769, Register16{128, 1}},
	{32768, Register16{128, 0}},
	{32896, Register16{128, 128}},
	{65535, Register16{255, 255}},
}

func TestRegister16SetWord(t *testing.T) {
	for _, tt := range register16SetWordTests {
		r := Register16{}

		t.Run(fmt.Sprintf("in=%d", tt.in), func(t *testing.T) {
			if r.SetWord(tt.in); !r.Equals(tt.out) {
				t.Errorf("got %v, expected %v", r, tt.out)
			}
		})
	}
}

var registerAFEqualsTests = []struct {
	r1  CompoundRegisterHandler
	r2  CompoundRegisterHandler
	out bool
}{
	{&RegisterAF{0, 0}, &RegisterAF{0, 0}, true},
	{&RegisterAF{0, 1}, &RegisterAF{0, 0}, false},
	{&RegisterAF{128, 0}, &RegisterAF{0, 128}, false},
	{&RegisterAF{255, 0}, &RegisterAF{255, 0}, true},
	{&RegisterAF{255, 255}, &RegisterAF{255, 255}, true},
	{&RegisterAF{255, 255}, &Register16{255, 255}, false},
}

func TestRegisterAFEquals(t *testing.T) {
	for _, tt := range registerAFEqualsTests {
		t.Run(fmt.Sprintf("r1=%v r2=%v", tt.r1, tt.r2), func(t *testing.T) {
			if out := tt.r1.Equals(tt.r1) && tt.r1.Equals(tt.r1) && tt.r1.Equals(tt.r2) && tt.r2.Equals(tt.r1); out != tt.out {
				t.Errorf("got %t, expected %t", out, tt.out)
			}
		})
	}
}

var registerAFHiTests = []struct {
	r   RegisterAF
	out uint8
}{
	{RegisterAF{0, 0}, 0},
	{RegisterAF{120, 0}, 120},
	{RegisterAF{255, 0}, 255},
}

func TestRegisterAFHi(t *testing.T) {
	for _, tt := range registerAFHiTests {
		t.Run(fmt.Sprintf("r=%v", tt.r), func(t *testing.T) {
			if out := tt.r.Hi(); out != tt.out {
				t.Errorf("got %d, expected %d", out, tt.out)
			}
		})
	}
}

var registerAFSetHiTests = []struct {
	val uint8
	out RegisterAF
}{
	{0, RegisterAF{0, 0}},
	{120, RegisterAF{120, 0}},
	{255, RegisterAF{255, 0}},
}

func TestRegisterAFSetHi(t *testing.T) {
	for _, tt := range registerAFSetHiTests {
		r := RegisterAF{}

		t.Run(fmt.Sprintf("val=%d", tt.val), func(t *testing.T) {
			r.SetHi(tt.val)

			if !r.Equals(tt.out) {
				t.Errorf("got %v, expected %v", r, tt.out)
			}
		})
	}
}

var registerAFLoTests = []struct {
	r   RegisterAF
	out uint8
}{
	{RegisterAF{0, 0}, 0},
	{RegisterAF{0, 120}, 120},
	{RegisterAF{0, 255}, 255},
}

func TestRegisterAFLo(t *testing.T) {
	for _, tt := range registerAFLoTests {
		t.Run(fmt.Sprintf("r=%v", tt.r), func(t *testing.T) {
			if out := tt.r.Lo(); out != tt.out {
				t.Errorf("got %d, expected %d", out, tt.out)
			}
		})
	}
}

var registerAFSetLoTests = []struct {
	val uint8
	out RegisterAF
}{
	{0, RegisterAF{0, 0}},
	{120, RegisterAF{0, 120}},
	{255, RegisterAF{0, 255}},
}

func TestRegisterAFSetLo(t *testing.T) {
	for _, tt := range registerAFSetLoTests {
		r := RegisterAF{}

		t.Run(fmt.Sprintf("val=%d", tt.val), func(t *testing.T) {
			r.SetLo(tt.val)

			if !r.Equals(tt.out) {
				t.Errorf("got %v, expected %v", r, tt.out)
			}
		})
	}
}

var registerAFWordTests = []struct {
	in  RegisterAF
	out uint16
}{
	{RegisterAF{1, 1}, 257},
	{RegisterAF{128, 1}, 32769},
	{RegisterAF{128, 0}, 32768},
	{RegisterAF{128, 128}, 32896},
	{RegisterAF{255, 255}, 65535},
}

func TestRegisterAFWord(t *testing.T) {
	for _, tt := range registerAFWordTests {
		t.Run(fmt.Sprintf("in=%v", tt.in), func(t *testing.T) {
			if r := tt.in.Word(); r != tt.out {
				t.Errorf("got %d, expected %d", r, tt.out)
			}
		})
	}
}

var registerAFSetWordTests = []struct {
	in  uint16
	out RegisterAF
}{
	{257, RegisterAF{1, 1}},
	{32769, RegisterAF{128, 1}},
	{32768, RegisterAF{128, 0}},
	{32896, RegisterAF{128, 128}},
	{65535, RegisterAF{255, 255}},
}

func TestRegisterAFSetWord(t *testing.T) {
	for _, tt := range registerAFSetWordTests {
		r := RegisterAF{}

		t.Run(fmt.Sprintf("in=%d", tt.in), func(t *testing.T) {
			if r.SetWord(tt.in); !r.Equals(tt.out) {
				t.Errorf("got %v, expected %v", r, tt.out)
			}
		})
	}
}
