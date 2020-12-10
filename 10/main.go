package main

import (
	"fmt"
	"sort"

	"github.com/dstokes/advent-of-code-2020/pkg/input"
)

// removes item from a slice
func remove(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

// calculates absolute diff between ints
func difference(a int, b int) (c int) {
	c = a - b
	if c < 0 {
		c *= -1
	}
	return
}

// evaluate adapter compatibility
func compatible(adapter int, jolts int) (bool, int) {
	diff := difference(adapter, jolts)
	return diff > 0 && diff <= 3, diff
}

func part1() int {
	adapters, err := input.ToInts()
	if err != nil {
		panic(err)
	}
	sort.Ints(adapters)

	jolt := 0
	deviceJolts := adapters[len(adapters)-1] + 3
	differences := map[int]int{}

	for {
		for i := 0; i < len(adapters); i++ {
			compat, diff := compatible(adapters[i], jolt)
			if compat {
				jolt = adapters[i]
				differences[diff] += 1
				adapters = remove(adapters, i)
				break
			}
		}
		if len(adapters) <= 0 {
			break
		}
	}

	differences[difference(jolt, deviceJolts)] += 1
	return differences[1] * differences[3]
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}
