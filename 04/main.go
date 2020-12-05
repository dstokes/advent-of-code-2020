package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dstokes/advent-of-code-2020/lib/input"
)

var rgx = map[string]*regexp.Regexp{
	"hcl": regexp.MustCompile("^#[0-9a-f]{6}$"),
	"ecl": regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$"),
	"pid": regexp.MustCompile("^[0-9]{9}$"),
}

var spec = map[string]func(string) (bool, error){
	"byr": func(v string) (bool, error) { return between(v, 1920, 2002) },
	"iyr": func(v string) (bool, error) { return between(v, 2010, 2020) },
	"eyr": func(v string) (bool, error) { return between(v, 2020, 2030) },
	"hgt": func(v string) (bool, error) { return validHeight(v) },
	"hcl": func(v string) (bool, error) { return match(v, rgx["hcl"]) },
	"ecl": func(v string) (bool, error) { return match(v, rgx["ecl"]) },
	"pid": func(v string) (bool, error) { return match(v, rgx["pid"]) },
	"cid": func(v string) (bool, error) { return true, nil },
}

func match(s string, r *regexp.Regexp) (bool, error) {
	return r.Match([]byte(s)), nil
}

func between(n string, min int, max int) (bool, error) {
	i, err := strconv.Atoi(n)
	if err != nil {
		return false, err
	}
	return i >= min && i <= max, err
}

func validHeight(s string) (bool, error) {
	suffix := string(s[len(s)-2:])
	if suffix == "cm" {
		n, err := strconv.Atoi(s[:len(s)-2])
		if err != nil {
			return false, err
		}
		if n >= 150 && n <= 193 {
			return true, nil
		}
	}
	if suffix == "in" {
		n, err := strconv.Atoi(s[:len(s)-2])
		if err != nil {
			return false, err
		}
		if n >= 59 && n <= 76 {
			return true, nil
		}
	}
	return false, nil
}

func hasRequiredFields(p map[string]struct{}) bool {
	for r := range spec {
		if r == "cid" {
			continue
		}
		if _, ok := p[r]; !ok {
			return false
		}
	}
	return true
}

func part1() (valid int) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	buf := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			buf = append(buf, strings.Split(line, " ")...)
			continue
		}

		p := make(map[string]struct{})
		for _, kv := range buf {
			p[kv[:3]] = struct{}{}
		}
		if hasRequiredFields(p) {
			valid++
		}
		buf = nil
	}
	return
}

func part2() (valid int) {
	scanner, err := input.ScannerFromFile()
	if err != nil {
		panic(err)
	}

	buf := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			buf = append(buf, strings.Split(line, " ")...)
			continue
		}

		p := make(map[string]struct{})
		for _, kv := range buf {
			key, val := kv[:3], kv[4:]
			ok, err := spec[key](val)
			if err != nil {
				panic(err)
			}
			if !ok {
				break
			}
			p[key] = struct{}{}
		}
		if hasRequiredFields(p) {
			valid++
		}
		buf = nil
	}
	return
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
