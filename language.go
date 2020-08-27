package go_utils

import "regexp"

// IsValidFullWidthKatakana checks if the input is valid full width katakana
func IsValidFullWidthKatakana(source string) bool {
	var validKana = regexp.MustCompile(`[ァ-ン]`)
	return validKana.MatchString(source)
}
