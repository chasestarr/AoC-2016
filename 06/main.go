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

func main() {
	input := readFile()
	memo := make(map[rune][]int)

	for _, line := range input {
		for i, char := range line {
			if _, ok := memo[char]; ok {
				memo[char][i]++
			} else {
				memo[char] = make([]int, 8)
			}
		}
	}

	highestCounts := make([]int, 8)
	runes := make([]rune, 8)

	for key := range memo {
		for i, count := range memo[key] {
			if count > highestCounts[i] {
				highestCounts[i] = count
				runes[i] = key
			}
		}
	}

	var output string
	for _, r := range runes {
		output += string(r)
	}
	fmt.Println(output)
}
