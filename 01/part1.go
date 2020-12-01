package main

func part1(nums []int, target int) int {
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
