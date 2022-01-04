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

	var all = make([]string, 10)
	var digits = make([]string, 4)
	var count int

	for scanner.Scan() {
		_, _ = fmt.Sscanf(scanner.Text(), "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&all[0], &all[1], &all[2], &all[3], &all[4], &all[5], &all[6], &all[7], &all[8], &all[9],
			&digits[0], &digits[1], &digits[2], &digits[3])

		for _, d := range digits {
			l := len(d)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}

	fmt.Printf("Day 8, part 1: %d\n", count)
}
