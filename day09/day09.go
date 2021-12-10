package day09

import (
	"AdventCode2021/util"
	"sort"
)

type grid [][]int

func buildGrid(lines []string) grid {
	g := make(grid, len(lines))
	cols := len(lines[0])
	for r, l := range lines {
		g[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			g[r][c] = util.TryParseInt(string(l[c]))
		}
	}

	return g
}

type gridPos struct {
	r, c int
}

type gridVisitor func(p gridPos, v int)

type gridVisitPath [][]int

var (
	surrounding = gridVisitPath{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, 0},
		{0, -1},
		{-1, -1},
	}

	immediate = gridVisitPath{
		{0, 0},
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
)

func (g grid) visit(rowStart int, colStart int, path gridVisitPath, f gridVisitor) {
	for _, p := range path {
		r := rowStart + p[0]
		if r < 0 || r >= len(g) {
			continue
		}

		c := colStart + p[1]
		if c < 0 || c >= len(g[r]) {
			continue
		}

		pos := gridPos{r, c}
		f(pos, g[r][c])
	}
}

func GetRiskLevel(input []string) (int, int) {
	lows, basins, g := make([]int, 0), make([]int, 0), buildGrid(input)
	for r := range g {
		for c := range g[r] {
			v, lowest := g[r][c], true
			g.visit(r, c, surrounding, func(p gridPos, compare int) {
				if v > compare {
					lowest = false
				}
			})

			if lowest {
				lows = append(lows, 1+v)
				basins = append(basins, basinSize(g, r, c))
			}
		}
	}

	return util.Sum(lows), basinScore(basins)
}

func basinSize(g grid, r int, c int) int {
	checked := make(map[gridPos]bool)
	findBasin(checked, g, r, c)

	size := 0
	for _, isBasin := range checked {
		if isBasin {
			size++
		}
	}

	return size
}

func findBasin(checked map[gridPos]bool, g grid, r int, c int) {
	g.visit(r, c, immediate, func(p gridPos, v int) {
		if _, ok := checked[p]; ok {
			return
		}
		if v != 9 {
			checked[p] = true
			findBasin(checked, g, p.r, p.c)
		} else {
			checked[p] = false
		}
	})
}

func basinScore(b []int) int {
	sort.Ints(b)
	top := b[len(b)-3:]
	return top[0] * top[1] * top[2]
}
