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
		x1, x2, y1, y2 int
		vents          = Vents{vents: map[int]map[int]int{}}
	)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if x1 == x2 {
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				vents.add(x1, y)
			}
		} else if y1 == y2 {
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				vents.add(x, y1)
			}
		} else { // comment out for solution to part 1
			x := x1
			y := y1
			for i := 0; i <= max(x1, x2)-min(x1, x2); i++ {
				vents.add(x, y)
				if x1 < x2 {
					x++
				} else {
					x--
				}
				if y1 < y2 {
					y++
				} else {
					y--
				}
			}
		}
	}

	fmt.Printf("%v\n", vents)

	fmt.Printf("Day 5, part 2: %d\n", vents.count())
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Vents struct {
	vents map[int]map[int]int
}

func (v Vents) add(x, y int) {
	ys, ok := v.vents[x]
	if !ok {
		ys = map[int]int{}
		v.vents[x] = ys
	}
	ys[y]++
}

func (v Vents) count() (count int) {
	for _, ys := range v.vents {
		for _, y := range ys {
			if y > 1 {
				count++
			}
		}
	}
	return
}
