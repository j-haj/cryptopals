package main

import (
	"encoding/hex"
	"fmt"
	"github.com/j-haj/cryptopals/sets/go/set1"
)

func TestSet1Challenge3() {
	const input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	res, err := set1.XorCharMap(input)
	if err != nil {
		fmt.Printf("error - %v\n", err)
		return
	}
	for _, x := range res {
		decoded, _ := hex.DecodeString(x)
		fmt.Printf("%v\n", string(decoded))
	}
}

func main() {
	TestSet1Challenge3()
}
