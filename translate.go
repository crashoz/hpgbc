package main

import (
	"github.com/crashoz/hpgbc/x86"
	"github.com/crashoz/hpgbc/z80"
)

var byteRegisterMap = map[z80.ByteRegister]x86.ByteRegister{
	z80.A: x86.AH,
	z80.B: x86.CH,
	z80.C: x86.CL,
	z80.D: x86.DH,
	z80.E: x86.DL,
	z80.H: x86.BH,
	z80.L: x86.BL,
}

var wordRegisterMap = map[z80.WordRegister]x86.WordRegister{
	z80.AF: x86.AX,
	z80.BC: x86.CX,
	z80.DE: x86.DX,
	z80.HL: x86.BX,
	z80.SP: x86.BP,
}

var wordToQuadMap = map[z80.WordRegister]x86.QWordRegister{
	z80.AF: x86.RAX,
	z80.BC: x86.RCX,
	z80.DE: x86.RDX,
	z80.HL: x86.RBX,
}

func trans_LDr8r8(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr8r8,
		Arg1: byteRegisterMap[inst.Arg1.(z80.ByteRegister)],
		Arg2: byteRegisterMap[inst.Arg2.(z80.ByteRegister)],
	}
	return res
}

func trans_LDr8n8(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr8n8,
		Arg1: byteRegisterMap[inst.Arg1.(z80.ByteRegister)],
		Arg2: inst.Arg2.(int),
	}
	return res
}

func trans_LDr16n16(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr16n16,
		Arg1: wordRegisterMap[inst.Arg1.(z80.WordRegister)],
		Arg2: inst.Arg2.(int),
	}
	return res
}

func trans_LDHLr8(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op: x86.MOVmemr8,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[z80.HL],
			Scale:  1,
			Offset: 0,
		},
		Arg2: byteRegisterMap[inst.Arg1.(z80.ByteRegister)],
	}
	return res
}

func trans_LDHLn8(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op: x86.MOVmemn8,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[z80.HL],
			Scale:  1,
			Offset: 0,
		},
		Arg2: inst.Arg1.(int),
	}
	return res
}

func trans_LDr8HL(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr8mem,
		Arg1: byteRegisterMap[inst.Arg1.(z80.ByteRegister)],
		Arg2: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[z80.HL],
			Scale:  1,
			Offset: 0,
		},
	}
	return res
}

func trans_LDr16A(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op: x86.MOVmemr8,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[inst.Arg1.(z80.WordRegister)],
			Scale:  1,
			Offset: 0,
		},
		Arg2: byteRegisterMap[z80.A],
	}
	return res
}

func trans_LDn16A(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op: x86.MOVmemr8,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  x86.NIL,
			Scale:  1,
			Offset: inst.Arg1.(int),
		},
		Arg2: byteRegisterMap[z80.A],
	}
	return res
}

func trans_LDHn8A(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op: x86.MOVmemr8,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  x86.NIL,
			Scale:  1,
			Offset: 0xff00 + inst.Arg1.(int),
		},
		Arg2: byteRegisterMap[z80.A],
	}
	return res
}

func trans_LDHCA(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 2)
	res[0] = x86.Instruction{
		Op:   x86.MOVr64r8,
		Arg1: x86.RSI,
		Arg2: byteRegisterMap[z80.C],
	}
	res[1] = x86.Instruction{
		Op: x86.MOVmemr8,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  x86.RSI,
			Scale:  1,
			Offset: 0xff00,
		},
		Arg2: byteRegisterMap[z80.A],
	}
	return res
}

func trans_LDAr16(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr8mem,
		Arg1: byteRegisterMap[z80.A],
		Arg2: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[inst.Arg1.(z80.WordRegister)],
			Scale:  1,
			Offset: 0,
		},
	}
	return res
}

func trans_LDAn16(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr8mem,
		Arg1: byteRegisterMap[z80.A],
		Arg2: x86.MemOperand{
			Base:   x86.RDI,
			Index:  x86.NIL,
			Scale:  1,
			Offset: inst.Arg1.(int),
		},
	}
	return res
}

func trans_LDHAn16(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr8mem,
		Arg1: byteRegisterMap[z80.A],
		Arg2: x86.MemOperand{
			Base:   x86.RDI,
			Index:  x86.NIL,
			Scale:  1,
			Offset: 0xff00 + inst.Arg1.(int),
		},
	}
	return res
}

func trans_LDHAC(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 2)
	res[0] = x86.Instruction{
		Op:   x86.MOVr64r8,
		Arg1: x86.RSI,
		Arg2: byteRegisterMap[z80.C],
	}
	res[1] = x86.Instruction{
		Op:   x86.MOVr8mem,
		Arg1: byteRegisterMap[z80.A],
		Arg2: x86.MemOperand{
			Base:   x86.RDI,
			Index:  x86.RSI,
			Scale:  1,
			Offset: 0xff00,
		},
	}
	return res
}

