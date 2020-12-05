package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var ids []int
		part1(&ids)
	}
}
