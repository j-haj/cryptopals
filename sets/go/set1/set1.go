package set1

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

// GetBytesFromHex takes a string representation of a hex number and
// returns an array of the corresponding bytes
func GetBytesFromHex(s string) ([]byte, error) {
	base16_map := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11,
		'C': 12, 'D': 13, 'E': 14, 'F': 15}
	strLen := len(s)
	if strLen%2 != 0 {
		return make([]byte, 0), errors.New("Hex string must contain an" +
			" even multiple of characters.")
	}
	bytes := make([]byte, strLen/2)
	strRunes := []rune(s)
	for i := strLen; i > 0; i -= 2 {
		curChars := strRunes[i-2 : i]
		bytes[i/2-1] = byte(16*base16_map[unicode.ToUpper(curChars[0])] +
			base16_map[unicode.ToUpper(curChars[1])])
	}
	return bytes, nil
}

// base64Pad returns the appropriate padding and padding size for a given input
// string to be converted to base64
func base64Pad(s string) (string, int) {
	length := len(s)
	padLength := 0
	inputPad := ""
	if length%3 == 1 {
		padLength = 2
		inputPad = "00"
	} else if length%3 == 2 {
		padLength = 1
		inputPad = "0"
	}
	return inputPad, padLength
}

// HexToBase64 takes an input string that is a hex representation and returns
// its base64 representation
func HexToBase64(s string) (string, error) {
	//base64_chars := [64]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	//	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W',
	//	'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	//	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w',
	//	'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	//	'+', '/'}

	const base64_chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwx" +
		"yz0123456789+/"
	length := utf8.RuneCountInString(s)

	// Pad input
	inputPad, padLength = base64Pad(s)

	bytes, err := GetBytesFromHex(s + inputPad)
	if err != nil {
		return "", err
	}

	numBytes := len(bytes)
	result := make([]rune, (length+padLength)*4/6)
	byteIndex := len(result) - 1
	for i := numBytes; i > 0; i -= 3 {
		byteTriple := bytes[i-3 : i]
		val := (2<<15)*int(byteTriple[0]) + (2<<7)*int(byteTriple[1]) +
			int(byteTriple[2])
		for j := 0; j < 4; j++ {
			c := (val >> uint(6*j)) & 0x3F
			result[byteIndex-j] = rune(base64_chars[c])
		}
		byteIndex -= 4
	}
	return string(result), nil
}
