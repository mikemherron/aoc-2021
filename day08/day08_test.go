package day08

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name   string
	input  []string
	digits int
}

func TestGetDigits(t *testing.T) {

	cases := []testCase{
		{
			"Example Part 2 single line",
			[]string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"},
			5353,
		},
		{
			"Example Part 2",
			puzzleinput.ReadLines("08_example_input.txt"),
			61229,
		},
		{
			"Real Part 2",
			puzzleinput.ReadLines("08_input.txt"),
			983030,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := GetDigits(c.input)
			if actual != c.digits {
				t.Errorf("expected %d, got %d", c.digits, actual)
			}
		})
	}
}

func TestGetDigitsSimple(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("08_example_input.txt"),
			26,
		},
		{
			"Real",
			puzzleinput.ReadLines("08_input.txt"),
			355,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := GetDigitsSimple(c.input)
			if actual != c.digits {
				t.Errorf("expected %d, got %d", c.digits, actual)
			}
		})
	}
}
