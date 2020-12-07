package input

import (
	"bufio"
	"os"
	"strconv"
)

func ScannerFromFile() (*bufio.Scanner, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return &bufio.Scanner{}, err
	}

	scanner := bufio.NewScanner(file)
	return scanner, scanner.Err()
}

func ToInts() ([]int, error) {
	var ret []int

	scanner, err := ScannerFromFile()
	if err != nil {
		return ret, err
	}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return ret, err
		}
		ret = append(ret, num)
	}
	return ret, err
}

func ToStrings() ([]string, error) {
	var ret []string

	scanner, err := ScannerFromFile()
	if err != nil {
		return ret, err
	}

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, err
}
