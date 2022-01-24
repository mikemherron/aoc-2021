package day24

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

func TestFindValidModelNumber(t *testing.T) {

	type testCase struct {
		name     string
		source   []string
		c        comparator
		expected int
	}

	cases := []testCase{
		{
			"Highest",
			puzzleinput.ReadLines("24_input.txt"),
			Highest,
			12934998949199,
		},
		{
			"Lowest",
			puzzleinput.ReadLines("24_input.txt"),
			Lowest,
			11711691612189,
		},
	}

	for _, c := range cases {
		actual := FindValidModelNumber(c.c, c.source)
		if actual != c.expected {
			t.Errorf("Expected %d, was %d", c.expected, actual)
		}
	}
}
