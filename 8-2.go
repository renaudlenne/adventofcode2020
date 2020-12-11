package main

import (
	L "./lib"
	"fmt"
	"os"
)

func executeInstructionSet(instructions []L.Instruction) (int, bool) {
	var executed = make(map[int]bool)
	pointer, acc := instructions[0].Execute(0, 0)

	for !executed[pointer] {
		if pointer == len(instructions) {
			return acc, true
		}
		executed[pointer] = true
		pointer, acc = instructions[pointer].Execute(pointer, acc)
	}

	return acc, false
}

func main() {
	filename := os.Args[1]
	instructions := L.ParseInstructions(filename)
	lastModifiedInstruction := -1

	acc, lastExectionCompleted := executeInstructionSet(instructions)

	for !lastExectionCompleted {
		if lastModifiedInstruction >= 0 {
			instructions[lastModifiedInstruction], _ = instructions[lastModifiedInstruction].SwitchOperation()
		}
		for idx := lastModifiedInstruction + 1; ; idx++ {
			instruction, switched := instructions[idx].SwitchOperation()
			if switched {
				lastModifiedInstruction = idx
				instructions[idx] = instruction
				break
			}
		}
		acc, lastExectionCompleted = executeInstructionSet(instructions)
	}

	fmt.Printf("%d\n", acc)
}
