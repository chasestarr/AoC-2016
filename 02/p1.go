package p1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type digit struct {
	X int
	Y int
}

func (d *digit) move(x int, y int) {
	tempX := d.X + x
	tempY := d.Y + y

	if tempX <= 2 && tempX >= 0 {
		d.X = tempX
	}

	if tempY <= 2 && tempY >= 0 {
		d.Y = tempY
	}
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

func keypad(x int, y int) string {
	keys := [][]int{}
	keys = append(keys, []int{1, 2, 3})
	keys = append(keys, []int{4, 5, 6})
	keys = append(keys, []int{7, 8, 9})

	return strconv.Itoa(keys[x][y])
}

func readInput() []string {
	input, _ := ioutil.ReadFile("input")
	return strings.Split(string(input), "\n")
}

func main() {
	var output string
	input := readInput()
	d := digit{X: 1, Y: 1}

	for _, line := range input {
		for _, char := range line {
			d.move(direction(string(char)))
		}
		output += keypad(d.X, d.Y)
	}
	fmt.Println(output)
}
