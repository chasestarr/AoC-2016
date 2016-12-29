package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func readFile() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), "\n")
}

func validTriangle(ints []int) bool {
	sort.Ints(ints)
	if ints[0]+ints[1] > ints[2] {
		return true
	}
	return false
}

func isTriangle(line string) bool {
	values := strings.Fields(line)
	ints := make([]int, 3)
	for i, val := range values {
		ints[i], _ = strconv.Atoi(val)
	}

	return validTriangle(ints)
}

func isColumnTriangle(lines []string) int {
	triangles := make(map[int][]int)
	for _, line := range lines {
		values := strings.Fields(line)
		for j, val := range values {
			length, _ := strconv.Atoi(val)
			triangles[j] = append(triangles[j], length)
		}
	}

	var count int
	for i := 0; i < 3; i++ {
		if validTriangle(triangles[i]) {
			count++
		}
	}
	return count
}

func main() {
	input := readFile()
	var count int

	for _, line := range input {
		if isTriangle(string(line)) {
			count++
		}
	}

	fmt.Println("======== part 1 =========")
	fmt.Println(count)
	count = 0

	for i := 0; i < len(input); i += 3 {
		lines := make([]string, 3)
		for j := 0; j < 3; j++ {
			lines[j] = input[i+j]
		}

		count += isColumnTriangle(lines)
	}

	fmt.Println("======== part 2 =========")
	fmt.Println(count)
}
