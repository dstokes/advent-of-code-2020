package main

import (
	"fmt"
	"sort"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

func search(coords []byte, haystack []int) int {
	h := haystack
	for _, v := range coords {
		if v == 'L' || v == 'F' {
			h = h[:len(h)/2]
		}
		if v == 'R' || v == 'B' {
			h = h[len(h)/2:]
		}
	}
	return h[0]
}

func part1(ids *[]int) (max int) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	cols := []int{0, 1, 2, 3, 4, 5, 6, 7}
	rows := make([]int, 128)
	for i := 0; i < 128; i++ {
		rows[i] = i
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		row := search(line[:len(line)-3], rows)
		col := search(line[len(line)-3:], cols)
		id := row*8 + col
		*ids = append(*ids, id)
		if id > max {
			max = id
		}
	}

	return
}

func part2(ids []int) int {
	sort.Ints(ids)
	off := ids[0]
	for _, v := range ids {
		if off != v {
			return off
		}
		off++
	}
	return off
}

func main() {
	var ids []int
	max := part1(&ids)
	fmt.Printf("Part 1: %d\n", max)
	fmt.Printf("Part 2: %d\n", part2(ids))
}
