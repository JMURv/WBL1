package main

import (
	"fmt"
	"strings"
)

func checkChars(str string) bool {
	charSet := make(map[rune]bool)
	for _, char := range strings.ToLower(str) {
		if _, ok := charSet[char]; ok {
			return false
		}
		charSet[char] = true
	}
	return true
}

func main() {
	fmt.Println(checkChars("abcd"))
	fmt.Println(checkChars("abCdefAaf"))
	fmt.Println(checkChars("aabcd"))
}
