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
BenchmarkPart1-12           2898            362095 ns/op
BenchmarkPart2-12            256           4699865 ns/op
PASS
ok      github.com/dstokes/advent-of-code-2020/04       2.808s

Type: alloc_space
Time: Dec 4, 2020 at 2:52pm (PST)
Showing nodes accounting for 1619.72MB, 89.30% of 1813.78MB total
Dropped 17 nodes (cum <= 9.07MB)
Showing top 10 nodes out of 45
      flat  flat%   sum%        cum   cum%
  424.26MB 23.39% 23.39%   424.26MB 23.39%  regexp/syntax.(*compiler).inst
  276.33MB 15.24% 38.63%   276.33MB 15.24%  regexp.onePassCopy
  249.53MB 13.76% 52.38%   249.53MB 13.76%  regexp/syntax.(*parser).newRegexp
  186.52MB 10.28% 62.67%   357.57MB 19.71%  github.com/dstokes/advent-of-code-2020/04.part1
  111.01MB  6.12% 68.79%   111.01MB  6.12%  strings.genSplit
  106.55MB  5.87% 74.66%   259.06MB 14.28%  regexp.makeOnePass
   72.51MB  4.00% 78.66%    72.51MB  4.00%  regexp/syntax.(*Regexp).Simplify
      68MB  3.75% 82.41%       68MB  3.75%  bufio.(*Scanner).Text (inline)
      65MB  3.58% 85.99%    92.50MB  5.10%  regexp.makeOnePass.func1
      60MB  3.31% 89.30%       60MB  3.31%  regexp.newQueue (inline)
```
