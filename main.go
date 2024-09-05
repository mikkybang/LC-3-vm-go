package main

import "fmt"

const MEMORY_MAX = (1 << 16)
var memory [MEMORY_MAX]uint16
func main() {
	fmt.Println("Hello From VM")
}
