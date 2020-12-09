package main

import (
	"fmt"

	"github.com/dstokes/advent-of-code-2020/pkg/input"
)

func combinations(i []int) map[int]struct{} {
	ret := map[int]struct{}{}
	for _, v := range i {
		for _, vv := range i {
			if v == vv {
				continue
			}
			ret[v+vv] = struct{}{}
		}
	}
	return ret
}

func part1() int {
	ints, err := input.ToInts()
	if err != nil {
		panic(err)
	}

	for i := 25; i < len(ints); i++ {
		c := combinations(ints[i-25 : i])
		if _, ok := c[ints[i]]; !ok {
			return ints[i]
		}
	}
	return 0
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}
