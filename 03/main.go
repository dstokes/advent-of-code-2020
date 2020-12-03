package main

import (
	"fmt"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

const tree = '#'

func part1(x int) (trees int) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	scanner.Scan() // skip the first row
	for scanner.Scan() {
		row := scanner.Bytes()
		if row[x%len(row)] == tree {
			trees++
		}
		x += 3
	}
	return
}

func main() {
	fmt.Printf("Part 1: %d\n", part1(3))
}
