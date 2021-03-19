package z80

//https://rgbds.gbdev.io/docs/v0.4.2/gbz80.7#LD_r8,r8

type ByteRegister int

const (
	A ByteRegister = iota
	B ByteRegister = iota
	C ByteRegister = iota
	D ByteRegister = iota
	E ByteRegister = iota
	H ByteRegister = iota
	L ByteRegister = iota
)

type WordRegister int

const (
	AF WordRegister = iota
	BC WordRegister = iota
	DE WordRegister = iota
	HL WordRegister = iota
	SP WordRegister = iota
)

type Operation int

const (
	NOP      Operation = iota
	LDr8r8   Operation = iota
	LDr8n8   Operation = iota
	LDr16n16 Operation = iota
	LDHLr8   Operation = iota
	LDHLn8   Operation = iota
	LDr8HL   Operation = iota
	LDr16A   Operation = iota
	LDn16A   Operation = iota
	LDHn8A   Operation = iota
	LDHCA    Operation = iota
	LDAr16   Operation = iota
	LDAn16   Operation = iota
	LDHAn16  Operation = iota
	LDHAC    Operation = iota
	LDHLIA   Operation = iota
	LDHLDA   Operation = iota
	LDAHLD   Operation = iota
	LDAHLI   Operation = iota
	LDSPn16  Operation = iota
	LDn16SP  Operation = iota
	LDHLSPe8 Operation = iota
	LDSPHL   Operation = iota
)

var OpSize = map[Operation]int{
	NOP:      1,
	LDr8r8:   1,
	LDr8n8:   2,
	LDr16n16: 3,
	LDHLr8:   1,
	LDHLn8:   2,
	LDr8HL:   1,
	LDr16A:   1,
	LDn16A:   3,
	LDHn8A:   2,
	LDHCA:    1,
	LDAr16:   1,
	LDAn16:   3,
	LDHAn16:  2,
	LDHAC:    1,
	LDHLIA:   1,
	LDHLDA:   1,
	LDAHLD:   1,
	LDAHLI:   1,
	LDSPn16:  3,
	LDn16SP:  3,
	LDHLSPe8: 2,
	LDSPHL:   1,
}

var OpCycles = map[Operation]int{
	NOP:      1,
	LDr8r8:   1,
	LDr8n8:   2,
	LDr16n16: 3,
	LDHLr8:   2,
	LDHLn8:   3,
	LDr8HL:   2,
	LDr16A:   2,
	LDn16A:   4,
	LDHn8A:   3,
	LDHCA:    2,
	LDAr16:   2,
	LDAn16:   4,
	LDHAn16:  3,
	LDHAC:    2,
	LDHLIA:   2,
	LDHLDA:   2,
	LDAHLD:   2,
	LDAHLI:   2,
	LDSPn16:  3,
	LDn16SP:  5,
	LDHLSPe8: 3,
	LDSPHL:   2,
}

type Instruction struct {
	Op   Operation
	Arg1 interface{}
	Arg2 interface{}
}

type Program []Instruction
