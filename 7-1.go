package main

import (
	L "./lib"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	bags := L.ParseBagsData(filename)

	res := 0
	for _, bag := range bags {
		if bag.CanContain("shiny gold") {
			res += 1
		}
	}

	fmt.Printf("%d\n", res)
}
