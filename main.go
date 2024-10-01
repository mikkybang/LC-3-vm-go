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
			break
		default:
			fmt.Print("invalid OP code")
		}
	}
	fmt.Print(11 >> 1)
}
