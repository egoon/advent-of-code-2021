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

	type loc struct {
		x, y int
	}

	var (
		risk      [][]int
		accRisk   [][]int
		locations = []loc{{x: 0, y: 0}}
		x         int
	)

	for scanner.Scan() {
		risk = append(risk, make([]int, len(scanner.Bytes())))
		accRisk = append(accRisk, make([]int, len(scanner.Bytes())))
		for y, r := range scanner.Bytes() {
			risk[x][y] = int(r - '0')
		}
		x++
	}

	for len(locations) > 0 {
		l := locations[0]
		locations = locations[1:]
		ar := accRisk[l.x][l.y]
		if l.x > 0 && (accRisk[l.x-1][l.y] == 0 || ar+risk[l.x-1][l.y] < accRisk[l.x-1][l.y]) {
			accRisk[l.x-1][l.y] = ar + risk[l.x-1][l.y]
			locations = append(locations, loc{x: l.x - 1, y: l.y})
		}
		if l.x+1 < len(risk) && (accRisk[l.x+1][l.y] == 0 || ar+risk[l.x+1][l.y] < accRisk[l.x+1][l.y]) {
			accRisk[l.x+1][l.y] = ar + risk[l.x+1][l.y]
			locations = append(locations, loc{x: l.x + 1, y: l.y})
		}
		if l.y > 0 && (accRisk[l.x][l.y-1] == 0 || ar+risk[l.x][l.y-1] < accRisk[l.x][l.y-1]) {
			accRisk[l.x][l.y-1] = ar + risk[l.x][l.y-1]
			locations = append(locations, loc{x: l.x, y: l.y - 1})
		}
		if l.y+1 < len(risk[0]) && (accRisk[l.x][l.y+1] == 0 || ar+risk[l.x][l.y+1] < accRisk[l.x][l.y+1]) {
			accRisk[l.x][l.y+1] = ar + risk[l.x][l.y+1]
			locations = append(locations, loc{x: l.x, y: l.y + 1})
		}
	}

	fmt.Println("Day 15, part 1:", accRisk[len(risk)-1][len(risk[0])-1])

}
