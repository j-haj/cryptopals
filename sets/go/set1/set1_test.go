package set1

import (
	"testing"
)

func TestHexToString(t *testing.T) {
	if ConvertHexStringToVal("A") != 10 {
		t.Error("Expected A converted to 10")
	}

	base16_map := map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11,
		'C': 12, 'D': 13, 'E': 14, 'F': 15}

	for k, v := range base16_map {
		if ConvertHexStringToVal(string(k)) != v {
			t.Errorf("Expected %s to be converted to %d\n", k, v)
		}
	}

	if ConvertHexStringToVal("12F") != 303 {
		t.Errorf("Expected 12F to be converted to 303, Got: %d\n", ConvertHexStringToVal("12F"))
	}
}

func TestHextToBase64(t *testing.T) {
	
}

func TestSet1Challenge1(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	actual := HexToBase64(input)
	if expected != actual {
		t.Errorf("Attempted to convert %s to base64\n\tExpected: %s\n\tGot: %s\n",
			input, expected, actual)
	}
}
