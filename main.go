package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aronhoyer/adventofcode/input"
	"github.com/aronhoyer/adventofcode/solutions"
)

var day = flag.Int("day", 0, "day to solve")

func main() {
	flag.Parse()

	if *day <= 0 {
		log.Fatal("day must be greater than 0")
	}

	var lines []string
	var err error

	if len(flag.Args()) == 0 {
		lines, err = input.ReadFromStdin()
	} else {
		lines, err = input.ReadFromFile(flag.Arg(0))
	}

	if err != nil {
		log.Fatal(err)
	}

	s, err := solutions.SolveDay(*day, lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
