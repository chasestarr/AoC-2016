package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), "\n")
}

func getBounds(input string) (int, int) {
	leftBound := strings.Index(input, "[")
	rightBound := strings.Index(input, "]")
	return leftBound, rightBound
}

// implementation needs work.. didn't realize line could have > 1 bracket sets
func cycleForPattern(s string, left, right int) bool {
	fmt.Println(left, right)
	fmt.Println(s[left : right-1])
	for i := left; i < right; i++ {
		if i >= len(s)-3 {
			// bail early before trying to access out-of-bound array element
			return false
		}

		if s[i] == s[i+1] {
			// cannot have an "aaaa" pattern
			return false
		}

		// "abba" --> "ab" == "ab"
		if s[i]+s[i+1] == s[i+3]+s[i+2] {
			return true
		}
	}

	return false
}

func p1(input []string) {
	var count int
	for _, line := range input {
		leftBound, rightBound := getBounds(line)

		// bail early if "abba" pattern occurs between brackets
		isBetween := cycleForPattern(line, leftBound+1, rightBound-1)
		if isBetween {
			continue
		}

		isLeft := cycleForPattern(line, 0, leftBound)
		isRight := cycleForPattern(line, rightBound+1, len(line)-1)
		if isLeft || isRight {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	input := readFile()
	p1(input)
}
