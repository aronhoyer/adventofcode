package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

var digits [18]string = [18]string{
	"1", "one",
	"2", "two",
	"3", "three",
	"4", "four",
	"5", "five",
	"6", "six",
	"7", "seven",
	"8", "eight",
	"9", "nine",
}

func parseDigit(digit string) int {
	switch digit {
	case "1", "one":
		return 1
	case "2", "two":
		return 2
	case "3", "three":
		return 3
	case "4", "four":
		return 4
	case "5", "five":
		return 5
	case "6", "six":
		return 6
	case "7", "seven":
		return 7
	case "8", "eight":
		return 8
	case "9", "nine":
		return 9
	}

	return 0
}

func getFirstDigit(l string) int {
	for i := 0; i < len(l); i++ {
		for _, digit := range digits {
			if strings.HasPrefix(l[i:], digit) {
				return parseDigit(digit)
			}
		}
	}

	return 0
}

func getLastDigit(l string) int {
	for i := len(l); i > 0; i-- {
		for _, digit := range digits {
			if strings.HasSuffix(l[0:i], digit) {
				return parseDigit(digit)
			}
		}
	}

	return 0
}

func solveDayOne(input []string) int {
	ans := 0

	for _, l := range input {
		first := getFirstDigit(l)
		last := getLastDigit(l)

		sub, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		ans += sub
	}

	return ans
}
