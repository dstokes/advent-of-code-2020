package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var nums []int

func read() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	read()
	fmt.Println(part1())
	fmt.Println(part2())
}
