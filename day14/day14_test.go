package day14

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name   string
	input  []string
	steps  int
	result int
}

func TestRunTemplate(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("14_example_input.txt"),
			10,
			1588,
		},
		{
			"Real",
			puzzleinput.ReadLines("14_input.txt"),
			10,
			2947,
		},
		{
			"Example 40 steps",
			puzzleinput.ReadLines("14_example_input.txt"),
			40,
			2188189693529,
		},
		{
			"Real 40 steps",
			puzzleinput.ReadLines("14_input.txt"),
			40,
			3232426226464,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualResult := RunTemplate(c.input, c.steps)
			if actualResult != c.result {
				t.Errorf("expected result %d for steps %d, got %d", c.result, c.steps, actualResult)
			}
		})
	}
}
