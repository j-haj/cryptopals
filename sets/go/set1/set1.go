package set1

import (
	"fmt"
	"math"
	"unicode"
	"unicode/utf8"
)

// Converts a hex input string into its integer representation
func ConvertHexStringToVal(s string) int {
	base16_map := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11,
		'C': 12, 'D': 13, 'E': 14, 'F': 15}
	exponent := float64(len(s) - 1)
	result := 0
	for _, runeValue := range s {
		result += base16_map[unicode.ToUpper(runeValue)] * int(math.Pow(16.0, exponent))
		exponent -= 1.0
	}
	return result
}

// HexToBase64 takes an input string that is a hex representation and returns
// its base64 representation
func HexToBase64(s string) string {
	base64_chars := [64]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W',
		'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
		'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'0', '+', '/'}

	length := utf8.RuneCountInString(s)
	// Pad input
	padLength := 0
	if length%3 == 1 {
		s += "00"
		length += 2
		padLength = 2
	} else if length%3 == 2 {
		s += "0"
		length += 1
		padLength = 1
	}
	val := ConvertHexStringToVal(s)
	// Construct base64 string
	output := make([]rune, length/6)
	if padLength == 2 {
		output[length/6 - 1] = '='
		output[length/6 - 2] = '='
	} else if padLength == 1 {
		output[length/6 - 1] = '='
	}
	for i := 0; i < length/6 - padLength; i++ {
		currentBits := val & (2 << 6)
		fmt.Printf("[%d]val: %d\tcurrentBits: %d\n",i, val, currentBits)
		val = val >> 6
		output[length/6-i-1] = base64_chars[currentBits]
	}
	return string(output)
}

