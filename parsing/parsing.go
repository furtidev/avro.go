package parsing

import (
	_ "embed"
	"strings"
)

var (
	//go:embed avrodict.json
	data []byte
)

// eughhhhh i hate this
func (ad *AvroDict) Parse(text *string) string {
	fixed_text := ad.FixStringCase(text)
	var output string
	for cur := 0; cur < len(fixed_text); cur++ {
		matched := false
		replace := true
		for _, pat := range ad.Data.Patterns {
			end := cur + len(pat.Find)
			if (end <= len(fixed_text)) && fixed_text[cur:end] == pat.Find {
				previous := cur - 1
				for _, rule := range pat.Rules {
					chk := 0
					for _, match := range rule.Matches {
						if match.Type == "suffix" {
							chk = end
						} else {
							chk = previous
						}
						negative := false
						if strings.HasPrefix(match.Scope, "!") {
							match.Scope = match.Scope[1:]
							negative = true
						}

						// And let the matching begin! (Go edition)
						switch match.Scope {
						case "punctuation":
							if ((chk < 0 && match.Type == "prefix") || (chk >= len(fixed_text) && match.Type == "suffix") && ad.IsVowel(rune(fixed_text[chk]))) == negative {
								replace = false
							}
						case "vowel":
							if (((chk >= 0 && match.Type == "prefix") || (chk < len(fixed_text) && match.Type == "suffix")) && ad.IsVowel(rune(fixed_text[chk]))) == negative {
								replace = false
							}
						case "consonant":
							if (((chk >= 0 && match.Type == "prefix") || (chk < len(fixed_text) && match.Type == "suffix")) && ad.IsConsonant(rune(fixed_text[chk]))) == negative {
								replace = false
							}
						case "exact":
							var exact_cur, exact_end int
							if match.Type == "prefix" {
								exact_cur = cur - len(match.Value)
								exact_end = cur
							} else {
								exact_cur = end
								exact_end = end + len(match.Value)
							}

							if !ad.IsExact(match.Value, fixed_text, exact_cur, exact_end, negative) {
								replace = false
							}
						}
					}
					if replace {
						output += rule.Replace
						cur = end - 1
						matched = true
						break
					}
				}
				if matched {
					break
				}

				output += pat.Replace
				cur = end - 1
				matched = true
				break
			}
		}
		if !matched {
			output += string(fixed_text[cur])
		}
	}
	return output
}
