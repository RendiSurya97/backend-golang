package online_test_answer

import "strings"

// func findFirstStringInBracket(str string) string {
// 	if len(str) > 0 {
// 		indexFirstBracketFound := strings.Index(str, "(")
// 		if indexFirstBracketFound >= 0 {
// 			runes := []rune(str)
// 			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
// 			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
// 			if indexClosingBracketFound >= 0 {
// 				runes := []rune(wordsAfterFirstBracket)
// 				return string(runes[1 : indexClosingBracketFound-1])
// 			} else {
// 				return ""
// 			}
// 		} else {
// 			return ""
// 		}
// 	} else {
// 		return ""
// 	}
// 	return ""
// }

func findFirstStringInBracket(str string) string {

	if len(str) == 0 {
		return ""
	}

	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	runes := []rune(str)
	wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	runes = []rune(wordsAfterFirstBracket)
	return string(runes[1 : indexClosingBracketFound-1])
}
