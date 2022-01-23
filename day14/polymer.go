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
		pairs  = map[string]int{}
		rules  = map[string]byte{}
		counts = map[byte]int{}
	)

	if scanner.Scan() {
		polymer := scanner.Text()
		for i, b := range []byte(polymer) {
			counts[b]++
			if i+1 < len(polymer) {
				pairs[polymer[i:i+2]]++
			}
		}
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		rule := scanner.Text()
		rules[rule[0:2]] = rule[6]
	}

	for i := 0; i < 40; i++ {
		var newPairs = map[string]int{}
		for pair, count := range pairs {
			b, _ := rules[pair]
			counts[b] += count
			newPairs[string([]byte{pair[0], b})] += count
			newPairs[string([]byte{b, pair[1]})] += count
		}
		pairs = newPairs

	}

	var min, max int
	for _, count := range counts {
		if count < min || min == 0 {
			min = count
		}
		if count > max || max == 0 {
			max = count
		}
	}

	fmt.Printf("Day 14, part 2: %d\n", max-min)

}
