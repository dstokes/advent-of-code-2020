package main

import (
	"fmt"

	"github.com/dstokes/advent-of-code-2020/pkg/input"
)

const tree = '#'

type route struct {
	x, y, trees int
}

func countTrees(routes []*route) (err error) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		return
	}

	for y := 1; scanner.Scan(); y++ {
		if y == 1 {
			continue
		}
		row := scanner.Bytes()
		for _, r := range routes {
			// skip appropriate rows when y > 1
			if r.y > 1 && y%r.y != 1 {
				continue
			}
			// calculate number of steps based on r.y
			steps := r.x * (y - 1) / r.y
			if row[steps%len(row)] == tree {
				r.trees++
			}
		}
	}
	return
}

func part1() int {
	routes := []*route{{x: 3, y: 1}}
	err := countTrees(routes)
	if err != nil {
		panic(err)
	}
	return routes[0].trees
}

func part2() int {
	routes := []*route{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}
	err := countTrees(routes)
	if err != nil {
		panic(err)
	}
	trees := 1
	for _, r := range routes {
		trees *= r.trees
	}
	return trees
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
