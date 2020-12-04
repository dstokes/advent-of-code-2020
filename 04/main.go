package main

import (
	"fmt"
	"strings"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

type passport map[string]string

var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func hasRequiredFields(p passport) (valid bool) {
	for _, r := range required {
		if _, ok := p[r]; !ok {
			return
		}
	}
	return true
}

func part1() (valid int) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	buf := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			buf = append(buf, strings.Split(line, " ")...)
			continue
		}

		p := passport{}
		for _, kv := range buf {
			p[kv[:3]] = kv[4:]
		}
		if hasRequiredFields(p) == true {
			valid++
		}
	}
	return
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}
