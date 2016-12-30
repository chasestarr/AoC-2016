package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	key rune
	val int
}

type byVal []pair

func (b byVal) Len() int      { return len(b) }
func (b byVal) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byVal) Less(i, j int) bool {
	if b[i].val == b[j].val {
		return b[i].key < b[j].key
	}
	return b[i].val > b[j].val
}

func readFile() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), "\n")
}

func checkSum(s string) string {
	memo := make(map[rune]int)
	for _, c := range s {
		if _, ok := memo[c]; ok {
			memo[c] = memo[c] + 1
		} else {
			memo[c] = 0
		}
	}

	rank := []pair{}
	for key := range memo {
		rank = append(rank, pair{key: key, val: memo[key]})
	}

	sort.Sort(byVal(rank))

	var output string
	for i := 0; i < 5; i++ {
		output += string(rank[i].key)
	}

	return output
}

func main() {
	input := readFile()
	var sum int

	for _, line := range input {
		sections := strings.Split(line, "-")
		tail := sections[len(sections)-1]
		head := sections[:len(sections)-1]

		var room string
		for _, s := range head {
			room += s
		}

		cs := checkSum(room)
		check := tail[4:9]
		if cs == check {
			i, _ := strconv.Atoi(tail[:3])
			sum += i
		}
	}
	fmt.Println("==== part 01 ====")
	fmt.Println(sum)
}
