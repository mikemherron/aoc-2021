package util

import "strconv"

func Sum(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}

	return sum
}

func TryParseInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}

	return i
}
