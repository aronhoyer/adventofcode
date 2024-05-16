package solutions

import (
	"fmt"
)

type solverfn func([]string) int

var solutions []solverfn = []solverfn{
	solveDayOne,
}

func SolveDay(day int, input []string) (int, error) {
	if int(day) > len(solutions) {
		return 0, fmt.Errorf("there are no solutions for day %d", day)
	}

	return solutions[day-1](input), nil
}
