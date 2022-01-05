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
		prev, cur, next string
		risk            int
	)

	for {
		prev = cur
		cur = next
		if scanner.Scan() {
			next = scanner.Text()
		} else {
			next = ""
		}

		if cur == "" {
			if next == "" {
				break
			}
			continue
		}

		risk += calcRisk(prev, cur, next)
	}

	fmt.Printf("Day 9, part 1: %d\n", risk)
}

func calcRisk(prev, cur, next string) int {
	var risk int
	for i := range cur {
		c := cur[i]
		if (prev != "" && prev[i] <= c) ||
			(next != "" && next[i] <= c) ||
			(i > 0 && cur[i-1] <= c) ||
			(i+1 < len(cur) && cur[i+1] <= c) {
			continue
		}
		fmt.Println(c, int(c)-'0'+1)
		risk += int(c) - '0' + 1
	}
	return risk
}
