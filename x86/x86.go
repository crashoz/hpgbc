package x86

import (
	"strconv"
	"strings"
)

type ByteRegister int

const (
	AH ByteRegister = iota
	CH ByteRegister = iota
	CL ByteRegister = iota
	DH ByteRegister = iota
	DL ByteRegister = iota
	BH ByteRegister = iota
	BL ByteRegister = iota
)

var ByteRegisterName = map[ByteRegister]string{
	AH: "ah",
	CH: "ch",
	CL: "cl",
	DH: "dh",
	DL: "dl",
	BH: "bh",
	BL: "bl",
}

type WordRegister int

const (
	AX WordRegister = iota
	CX WordRegister = iota
	DX WordRegister = iota
	BX WordRegister = iota
	BP WordRegister = iota
)

var WordRegisterName = map[WordRegister]string{
	AX: "ax",
	CX: "cx",
	DX: "dx",
	BX: "bx",
	BP: "bp",
}

type QWordRegister int

const (
	NIL QWordRegister = iota
	RAX QWordRegister = iota
	RBX QWordRegister = iota
	RCX QWordRegister = iota
	RDX QWordRegister = iota
	RDI QWordRegister = iota
	RSI QWordRegister = iota
	RBP QWordRegister = iota
	R12 QWordRegister = iota
	R13 QWordRegister = iota
	R14 QWordRegister = iota
	R15 QWordRegister = iota
)

var QWordRegisterName = map[QWordRegister]string{
	NIL: "",
	RAX: "rax",
	RBX: "rbx",
	RCX: "rcx",
	RDX: "rdx",
	RDI: "rdi",
	RSI: "rsi",
	RBP: "rbp",
	R12: "r12",
	R13: "r13",
	R14: "r14",
	R15: "r15",
}

type MemOperand struct {
	Base   QWordRegister
	Index  QWordRegister
	Scale  int
	Offset int
}

type Operation int

const (
	NOP       Operation = iota
	MOVr8r8   Operation = iota
	MOVr8n8   Operation = iota
	MOVr16r16 Operation = iota
	MOVr16n16 Operation = iota
	MOVmemr8  Operation = iota
	MOVmemn8  Operation = iota
	MOVr8mem  Operation = iota
	MOVmemr16 Operation = iota
	MOVr64r8  Operation = iota

	INCr8  Operation = iota
	INCr16 Operation = iota
	DECr8  Operation = iota
	DECr16 Operation = iota

	ADDr8n8   Operation = iota
	ADDr16n16 Operation = iota
)

var OpSize = map[Operation]int{
	NOP:       1,
	MOVr8r8:   1,
	MOVr8n8:   1,
	MOVr16r16: 1,
	MOVr16n16: 1,
	MOVmemr8:  1,
	MOVmemn8:  1,
	MOVr8mem:  1,
	MOVr64r8:  1,
	MOVmemr16: 1,
	INCr8:     1,
	INCr16:    1,
	DECr8:     1,
	DECr16:    1,
	ADDr8n8:   1,
	ADDr16n16: 1,
}

type Instruction struct {
	Op   Operation
	Arg1 interface{}
	Arg2 interface{}
}

type Program []Instruction

func comp_MemOperand(mem MemOperand) string {
	s := "["
	if mem.Base != NIL {
		s += QWordRegisterName[mem.Base]
	}
	if mem.Index != NIL {
		s += " + " + QWordRegisterName[mem.Index]
		if mem.Scale != 1 {
			s += "*" + strconv.Itoa(mem.Scale)
		}
	}
	if mem.Offset != 0 {
		if mem.Offset > 0 {
			s += " + " + strconv.Itoa(mem.Offset)
		} else {
			s += " - " + strconv.Itoa(-mem.Offset)
		}
	}
	return s + "]"
}

func comp_NOP(inst Instruction) string {
	return "nop"
}

