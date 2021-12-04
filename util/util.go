package util

import (
	"strconv"
	"strings"
)

func Sum(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}

	return sum
}

func SplitStringToInts(s, sep string) []int {
	nums := make([]int, 0)
	for _, val := range strings.Split(s, sep) {
		if val != "" {
			nums = append(nums, TryParseInt(val))
		}
	}

	return nums
}

func TryParseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}

	return i
}

func TryParseBinary(s string) int {
	i, e := strconv.ParseInt(s, 2, 32)
	if e != nil {
		panic(e)
	}

	return int(i)

}
