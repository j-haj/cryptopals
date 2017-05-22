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
		if ConvertHexStringToVal(k) != v {
			t.Errorf("Expected %s to be converted to %d\n", k, v)
		}
	}

	if ConvertHexStringToVal("12F") != 303 {
		t.Error("Expected 12F to be converted to 303")
	}
}
