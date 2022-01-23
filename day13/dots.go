package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)

	var (
		ds         = dots{}
		maxX, maxY int
		first      = true
	)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		var x, y int
		_, _ = fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		ds.Add(x, y)

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		dir := scanner.Text()[11]
		line, _ := strconv.Atoi(scanner.Text()[13:])

		switch dir {
		case 'x':
			ds.FoldX(line)
			maxX = line
		case 'y':
			ds.FoldY(line)
			maxY = line
		}

		if first {
			fmt.Printf("Day 13, part 1: %d\n", ds.Count())
			first = false
		}

	}

	ds.Print(maxX, maxY)
}

type dots struct {
	dots map[int]map[int]bool
}

func (d *dots) Add(x, y int) {
	if d.dots == nil {
		d.dots = map[int]map[int]bool{}
	}
	ys, ok := d.dots[x]
	if !ok {
		ys = map[int]bool{}
	}
	ys[y] = true
	d.dots[x] = ys
}

func (d *dots) Count() (count int) {
	for _, ys := range d.dots {
		for _, y := range ys {
			if y {
				count++
			}
		}
	}
	return count
}

func (d *dots) FoldY(fold int) {
	for x, ys := range d.dots {
		for y, dot := range ys {
			if dot && y > fold {
				ys[y] = false
				d.Add(x, y-2*(y-fold))
			}
		}
	}
}

func (d *dots) FoldX(fold int) {
	for x, ys := range d.dots {
		if x > fold {
			for y, dot := range ys {
				if dot {
					ys[y] = false
					d.Add(x-2*(x-fold), y)
				}
			}

		}
	}
}

func (d *dots) Print(maxX, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if d.dots[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
