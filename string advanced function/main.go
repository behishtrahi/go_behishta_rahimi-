package main // P1Q1

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pickVocals("alterra academy"))                    // aea aae
	fmt.Println(pickVocals("i love programming"))                 // i oe oai
	fmt.Println(pickVocals("go is awesome programming language")) // o i aeoe oai auae
}

func pickVocals(sentence string) string {
	vowels := "aeiouAEIOU"
	result := ""

	for _, char := range sentence {
		if strings.ContainsRune(vowels, char) {
			result += string(char)
		}
	}
	return ""
}
