package main

import (
	"bufio"
	"fmt"
	"os"
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
		cave      = map[string][]string{}
		paths     = []path{{next: "start"}}
		completed []path
	)

	for scanner.Scan() {
		p := strings.Split(scanner.Text(), "-")

		cave[p[0]] = append(cave[p[0]], p[1])
		cave[p[1]] = append(cave[p[1]], p[0])
	}

	for len(paths) > 0 {
		p := paths[0]
		paths = paths[1:]
		valid, revisited := p.IsValid()
		if valid {
			p.caves = append(p.caves, p.next)
			if p.next == "end" {
				completed = append(completed, path{caves: p.caves})
			} else {
				for _, c := range cave[p.next] {
					paths = append(paths, newPath(p.caves, c, revisited))
				}
			}
		}
	}

	fmt.Println("Day 12, part 2: ", len(completed))

}

type path struct {
	caves   []string
	next    string
	revisit bool
}

func newPath(caves []string, next string, revisit bool) path {
	cs := make([]string, len(caves))
	copy(cs, caves)
	return path{caves: cs, next: next, revisit: revisit}
}

func (p path) IsValid() (bool, bool) {
	for _, c := range p.caves {
		if c == p.next {
			if c[0] >= 'a' && c[0] <= 'z' {
				if c != "start" && !p.revisit {
					return true, true
				}
				return false, p.revisit
			}
			return true, p.revisit
		}
	}
	return true, p.revisit
}
