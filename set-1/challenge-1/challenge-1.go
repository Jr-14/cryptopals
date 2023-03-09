package main

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Printf("String to Encode:\n%s\n", hexString)
	fmt.Printf("String length %d\n\n", len(hexString))
	decodedString, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("String in bytes:\n%v\n", decodedString)
	fmt.Printf("Byte length: %d\n\n", len(decodedString))

	base64String := base64.RawStdEncoding.EncodeToString(decodedString)
	fmt.Printf("Base 64 String: %s\n", base64String)
}
