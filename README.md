# Advent of Code 2020

Day 1
-----

Day 2
-----

### Benchmarks

```
goos: darwin
goarch: amd64
pkg: github.com/dstokes/advent-of-code-2020/02
BenchmarkPart1-12           3990            252644 ns/op
BenchmarkPart2-12           4314            279405 ns/op
PASS
ok      github.com/dstokes/advent-of-code-2020/02       2.358s
```

Day 3
-----

This one was fun. Part two got me with an off-by-one error but it was
easily fixed.

### Benchmarks

```
goos: darwin
goarch: amd64
pkg: github.com/dstokes/advent-of-code-2020/03
BenchmarkPart1-12          29230             41723 ns/op
BenchmarkPart2-12          17044             70384 ns/op
PASS
ok      github.com/dstokes/advent-of-code-2020/03       3.635s
```

Day 4
-----

This one felt like work.

### Benchmarks

```
04 #master ✗ go test -memprofile mem.prof -bench .; echo; go tool pprof -text -nodecount 10 mem.prof
goos: darwin
goarch: amd64
pkg: github.com/dstokes/advent-of-code-2020/04
BenchmarkPart1-12        3141       381199 ns/op
BenchmarkPart2-12           1914            581040 ns/op
PASS
ok      github.com/dstokes/advent-of-code-2020/04       3.616s

Type: alloc_space
Time: Dec 4, 2020 at 4:53pm (PST)
Showing nodes accounting for 1006.70MB, 99.75% of 1009.21MB total
Dropped 26 nodes (cum <= 5.05MB)
Showing top 10 nodes out of 11
      flat  flat%   sum%        cum   cum%
  386.05MB 38.25% 38.25%   743.13MB 73.64%  github.com/dstokes/advent-of-code-2020/04.part1
  268.01MB 26.56% 64.81%   268.01MB 26.56%  strings.genSplit
  183.01MB 18.13% 82.94%   183.01MB 18.13%  bufio.(*Scanner).Text (inline)
  132.02MB 13.08% 96.02%   265.56MB 26.31%  github.com/dstokes/advent-of-code-2020/04.part2
   28.61MB  2.84% 98.86%    28.61MB  2.84%  bufio.(*Scanner).Scan
       9MB  0.89% 99.75%        9MB  0.89%  github.com/dstokes/advent-of-code-2020/04.match
         0     0% 99.75%   743.13MB 73.64%  github.com/dstokes/advent-of-code-2020/04.BenchmarkPart1
         0     0% 99.75%   265.56MB 26.31%  github.com/dstokes/advent-of-code-2020/04.BenchmarkPart2
         0     0% 99.75%   268.01MB 26.56%  strings.Split (inline)
         0     0% 99.75%  1008.20MB 99.90%  testing.(*B).launch
```
