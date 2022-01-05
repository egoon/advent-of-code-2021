package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	var (
		floor      []string
		basinSizes []int
	)

	for scanner.Scan() {
		floor = append(floor, scanner.Text())
	}

	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			if isLowPoint(floor, i, j) {
				basinSizes = append(basinSizes, size(floor, i, j))
				fmt.Println(floor[i][j], i, j, basinSizes[len(basinSizes)-1])
			}
		}
	}

	sort.Ints(basinSizes)

	basinSizes = basinSizes[len(basinSizes)-3:]

	fmt.Printf("Day 9, part 2: %d\n", basinSizes[0]*basinSizes[1]*basinSizes[2])
}

func isLowPoint(floor []string, i, j int) bool {
	c := floor[i][j]
	return (i == 0 || floor[i-1][j] > c) &&
		(i+1 == len(floor) || floor[i+1][j] > c) &&
		(j == 0 || floor[i][j-1] > c) &&
		(j+1 == len(floor[i]) || floor[i][j+1] > c)
}

type loc struct {
	x, y int
}

func size(floor []string, i, j int) int {
	var basin = map[loc]bool{}
	var queue = []loc{{x: i, y: j}}

	for len(queue) > 0 {
		l := queue[0]
		queue = queue[1:]
		if l.x < 0 || l.x >= len(floor) ||
			l.y < 0 || l.y >= len(floor[0]) ||
			floor[l.x][l.y] == '9' || basin[l] {
			continue
		}
		basin[l] = true
		queue = append(queue, loc{x: l.x - 1, y: l.y}, loc{x: l.x + 1, y: l.y}, loc{x: l.x, y: l.y - 1}, loc{x: l.x, y: l.y + 1})
	}

	return len(basin)
}
