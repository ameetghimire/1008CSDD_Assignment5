package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func CountVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	testStrings := []string{"racecar", "hello", "csdd", ""}
	for _, str := range testStrings {
		fmt.Printf("IsPalindrome(%q) = %v\n", str, IsPalindrome(str))
		fmt.Printf("CountVowels(%q) = %d\n", str, CountVowels(str))
	}
}
