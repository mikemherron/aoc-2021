package day15

import (
	"AdventCode2021/util/grid"
	"container/heap"
	"math"
)

func FindMinRiskPath(input []string, expand bool) int {
	g := grid.NewIntGrid(input)
	if expand {
		g = expandGrid(g)
	}

	return findPath(&g)
}

type PosCost struct {
	grid.Pos
	cost  int
	index int
}

type PosCostHeap []*PosCost

func (h PosCostHeap) Len() int {
	return len(h)
}

func (h PosCostHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h PosCostHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PosCostHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*PosCost)
	item.index = n
	*h = append(*h, item)
}

func (h *PosCostHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findPath(g *grid.Grid) int {

	shortest := make(PosCostHeap, 0)
	shortestLookup := make(map[grid.Pos]*PosCost)
	g.VisitAll(func(p grid.Pos, v *int) {
		cost := math.MaxInt
		if p.Row == 0 && p.Col == 0 {
			cost = 0
		}

		posCost := &PosCost{
			Pos:  p,
			cost: cost,
		}

		heap.Push(&shortest, posCost)
		shortestLookup[p] = posCost
	})

	for len(shortest) > 0 {

		minNode := heap.Pop(&shortest).(*PosCost)

		g.VisitAdjacent(minNode.Row, minNode.Col, grid.Immediate, func(p grid.Pos, v *int) {
			posCost := shortestLookup[p]
			tentativeCost := minNode.cost + g.Value(p.Row, p.Col)
			if tentativeCost < posCost.cost {
				posCost.cost = tentativeCost
				for i, item := range shortest {
					if item == posCost {
						heap.Fix(&shortest, i)
						break
					}
				}
			}
		})

	}

	return shortestLookup[grid.Pos{Row: len(*g) - 1, Col: len((*g)[0]) - 1}].cost
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
