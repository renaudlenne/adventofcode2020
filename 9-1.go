package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func verifyChecksum(preamble []int, value int) bool {
	for idx, i := range preamble {
		for _, j := range preamble[idx+1:] {
			if i+j == value {
				return true
			}
		}
	}
	return false
}

const preambleSize = 25

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var xmas []int
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		xmas = append(xmas, value)
	}

	idx := preambleSize
	for verifyChecksum(xmas[idx-preambleSize:idx], xmas[idx]) {
		idx++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", xmas[idx])
}
