package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 3 {
		return
	}
	if isAnagram(os.Args[1], os.Args[2]) {
		fmt.Printf("%s and %s is anagram\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s and %s is not anagram\n", os.Args[1], os.Args[2])
	}

}

func isAnagram(s1, s2 string) bool {
	if utf8.RuneCountInString(s1) != utf8.RuneCountInString(s2) {
		return false
	}

	// Using runeCount to keep track of the count of rune occurrences
	runeCount := make(map[rune]int)

	// Range through the first string and count runes
	for _, r := range s1 {
		runeCount[r]++
	}
	// Range through the second string and remove rune counts
	for _, r := range s2 {
		runeCount[r]--
	}
	// Range through the runeCount map and determine if any values are not 0
	for _, v := range runeCount {
		if v != 0 {
			return false
		}
	}
	return true
}
