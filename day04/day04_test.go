package day04

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name   string
	input []string
	first int
	last int
}

func TestBingo(t *testing.T) {
	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLinesFrom("04_example_input.txt"),
			4512,
			1924,
		},
		{
			"Real",
			puzzleinput.ReadLinesFrom("04_input.txt"),
			72770,
			13912,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualFirst, actualLast := Bingo(c.input)
			if actualFirst != c.first {
				t.Errorf("expected %d, got %d", c.first, actualFirst)
			}
			if actualLast != c.last {
				t.Errorf("expected %d, got %d", c.last, actualLast)
			}
		})
	}
}
