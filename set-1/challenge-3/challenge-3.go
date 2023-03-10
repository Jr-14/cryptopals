package main

import (
	"encoding/hex"
	"fmt"
)

const englishCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "

func main() {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	byteString, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(englishCharacters); i++ {
		fmt.Printf("%c:\t", englishCharacters[i])
		charByte := byte(englishCharacters[i])
		cBytes := make([]byte, len(byteString))
		for j := 0; j < len(byteString); j += 1 {
			xoredBytes := charByte ^ byteString[j]
			fmt.Printf("%c", xoredBytes)
			cBytes = append(cBytes, xoredBytes)
		}
		fmt.Println("")
	}
}
