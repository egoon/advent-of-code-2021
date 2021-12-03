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
		inputWidth = 12
		zeroes     = make([]int, inputWidth)
		ones       = make([]int, inputWidth)
		gamma      int
		epsilon    int
	)

	for scanner.Scan() {
		input := scanner.Bytes()

		for i := 0; i < inputWidth; i++ {
			switch input[i] {
			case '0':
				zeroes[i]++
			case '1':
				ones[i]++
			default:
				panic("not binary: ")
			}
		}
	}

	bitValue := 1
	for i := inputWidth - 1; i >= 0; i-- {
		if zeroes[i] > ones[i] {
			epsilon += bitValue
		} else if zeroes[i] < ones[i] {
			gamma += bitValue
		}
		bitValue *= 2
	}

	fmt.Printf("%v, %v\n", zeroes, ones)

	fmt.Println("Day 3, part 1: ", epsilon*gamma)
}
