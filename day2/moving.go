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

	scanner := bufio.NewScanner(file)

	var (
		depth      int
		horizontal int
	)

	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		dir := move[0]
		amount, err := strconv.Atoi(move[1])
		if err != nil {
			panic(err)
		}

		switch dir {
		case "up":
			depth -= amount
		case "down":
			depth += amount
		case "forward":
			horizontal += amount
		}
		//fmt.Printf("%s %d, h: %d, d: %d\n", dir, amount, horizontal, depth)
	}

	fmt.Println("Day 2, part 1: ", depth*horizontal)
}
