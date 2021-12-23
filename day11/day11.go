package day11

import (
	"AdventCode2021/util/grid"
)

func CountFlashes(input []string, target int) (int, int) {
	var flashes, targetFlashes int
	g := grid.NewIntGrid(input)
	i := 1
	for {
		previous := flashes
		g.VisitAll(flash(g, &flashes))
		g.VisitAll(clear)

		if i == target {
			targetFlashes = flashes
		}

		if flashes-previous == g.Cells() {
			return targetFlashes, i
		}

		i++
	}
}

func flash(g grid.Grid, flashes *int) grid.Visitor {
	return func(p grid.Pos, v *int) {
		if *v == 10 {
			return
		}

		*v++
		if *v == 10 {
			*flashes++
			g.VisitAdjacent(p.Row, p.Col, grid.Surrounding, flash(g, flashes))
		}
	}
}

func clear(p grid.Pos, v *int) {
	if *v == 10 {
		*v = 0
	}
}
