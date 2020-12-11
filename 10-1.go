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
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, value)
	}
	sort.Ints(adapters)
	diffs := [3]int{0, 0, 0}
	for idx, val := range adapters {
		if idx == 0 {
			diffs[val-1]++
		} else {
			diffs[val-adapters[idx-1]-1]++
		}
	}

	fmt.Printf("%d\n", diffs[0]*(diffs[2]+1))
}
