package main

import (
	"fmt"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

const target = 2020

func part1() int {
	nums, err := input.ToInts()
	if err != nil {
		panic(err)
	}
	seen := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := seen[target-v]; ok {
			return v * (target - v)
		} else {
			seen[v] = struct{}{}
		}
	}
	return 0
}

func part2() int {
	nums, err := input.ToInts()
	if err != nil {
		panic(err)
	}
	seen := make(map[int]struct{})

	for i, v := range nums {
		itarget := target - v
		for ii := i + 1; ii < len(nums); ii++ {
			if _, ok := seen[itarget-nums[ii]]; ok {
				return v * nums[ii] * (itarget - nums[ii])
			} else {
				seen[nums[ii]] = struct{}{}
			}
		}
	}
	return 0
}

func main() {
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(), part2())
}
