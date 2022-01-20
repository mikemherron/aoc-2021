package day25

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

func TestStepsToNoMovement(t *testing.T) {

	type testCase struct {
		name     string
		input    []string
		expected int
	}

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("25_example_input.txt"),
			58,
		},
		{
			"Real",
			puzzleinput.ReadLines("25_real_input.txt"),
			549,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := StepsToNoMovement(c.input)
			if actual != c.expected {
				t.Errorf("expected %d, got %d", c.expected, actual)
			}
		})
	}
}
