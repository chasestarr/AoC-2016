package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type digit struct {
  Keypad [][]interface{}
	X int
	Y int
}

func newDigit() *digit {
  d := digit{X: 2, Y: 0}
  d.Keypad = append(d.Keypad, []interface{}{0, 0, 1, 0, 0})
  d.Keypad = append(d.Keypad, []interface{}{0, 2, 3, 4, 0})
  d.Keypad = append(d.Keypad, []interface{}{5, 6, 7, 8, 9})
  d.Keypad = append(d.Keypad, []interface{}{0, "A", "B", "C", 0})
  d.Keypad = append(d.Keypad, []interface{}{0, 0, "D", 0, 0})

  return &d
}

func (d *digit) move(x int, y int) {
	tempX := d.X + x
	tempY := d.Y + y

	if tempX > 4 || tempX < 0 {
		return
	}

	if tempY > 4 || tempY < 0 {
		return
	}

  dest := d.Keypad[tempX][tempY]
  if dest == 0 {
    return
  }

  d.X = tempX
  d.Y = tempY
}

func direction(char string) (int, int) {
	if char == "U" {
		return -1, 0
	} else if char == "D" {
		return 1, 0
	} else if char == "R" {
		return 0, 1
	}

	return 0, -1
}

func (d *digit) keypad(x int, y int) string {
  keys := d.Keypad

  if _, ok := keys[x][y].(int); ok {
    return strconv.Itoa(keys[x][y].(int))
  }
  return keys[x][y].(string)
}

func readInput() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), "\n")
}

func main() {
	var output string
	input := readInput()
  d := newDigit()

	for _, line := range input {
		for _, char := range line {
			d.move(direction(string(char)))
		}
		output +=  d.keypad(d.X, d.Y)
	}
	fmt.Println(output)
}
