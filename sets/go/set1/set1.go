package set1

import (
	"encoding/hex"
	"errors"
	"unicode/utf8"
)

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
	const base64_chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwx" +
		"yz0123456789+/"
	length := utf8.RuneCountInString(s)

	// Pad input
	inputPad, padLength := base64Pad(s)

	bytes, err := hex.DecodeString(s + inputPad)
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

func Xor(s1, s2 string) (string, error) {
	if len(s1) != len(s2) {
		return "", errors.New("input strings must be same length")
	}
	bytes1, err := hex.DecodeString(s1)
	if err != nil {
		return "", err
	}
	bytes2, err := hex.DecodeString(s2)
	if err != nil {
		return "", err
	}
	length := len(bytes1)
	resBytes := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		resBytes[i] = bytes1[i] ^ bytes2[i]
	}
	output := hex.EncodeToString(resBytes)
	if err != nil {
		return "", err
	}
	return output, nil
}
