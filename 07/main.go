package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dstokes/advent-of-code-2020/pkg/input"
)

type Bag struct {
	Name      string
	Contains  []*Bag
	Contained []*Bag
}

func tokenize(line string) (*Bag, map[*Bag]int) {
	words := strings.Split(line, " ")

	name := strings.Join(words[0:2], " ")
	others := words[4:]

	contains := map[*Bag]int{}

	for i := 0; i < len(others); i += 4 {
		count, cname := others[i], strings.Join(others[i+1:i+3], " ")
		ct, err := strconv.Atoi(count)
		if err != nil {
			continue
		}
		contains[&Bag{Name: cname}] = ct
	}

	return &Bag{Name: name}, contains
}

func graph() map[string]*Bag {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	bags := map[string]*Bag{}

	for scanner.Scan() {
		line := scanner.Text()
		bag, contains := tokenize(line)

		if _, ok := bags[bag.Name]; !ok {
			bags[bag.Name] = bag
		} else {
			bag = bags[bag.Name]
		}

		for inner, ct := range contains {
			if _, ok := bags[inner.Name]; !ok {
				bags[inner.Name] = inner
			}
			for i := 0; i < ct; i++ {
				bag.Contains = append(bag.Contains, bags[inner.Name])
			}
			bags[inner.Name].Contained = append(bags[inner.Name].Contained, bag)
		}
	}

	return bags
}

func part1() (total int) {

	bags := graph()

	queue := []string{"shiny gold"}
	visited := map[string]struct{}{}

	for {
		cur := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]

		for _, bag := range bags[cur].Contained {
			if _, ok := visited[bag.Name]; !ok {
				visited[bag.Name] = struct{}{}
				total += 1
				queue = append(queue, bag.Name)
			}
		}
		if len(queue) == 0 {
			break
		}
	}

	return
}

func part2() (total int) {
	bags := graph()

	queue := []*Bag{{Name: "shiny gold"}}
	visited := map[string]struct{}{}

	for {
		cur := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]

		total += len(bags[cur.Name].Contains)
		for _, bag := range bags[cur.Name].Contains {
			if _, ok := visited[bag.Name]; !ok {
				queue = append(queue, &Bag{Name: bag.Name})
			}
		}
		if len(queue) == 0 {
			break
		}
	}
	return
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
