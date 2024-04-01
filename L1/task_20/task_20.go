package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input)
	reversed := make([]string, len(words))

	for i, word := range words {
		reversed[len(words)-1-i] = word
	}

	return strings.Join(reversed, " ")
}

func main() {
	input := "snow dog sun"
	output := reverseWords(input)
	fmt.Println(output)
}
