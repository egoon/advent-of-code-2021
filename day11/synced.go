package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	var (
		octopus = make([][]int, 10)
		i       int
		step    = 0
		flashes int
	)

	for scanner.Scan() {
		octopus[i] = make([]int, 10)

		for j, c := range scanner.Bytes() {
			octopus[i][j] = int(c) - '0'
		}
		i++
	}

	for step = 0; flashes < 100; step++ {
		flashes = 0
		for i = 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				flashes += inc2(octopus, i, j)
			}
		}

		for i = 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octopus[i][j] > 9 {
					octopus[i][j] = 0
				}
			}
		}

	}

	fmt.Printf("Day 11, part 2: %d\n", step)
}

func inc2(octopus [][]int, i, j int) int {
	flashes := 0
	if i >= 0 && i < 10 && j >= 0 && j < 10 {
		octopus[i][j]++
		if octopus[i][j] == 10 {
			flashes++
			flashes += inc2(octopus, i-1, j-1)
			flashes += inc2(octopus, i-1, j)
			flashes += inc2(octopus, i-1, j+1)
			flashes += inc2(octopus, i, j-1)
			flashes += inc2(octopus, i, j+1)
			flashes += inc2(octopus, i+1, j-1)
			flashes += inc2(octopus, i+1, j)
			flashes += inc2(octopus, i+1, j+1)
		}
	}

	return flashes
}
