package main

import (
	"fmt"
	"sort"

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

func sum(nums []int) (ret int) {
	for _, num := range nums {
		ret += num
	}
	return ret
}

func part2(ret int) (weakness int) {
	ints, err := input.ToInts()
	if err != nil {
		panic(err)
	}

	visited := []int{}
	for _, num := range ints {
		// reset visited if num > ret
		if num > ret {
			visited = nil
			continue
		}
		visited = append(visited, num)
		// shift visited until sum is < ret
		for sum(visited) > ret && len(visited) > 0 {
			visited = visited[1:]
		}
		if sum(visited) == ret && len(visited) >= 2 {
			sort.Ints(visited)
			weakness = visited[0] + visited[len(visited)-1]
			break
		}
	}
	return
}

func main() {
	ret := part1()
	fmt.Printf("Part 1: %d\n", ret)
	fmt.Printf("Part 2: %d\n", part2(ret))
}
