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

func rotate(rotation int, x int, y int) (int, int) {
	switch rotation {
	case 1:
		return y, -x
	case 2:
		return -x, -y
	case 3:
		return -y, x
	default:
		return x, y
	}
}

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
	waypointX := 10
	waypointY := 1
	instructionRegex := regexp.MustCompile("^(?P<Action>\\w)(?P<Argument>\\d+)$")
	for scanner.Scan() {
		match := instructionRegex.FindStringSubmatch(scanner.Text())
		action := match[1][0]
		argument, _ := strconv.Atoi(match[2])
		switch action {
		case 'N':
			waypointY += argument
		case 'S':
			waypointY -= argument
		case 'E':
			waypointX += argument
		case 'W':
			waypointX -= argument
		case 'L':
			rotation := ((((-argument / 90) % 4) + 4) % 4) % 4
			waypointX, waypointY = rotate(rotation, waypointX, waypointY)
		case 'R':
			rotation := (argument / 90) % 4
			waypointX, waypointY = rotate(rotation, waypointX, waypointY)
		case 'F':
			x += waypointX * argument
			y += waypointY * argument
		}
	}
	distance := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Printf("%d (%d, %d)\n", int(distance), x, y)
}
