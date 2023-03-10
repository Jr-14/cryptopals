package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

const englishCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "

type ScoredString struct {
	BestString string
	Score      int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scoredBytesArray := make([]ScoredString, 0)

	for scanner.Scan() {
		textString := scanner.Text()
		bytesArray, byteScore := decryptString(textString)
		decryptedString := string(bytesArray)
		scoredByte := ScoredString{decryptedString, byteScore}
		scoredBytesArray = append(scoredBytesArray, scoredByte)
	}

	bestString := scoredBytesArray[0].BestString
	bestByteScore := scoredBytesArray[0].Score
	for i := 1; i < len(scoredBytesArray); i++ {
		if scoredBytesArray[i].Score >= bestByteScore {
			bestString = scoredBytesArray[i].BestString
			bestByteScore = scoredBytesArray[i].Score
		}
	}
	fmt.Println(bestString)
	fmt.Println(bestByteScore)
}

func decryptString(s string) (bestByteScore []byte, byteScore int) {
	byteString, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	bestByteScore = byteString
	currentScore := 0

	for i := 0; i < len(englishCharacters); i++ {
		charByte := byte(englishCharacters[i])
		decryptedByte := make([]byte, 0)
		for j := 0; j < len(byteString); j += 1 {
			xoredBytes := charByte ^ byteString[j]
			decryptedByte = append(decryptedByte, xoredBytes)
		}
		score := scoreBytesArray(decryptedByte)
		if score >= currentScore {
			bestByteScore = decryptedByte
			currentScore = score
		}
	}
	fmt.Printf("Best bytes: %d\n", bestByteScore)
	return
}

func scoreBytesArray(bytes []byte) (score int) {
	for _, b := range bytes {
		score += getByteScore(b)
	}
	return
}

func getByteScore(b byte) int {
	// ASCII Uppercase characters
	if b >= 'A' && b <= 'Z' {
		return 3
	}
	// ASCII for lowercase characters
	if b >= 'a' && b <= 'z' {
		return 4
	}
	// ASCII for single quote
	if b == '\'' {
		return 1
	}
	// ASCII for space character
	if b == ' ' {
		return 1
	}
	return 0
}
