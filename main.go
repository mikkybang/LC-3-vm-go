package main

import "fmt"

const MEMORY_MAX = (1 << 16)

var memory [MEMORY_MAX]uint16

// Register Enums
const (
	R_R0 = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC /* program counter */
	R_COND
	R_COUNT
)

var registers [R_COUNT]uint16

// Operand Enums
const (
	OP_BR   = iota /* branch */
	OP_ADD         /* add  */
	OP_LD          /* load */
	OP_ST          /* store */
	OP_JSR         /* jump register */
	OP_AND         /* bitwise and */
	OP_LDR         /* load register */
	OP_STR         /* store register */
	OP_RTI         /* unused */
	OP_NOT         /* bitwise not */
	OP_LDI         /* load indirect */
	OP_STI         /* store indirect */
	OP_JMP         /* jump */
	OP_RES         /* reserved (unused) */
	OP_LEA         /* load effective address */
	OP_TRAP        /* execute trap */
)

// Condition flags
const (
	FL_POS = 1 << 0 // P
	FL_ZRO = 1 << 1 // Z
	FL_NEG = 1 << 2 // N
)

const (
	PC_START = 0x3000
)

func memRead(address uint16) uint16 {
	return memory[address]
}

func main() {
	fmt.Println("Hello From VM")

	// Set the condition flag to the Z flag
	registers[R_COND] = FL_ZRO

	// Set the program counter to starting position
	registers[R_PC] = PC_START

	isRunning := true

	for isRunning {
		/* FETCH */
		instr := memRead(registers[R_PC])
		registers[R_PC]++
		op := instr >> 12

		switch op {
		case OP_ADD:
			r0 := (instr >> 9) & 0x7
			r1 := (instr >> 6) & 0x7
			isImmMode := bool((instr>>5)&0x1 == 1)
			if isImmMode {
				imm5 := signExtend(instr&0x1F, 5)
				registers[r0] = registers[r1] + imm5
			} else {
				r2 := instr & 0x7
				registers[r0] = registers[r1] + registers[r2]
			}
			updateFlag(registers[r0])
			break
		default:
			fmt.Print("invalid OP code")
		}
	}
	fmt.Print(11 >> 1)
}

func signExtend(x uint16, bitCount int) uint16 {
	if (x>>(bitCount-1))&1 == 1 {
		x |= 0xFFFF << bitCount
	}
	return x
}

func updateFlag(register uint16) {
	if registers[register] == 0 {
		registers[R_COND] = FL_ZRO
	} else if registers[register]>>15 == 1 {
		registers[R_COND] = FL_NEG
	} else {
		registers[R_COND] = FL_POS
	}
}
