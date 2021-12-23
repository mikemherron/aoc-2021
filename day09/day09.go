package day09

import (
	"AdventCode2021/util"
	"AdventCode2021/util/grid"
	"sort"
)

func GetRiskLevel(input []string) (int, int) {
	lows, basins, g := make([]int, 0), make([]int, 0), grid.NewIntGrid(input)
	for r := range g {
		for c := range g[r] {
			v, lowest := g[r][c], true
			g.VisitAdjacent(r, c, grid.Surrounding, func(p grid.Pos, compare *int) {
				if v > *compare {
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

func basinSize(g grid.Grid, r int, c int) int {
	checked := make(map[grid.Pos]bool)
	findBasin(checked, g, r, c)

	size := 0
	for _, isBasin := range checked {
		if isBasin {
			size++
		}
	}

	return size
}

func findBasin(checked map[grid.Pos]bool, g grid.Grid, r int, c int) {
	g.VisitAdjacent(r, c, grid.Immediate, func(p grid.Pos, v *int) {
		if _, ok := checked[p]; ok {
			return
		}
		if *v != 9 {
			checked[p] = true
			findBasin(checked, g, p.Row, p.Col)
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
