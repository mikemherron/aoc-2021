package day08

import (
	"strings"
)

func GetDigitsSimple(input []string) int {
	counts := map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}

	total := 0

	for _, i := range input {
		parts := strings.Split(i, "|")
		outputs := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, o := range outputs {
			if _, ok := counts[len(o)]; ok {
				total++
			}
		}
	}

	return total
}
