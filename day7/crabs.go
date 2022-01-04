package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	input := strings.Split(scanner.Text(), ",")
	positions := make([]int, len(input))

	for i, n := range input {
		number, _ := strconv.Atoi(n)
		positions[i] = number
	}

	sort.Ints(positions)

	pos := positions[len(positions)/2]
	fuel := 0
	fmt.Println(pos, positions[pos])
	for i := 0; i < len(positions)/2; i++ {
		fuel += pos - positions[i]
		fmt.Println(pos - positions[i])
	}
	for i := len(positions) / 2; i < len(positions); i++ {
		fuel += positions[i] - pos
		fmt.Println(positions[i] - pos)
	}

	fmt.Printf("Day 7, part 1: %d\n", fuel)
}
