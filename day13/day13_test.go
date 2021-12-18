package day13

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name  string
	input []string
	dots  int
}

func TestCount(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("13_example_input.txt"),
			16,
		},
		{
			"Real",
			puzzleinput.ReadLines("13_input.txt"),
			91,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualDots := CountDots(c.input)
			if actualDots != c.dots {
				t.Errorf("expected dots %d, got %d", c.dots, actualDots)
			}
		})
	}
}
