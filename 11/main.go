package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func adjacent(x int, y int) [][]int {
	return [][]int{
		{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
		{x - 1, y}, {x + 1, y},
		{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
	}
}

func occupied(coords [][]int, seats [][]byte) (occ int) {
	rows, cols := len(seats), len(seats[0])
	for _, v := range coords {
		if v[1] < 0 || v[1] >= rows {
			continue
		}
		if v[0] < 0 || v[0] >= cols {
			continue
		}
		if seats[v[1]][v[0]] == '#' {
			occ += 1
		}
	}
	return
}

func part1() (total int) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	data = bytes.Trim(data, "\n")
	orig := bytes.Split(data, []byte{'\n'})
	rows, cols := len(orig), len(orig[0])

	for {
		total = 0
		stable := true
		seats := make([][]byte, rows)

		for y := 0; y < rows; y++ {
			seats[y] = make([]byte, cols)
			for x := 0; x < cols; x++ {
				seat := orig[y][x]
				seats[y][x] = seat
				if seat == '.' {
					continue
				}

				count := occupied(adjacent(x, y), orig)
				if seat == 'L' && count == 0 {
					seats[y][x] = '#'
					stable = false
				} else if seat == '#' && count >= 4 {
					seats[y][x] = 'L'
					stable = false
				}

				if seats[y][x] == '#' {
					total += 1
				}
			}
		}

		if stable {
			break
		}
		orig = seats
	}
	return
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}
