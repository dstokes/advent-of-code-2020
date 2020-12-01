package main

func part2(nums []int, target int) int {
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
