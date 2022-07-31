package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	// String representation of a hexadecimal value
	const hexString string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(decoded)))
	base64.StdEncoding.Encode(dst, decoded)
	fmt.Println(string(dst))
}