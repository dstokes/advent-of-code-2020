package main

import (
	"fmt"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

const target = 2020

func main() {
	nums, err := input.ToInts()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n",
		part1(nums, target), part2(nums, target))
}
