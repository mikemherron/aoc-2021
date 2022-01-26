package day15

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name   string
	input  []string
	expand bool
	min    int
}

func TestFindMinRiskPath(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("15_example_input.txt"),
			false,
			40,
		},
		{
			"Small",
			puzzleinput.ReadLines("15_small_input.txt"),
			false,
			18,
		},
		{
			"Real",
			puzzleinput.ReadLines("15_input.txt"),
			false,
			769,
		},
		{
			"Example Expanded",
			puzzleinput.ReadLines("15_example_input.txt"),
			true,
			315,
		},
		{
			"Up",
			puzzleinput.ReadLines("15_up_input.txt"),
			false,
			8,
		},
		{
			"Real Expanded",
			puzzleinput.ReadLines("15_input.txt"),
			true,
			2963,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualResult := FindMinRiskPath(c.input, c.expand)
			if actualResult != c.min {
				t.Errorf("expected min %d, got %d", c.min, actualResult)
			}
		})
	}
}
