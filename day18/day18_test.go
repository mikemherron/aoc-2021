package day18

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

func TestResolve(t *testing.T) {

	type testCases struct {
		name      string
		input     []string
		magnitude int
		sum       string
	}

	cases := []testCases{
		{
			"No Reduce Example",
			puzzleinput.ReadLines("18_no_reduce_example.txt"),
			0,
			"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			"Long Mangnitude Example",
			puzzleinput.ReadLines("18_magnitude_example.txt"),
			3488,
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			"Explode #1",
			[]string{"[[[[[9,8],1],2],3],4]"},
			0,
			"[[[[0,9],2],3],4]",
		},
		{
			"Explode #2",
			[]string{"[7,[6,[5,[4,[3,2]]]]]"},
			0,
			"[7,[6,[5,[7,0]]]]",
		},
		{
			"Explode #3",
			[]string{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
			0,
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			"Example",
			puzzleinput.ReadLines("18_example.txt"),
			4140,
			"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		},
		{
			"Example 2",
			puzzleinput.ReadLines("18_example_2.txt"),
			0,
			"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},

		{
			"Example Simpler",
			puzzleinput.ReadLines("18_example_simpler.txt"),
			0,
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			"Real",
			puzzleinput.ReadLines("18_input.txt"),
			3359,
			"",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualSum, actualMagnitude := Resolve(c.input)
			if len(c.sum) > 0 && actualSum != c.sum {
				t.Errorf("expected sum %s, got %s", c.sum, actualSum)
			}
			if c.magnitude > 0 && actualMagnitude != c.magnitude {
				t.Errorf("expected magnitude %d, got %d", c.magnitude, actualMagnitude)
			}
		})
	}
}

func TestLargest(t *testing.T) {

	type testCases struct {
		name      string
		input     []string
		magnitude int
	}

	cases := []testCases{
		{
			"Example",
			puzzleinput.ReadLines("18_example.txt"),
			3993,
		},
		{
			"Real",
			puzzleinput.ReadLines("18_input.txt"),
			4616,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualMagnitude := LargestMagnitude(c.input)
			if actualMagnitude != c.magnitude {
				t.Errorf("expected magnitude %d, got %d", c.magnitude, actualMagnitude)
			}
		})
	}
}
