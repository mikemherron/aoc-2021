package util

import (
	"strconv"
	"strings"
)

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}
func Sum(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}

	return sum
}

func Sum64(i []int64) int64 {
	sum := int64(0)
	for _, v := range i {
		sum += v
	}

	return sum
}

func Product64(v []int64) int64 {
	prod := v[0]
	for i := 1; i < len(v); i++ {
		prod *= v[i]
	}

	return prod
}

func Min64(v []int64) int64 {
	min := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] < min {
			min = v[i]
		}
	}

	return min
}

func Max64(v []int64) int64 {
	max := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > max {
			max = v[i]
		}
	}

	return max
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
func AsInts(s []string) []int {
	ints := make([]int, len(s))
	for i, v := range s {
		ints[i] = TryParseInt(v)
	}

	return ints
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

func TryParseBinary64(s string) int64 {
	i, e := strconv.ParseInt(s, 2, 64)
	if e != nil {
		panic(e)
	}

	return i
}

func CopyMap(m map[string]bool) map[string]bool {
	n := make(map[string]bool)
	for k, v := range m {
		n[k] = v
	}

	return n
}
