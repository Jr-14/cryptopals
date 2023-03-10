package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	sOne := "1c0111001f010100061a024b53535009181c"
	bytesOne, err := hex.DecodeString(sOne)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%b\n", bytesOne)

	sTwo := "686974207468652062756c6c277320657965"
	bytesTwo, err := hex.DecodeString(sTwo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%b\n", bytesTwo)

	c := make([]byte, len(bytesOne))
	for i := range bytesOne {
		c[i] = bytesOne[i] ^ bytesTwo[i]
	}
	fmt.Printf("%b\n", c)

	decodedString := hex.EncodeToString(c)
	fmt.Printf("%s\n", decodedString)
}
