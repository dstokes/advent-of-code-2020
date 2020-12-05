package main

import (
	"fmt"
	"sort"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

func search(pass []byte, min int, max int) int {
	for _, v := range pass {
		mid := (min + max) >> 1
		if v == 'L' || v == 'F' {
			max = mid
		}
		if v == 'R' || v == 'B' {
			min = mid + 1
		}
	}
	return min
}

func part1() (max int) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		row := search(line[:len(line)-3], 0, 127)
		col := search(line[len(line)-3:], 0, 7)
		id := row*8 + col
		if id > max {
			max = id
		}
	}

	return
}

func part2() int {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	var ids []int
	for scanner.Scan() {
		line := scanner.Bytes()
		row := search(line[:len(line)-3], 0, 127)
		col := search(line[len(line)-3:], 0, 7)
		ids = append(ids, row*8+col)
	}

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
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
