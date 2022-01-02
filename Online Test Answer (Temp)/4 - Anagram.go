package online_test_answer

import "sort"

func SortWord(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(a, b int) bool {
		return runes[a] < runes[b]
	})
	return string(runes)
}

func FindAnagram(words []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, word := range words {
		sortedWord := SortWord(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	var groupedWords [][]string
	for _, group := range anagramMap {
		groupedWords = append(groupedWords, group)
	}
	return groupedWords
}
