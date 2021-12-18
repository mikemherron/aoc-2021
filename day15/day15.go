package day15

import (
	"AdventCode2021/util/grid"
	"math"
)

//var (
//	sols = make(map[int]map[grid.Pos]bool)
//)

func FindMinRiskPath(input []string, expand bool) int {
	g := grid.NewGrid(input)
	if expand {
		g = expandGrid(g)
	}
	return minRiskFrom(&g, 0, 0, make(map[grid.Pos]int))
}

func minRiskFrom(g *grid.Grid, row int, col int, solutions map[grid.Pos]int) int {
	pos := grid.Pos{Row: row, Col: col}

	if v, ok := solutions[pos]; ok {
		return v
	}

	risk := 0
	if row != 0 || col != 0 {
		risk = g.Value(row, col)
	}

	if row == len(*g)-1 && col == len((*g)[row])-1 {
		return risk
	}

	min := math.MaxInt

	g.VisitAdjacent(row, col, grid.DownAndRight, func(p grid.Pos, v *int) {
		minFromPath := risk + minRiskFrom(g, p.Row, p.Col, solutions)
		if minFromPath < min {
			min = minFromPath
		}
	})

	solutions[pos] = min
	return min
}

func expandGrid(g grid.Grid) grid.Grid {
	const ExpansionSize = 5
	size := len(g)
	expanded := make(grid.Grid, size*ExpansionSize)
	g.VisitAll(func(p grid.Pos, v *int) {
		for tr := 0; tr < ExpansionSize; tr++ {
			for tc := 0; tc < ExpansionSize; tc++ {
				val := *v + (tr + tc)
				if val > 9 {
					val = val - 9
				}

				row := p.Row + size*tr
				if expanded[row] == nil {
					expanded[row] = make([]int, len(expanded))
				}

				col := p.Col + size*tc
				expanded[row][col] = val
			}
		}
	})

	return expanded
}

//func copyMap(m map[grid.Pos]bool) map[grid.Pos]bool {
//	n := make(map[grid.Pos]bool)
//	for k, v := range m {
//		n[k] = v
//	}
//
//	return n
//}
