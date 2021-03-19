package main

import (
	"fmt"

	"github.com/crashoz/hpgbc/x86"
	"github.com/crashoz/hpgbc/z80"
)

func main() {
	p := z80.Program{
		z80.Instruction{Op: z80.LDr8r8, Arg1: z80.A, Arg2: z80.B},
		z80.Instruction{Op: z80.LDr8r8, Arg1: z80.C, Arg2: z80.A},
		z80.Instruction{Op: z80.NOP, Arg1: nil, Arg2: nil},
		z80.Instruction{Op: z80.LDn16A, Arg1: 12, Arg2: nil},
		z80.Instruction{Op: z80.LDr8r8, Arg1: z80.C, Arg2: z80.A},
		z80.Instruction{Op: z80.LDAr16, Arg1: z80.BC, Arg2: nil},
		z80.Instruction{Op: z80.LDAHLD, Arg1: nil, Arg2: nil},
		z80.Instruction{Op: z80.LDHLSPe8, Arg1: 24, Arg2: nil},
	}

	t := Translate(p)

	fmt.Println(x86.Compile(t))
}
