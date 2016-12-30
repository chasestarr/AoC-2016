package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	input := "ojvtpuvg"
	var output string
	var i int

	for len(output) < 8 {
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
		}

		i++
	}

	// part 01 => 4543c154
}
