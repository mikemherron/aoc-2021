package grid

import (
	"AdventCode2021/util"
	"fmt"
)

type Grid [][]int

func NewGrid(lines []string) Grid {
	g := make(Grid, len(lines))
	cols := len(lines[0])
	for r, l := range lines {
		g[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			g[r][c] = util.TryParseInt(string(l[c]))
		}
	}

	return g
}

type Pos struct {
	Row, Col int
}

func (p Pos) String() string {
	return fmt.Sprintf("[Row: %d, Col: %d]", p.Row, p.Col)
}

type Visitor func(p Pos, v *int)

type VisitPath [][]int

var (
	Surrounding = VisitPath{
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

	Immediate = VisitPath{
		{0, 0},
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
)

func (g Grid) VisitAdjacent(rowStart int, colStart int, path VisitPath, f Visitor) {
	for _, p := range path {
		r := rowStart + p[0]
		if r < 0 || r >= len(g) {
			continue
		}

		c := colStart + p[1]
		if c < 0 || c >= len(g[r]) {
			continue
		}

		pos := Pos{r, c}
		f(pos, &g[r][c])
	}
}

func (g Grid) VisitAll(f Visitor) {
	for r := range g {
		for c := range g[r] {
			f(Pos{r, c}, &g[r][c])
		}
	}
}

func (g Grid) Cells() int {
	return len(g) * len(g[0])
}

func (g *Grid) String() string {
	s, r := "", 0
	g.VisitAll(func(p Pos, v *int) {
		if p.Row > r {
			s += "\n"
			r = p.Row
		}
		if *v == 0 {
			s += "*"
		} else {
			s += fmt.Sprintf("%d", *v)
		}
	})

	return s
}
