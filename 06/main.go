package main

import (
	"fmt"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

func part1() (total int) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	count := 0
	answers := make(map[byte]struct{})

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			total += count

			count = 0
			answers = make(map[byte]struct{})
			continue
		}

		for _, v := range line {
			if _, ok := answers[v]; !ok {
				count++
			}
			answers[v] = struct{}{}
		}
	}

	return total + count
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}