func trans_LDHLIA(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 2)
	res[0] = x86.Instruction{
		Op: x86.MOVmemr8,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[z80.HL],
			Scale:  1,
			Offset: 0,
		},
		Arg2: byteRegisterMap[z80.A],
	}
	res[1] = x86.Instruction{
		Op:   x86.INCr16,
		Arg1: wordRegisterMap[z80.HL],
		Arg2: nil,
	}
	return res
}

func trans_LDHLDA(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 2)
	res[0] = x86.Instruction{
		Op: x86.MOVmemr8,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[z80.HL],
			Scale:  1,
			Offset: 0,
		},
		Arg2: byteRegisterMap[z80.A],
	}
	res[1] = x86.Instruction{
		Op:   x86.DECr16,
		Arg1: wordRegisterMap[z80.HL],
		Arg2: nil,
	}
	return res
}

func trans_LDAHLD(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 2)
	res[0] = x86.Instruction{
		Op:   x86.MOVr8mem,
		Arg1: byteRegisterMap[z80.A],
		Arg2: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[z80.HL],
			Scale:  1,
			Offset: 0,
		},
	}
	res[1] = x86.Instruction{
		Op:   x86.DECr16,
		Arg1: wordRegisterMap[z80.HL],
		Arg2: nil,
	}
	return res
}

func trans_LDAHLI(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 2)
	res[0] = x86.Instruction{
		Op:   x86.MOVr8mem,
		Arg1: byteRegisterMap[z80.A],
		Arg2: x86.MemOperand{
			Base:   x86.RDI,
			Index:  wordToQuadMap[z80.HL],
			Scale:  1,
			Offset: 0,
		},
	}
	res[1] = x86.Instruction{
		Op:   x86.INCr16,
		Arg1: wordRegisterMap[z80.HL],
		Arg2: nil,
	}
	return res
}

func trans_LDSPn16(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr16n16,
		Arg1: wordRegisterMap[z80.SP],
		Arg2: inst.Arg2.(int),
	}
	return res
}

func trans_LDn16SP(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op: x86.MOVmemr16,
		Arg1: x86.MemOperand{
			Base:   x86.RDI,
			Index:  x86.NIL,
			Scale:  1,
			Offset: inst.Arg1.(int),
		},
		Arg2: wordRegisterMap[z80.SP],
	}
	return res
}

func trans_LDHLSPe8(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 2)
	res[0] = x86.Instruction{
		Op:   x86.MOVr16r16,
		Arg1: wordRegisterMap[z80.HL],
		Arg2: wordRegisterMap[z80.SP],
	}
	res[1] = x86.Instruction{
		Op:   x86.ADDr16n16,
		Arg1: wordRegisterMap[z80.HL],
		Arg2: inst.Arg1.(int),
	}
	return res
}

func trans_LDSPHL(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{
		Op:   x86.MOVr16r16,
		Arg1: wordRegisterMap[z80.SP],
		Arg2: wordRegisterMap[z80.HL],
	}
	return res
}

func trans_NOP(inst z80.Instruction) []x86.Instruction {
	res := make([]x86.Instruction, 1)
	res[0] = x86.Instruction{Op: x86.NOP, Arg1: nil, Arg2: nil}
	return res
}

var translateDispatch = map[z80.Operation]func(z80.Instruction) []x86.Instruction{
	z80.NOP:      trans_NOP,
	z80.LDr8r8:   trans_LDr8r8,
	z80.LDr8n8:   trans_LDr8n8,
	z80.LDr16n16: trans_LDr16n16,
	z80.LDHLr8:   trans_LDHLr8,
	z80.LDHLn8:   trans_LDHLn8,
	z80.LDr8HL:   trans_LDr8HL,
	z80.LDr16A:   trans_LDr16A,
	z80.LDn16A:   trans_LDn16A,
	z80.LDHn8A:   trans_LDHn8A,
	z80.LDHCA:    trans_LDHCA,
	z80.LDAr16:   trans_LDAr16,
	z80.LDAn16:   trans_LDAn16,
	z80.LDHAn16:  trans_LDHAn16,
	z80.LDHAC:    trans_LDHAC,
	z80.LDHLIA:   trans_LDHLIA,
	z80.LDHLDA:   trans_LDHLDA,
	z80.LDAHLD:   trans_LDAHLD,
	z80.LDAHLI:   trans_LDAHLI,
	z80.LDSPn16:  trans_LDSPn16,
	z80.LDn16SP:  trans_LDn16SP,
	z80.LDHLSPe8: trans_LDHLSPe8,
	z80.LDSPHL:   trans_LDSPHL,
}

func Translate(program z80.Program) x86.Program {
	res := x86.Program{}

	for _, inst := range program {
		tr := translateDispatch[inst.Op](inst)
		res = append(res, tr...)
	}

	return res
}
