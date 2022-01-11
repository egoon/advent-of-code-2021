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

	var scores []int

	for scanner.Scan() {
		line := scanner.Bytes()

		var open []byte
		corrupted := false

		for _, b := range line {
			if isOpen(b) {
				open = append(open, b)
			} else if len(open) > 0 && isClose(b, open[len(open)-1]) {
				open = open[:len(open)-1]
			} else {
				corrupted = true
				break
			}
		}
		if !corrupted {
			score := 0
			for i := len(open) - 1; i >= 0; i-- {
				score *= 5
				score += points(open[i])
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	fmt.Printf("Day 10, part 2: %d\n", scores[len(scores)/2])
}

func isOpen(o byte) bool {
	return o == '(' || o == '[' || o == '{' || o == '<'
}

func isClose(c byte, o byte) bool {
	return (o == '(' && c == ')') ||
		(o == '[' && c == ']') ||
		(o == '{' && c == '}') ||
		(o == '<' && c == '>')
}

func points(b byte) int {
	switch b {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	default:
		panic(b)
	}
}
