package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readInput() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), ", ")
}

type taxi struct {
	Direction int
	X         int
	Y         int
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

func (t *taxi) move(dist int) {
	// north
	if t.Direction == 0 {
		t.Y += dist

		// east
	} else if t.Direction == 1 {
		t.X += dist

		// south
	} else if t.Direction == 2 {
		t.Y -= dist

		// west
	} else {
		t.X -= dist
	}
}

func (t *taxi) distance() int {
	fX := float64(t.X)
	fY := float64(t.Y)
	return int(math.Abs(fX) + math.Abs(fY))
}

func main() {
	t := taxi{Direction: 0}
	for _, i := range readInput() {
		turn := i[:1]
		dist, _ := strconv.Atoi(i[1:])

		t.rotate(turn)
		t.move(dist)
	}

	fmt.Println(t.distance())
}
