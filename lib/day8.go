package lib

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	Operation string
	Argument  int
}

type InstructionSet struct {
	instructions []Instruction
}

func (i Instruction) SwitchOperation() (Instruction, bool) {
	switch i.Operation {
	case "jmp":
		return Instruction{
				Operation: "nop",
				Argument:  i.Argument,
			},
			true
	case "nop":
		return Instruction{
				Operation: "jmp",
				Argument:  i.Argument,
			},
			true
	default:
		return i, false
	}
}

func (i Instruction) Execute(pointer int, acc int) (int, int) {
	//fmt.Printf("executing %s %d on (%d, %d)\n", i.operation, i.argument, pointer, acc)

	switch i.Operation {
	case "acc":
		return pointer + 1, acc + i.Argument
	case "jmp":
		return pointer + i.Argument, acc
	default:
		return pointer + 1, acc
	}
}

func ParseInstructions(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instrutionRegexp := regexp.MustCompile("^(?P<Operation>\\w{3}) (?P<Argument>[+-]\\d+)$")
	var instructions []Instruction
	for scanner.Scan() {
		match := instrutionRegexp.FindStringSubmatch(scanner.Text())
		operation := match[1]
		argument, _ := strconv.Atoi(match[2])
		instructions = append(instructions, Instruction{
			Operation: operation,
			Argument:  argument,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return instructions
}
