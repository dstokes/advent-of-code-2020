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

// thanks to https://github.com/timzero for his
// suggestion to add caching here
func walk(adapter int, adapters []int, max int, cache map[int]int) int {
	valid := false
	for _, v := range adapters {
		if v == adapter {
			valid = true
		}
	}
	if valid == false {
		return 0
	} else if adapter == max {
		return 1
	}

	if _, ok := cache[adapter]; !ok {
		total := 0
		for i := 1; i <= 3; i++ {
			total += walk(adapter+i, adapters, max, cache)
		}
		cache[adapter] = total
	}
	return cache[adapter]
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

func part2() int {
	adapters, err := input.ToInts()
	if err != nil {
		panic(err)
	}

	adapters = append(adapters, 0)
	sort.Ints(adapters)
	max := adapters[len(adapters)-1]

	cache := map[int]int{}
	return walk(0, adapters, max, cache)
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
