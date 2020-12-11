package main

import (
	L "./lib"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	bags := L.ParseBagsData(filename)

	res := bags["shiny gold"].CountInnerBags()
	fmt.Printf("%d\n", res)
}
