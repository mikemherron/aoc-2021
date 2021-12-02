package puzzleinput

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadIntsFrom(filePath string) []int {
	lines := ReadLinesFrom(filePath)
	ints := make([]int, 0, len(lines))
	for _, line := range lines {
		i, e := strconv.Atoi(line)
		if e != nil {
			panic(e)
		}
		ints = append(ints, i)
	}

	return ints
}

func ReadLinesFrom(fileName string) []string {
	pwd, e := os.Getwd()
	if e != nil {
		panic(e)
	}

	all, e := ioutil.ReadFile(filepath.Join(pwd,fileName))
	if e != nil {
		panic(e)
	}

	fmt.Println(filepath.Join(pwd,fileName))

	return strings.Split(string(all), "\n")
}
