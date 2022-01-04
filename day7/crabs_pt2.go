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
	fuel := cost(positions, pos)
	pFuel := cost(positions, pos+1)

	if fuel != pFuel {

		move := func(i int) int {
			return i - 1
		}
		if fuel > pFuel {
			move = func(i int) int {
				return i + 1
			}
		}
		pFuel = fuel + 1
		pos = move(pos)
		for ; fuel < pFuel; pos = move(pos) {
			pFuel = fuel
			fuel = cost(positions, pos)
			//fmt.Printf("pos: %d, %d (%d)\n", pos, fuel, pFuel)
		}
		fuel = pFuel
	}

	fmt.Printf("Day 7, part 2: %d\n", fuel)
}

var factorials = []int{0}

func fact(i int) int {
	if i < 0 {
		i = -i
	}
	if i < len(factorials) {
		return factorials[i]
	}
	for n := len(factorials); n <= i; n++ {
		factorials = append(factorials, factorials[n-1]+n)
	}
	return factorials[i]
}

func cost(positions []int, depth int) int {
	cost := 0
	for _, p := range positions {
		cost += fact(p - depth)
	}
	return cost
}
