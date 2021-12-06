package day06

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name  string
	input []int
	days  int
	total int64
}

func TestSimulateLanternFish(t *testing.T) {

	cases := []testCase{
		{
			"Example 18 days",
			[]int{3, 4, 3, 1, 2},
			18,
			26,
		},
		{
			"Example 80 days",
			[]int{3, 4, 3, 1, 2},
			80,
			5934,
		},
		{
			"Real 80",
			puzzleinput.ReadCommaSeperatedInts("06_input.txt"),
			80,
			383160,
		},
		{
			"Example 256",
			[]int{3, 4, 3, 1, 2},
			256,
			26984457539,
		},
		{
			"Real 256",
			puzzleinput.ReadCommaSeperatedInts("06_input.txt"),
			256,
			1721148811504,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := SimulateLanternFish(c.input, c.days)
			if actual != c.total {
				t.Errorf("expected %d, got %d", c.total, actual)
			}
		})
	}
}
