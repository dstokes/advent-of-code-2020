package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/dstokes/advent-of-code-2020/pkg/input"
	"github.com/dstokes/advent-of-code-2020/pkg/util"
)

func part1() (total int) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	count := 0
	answers := make(map[byte]struct{})

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			total += count

			count = 0
			answers = make(map[byte]struct{})
			continue
		}

		for _, v := range line {
			if _, ok := answers[v]; !ok {
				count++
			}
			answers[v] = struct{}{}
		}
	}

	return total + count
}

func part2() (total int) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	data = bytes.Trim(data, "\n")
	groups := bytes.Split(data, []byte{'\n', '\n'})

	for _, group := range groups {
		people := bytes.Split(group, []byte{'\n'})

		answers := util.NewSet()
		answers.AddBytes(people[0])
		for _, person := range people[1:] {
			p := util.NewSet()
			p.AddBytes(person)
			answers = answers.Intersect(p)
		}
		total += answers.Cardinality()
	}

	return
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
