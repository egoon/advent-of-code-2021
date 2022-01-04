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
	var sum int
	var val int

	for scanner.Scan() {
		_, _ = fmt.Sscanf(scanner.Text(), "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&all[0], &all[1], &all[2], &all[3], &all[4], &all[5], &all[6], &all[7], &all[8], &all[9],
			&digits[0], &digits[1], &digits[2], &digits[3])

		var one, four string
		for _, d := range all {
			switch len(d) {
			case 2:
				one = d
			case 4:
				four = d
			}
		}

		val = 0
		for _, d := range digits {
			val *= 10
			switch len(d) {
			case 2:
				val += 1
			case 3:
				val += 7
			case 4:
				val += 4
			case 5:
				if intersect(d, one) == 2 {
					val += 3
				} else if intersect(d, four) == 2 {
					val += 2
				} else {
					val += 5
				}
			case 6:
				if intersect(d, four) == 4 {
					val += 9
				} else if intersect(d, one) == 2 {
					val += 0
				} else {
					val += 6
				}
			case 7:
				val += 8
			}
		}
		sum += val
	}

	fmt.Printf("Day 8, part 2: %d\n", sum)
}

func intersect(d, n string) int {
	count := 0
	for _, c1 := range []byte(d) {
		for _, c2 := range []byte(n) {
			if c1 == c2 {
				count++
			}
		}
	}
	return count
}
