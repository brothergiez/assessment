package anagram

import (
	"fmt"
	"strings"
)

func GroupAnagram(words []string) [][]string {
	anagramGroup := make(map[string][]string)

	for _, word := range words {
		charCount := make([]int, 26)
		for _, ch := range word {
			charCount[ch-'a']++
		}

		var key strings.Builder
		for _, count := range charCount {
			key.WriteString(fmt.Sprintf("%d#", count))
		}

		anagramGroup[key.String()] = append(anagramGroup[key.String()], word)
	}

	result := make([][]string, 0, len(anagramGroup))
	for _, group := range anagramGroup {
		result = append(result, group)
	}

	return result
}
