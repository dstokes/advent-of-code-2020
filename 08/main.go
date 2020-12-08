package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Instruction struct {
	Operation string
	Argument  int
}

func Execute(i Instruction) (offset int, delta int) {
	switch i.Operation {
	case "acc":
		return 1, i.Argument
	case "jmp":
		return i.Argument, 0
	case "nop":
		return 1, 0
	}
	return
}

func NewInstruction(i []byte) (Instruction, error) {
	operation := string(i[0:3])
	arg, err := strconv.Atoi(string(i[4:]))
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{Operation: operation, Argument: arg}, nil
}

func part1() (accumulator int) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	offset := 0
	data = bytes.Trim(data, "\n")
	lines := bytes.Split(data, []byte{'\n'})
	visited := map[int]struct{}{}

	for {
		instruction, err := NewInstruction(lines[offset])
		if err != nil {
			panic(err)
		}

		if _, ok := visited[offset]; ok {
			break
		}
		visited[offset] = struct{}{}

		off, acc := Execute(instruction)

		offset += off
		accumulator += acc

		if offset == len(lines) {
			break
		}
	}

	return
}

func part2() (accumulator int) {
	return
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}
