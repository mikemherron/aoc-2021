package day16

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

func TestGetVersionNumbers(t *testing.T) {

	type testCase struct {
		name       string
		input      string
		versionSum int
	}

	cases := []testCase{
		{
			"Example 1",
			"8A004A801A8002F478",
			16,
		},
		{
			"Example 2",
			"620080001611562C8802118E34",
			12,
		},
		{
			"Example 3",
			"C0015000016115A2E0802F182340",
			23,
		},
		{
			"Example 4",
			"A0016C880162017C3686B18A3D4780",
			31,
		},
		{
			"Real",
			puzzleinput.ReadLines("16_input.txt")[0],
			953,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualResult, _ := ParseHex(c.input)
			if actualResult != c.versionSum {
				t.Errorf("expected version sum %d, got %d", c.versionSum, actualResult)
			}
		})
	}
}

func TestValues(t *testing.T) {

	type testCase struct {
		name  string
		input string
		value int64
	}

	cases := []testCase{
		{
			"Example 1",
			"C200B40A82",
			3,
		},
		{
			"Example 2",
			"04005AC33890",
			54,
		},
		{
			"Example 3",
			"880086C3E88112",
			7,
		},
		{
			"Example 4",
			"CE00C43D881120",
			9,
		},
		{
			"Example 5",
			"D8005AC2A8F0",
			1,
		},
		{
			"Example 6",
			"F600BC2D8F",
			0,
		},
		{
			"Example 7",
			"9C005AC2F8F0",
			0,
		},
		{
			"Example 8",
			"9C0141080250320F1802104A08",
			1,
		},
		{
			"Example 9",
			"EE00D40C823060",
			3,
		},
		{
			"Real",
			puzzleinput.ReadLines("16_input.txt")[0],
			246225449979,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, actualResult := ParseHex(c.input)
			if actualResult != c.value {
				t.Errorf("expected value %d, got %d", c.value, actualResult)
			}
		})
	}
}
