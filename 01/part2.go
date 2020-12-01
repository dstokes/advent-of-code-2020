package main

func part2() int {
	for i, v := range nums {
		for ii, vv := range nums {
			if ii == i {
				continue
			}
			for iii, vvv := range nums {
				if iii == ii || iii == i {
					continue
				}
				if v+vv+vvv == 2020 {
					return v * vv * vvv
				}
			}
		}
	}
	return 0
}
