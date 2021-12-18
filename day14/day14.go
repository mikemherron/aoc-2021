package day14

import (
	"math"
)

type pair struct {
	s uint8
	e uint8
}

func RunTemplate(input []string, steps int) int {
	template := input[0]

	counts := make(map[uint8]int)
	for i := 0; i < len(template); i++ {
		counts[template[i]]++
	}

	rules := make(map[pair]uint8)
	for i := 2; i < len(input); i++ {
		rules[pair{input[i][0], input[i][1]}] = input[i][len(input[i])-1]
	}

	pairs := make(map[pair]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[pair{template[i], template[i+1]}]++
	}

	for i := 0; i < steps; i++ {
		newPairs := make(map[pair]int)
		for p, count := range pairs {
			if v, ok := rules[p]; ok {
				newPairs[pair{p.s, v}] += count
				newPairs[pair{v, p.e}] += count
				counts[v] += count
			} else {
				newPairs[p] = count
			}
		}

		pairs = newPairs
	}

	min, max := math.MaxInt, 0
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}
