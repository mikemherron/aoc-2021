package day12

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name              string
	input             []string
	allowSmallRevisit bool
	paths             int
}

func TestCountPaths(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("12_example_input.txt"),
			false,
			10,
		},
		{
			"Real",
			puzzleinput.ReadLines("12_input.txt"),
			false,
			5178,
		},
		{
			"Example with single small revisit",
			puzzleinput.ReadLines("12_example_input.txt"),
			true,
			36,
		},
		{
			"Real with single small revisit",
			puzzleinput.ReadLines("12_input.txt"),
			true,
			130094,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualPaths := CountPaths(c.input, c.allowSmallRevisit)
			if actualPaths != c.paths {
				t.Errorf("expected paths %d, got %d", c.paths, actualPaths)
			}
		})
	}
}
