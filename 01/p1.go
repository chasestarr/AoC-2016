package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readInput() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), ", ")
}

type taxi struct {
	Direction int
}

func (t *taxi) rotate(turn string) int {
	// change 'direction' from rotation
	if turn == "R" {
		t.Direction++
	} else if turn == "L" {
		t.Direction--
	}

	// reset back to 'north'
	if t.Direction > 3 {
		t.Direction = 0
	}

	// reset back to 'west'
	if t.Direction < 0 {
		t.Direction = 3
	}

	return t.Direction
}

func main() {
	t := taxi{Direction: 0}
	for _, i := range readInput() {
		turn := i[:1]
		dist := i[1:]
		fmt.Println(dist)

		r := t.rotate(turn)
		fmt.Printf("index %d \n", r)
	}
}
