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

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func minMax(array []int) (int, int) {
	min := array[0]
	max := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
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

	weakValue := xmas[idx]
	i := 0
	j := 1
	currentSum := sum(xmas[i:j])
	for currentSum != weakValue {
		if currentSum > weakValue {
			i++
		} else if currentSum < weakValue {
			j++
		}
		currentSum = sum(xmas[i:j])
	}
	min, max := minMax(xmas[i:j])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", min+max)
}
