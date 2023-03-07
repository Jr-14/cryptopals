package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	byteString := []byte(hexString)
	fmt.Printf("mystr:\t %v \n", byteString)
	decodedString, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded string:\t %v \n", decodedString)
}
