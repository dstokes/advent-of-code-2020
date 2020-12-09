package main

import (
	"testing"
)

var target int

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ret := part1()
		if target == 0 {
			target = ret
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(target)
	}
}
