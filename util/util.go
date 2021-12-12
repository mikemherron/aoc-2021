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

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func SplitToInt(s, sep string) []int {
	nums := make([]int, 0)
	for _, val := range strings.Split(s, sep) {
		if val != "" {
			nums = append(nums, TryParseInt(val))
		}
	}

	return nums
}

func Filter(s []string, f func(s string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range s {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func Copy(s []string) []string {
	c := make([]string, len(s))
	copy(c, s)
	return c
}

func SplitByCommaToInt(s string) []int {
	return SplitToInt(s, ",")
}

func TryParseInt(s string) int {
	i, e := strconv.Atoi(strings.TrimSpace(s))
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

func CopyMap(m map[string]bool) map[string]bool {
	n := make(map[string]bool)
	for k, v := range m {
		n[k] = v
	}

	return n
}
