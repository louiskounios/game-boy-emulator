package cpu

type instructionSet [256]struct {
	opcode      uint8
	clockCycles uint8
	mnemonic    string
	execute     func() bool
}

var instructions = instructionSet{
	{0x00, 0, "mnemonic", func() bool { return false }},
	{0x01, 0, "mnemonic", func() bool { return true }},
}
