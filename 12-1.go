package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	x := 0
	y := 0
	direction := 0
	instructionRegex := regexp.MustCompile("^(?P<Action>\\w)(?P<Argument>\\d+)$")
	for scanner.Scan() {
		match := instructionRegex.FindStringSubmatch(scanner.Text())
		action := match[1][0]
		argument, _ := strconv.Atoi(match[2])
		switch action {
		case 'N':
			y += argument
		case 'S':
			y -= argument
		case 'E':
			x += argument
		case 'W':
			x -= argument
		case 'L':
			direction = ((((direction - (argument / 90)) % 4) + 4) % 4) % 4
		case 'R':
			direction = (direction + (argument / 90)) % 4
		case 'F':
			switch direction {
			case 0:
				x += argument
			case 1:
				y -= argument
			case 2:
				x -= argument
			case 3:
				y += argument
			}
		}
	}
	distance := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Printf("%d (%d, %d)\n", int(distance), x, y)
}
