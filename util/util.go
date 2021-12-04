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

func TryParseBinary(s string) int {
	i, e := strconv.ParseInt(s, 2, 32)
	if e != nil {
		panic(e)
	}

	return int(i)

}
