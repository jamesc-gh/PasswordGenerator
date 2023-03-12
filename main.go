package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func GetAllRunes() string {
	runes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	runes += "abcdefghijklmnopqrstuvwxyz"
	runes += "0123456789"
	runes += "!@#$%^&*()_-+={}[]|\\:;\"'<>,.?/`~"

	return runes
}

func GetExcludedRunes(exclusionString string) map[rune]bool {
	fmt.Println("excludedRunes=" + exclusionString)

	excludedRunes := make(map[rune]bool)
	for _, r := range exclusionString {
		excludedRunes[r] = true
	}

	return excludedRunes
}

func GetIncludedRunes(allRunes string, excludedRunes map[rune]bool) string {
	if len(excludedRunes) == 0 {
		return allRunes
	}

	var includedRunes []rune
	for _, r := range allRunes {
		if excludedRunes[r] {
			continue
		}
		includedRunes = append(includedRunes, r)
	}

	return string(includedRunes)
}

func GetRandRuneSequences(runes string, linesToGenerate int, runesPerLine int) []string {
	sequences := make([]string, linesToGenerate+2)
	numRunes := big.NewInt(int64(len(runes)))

	sequences[0] = GetSequenceMarkers(runesPerLine)
	for i := 0; i < linesToGenerate; i++ {
		sequence := make([]rune, runesPerLine)
		for j := 0; j < runesPerLine; j++ {
			nBig, err := rand.Int(rand.Reader, numRunes)
			if err != nil {
				panic(err)
			}
			sequence[j] = rune(runes[nBig.Int64()])
		}
		sequences[i+1] = string(sequence)
	}
	sequences[linesToGenerate+1] = GetSequenceMarkers(runesPerLine)

	return sequences
}

func GetSequenceMarkers(runesPerLine int) string {
	result := ""
	for i := 1; i <= runesPerLine; i++ {
		if i%5 == 0 {
			result += "|"
		} else {
			result += "."
		}
	}
	return result
}

func main() {
	runes := GetAllRunes()
	if len(os.Args) > 1 {
		excludedRunes := GetExcludedRunes(os.Args[1])
		runes = GetIncludedRunes(runes, excludedRunes)
	}
	fmt.Println("includedRunes=" + runes)

	randomSequences := GetRandRuneSequences(runes, 5, 100)
	fmt.Println("randomSequences=")
	for _, s := range randomSequences {
		fmt.Println(s)
	}
}
