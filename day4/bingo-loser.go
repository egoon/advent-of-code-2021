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
		numbers       []int
		boards        [][][]square
		board         int
		row           int
		column        int
		winner        [][]square
		winningNumber int
		score         int
	)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	for _, n := range strings.Split(scanner.Text(), ",") {
		number, _ := strconv.Atoi(n)
		numbers = append(numbers, number)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			row = 0
			boards = append(boards, make([][]square, 5))
			board = len(boards) - 1
			continue
		}

		for _, n := range strings.Split(line, " ") {
			if column == 0 {
				boards[board][row] = make([]square, 5)
			}
			number, err := strconv.Atoi(n)
			if err != nil {
				continue
			}
			boards[board][row][column] = square{val: number}
			column++
		}
		column = 0
		row++
	}

numbers:
	for _, n := range numbers {
		for i := 0; i < len(boards); i++ {
			b := boards[i]
		rows:
			for r := 0; r < 5; r++ {
				for c := 0; c < 5; c++ {
					if b[r][c].val == n {
						b[r][c].marked = true
						if bingo(b) {
							if len(boards) == 1 {
								winningNumber = n
								winner = b
								break numbers
							} else {
								boards = append(boards[:i], boards[i+1:]...)
								i--
								fmt.Println(len(boards))
							}
						}
						break rows
					}
				}
			}
		}
	}

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !winner[r][c].marked {
				score += winner[r][c].val
			}
		}
	}
	fmt.Printf("Day 4, part 2: %d\n", winningNumber*score)
}

type square struct {
	val    int
	marked bool
}

func (s square) String() string {
	if s.marked {
		return "*"
	}
	return fmt.Sprintf("%2d", s.val)
}

func bingo(b [][]square) bool {
	for i := 0; i < 5; i++ {
		badRow := false
		badCol := false
		for j := 0; j < 5; j++ {
			badRow = badRow || !b[i][j].marked
			badCol = badCol || !b[j][i].marked
		}
		if !badRow || !badCol {
			return true
		}
	}
	return false
}
