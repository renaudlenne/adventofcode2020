package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var adapters []int
	adapters = append(adapters, 0)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, value)
	}
	sort.Ints(adapters)
	end := adapters[len(adapters)-1] + 3
	adapters = append(adapters, end)

	var nbArrangements = make([]int, len(adapters))
	for idx, val := range adapters {
		if idx == 0 {
			nbArrangements[0] = 1
		} else {
			nbArrangements[idx] = nbArrangements[idx-1]
		}
		if idx > 1 && val-adapters[idx-2] <= 3 {
			nbArrangements[idx] += nbArrangements[idx-2]
		}
		if idx > 2 && val-adapters[idx-3] <= 3 {
			nbArrangements[idx] += nbArrangements[idx-3]
		}
	}

	fmt.Printf("%d\n", nbArrangements[len(nbArrangements)-1])
}
