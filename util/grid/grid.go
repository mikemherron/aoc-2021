package grid

import (
	"AdventCode2021/util"
	"fmt"
	"strconv"
	"strings"
)

type Grid [][]int

func NewIntGrid(lines []string) Grid {
	return NewGrid(lines, func(s string) int {
		return util.TryParseInt(s)
	})
}

func NewGrid(lines []string, reader CellReader) Grid {
	g := make(Grid, len(lines))
	cols := len(lines[0])
	for r, l := range lines {
		g[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			g[r][c] = reader(string(l[c]))
		}
	}

	return g
}

type CellReader func(string) int

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
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	Immediate = VisitPath{
		{0, 0},
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	DownAndRight = VisitPath{
		{1, 0},
		{0, 1},
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

func (g Grid) VisitAdjacentWithFiller(rowStart int, colStart int, filler int, path VisitPath, f Visitor) {
	for _, p := range path {
		r := rowStart + p[0]
		c := colStart + p[1]
		var val *int
		if r < 0 || r >= len(g) || c < 0 || c >= len(g[r]) {
			val = &filler
		} else {
			val = &g[r][c]
		}

		pos := Pos{r, c}
		f(pos, val)
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

func (g Grid) Value(r int, c int) int {
	return g[r][c]
}

func (g *Grid) String() string {
	return g.Render(func(i int) string {
		return fmt.Sprintf("%d", i)
	})
}

func (g *Grid) Render(renderer func(int) string) string {
	s, r := "", 0
	g.VisitAll(func(p Pos, v *int) {
		if p.Row > r {
			s += "\n"
			r = p.Row
		}

		s += renderer(*v)
	})

	return s
}

func (g *Grid) Copy() *Grid {
	c := NewIntGrid(strings.Split(g.String(), "\n"))
	return &c
}

func (g *Grid) Grow(val int) *Grid {
	valString := strconv.Itoa(val)
	lines := strings.Split(g.String(), "\n")
	cols := len(lines[0])
	emptyLine := strings.Repeat(valString, cols)

	lines = append([]string{emptyLine, emptyLine}, lines...)
	lines = append(lines, emptyLine, emptyLine)
	for i, line := range lines {
		lines[i] = fmt.Sprintf("%s%s%s", strings.Repeat(valString, 2), line, strings.Repeat(valString, 2))
	}

	grown := NewIntGrid(lines)
	return &grown
}
