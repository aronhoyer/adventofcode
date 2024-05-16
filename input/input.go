package input

import (
	"bufio"
	"os"
	"strings"
)

func scanToLines(s *bufio.Scanner) ([]string, error) {
	lines := []string{}

	for s.Scan() {
		line := s.Text()

		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		if err := s.Err(); err != nil {
			return []string{}, err
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func ReadFromStdin() ([]string, error) {
	s := bufio.NewScanner(os.Stdin)

	return scanToLines(s)
}

func ReadFromFile(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return []string{}, err
	}

	s := bufio.NewScanner(f)
	return scanToLines(s)
}
