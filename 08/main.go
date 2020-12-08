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

func NewInstruction(i []byte) (Instruction, error) {
	operation := string(i[0:3])
	arg, err := strconv.Atoi(string(i[4:]))
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{Operation: operation, Argument: arg}, nil
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

func Run(instructions [][]byte, offset int, overrides map[int]struct{}) (accumulator int, exited bool) {
	visited := map[int]struct{}{}

	for {
		var instruction Instruction
		instruction, _ = NewInstruction(instructions[offset])

		if _, ok := overrides[offset]; ok {
			if instruction.Operation == "jmp" {
				instruction.Operation = "nop"
			} else {
				instruction.Operation = "jmp"
			}
		}

		// infinite loop
		if _, ok := visited[offset]; ok {
			break
		}
		visited[offset] = struct{}{}

		off, acc := Execute(instruction)

		offset += off
		accumulator += acc

		// beyond instuction set
		if offset == len(instructions) {
			exited = true
			break
		}
	}

	return
}

func part1() (accumulator int) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	offset := 0
	data = bytes.Trim(data, "\n")
	lines := bytes.Split(data, []byte{'\n'})

	accumulator, _ = Run(lines, offset, map[int]struct{}{})
	return
}

func part2() (accumulator int) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	offset := 0
	data = bytes.Trim(data, "\n")
	lines := bytes.Split(data, []byte{'\n'})

	fix := []int{}

	// find the nops / jmps
	for i, l := range lines {
		if string(l[:3]) == "jmp" || string(l[:3]) == "nop" {
			fix = append(fix, i)
		}
	}

	for {
		overrides := map[int]struct{}{}

		overrides[fix[0]] = struct{}{}
		fix = fix[1:]

		acc, exited := Run(lines, offset, overrides)
		if exited {
			accumulator = acc
			break
		}
	}

	return
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
