package main

import (
	"fmt"
	"strconv"
)

func main() {
	message := [...]string{
		"01001000011010010010110000100000011011010111100100100000011011100110000101101101",
		"01100101001000000110100101110011001000000100001001100001011001110110100101101100",
		"01101100001000010010000001001001001000000110101101100101011001010111000000100000",
		"01101001011101000010000001110010011001010110000101101100001000010110100001100001",
		"01110000011100000111100100100000011000100110100101110010011101000110100001100100",
		"0110000101111001001000000110001101100001011100100110001001110011",
	}

	for _, line := range message {
		pos := 0
		for pos < len(line) {
			character := line[pos : pos+8]
			i, _ := strconv.ParseInt(character, 2, 8)
			fmt.Printf("%s %c\n", character, i)
			pos += 8
		}
	}
}
