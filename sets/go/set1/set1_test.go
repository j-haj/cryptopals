package set1

import (
	"bytes"
	"testing"
)

func TestGetBytesFromHex(t *testing.T) {
	base16_map := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11,
		'C': 12, 'D': 13, 'E': 14, 'F': 15}

	for k, v := range base16_map {
		expected := make([]byte, 1)
		expected[0] = byte(v)
		actual, err := GetBytesFromHex("0" + string(k))
		if err != nil {
			t.Errorf("error - %s\n", err)
		}
		if !bytes.Equal(actual, expected) {
			t.Errorf("Expected %v to be converted to %v. Got: %v\n",
				k, expected, actual)
		}
	}

	actual, err := GetBytesFromHex("012F")
	if err != nil {
		t.Errorf("error - %s\n", err)
	}
	expected := make([]byte, 2)
	expected[0] = byte(1)
	expected[1] = byte(47)
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected 12F to be converted to %v. Got: %v\n",
			expected, actual)
	}
}

func TestSet1Challenge1(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b6520612" +
		"0706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11" +
		"c2hyb29t"
	actual, err := HexToBase64(input)
	if err != nil {
		t.Errorf("error - %s\n", err)
	}
	if expected != actual {
		t.Errorf("Attempted to convert %s to base64\n\tExpected: %s\n\tGot: %s\n",
			input, expected, actual)
	}
}
