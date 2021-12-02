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
		aim        int
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
			aim -= amount
		case "down":
			aim += amount
		case "forward":
			horizontal += amount
			depth += amount * aim
		}
		//fmt.Printf("%7s %3d, a: %3d, h: %3d, d: %5d\n", dir, amount, aim, horizontal, depth)
	}

	fmt.Println("Day 2, part 2: ", depth*horizontal)
}