func comp_MOVr8r8(inst Instruction) string {
	return "mov " + ByteRegisterName[inst.Arg1.(ByteRegister)] + ", " + ByteRegisterName[inst.Arg2.(ByteRegister)]
}

func comp_MOVr8n8(inst Instruction) string {
	return "mov " + ByteRegisterName[inst.Arg1.(ByteRegister)] + ", " + strconv.Itoa(inst.Arg2.(int))
}

func comp_MOVr16r16(inst Instruction) string {
	return "mov " + WordRegisterName[inst.Arg1.(WordRegister)] + ", " + WordRegisterName[inst.Arg2.(WordRegister)]
}

func comp_MOVr16n16(inst Instruction) string {
	return "mov " + WordRegisterName[inst.Arg1.(WordRegister)] + ", " + strconv.Itoa(inst.Arg2.(int))
}

func comp_MOVmemr8(inst Instruction) string {
	return "mov byte " + comp_MemOperand(inst.Arg1.(MemOperand)) + ", " + ByteRegisterName[inst.Arg2.(ByteRegister)]
}

func comp_MOVmemn8(inst Instruction) string {
	return "mov byte " + comp_MemOperand(inst.Arg1.(MemOperand)) + ", " + strconv.Itoa(inst.Arg2.(int))
}

func comp_MOVr8mem(inst Instruction) string {
	return "mov " + ByteRegisterName[inst.Arg1.(ByteRegister)] + ", byte " + comp_MemOperand(inst.Arg2.(MemOperand))
}

func comp_MOVmemr16(inst Instruction) string {
	return "mov word " + comp_MemOperand(inst.Arg1.(MemOperand)) + ", " + WordRegisterName[inst.Arg2.(WordRegister)]
}

func comp_MOVr64r8(inst Instruction) string {
	return "movzx " + QWordRegisterName[inst.Arg1.(QWordRegister)] + ", " + ByteRegisterName[inst.Arg2.(ByteRegister)]
}

func comp_INCr8(inst Instruction) string {
	return "inc " + ByteRegisterName[inst.Arg1.(ByteRegister)]
}

func comp_INCr16(inst Instruction) string {
	return "inc " + WordRegisterName[inst.Arg1.(WordRegister)]
}

func comp_DECr8(inst Instruction) string {
	return "dec " + ByteRegisterName[inst.Arg1.(ByteRegister)]
}

func comp_DECr16(inst Instruction) string {
	return "dec " + WordRegisterName[inst.Arg1.(WordRegister)]
}

func comp_ADDr8n8(inst Instruction) string {
	return "add " + ByteRegisterName[inst.Arg1.(ByteRegister)] + ", " + strconv.Itoa(inst.Arg2.(int))
}

func comp_ADDr16n16(inst Instruction) string {
	return "add " + WordRegisterName[inst.Arg1.(WordRegister)] + ", " + strconv.Itoa(inst.Arg2.(int))
}

var compileDispatch = map[Operation]func(Instruction) string{
	NOP:       comp_NOP,
	MOVr8r8:   comp_MOVr8r8,
	MOVr8n8:   comp_MOVr8n8,
	MOVr16r16: comp_MOVr16r16,
	MOVr16n16: comp_MOVr16n16,
	MOVmemr8:  comp_MOVmemr8,
	MOVmemn8:  comp_MOVmemn8,
	MOVr8mem:  comp_MOVr8mem,
	MOVmemr16: comp_MOVmemr16,
	MOVr64r8:  comp_MOVr64r8,
	INCr8:     comp_INCr8,
	INCr16:    comp_INCr16,
	DECr8:     comp_DECr8,
	DECr16:    comp_DECr16,
	ADDr8n8:   comp_ADDr8n8,
	ADDr16n16: comp_ADDr16n16,
}

func Compile(program Program) string {
	lines := make([]string, 0)
	for _, inst := range program {
		instStr := compileDispatch[inst.Op](inst)
		lines = append(lines, instStr)
	}

	return strings.Join(lines, "\n")
}
