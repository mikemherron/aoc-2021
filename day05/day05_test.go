package day05

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name      string
	input     []string
	total     int
	diagonals bool
}

func TestFindDangerousVents(t *testing.T) {
	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("05_example_input.txt"),
			5,
			false,
		},
		{
			"Example with Diagonals",
			puzzleinput.ReadLines("05_example_input.txt"),
			12,
			true,
		},
		{
			"Real",
			puzzleinput.ReadLines("05_input.txt"),
			7674,
			false,
		},
		{
			"Real with Diagonals",
			puzzleinput.ReadLines("05_input.txt"),
			20898,
			true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := FindDangerousVents(c.input, c.diagonals)
			if actual != c.total {
				t.Errorf("expected %d, got %d", c.total, actual)
			}
		})
	}
}
