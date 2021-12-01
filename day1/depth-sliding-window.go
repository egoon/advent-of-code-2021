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
		windowSize = 3
		lastWindowDepth int
		scans int
		window = make([]int, windowSize)
		increments int
	)


	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}


		for i := 0; i < windowSize; i ++ {
			window[i] += depth
		}
		currentWindowDepth := window[scans % windowSize]
		window[scans % windowSize] = 0

		if scans >= windowSize  {
			if lastWindowDepth < currentWindowDepth {
				increments++
			}
		}
		//fmt.Printf("c: %d, l: %d, %v\n", currentWindowDepth, lastWindowDepth, window)

		scans++
		lastWindowDepth = currentWindowDepth
	}

	fmt.Println("Day 1, part 2: ", increments)
}
