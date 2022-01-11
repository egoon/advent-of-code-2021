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

	var score int

	for scanner.Scan() {
		line := scanner.Bytes()

		var open []byte

		for _, b := range line {
			if isOpen(b) {
				open = append(open, b)
			} else if len(open) > 0 && isClose(b, open[len(open)-1]) {
				open = open[:len(open)-1]
			} else { //corrupted
				score += points(b)
				break
			}
		}
	}
	fmt.Printf("Day 10, part 1: %d\n", score)
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
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		panic(b)
	}
}
