package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

func tokenize(line string) (min int, max int, char string, pass string, err error) {
	var rule string

	chunks := strings.Split(line, " ")
	rule, char, pass = chunks[0], string(chunks[1][0]), chunks[2]

	sep := strings.Index(rule, "-")
	min, err = strconv.Atoi(rule[:sep])
	if err != nil {
		return min, max, char, pass, err
	}
	max, err = strconv.Atoi(rule[sep+1:])
	if err != nil {
		return min, max, char, pass, err
	}

	return min, max, char, pass, err
}

func part1(lines []string) int {
	valid := 0
	for _, line := range lines {
		min, max, char, pass, err := tokenize(line)
		if err != nil {
			panic(err)
		}
		count := strings.Count(pass, char)
		if count >= min && count <= max {
			valid++
		}
	}
	return valid
}

func part2(lines []string) int {
	valid := 0
	for _, line := range lines {
		min, max, char, pass, err := tokenize(line)
		if err != nil {
			panic(err)
		}
		matches := 0

		if string(pass[min-1]) == char {
			matches++
		}
		if string(pass[max-1]) == char {
			matches++
		}
		if matches > 0 && matches < 2 {
			valid++
		}
	}
	return valid
}

func main() {
	lines, err := input.ToStrings()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
