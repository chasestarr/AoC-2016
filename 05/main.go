package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	input := "ojvtpuvg"
	var count int
	var output string
	var i int
	storage := make([]interface{}, 8)

	for count < 8 {
		data := input + strconv.Itoa(i)
		h := md5.Sum([]byte(data))
		result := fmt.Sprintf("%x", h)

		isValid := true
		for j := 0; j < 5; j++ {
			if string(result[j]) != "0" {
				isValid = false
			}
		}

		if isValid {
			output += string(result[5])
			fmt.Println(output)

			position, _ := strconv.Atoi(string(result[5]))
			value := string(result[6])

			if position < 8 {
				if storage[position] == nil {
					count++
				}
				storage[position] = value
				fmt.Println(storage)
			}
		}

		i++
	}

	// part 01 => 4543c154
	// part 02 => 4050cbbd 0550c4bd
}
