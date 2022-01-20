package day25

import (
	"AdventCode2021/util/grid"
	"fmt"
)

const (
	East  = 1
	South = 2
	None  = 3
)

func StepsToNoMovement(input []string) int {

	g := grid.NewGrid(input, func(s string) int {
		if s == ">" {
			return East
		} else if s == "v" {
			return South
		} else if s == "." {
			return None
		} else {
			panic(fmt.Sprintf("Unexpected input: %s", s))
		}
	})

	steps := 0
	for {
		eastMoves := make(map[grid.Pos]int)
		g.VisitAll(func(p grid.Pos, v *int) {
			if *v != East {
				return
			}

			nextCol := p.Col + 1
			if nextCol == len(g[0]) {
				nextCol = 0
			}

			if g[p.Row][nextCol] == None {
				eastMoves[p] = nextCol
			}
		})

		for pos, nextCol := range eastMoves {
			g[pos.Row][pos.Col] = None
			g[pos.Row][nextCol] = East
		}

		southMoves := make(map[grid.Pos]int)
		g.VisitAll(func(p grid.Pos, v *int) {
			if *v != South {
				return
			}

			nextRow := p.Row + 1
			if nextRow == len(g) {
				nextRow = 0
			}

			if g[nextRow][p.Col] == None {
				southMoves[p] = nextRow
			}
		})

		for pos, nextRow := range southMoves {
			g[pos.Row][pos.Col] = None
			g[nextRow][pos.Col] = South
		}

		steps++

		if len(southMoves) == 0 && len(eastMoves) == 0 {
			break
		}
	}

	return steps
}
