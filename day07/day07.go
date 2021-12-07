package day07

import (
	"math"
	"sort"
)

func AlignCrabs(positions []int, scalingFuel bool) int {

	sort.Ints(positions)

	lower, upper := positions[0], positions[len(positions)-1]

	maxMove := upper - lower
	fuelCosts := make([]int, maxMove+1)
	for i := 0; i <= maxMove; i++ {
		if i == 0 || !scalingFuel {
			fuelCosts[i] = i
		} else {
			fuelCosts[i] = fuelCosts[i-1] + i
		}
	}

	min := math.MaxInt
	for i := lower; i < upper; i++ {
		fuel := 0
		for _, v := range positions {
			distance := v - i
			if distance < 0 {
				distance *= -1
			}
			fuel += fuelCosts[distance]
		}

		if fuel < min {
			min = fuel
		}
	}

	return min

}
