package day12

import (
	"AdventCode2021/util"
	"strings"
)

const (
	start = "start"
	end   = "end"
	sep   = "-"
)

type cave struct {
	name        string
	end         bool
	start       bool
	small       bool
	connections []*cave
}

func (n *cave) connect(o *cave) {
	n.connections = append(n.connections, o)
	o.connections = append(o.connections, n)
}

func newCave(n string) *cave {
	return &cave{
		name:        n,
		start:       n == start,
		end:         n == end,
		small:       strings.ToLower(n) == n,
		connections: make([]*cave, 0),
	}
}

func findOrCreateCave(name string, caves map[string]*cave) *cave {
	if c, ok := caves[name]; ok {
		return c
	}

	c := newCave(name)
	caves[name] = c

	return c
}

func CountPaths(input []string, canRevisit bool) int {
	return traverse(buildCaveGraph(input), make(map[string]bool), canRevisit)
}

func buildCaveGraph(input []string) *cave {
	caves := make(map[string]*cave)
	for _, s := range input {
		names := strings.Split(s, sep)
		cave1 := findOrCreateCave(names[0], caves)
		cave2 := findOrCreateCave(names[1], caves)
		cave1.connect(cave2)
	}

	return caves[start]
}

func traverse(cave *cave, visited map[string]bool, canRevisit bool) int {

	visited[cave.name] = true
	if cave.end {
		return 1
	}

	p := 0
	for _, next := range cave.connections {
		if next.start {
			continue
		}
		if next.small {
			if !visited[next.name] {
				p += traverse(next, util.CopyMap(visited), canRevisit)
			} else if canRevisit {
				p += traverse(next, util.CopyMap(visited), false)
			}
		} else {
			p += traverse(next, util.CopyMap(visited), canRevisit)
		}
	}

	return p
}
