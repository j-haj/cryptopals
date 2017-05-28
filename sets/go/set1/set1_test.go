package set1

import (
	"fmt"
	"testing"
)

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

func TestSet1Challeng2(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	expected := "746865206b696420646f6e277420706c6179"
	actual, err := Xor(input1, input2)
	if err != nil {
		t.Errorf("error - %s\n", err)
	}
	if actual != expected {
		t.Errorf("XORed %s with %s.\n\tExpected: %s\n\tGot: %s\n",
			input1, input2, expected, actual)
	}
}

func TestSet1Challenge3(t *testing.T) {
	const input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	res, _ := XorCharMap(input)
	for x, _ := range res {
		fmt.Printf("%v\n", x)
	}
}
