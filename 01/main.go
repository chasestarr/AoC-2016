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
	Memo      map[string]bool
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

func (t *taxi) move(dist int) (int, int) {
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

	return t.X, t.Y
}

func (t *taxi) distance() int {
	fX := float64(t.X)
	fY := float64(t.Y)
	return int(math.Abs(fX) + math.Abs(fY))
}

func (t *taxi) memo(x int, y int) bool {
	key := strconv.Itoa(x) + "," + strconv.Itoa(y)

	if _, ok := t.Memo[key]; ok {
		return true
	}

	t.Memo[key] = true
	return false
}

func main() {
	t := taxi{Direction: 0, Memo: make(map[string]bool)}
	for e, i := range readInput() {
		turn := i[:1]
		dist, _ := strconv.Atoi(i[1:])

		t.rotate(turn)
		x, y := t.move(dist)

		if t.memo(x, y) {
			fmt.Printf("visited twice: %d, %d, %d \n", x, y, e)
			fmt.Println(t.distance())
		}
	}

	fmt.Println(t.distance())
}
