package day22

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

func TestPowerOn(t *testing.T) {

	type testCases struct {
		name  string
		input []string
		on    int
	}

	cases := []testCases{
		{
			"Simple",
			puzzleinput.ReadLines("22_simple_input.txt"),
			39,
		},
		{
			"Simple Negatives",
			puzzleinput.ReadLines("22_simple_input_with_negatives.txt"),
			39,
		},
		{
			"Simple Off Test",
			puzzleinput.ReadLines("22_off_test_case.txt"),
			1,
		},
		{
			"Simple On Test",
			puzzleinput.ReadLines("22_on_test_case.txt"),
			8,
		},
		{
			"Example Part One",
			puzzleinput.ReadLines("22_example_input_part_one.txt"),
			590784,
		},
		{
			"Real Part One",
			puzzleinput.ReadLines("22_input_part_one.txt"),
			547648,
		},
		{
			"Example Part Two",
			puzzleinput.ReadLines("22_example_input_part_two.txt"),
			2758514936282235,
		},
		{
			"Real Part Two",
			puzzleinput.ReadLines("22_input_part_two.txt"),
			1206644425246111,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := PowerOn(c.input)
			if actual != c.on {
				t.Errorf("expected %d, got %d", c.on, actual)
			}
		})
	}
}

func TestCuboidSize(t *testing.T) {

	type testCases struct {
		c        Cuboid
		expected int
	}

	cases := []testCases{
		{Cuboid{from: &Pos{x: 1, y: 1, z: 1}, to: &Pos{x: 1, y: 1, z: 1}}, 1},
		{Cuboid{from: &Pos{x: -1, y: 1, z: 1}, to: &Pos{x: 1, y: 1, z: 1}}, 3},
		{Cuboid{from: &Pos{x: -20, y: -36, z: -47}, to: &Pos{x: 26, y: 17, z: 7}}, 139590},
		{Cuboid{from: &Pos{x: -20, y: -31, z: -26}, to: &Pos{x: 33, y: 23, z: 28}}, 163350},
	}

	for _, c := range cases {
		actual := c.c.size()
		if actual != c.expected {
			t.Errorf("for %v expected size %d, got %d", c.c, c.expected, actual)
		}
	}
}

func TestAxisIntersection(t *testing.T) {

	type testCases struct {
		a        []int
		b        []int
		expected []int
	}

	cases := []testCases{
		{[]int{-10, 10}, []int{-5, 5}, []int{-5, 5}},
		{[]int{-5, 5}, []int{-10, 10}, []int{-5, 5}},

		{[]int{-10, 10}, []int{-5, 15}, []int{-5, 10}},
		{[]int{-5, 15}, []int{-10, 10}, []int{-5, 10}},

		{[]int{-10, 10}, []int{-15, 15}, []int{-10, 10}},
		{[]int{-15, 15}, []int{-10, 10}, []int{-10, 10}},

		{[]int{-10, 10}, []int{-15, 5}, []int{-10, 5}},
		{[]int{-15, 5}, []int{-10, 10}, []int{-10, 5}},
	}

	for _, c := range cases {
		actualA, actualB := axisIntersection(c.a[0], c.a[1], c.b[0], c.b[1])
		if actualA != c.expected[0] || actualB != c.expected[1] {
			t.Errorf("for %v, %v, expected %d, %d, got %d, %d", c.a, c.b, c.expected[0], c.expected[1], actualA, actualB)
		}
	}
}
