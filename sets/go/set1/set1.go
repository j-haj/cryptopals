package set1

import (
	"fmt"
	"math"
	"unicode/utf8"
)

// HexToBase64 takes an input string that is a hex representation and returns
// its base64 representation
func HexToBase64(s string) string {
	base64_chars := [64]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W',
		'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
		'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'0', '+', '/'}

	val = ConvertHexToStringVal(s)
	length := utf8.RuneCountInString(s)
	// Pad input
	if length%3 == 1 {
		s += "00"
		length += 2
	} else if length%3 == 2 {
		s += "0"
		length += 1
	}

	// Construct base64 string
	output := make([]rune, length/6)
	for i := 0; i < length/6; i++ {
		current_bits = val & (2 << 6)
		val >> 6
		output[length/6-i-1] = base64_chars[current_bits]
	}
	return string(output)
}

// Converts a hex input string into its integer representation
func ConvertHexStringToVal(s string) int {
	base16_map := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11,
		'C': 12, 'D': 13, 'E': 14, 'F': 15}
	exponent := 0
	result := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		result += base16_map[c] * math.Pow(16, exponent)
	}
	return result
}
