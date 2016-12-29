package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readFile() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), "\n")
}

func isTriangle(line string) bool {
	values := strings.Fields(line)
	ints := make([]int, 3)
	for i, val := range values {
		ints[i], _ = strconv.Atoi(val)
	}

	sort.Ints(ints)
	if ints[0]+ints[1] > ints[2] {
		return true
	}
	return false
}

func main() {
	input := readFile()
	var count int
	for _, line := range input {
		if isTriangle(string(line)) {
			count++
		}
	}

	println(count)
}
