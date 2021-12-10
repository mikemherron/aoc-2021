package day08

import (
	"AdventCode2021/util"
	"fmt"
	"strings"
)

const (
	T  = 0
	TL = 1
	TR = 2
	M  = 3
	BL = 4
	BR = 5
	B  = 6

	Symbols = "abcdefg"
)

var digitSegments = map[int][]int{
	1: {TR, BR},
	2: {T, TR, M, BL, B},
	3: {T, TR, M, BR, B},
	4: {TR, TL, M, BR},
	5: {T, TL, M, BR, B},
	6: {T, TL, M, BR, BL, B},
	7: {T, TR, BR},
	8: {T, TR, TL, M, BL, BR, B},
	9: {T, TR, TL, M, BR, B},
	0: {T, TR, TL, BL, BR, B},
}

var knownLengths = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}

func digit(mapping string, sequence string) int {
	for n, segmentsOn := range digitSegments {
		if len(segmentsOn) != len(sequence) {
			continue
		}
		c := 0
		for _, pos := range segmentsOn {
			if strings.Contains(sequence, string(mapping[pos])) {
				c++
			}
		}

		if c == len(segmentsOn) {
			return n
		}
	}

	return -1
}

func GetDigits(input []string) int {

	allMappings := make([]string, 0)
	permute(Symbols, 0, len(Symbols)-1, &allMappings)

	total := 0

	for _, i := range input {

		possibleMappings := make([]string, len(allMappings))
		copy(possibleMappings, allMappings)

		parts := strings.Split(i, "|")
		patterns := strings.Split(strings.TrimSpace(parts[0]), " ")
		for _, p := range patterns {
			if _, ok := knownLengths[len(p)]; ok {
				continue
			}

			v := knownLengths[len(p)]
			for i := 0; i < len(possibleMappings); i++ {
				if digit(possibleMappings[i], p) != v {
					possibleMappings = append(possibleMappings[:i], possibleMappings[i+1:]...)
					i--
					break
				}
			}
		}

		for i := 0; i < len(possibleMappings); i++ {
			for _, p := range patterns {
				if digit(possibleMappings[i], p) == -1 {
					possibleMappings = append(possibleMappings[:i], possibleMappings[i+1:]...)
					i--
					break
				}
			}
		}

		reading := ""
		for _, o := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			reading += fmt.Sprintf("%d", digit(possibleMappings[0], o))
		}

		total += util.TryParseInt(reading)
	}

	return total
}

func permute(s string, l int, r int, out *[]string) {
	if l == r {
		*out = append(*out, s)
	} else {
		for i := l; i <= r; i++ {
			s = swap(s, l, i)
			permute(s, l+1, r, out)
			s = swap(s, l, i)
		}
	}
}

func swap(s string, i int, j int) string {
	b := []byte(s)
	t := b[i]
	b[i] = b[j]
	b[j] = t
	return string(b)
}
