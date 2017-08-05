// Solutions to set1 of Cryptopals
package set1

import (
	"encoding/hex"
	"errors"
	"math"
	"strings"
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

// Xor XORs the bytes of the two strings and returns the resulting string
// after converting the bytes back to a string.
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

// SingleCharXor takes the given rune, replicates it so that it is as long
// as the input string, and XORs the two strings together, returning the
// resulting string
func SingleCharXor(s string, r rune) (string, error) {

}

// XorCharMap takes a hex encoded input string that has been XOR'ed with
// a single character and returns the most likely original string, based
// on a character frequency analysis of the decoded string.
func XorCharMap(s string) (string, error) {
	charFrequencies := [26]rune{'e', 't', 'a', 'o', 'i', 'n', 's',
		'h', 'd', 'l', 'c', 'u', 'm', 'w',
		'f', 'g', 'y', 'p', 'b', 'v', 'k',
		'j', 'x', 'q', 'z'}

	// From Wikipedia page https:://en.wikipedia.org/wiki/Letter_frequency
	relativeFrequencies := map[rune]float64{
		'e': 0.12702, 't': 0.09056, 'a': 0.08167, 'o': 0.07507, 'i': 0.06966,
		'n': 0.06749, 's': 0.06327, 'h': 0.06094, 'r': 0.05987, 'd': 0.04253,
		'l': 0.04025, 'c': 0.02782, 'u': 0.02758, 'm': 0.02406, 'w': 0.02360,
		'f': 0.02228, 'g': 0.02015, 'y': 0.01974, 'p': 0.01929, 'b': 0.01492,
		'v': 0.00978, 'k': 0.00772, 'j': 0.00153, 'x': 0.00150, 'q': 0.00095,
		'z': 0.00074}

	decodedStrings := make([]string, 26)
	for i, c := range charFrequencies {
		decodedInput, err := hex.DecodeString(s)
		if err != nil {
			return "", err
		}

		// Here we make a string of the given character repeated enough times so
		// that the resulting string has the same length as `s`. We then XOR the
		// two strings and store there result
		cString := strings.Repeat(string(c), utf8.RuneCountInString(string(decodedInput)))
		hexCString := hex.EncodeToString([]byte(cString))
		decodedString, err := Xor(s, hexCString)
		decodedStrings[i] = decodedString
		if err != nil {
			return "", err
		}
	}

	// Now we want to perform some frequency analysis on the `decodedStrings`
	// The basic idea here is to rank each string by its absolute difference (in
	// terms of distance from expected frequency distribution) from the predefined
	// distribution above (`relativeFrequencies`)
	scoredStrings := make(map[string]float64)
	for _, s := range decodedStrings {
		bs, err := hex.DecodeString(s)
		ds := strings.ToLower(string(bs))
		if err != nil {
			return "", err
		}

		// First get a frequency list for `s`
		sFreqMap := make(map[rune]float64)
		for _, c := range ds {
			sFreqMap[c] += 1
		}

		// Actually calc the frequencies here
		sLen := utf8.RuneCountInString(ds)
		for key, value := range sFreqMap {
			sFreqMap[key] = value / float64(sLen)
		}

		// Then compare the distance from s to `relativeFrequencies`
		score := 0.0
		for k, v := range sFreqMap {
			if _, ok := relativeFrequencies[k]; !ok {
				score += 0.1
			}
			score += math.Abs(v - relativeFrequencies[k])
		}
		scoredStrings[ds] = score
	}

	// Get best result
	result := ""
	minScore := math.MaxFloat64
	for k, v := range scoredStrings {
		if v < minScore {
			minScore = v
			result = k
		}
	}

	return result, nil
}
