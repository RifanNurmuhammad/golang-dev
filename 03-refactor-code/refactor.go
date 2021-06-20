package refactor

import (
	"strings"
)

//findFirstStringInBracket
func findFirstStringInBracket(str string) string {
	lenStr := len(str)

	// validate length
	if lenStr < 0 {
		return ""
	}

	// get index from first bracket
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	runes := []rune(str)
	wordsAfterFirstBracket := string(runes[indexFirstBracketFound:lenStr])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound >= 0 {
		runes := []rune(wordsAfterFirstBracket)
		return string(runes[1:indexClosingBracketFound])
	}
	return ""
}
