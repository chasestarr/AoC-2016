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

func matchPatternTLS(s string) bool {
	var isBetweenBrackets bool
	var output bool

	// bail early before trying to access out-of-bound array element
outer:
	for i := 0; i < len(s)-3; i++ {
		// look ahead at next 4 chars... flip bool
		for j := i; j < i+4; j++ {
			if char := rune(s[j]); char == '[' {
				isBetweenBrackets = true
				continue outer
			} else if char == ']' {
				isBetweenBrackets = false
				continue outer
			}
		}

		// cannot have an "aaaa" pattern
		if s[i] == s[i+1] {
			continue outer
		}

		// "abba" --> "ab" == "ab"
		if string(s[i])+string(s[i+1]) == string(s[i+3])+string(s[i+2]) {
			// if "abba" pattern inBetween, bail and return false
			if isBetweenBrackets {
				return false
			}
			output = true
		}
	}

	return output
}

func matchPatternSSL(s string) bool {
	fmt.Println(s)
	var isBetweenBrackets bool
	abaPatterns := make(map[string]bool)
	babPatterns := make(map[string]bool)

outer:
	for i := 0; i < len(s)-3; i++ {

		// look ahead at next 4 chars...  ignore brackets
		for j := i; j < i+3; j++ {
			if char := rune(s[j]); char == '[' {
				isBetweenBrackets = true
				continue outer
			} else if char == ']' {
				isBetweenBrackets = false
				continue outer
			}
		}

		// cannot have an "aaa" pattern
		if s[i] == s[i+1] {
			continue outer
		}

		// "abc" --> "a" == "c" ?
		if s[i] == s[i+2] {
			var pattern string
			for j := i; j < i+3; j++ {
				pattern = pattern + string(s[j])
			}

			opposite := string(s[i+1]) + string(s[i]) + string(s[i+1])

			if isBetweenBrackets {
				if _, ok := abaPatterns[opposite]; ok {
					return true
				}
				babPatterns[pattern] = true
			} else {
				if _, ok := babPatterns[opposite]; ok {
					return true
				}
				abaPatterns[pattern] = true
			}

			fmt.Println("aba", abaPatterns)
			fmt.Println("bab", babPatterns)
			continue outer
		}
	}
	return false
}

func p1(input []string) {
	var count int
	for _, line := range input {
		if matchPatternTLS(line) {
			count++
		}
	}
	fmt.Println(count)
}

func p2(input []string) {
	var count int
	for _, line := range input {
		if matchPatternSSL(line) {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	input := readFile()
	p1(input)
	p2(input)
}
