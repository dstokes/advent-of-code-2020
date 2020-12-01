package main

func part1() int {
	for i, v := range nums {
		for ii, vv := range nums {
			if ii == i {
				continue
			}
			if v+vv == 2020 {
				return v * vv
			}
		}
	}
	return 0
}
