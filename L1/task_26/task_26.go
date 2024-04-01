package main

import (
	"fmt"
	"strings"
)

func uniqueCharacters(str string) bool {

	str = strings.ToLower(str)
	charMap := make(map[rune]bool)

	for _, char := range str {
		// Если символ уже встречался, то возвращаем false
		if charMap[char] {
			return false
		}
		// В противном случае отмечаем символ как встреченный
		charMap[char] = true
	}
	// Если прошли по всей строке и не встретили повторяющихся символов, возвращаем true
	return true
}

func main() {

	testStrings := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range testStrings {
		fmt.Printf("%s - %t\n", str, uniqueCharacters(str))
	}
}
