package day21

import (
	"testing"
)

func TestPlayDeterministic(t *testing.T) {

	type testCases struct {
		name    string
		p1Start int
		p2Start int
		result  int
	}

	cases := []testCases{
		{
			"Example",
			4,
			8,
			739785,
		},
		{
			"Real",
			9,
			3,
			1073709,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := PlayDeterministic(c.p1Start, c.p2Start)
			if actual != c.result {
				t.Errorf("expected %d, got %d", c.result, actual)
			}
		})
	}
}

func TestPlayQuantum(t *testing.T) {

	type testCases struct {
		name    string
		p1Start int
		p2Start int
		result  int
	}

	cases := []testCases{
		{
			"Example",
			4,
			8,
			444356092776315,
		},
		{
			"Real",
			9,
			3,
			148747830493442,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := PlayQuantum(c.p1Start, c.p2Start)
			if actual != c.result {
				t.Errorf("expected %d, got %d", c.result, actual)
			}
		})
	}
}
