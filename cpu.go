package main

type CPU struct {
	Memory [1 << 16]uint16
}

type Instruction uint16

func (instruction Instruction) Add() {
}
