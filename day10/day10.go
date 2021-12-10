package day10

import (
	"sort"
)

var (
	chunks = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	syntaxScores = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	completionScores = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
)

type stack []string

func (s *stack) pop() string {
	l := len(*s) - 1
	v := (*s)[l]
	*s = (*s)[:l]

	return v
}

func (s *stack) push(v string) {
	*s = append(*s, v)
}

func GetSyntaxErrors(input []string) (int, int) {
	syntaxTotal, completions := 0, make([]int, 0)
	for _, s := range input {
		syntax, completion := check(s)
		syntaxTotal += syntax
		if completion > 0 {
			completions = append(completions, completion)
		}
	}

	sort.Ints(completions)
	midCompletion := 0
	if len(completions) > 0 {
		midCompletion = completions[len(completions)/2]
	}

	return syntaxTotal, midCompletion
}

func check(line string) (int, int) {
	chunkStack := make(stack, 0)
	for i := range line {
		chunk := string(line[i])
		if closing, ok := chunks[chunk]; ok {
			chunkStack.push(closing)
		} else {
			expected := chunkStack.pop()
			if expected != chunk {
				return syntaxScores[chunk], 0
			}
		}
	}

	completionScore := 0

	for len(chunkStack) > 0 {
		expected := chunkStack.pop()
		completionScore *= 5
		completionScore += completionScores[expected]
	}

	return 0, completionScore
}
