package parsing

import (
	"strings"
	"unicode"
)

func (ad *AvroDict) IsVowel(r rune) bool {
	return strings.ContainsRune(ad.Data.Vowel, unicode.ToLower(r))
}

func (ad *AvroDict) IsConsonant(r rune) bool {
	return strings.ContainsRune(ad.Data.Consonant, unicode.ToLower(r))
}

func (ad *AvroDict) IsNumber(r rune) bool {
	return strings.ContainsRune(ad.Data.Number, unicode.ToLower(r))
}

func (ad *AvroDict) IsPunctuation(r rune) bool {
	return !(ad.IsVowel(r) || ad.IsConsonant(r))
}

func (ad *AvroDict) IsCaseSensitive(r rune) bool {
	return strings.ContainsRune(ad.Data.CaseSensitive, unicode.ToLower(r))
}

// I have zero clue what this is for.
func (ad *AvroDict) IsExact(needle string, haystack string, start int, end int, matchnot bool) bool {
	return (start >= 0 && end < len(haystack) && haystack[start:end] == (needle)) != matchnot
}

func (ad *AvroDict) FixStringCase(text *string) string {
	fixed := make([]rune, 0)
	// ignore the fecking index
	for _, char := range *text {
		if ad.IsCaseSensitive(char) {
			fixed = append(fixed, char)
		} else {
			fixed = append(fixed, unicode.ToLower(char))
		}
	}
	return string(fixed)
}
