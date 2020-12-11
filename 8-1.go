package main

import (
	L "./lib"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	instructions := L.ParseInstructions(filename)

	var executed = make(map[int]bool)
	pointer, acc := instructions[0].Execute(0, 0)

	for !executed[pointer] {
		executed[pointer] = true
		pointer, acc = instructions[pointer].Execute(pointer, acc)
	}

	fmt.Printf("%d\n", acc)
}
