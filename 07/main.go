package main

import (
	"io/ioutil"
	"strings"
)

func readFile() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), "\n")
}

func getBounds(input string) (int, int) {
	leftBound := strings.Index(input, "[")
	rightBound := strings.Index(input, "[")
	return leftBound, rightBound
}

func cycleForPattern(s string, left, right int) bool {
	for i := left; i < right; i++ {
		if i >= len(s)-3 {
			// bail early before trying to access out-of-bound array element
			return false
		}

		if s[i] == s[i+1] {
			// cannot have an "aaaa" pattern
			return false
		}

	}
}

func p1(input []string) {
	for _, line := range input {
		leftBound, rightBound := getBounds(line)
		isLeft := cycleForPattern(line, 0, leftBound)
		isRight := cycleForPattern(line, rightBound+1, len(line)-1)
	}
}

func main() {
	input := readFile()
	p1(input)
}
