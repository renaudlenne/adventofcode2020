package lib

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MultipleBag struct {
	bag      Bag
	bagCount int
}

type Bag struct {
	name              string
	contentDefinition string
	innerBags         map[string]MultipleBag
}

func (bag Bag) CanContain(bagName string) bool {
	for _, inner := range bag.innerBags {
		if inner.bag.name == bagName {
			return true
		}
		if inner.bag.CanContain(bagName) {
			return true
		}
	}
	return false
}

func (bag Bag) CountInnerBags() int {
	count := 0
	for _, inner := range bag.innerBags {
		count += inner.bagCount * (1 + inner.bag.CountInnerBags())
	}
	return count
}

func ParseBagsData(filename string) map[string]Bag {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bagLineRegexp := regexp.MustCompile("^(?P<BagName>[\\w\\s]+) bags contain(?P<BagContent>[\\w\\s\\d,]+ bags?)\\.$")
	var bags = make(map[string]Bag)
	for scanner.Scan() {
		match := bagLineRegexp.FindStringSubmatch(scanner.Text())
		bagName := match[1]
		bags[bagName] = Bag{
			name:              match[1],
			contentDefinition: match[2],
			innerBags:         make(map[string]MultipleBag),
		}
	}

	bagContentRegexp := regexp.MustCompile("^ ?(?P<BagCount>\\d+) (?P<BagName>[\\w\\s]+) bags?\\.?$")
	for _, bag := range bags {
		if bag.contentDefinition != " no other bags" {
			for _, innerDef := range strings.Split(bag.contentDefinition, ", ") {
				match := bagContentRegexp.FindStringSubmatch(innerDef)
				innerBagCount, _ := strconv.Atoi(match[1])
				innerBagName := match[2]
				bag.innerBags[innerBagName] = MultipleBag{
					bag:      bags[innerBagName],
					bagCount: innerBagCount,
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bags
}
