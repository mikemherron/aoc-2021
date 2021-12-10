package day10

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name        string
	input       []string
	errors      int
	completions int
}

func TestGetRiskLevel(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("10_example_input.txt"),
			26397,
			288957,
		},
		{
			"Real",
			puzzleinput.ReadLines("10_input.txt"),
			290691,
			2768166558,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualErrors, actualCompletions := GetSyntaxErrors(c.input)
			if actualErrors != c.errors {
				t.Errorf("expected errors %d, got %d", c.errors, actualErrors)
			}
			if actualCompletions != c.completions {
				t.Errorf("expected completions %d, got %d", c.completions, actualCompletions)
			}
		})
	}
}
