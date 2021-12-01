package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func() {_ = file.Close()}()

	scanner := bufio.NewScanner(file)

	var (
		lastDepth int
		firstScan = true
		increments int
	)


	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		if !firstScan {
			if lastDepth < depth {
				increments++
			}
		}
		firstScan = false
		lastDepth = depth
	}

	fmt.Println("Day 1, part 1: ", increments)
}
