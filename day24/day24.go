package day24

import (
	"AdventCode2021/day24/alu"
	"AdventCode2021/util"
	"fmt"
	"strconv"
)

type pair struct {
	i, j int
}

type comparator func(new int, existing int) bool

var Lowest comparator = func(new int, existing int) bool {
	return new < existing
}

var Highest comparator = func(new int, existing int) bool {
	return new > existing
}

func FindValidModelNumber(versionComparator comparator, source []string) int {

	a := alu.NewALU()
	program := a.Compile(source)

	candidates := findPairCandidates(versionComparator, a, program)
	valid, number := findValid(make([]int, 14), candidates, a, program)
	if !valid {
		panic("no valid model number")
	}

	return number
}

func findPairCandidates(c comparator, a *alu.ALU, p *alu.Program) map[pair]pair {
	pairCandidates := make(map[pair]pair, 0)
	for i := 1; i <= 13; i++ {
		iLabel := label(i)
		for j := i + 1; j <= 14; j++ {
			jLabel := label(j)

			position := pair{i, j}
			for n1 := 1; n1 <= 9; n1++ {
				for n2 := 1; n2 <= 9; n2++ {

					a.RunLabel(p, iLabel, 0, 0, 0, []int{n1})
					a.RunLabel(p, jLabel, a.X, a.Y, a.Z, []int{n2})

					if a.Z != 0 {
						continue
					}

					value := pair{n1, n2}
					if existing, ok := pairCandidates[position]; ok {
						if c(value.i, existing.i) {
							pairCandidates[position] = value
						}
					} else {
						pairCandidates[position] = value
					}
				}
			}
		}
	}

	return pairCandidates
}

func findValid(version []int, candidates map[pair]pair, a *alu.ALU, p *alu.Program) (bool, int) {
	isComplete := true
	for i, v := range version {

		if v > 0 {
			continue
		}

		//This index has not had a candidate pair created yet
		//the number is not complete, don't validate at the end
		isComplete = false

		//Find a candidate pair to fill this slot
		for digits, values := range candidates {

			//Candidates are indexed from 1, but version number is from 0
			//If this pair is a candidate for empty digit.
			if digits.i == i+1 {

				//Copy candidates array, only take those that don't
				//conflict with the pairs we've just selected
				pairsCopy := make(map[pair]pair)
				for k, v := range candidates {
					if k.i != digits.i && k.i != digits.j && k.j != digits.i && k.j != digits.j {
						pairsCopy[k] = v
					}
				}

				//Copy the version number
				versionCopy := make([]int, len(version))
				copy(versionCopy, version)

				//Apply the pair candidate to the number
				versionCopy[digits.i-1] = values.i
				versionCopy[digits.j-1] = values.j

				//Recurse to fill other slots with remaining candidates
				found, value := findValid(versionCopy, pairsCopy, a, p)
				if found {
					return true, value
				}
			}
		}

		break
	}

	if isComplete {
		a.Run(p, version)
		if a.Z == 0 {
			return true, join(version)
		}
	}

	return false, 0
}

func join(n []int) int {
	var s string
	for _, v := range n {
		s += strconv.Itoa(v)
	}

	return util.TryParseInt(s)
}

func label(n int) string {
	return fmt.Sprintf("%d:", n)
}
