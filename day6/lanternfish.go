package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	var (
		ageCount = make([]int, 9)
	)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	for _, n := range strings.Split(scanner.Text(), ",") {
		number, _ := strconv.Atoi(n)
		ageCount[number]++
	}

	for day := 0; day < 256; day++ { // change to 80 for part 1 answer
		breeders := ageCount[0]
		for age := 1; age < 9; age++ {
			ageCount[age-1] = ageCount[age]
		}
		ageCount[6] += breeders
		ageCount[8] = breeders
		//fmt.Printf("%v\n", ageCount)
	}

	fishCount := 0
	for _, c := range ageCount {
		fishCount += c
	}

	fmt.Printf("Day 6, part 1: %d\n", fishCount)

}
