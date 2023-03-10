// Package refine ...
package refine

import (
	"unicode/utf8"

	"github.com/TessorNetwork/cerberus/pkg/schema"
)

// Cookie refines the given cookie is not autogenerated.
func Cookie(cookie *schema.V1Cookie) bool {
	runeCount := utf8.RuneCount([]byte(cookie.Value))

	if runeCount < 3 {
		return false
	}

	if runeCount >= 6 && validateRunesSame(cookie.Value) {
		return false
	}

	return true
}

// SearchHistory refines the given search history item is not autogenerated.
func SearchHistory(search *schema.V1SearchHistory) bool {
	runeCount := utf8.RuneCount([]byte(search.Query))

	if runeCount < 3 {
		return false
	}

	if runeCount >= 6 && validateRunesSame(search.Query) {
		return false
	}

	return true
}

func validateRunesSame(str string) bool {
	var first rune

	for pos, char := range str {
		if pos == 0 {
			first = char
		}
		if first != char {
			return false
		}
	}

	return true
}
